  package routes

  import (
    "github.com/gin-gonic/gin"
    controller "services/article/controller"

    "testing"
    "net/http"
    "net/http/httptest"
    "io/ioutil"
  )

  func Route() {

    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    api := router.Group("/v1")

    api.POST("/article", controller.Add)
    api.PATCH("/article", controller.Update)
    api.GET("/article", controller.Get)

    api.GET("/health-check", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "data": "OK",
      })
    })

    router.Run(":8080")
  }

  func article(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("health"))
  }
  
  func Test_article(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/v1/article", nil)
    res := httptest.NewRecorder()
   
    article(res, req)
   
    // Instead of lines below, you can just use res.Code and res.Body.String() directly.
   
    result := res.Result()
   
    body, err := ioutil.ReadAll(result.Body)
    if err != nil {
      t.Fatal(err)
    }
    result.Body.Close()
   
    if http.StatusOK != result.StatusCode  {
      t.Error("expected", http.StatusOK, "got", result.StatusCode)
    }
  
    if "health" != string(body) {
      t.Error("expected health got", string(body))
    }
  }