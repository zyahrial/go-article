package routes

import (
//   "time"
//   "fmt"
  // "os"
  "github.com/gin-gonic/gin"
  message "query/article/message_broker"
  unitTest "github.com/Valiben/gin_unit_test"

  controller "query/article/controller"
  "net/http"
)

func Route() {
	router := gin.New()
	api := router.Group("query")

	api.GET("listen/message", message.Consume)
	api.GET("article", controller.Get)

	api.GET("health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"data": "OK",
		})
	})

	router.Run(":8080")
	unitTest.SetRouter(router)	
}