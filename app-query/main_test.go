// main_test.go
package main_test

import (
    "testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"github.com/kamva/mgm/v3"
	"fmt"
	models "query/article/models"

)

func article(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("health"))
}

func Test_article(t *testing.T) {

	a := models.Article{}
	coll := mgm.Coll(&a)
	fmt.Println(coll)

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