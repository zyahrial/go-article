package main

import (
    // "github.com/gin-gonic/gin"
    // controller "query/article/controller"
    routes "query/article/routes"
    // "os/exec"

    "testing"
    "net/http"
    "net/http/httptest"
    "io/ioutil"
	  // "fmt"
    // "context"
    // message "query/article/message_broker"
    // "github.com/jasonlvhit/gocron"
  )

func main() {
    routes.Route()
    http.ListenAndServe("localhost:8080/query/listen/message", nil)

}
  
  func article(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("health"))
  }
  
  func Test_article(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/query/article", nil)
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