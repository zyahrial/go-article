package controller

import (
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	models "services/article/models"
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

	store := models.NewArictle(article.Author, article.Tittle, article.Body)
	succ := mgm.Coll(store).Create(store)
	
	if succ != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insert data has failed, connection aborted!"})
        return
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
    })

	res, err := json.Marshal(store)

	var m models.ShowArticle
	
	if err := json.Unmarshal(res, &m); err != nil {
		panic(err)
	}

	id := m.Id

	fmt.Println(id)

	err = client.Set(id, res, 0).Err()
    if err != nil {
        fmt.Println(err)
    }

	c.JSON(http.StatusOK, gin.H{"status": "success", "data" : store})
}