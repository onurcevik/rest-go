package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

var Conn *sql.DB

func GetPostgresDataSource() string {

	host := os.Getenv("pghost")
	port, err := strconv.Atoi(os.Getenv("pgport"))
	if err != nil {
		log.Fatalln()
	}
	user := os.Getenv("pguser")
	password := os.Getenv("pgpass")
	dbname := os.Getenv("pgdbname")

	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

// //
// func StartDatabase() {
// 	psqlInfo := getPostgresDataSource()

// 	Conn, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = Conn.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// }
