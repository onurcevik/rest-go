package main

import (
	"net/http"

	"github.com/onurcevik/restful/internal/middlewares"

	"github.com/onurcevik/restful/internal/db"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/onurcevik/restful/internal/handlers"
)

func main() {

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// var err error
	// db.Conn, err = sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Conn.Close()
	// err = db.Conn.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	db.StartDatabase()
	router := mux.NewRouter()
	router.Handle("/", middlewares.JWTmiddleware(http.HandlerFunc(handlers.IndexHandler))).Methods("GET")

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.Handle("/notes", middlewares.JWTmiddleware(http.HandlerFunc(handlers.ListNotesHandler))).Methods("GET")
	router.HandleFunc("/note", handlers.NewNoteHandler).Methods("POST")
	router.HandleFunc("/note/{id}", handlers.GetNoteHandler).Methods("GET")
	router.HandleFunc("/note/{id}", handlers.DeleteNoteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
