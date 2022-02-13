package main

import (
    "github.com/gin-gonic/gin"
    controller "services/article/controller"

    "testing"
    "net/http"
    "net/http/httptest"
    "io/ioutil"
	"fmt"
)

func main() {
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		fmt.Println("Unable to start:", err)
	}
}

func engine() *gin.Engine {
	
	api := gin.New()
    api.GET("query/article", controller.Get)

    api.GET("query/health-check", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "data": "OK",
      })
    })

	return api
}
  
  func article(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("health"))
  }
  
  func Test_article(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/query/article", nil)
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