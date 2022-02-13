package main

import (
    "github.com/gin-gonic/gin"
    controller "command/article/controller"
    "command/article/db/database"
    "command/article/db/migrations"

    _ "github.com/jinzhu/gorm/dialects/postgres"
    "time"
    "testing"
    "net/http"
    "net/http/httptest"
    "io/ioutil"
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
    	
	  api := gin.New()
    api.POST("command/article", controller.Add)

    api.GET("/health-check", func(c *gin.Context) {
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
    req := httptest.NewRequest(http.MethodGet, "command/article", nil)
    res := httptest.NewRecorder()
   
    article(res, req)
      
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