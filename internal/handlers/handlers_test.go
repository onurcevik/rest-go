package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/onurcevik/restful/internal/handlers"
	"github.com/onurcevik/restful/pkg/helpers"

	"github.com/onurcevik/restful/internal/model"
)

//testdata
type td struct {
	JWT    string //no expire date
	Userid int
	notes  []model.Note
}

var testdatas []td

func init() {

	testdatas = []td{
		td{
			JWT:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzEsInVzZXJuYW1lIjoidGVzdHVzZXIxIn0.7AmboDarHgOiIJ9c7jkzIgAs3d2p6S6ZR7I1l4jOWsI",
			Userid: 31,
			notes: []model.Note{
				model.Note{
					ID:      18,
					Content: "testuser1 first note",
				},
				model.Note{
					ID:      19,
					Content: "testuser1 second note",
				}, model.Note{
					ID:      20,
					Content: "testuser1 third note",
				},
			},
		},

		td{
			JWT:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzIsInVzZXJuYW1lIjoidGVzdHVzZXIyIn0.fVZiZf2me65hvJKd_p24ufMhKQ0RXsOblPd4oAqGaQk",
			Userid: 32,
			notes: []model.Note{
				model.Note{
					ID:      21,
					Content: "testuser2 one and only note",
				},
			},
		},
	}

}

//IndexHandler handles requests to / path
func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.IndexHandler)

	for k, _ := range testdatas {
		claims, err := helpers.GetJWTClaims(testdatas[k].JWT)

		if err != nil {
			t.Error("Error getting JWTClaims")
		}
		ctx := context.WithValue(req.Context(), "claims", claims)
		handler.ServeHTTP(rr, req.WithContext(ctx))
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}
