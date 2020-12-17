package db_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/onurcevik/restful/internal/db"
)

const userstablecreationquery = ` CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username VARCHAR ( 500 ) UNIQUE NOT NULL,  
	password VARCHAR ( 500 ) NOT NULL
);
`

const notestablecreationquery = ` CREATE TABLE IF NOT EXISTS notes (
	id SERIAL PRIMARY KEY,
	username VARCHAR ( 500 ) UNIQUE NOT NULL,  
	password VARCHAR ( 500 ) NOT NULL
);
`

func TestEnsureTablesExists(t *testing.T) {

	config := &db.PostgresConfig{
		Host:     os.Getenv("pghost"),
		Port:     5432,
		User:     os.Getenv("pguser"),
		Password: os.Getenv("pgpass"),
		DBname:   os.Getenv("testpgdbname"),
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBname)
	var err error
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		t.Errorf("couldn open connection")
	}
	err = conn.Ping()
	if err != nil {
		t.Errorf("db ping doesnt work")
	}
	if _, err := conn.Exec(userstablecreationquery); err != nil {
		t.Errorf("userstablecreationquery failed")
	}
	if _, err := conn.Exec(notestablecreationquery); err != nil {
		t.Errorf("notestablecreationquery failed")
	}
}
