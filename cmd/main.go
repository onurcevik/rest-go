package main

import (
	"net/http"
	"os"

	"github.com/onurcevik/rest-go/internal/db"

	"github.com/onurcevik/rest-go/internal/api"

	_ "github.com/lib/pq"
)

func main() {

	config := &db.PostgresConfig{
		Host:     os.Getenv("pghost"),
		Port:     os.Getenv("pgport"),
		User:     os.Getenv("pguser"),
		Password: os.Getenv("pgpass"),
		DBname:   os.Getenv("pgdbname"),
	}

	var a api.API
	d := db.NewDB(*config)
	a.Database = &d
	r := api.GetRouter(&a)

	http.ListenAndServe(":8080", r)
}
