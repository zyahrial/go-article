  package routes

  import (
    "github.com/gin-gonic/gin"
    controller "services/article/controller"
  )

  func Route() {

    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    api := router.Group("/v1")

    api.POST("/article", controller.Add)
    api.PATCH("/article", controller.Update)
    api.GET("/article", controller.Get)

    router.Run(":8080")
  }