package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/onurcevik/restful/internal/db"
	"github.com/onurcevik/restful/internal/handlers"
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

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/notes", handlers.ListNotesHandler).Methods("GET")
	router.HandleFunc("/note", handlers.NewNoteHandler).Methods("POST")
	router.HandleFunc("/note/{id}", handlers.GetNoteHandler).Methods("GET")
	router.HandleFunc("/note/{id}", handlers.DeleteNoteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
