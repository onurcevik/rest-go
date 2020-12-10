package main

import (
	"net/http"

	"github.com/onurcevik/restful/internal/db"
	"github.com/onurcevik/restful/internal/router"

	_ "github.com/lib/pq"
)

func main() {
	db.StartDatabase()
	r := router.GetRouter()
	http.ListenAndServe(":8080", r)
}
