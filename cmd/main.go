package main

import (
	"net/http"

	"github.com/onurcevik/restful/internal/db"

	"github.com/onurcevik/restful/internal/api"

	_ "github.com/lib/pq"
)

func main() {

	var a api.API
	d := db.NewDB()
	a.Database = &d
	r := api.GetRouter(&a)

	http.ListenAndServe(":8080", r)
}
