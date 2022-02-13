package controller

import (
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"

	models "command/article/models"
	"command/article/db/database"
	rmq "command/article/message_broker"

	"strconv"
	"github.com/go-redis/redis"
	"fmt"
)

func Add(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var article models.Article

    if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All field is required"})
        return
    }

	loc, _ := time.LoadLocation("Asia/Jakarta")
    now := time.Now().In(loc)
	
	addArticle := models.Article{Author: article.Author, Tittle: article.Tittle, Body: article.Body, Created_at: now}

	if err := database.DBCon.Create(&addArticle).Error; err != nil {
		fmt.Printf("error add : %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// d := models.ShowWallet{add.Author,add.Tittle,add.Body,add.CreatedAt}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
    })

	res, err := json.Marshal(addArticle)

	var m models.ShowArticle	
	if err := json.Unmarshal(res, &m); err != nil {
		panic(err)
	}

	id := m.ID
	s := strconv.FormatInt(id, 10)

	err = client.Set(s, res, 0).Err()
    if err != nil {
        fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
    }

	queue := "QueueAddArticle"
	send := rmq.Publish(res, queue)

	c.JSON(http.StatusOK, gin.H{"status": "success", "data" : send})
}