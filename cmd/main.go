package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/onurcevik/restful/internal/db"
)

var (
	host     = os.Getenv("HOST")
	port, _  = strconv.Atoi(os.Getenv("POSTGRESPORT"))
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	dbname   = os.Getenv("DBNAME")
)

func main() {

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable ", user, password, dbname)
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
