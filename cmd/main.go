package main

import (
	"net/http"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/onurcevik/rest-go/docs"
	"github.com/onurcevik/rest-go/internal/api"
	"github.com/onurcevik/rest-go/internal/db"
)

// @title Restful API with Go
// @version 1.0
// @description This is an example restful api writtin with go
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
