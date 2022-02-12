package main
import (
	"services/article/routes"
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"

)

func main() {
	routes.Route()
	
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