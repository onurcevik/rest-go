package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/onurcevik/restful/internal/db"
	"github.com/onurcevik/restful/internal/router"
)

func main() {

	psqlInfo := db.GetPostgresDataSource()
	var err error
	db.Conn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close()
	err = db.Conn.Ping()
	if err != nil {
		panic(err)
	}

	r := router.GetRouter()
	//db.StartDatabase()
	http.ListenAndServe(":8080", r)
}
