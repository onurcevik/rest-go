package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/onurcevik/restful/internal/api"
)

func TestLoginHandler(t *testing.T) {
	s := api.NewSuite()
	testMap := map[string]string{"username": "testuser", "password": "1234"}
	testUser, _ := json.Marshal(testMap)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(testUser))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.LoginController{API: &s.API}.Create)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
