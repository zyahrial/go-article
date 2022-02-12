package controller

import (
	// "time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "encoding/json"
	// "strconv"
	// "fmt"
	"net/http"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	models "services/article/models"
)

func Get(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	keyword := c.Query("query")
	author := c.Query("author")
	
	article := models.Article{}
	coll := mgm.Coll(&article)

	result := []models.ShowArticle{}

	if author != ""{
		_ = coll.SimpleFind(&result, bson.M{"author":author}, &options.FindOptions{
			Sort: bson.D{{"created_at", 1}},
		})
	}else if keyword != ""{
		_ = coll.SimpleFind(&result, bson.M{"$text": bson.M{"$search": keyword}}, &options.FindOptions{
			Sort: bson.D{{"created_at", 1}},
		})
	}else{
		var limit int64 = 100
		_ = coll.SimpleFind(&result, bson.M{} , &options.FindOptions{
			Limit: &limit,
			Sort: bson.D{{"created_at", 1}},
		})
	}

	c.JSON(http.StatusOK, gin.H{
	  "data": result ,
	})
}