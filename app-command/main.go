package main

import (
    "github.com/gin-gonic/gin"
    controller "command/article/controller"
    "command/article/db/database"
    "command/article/db/migrations"
    unitTest "github.com/Valiben/gin_unit_test"

    _ "github.com/jinzhu/gorm/dialects/postgres"
    "time"
    "net/http"
	  "fmt"
)

func main() {
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8081"); err != nil {
		fmt.Println("Unable to start:", err)
	}
}

func engine() *gin.Engine {

	  fmt.Printf("Started at : %3v \n", time.Now())

    database.InitDB()

    migrations.Migrate()
    // defer database.DBCon.Close()

    gin.SetMode(gin.ReleaseMode)
    	
	  api := gin.Default()
    api.POST("command/article", controller.Add)

    api.GET("/health-check", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "data": "OK",
      })
    })
  unitTest.SetRouter(api)
	return api
}