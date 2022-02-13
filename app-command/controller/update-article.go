package controller

// import (
// 	"time"
// 	"github.com/gin-gonic/gin"
// 	"fmt"

// 	"net/http"
// 	"github.com/kamva/mgm/v3"

// 	models "services/article/models"
// )

// func Update(c *gin.Context){
// 	c.Header("Content-Type", "application/json")
// 	c.Header("Access-Control-Allow-Origin", "*")
  
// 	var data models.ShowArticle
// 	c.BindJSON(&data)

// 	id := data.Id
// 	author := data.Author
// 	tittle := data.Tittle
// 	body := data.Body

// 	article := models.Article{}
// 	coll := mgm.Coll(&article)

// 	_ = coll.FindByID(id, &article)

// 	loc, _ := time.LoadLocation("Asia/Jakarta")
//     now := time.Now().In(loc)

// 	article.Author = author
// 	article.Tittle = tittle
// 	article.Body = body
// 	article.UpdatedAt = now
// 	err := coll.Update(&article)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 	  "data": "article has been updated" ,
// 	})
// }