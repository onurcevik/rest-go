package main

import (
	"net/http"
	"os"

	"github.com/onurcevik/restful/internal/db"

	"github.com/onurcevik/restful/internal/api"

	_ "github.com/lib/pq"
)

func main() {
	// p, err := strconv.Atoi(os.Getenv("pgport"))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	config := &db.PostgresConfig{
		Host:     os.Getenv("pghost"),
		Port:     5432,
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
