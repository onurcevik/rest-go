package main_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/onurcevik/restful/internal/db"
)

func initialize(t *testing.T) {
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable ", os.Getenv("USER"), password, dbname)
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

}

func testmain