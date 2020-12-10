package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onurcevik/restful/internal/handlers"
	"github.com/onurcevik/restful/internal/middlewares"
)

//GetRouter returns pointer to mux.Router after adding middleware wrapped handler functions
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", middlewares.JWTmiddleware(http.HandlerFunc(handlers.IndexHandler))).Methods("GET")

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	router.Handle("/notes", middlewares.JWTmiddleware(http.HandlerFunc(handlers.ListNotesHandler))).Methods("GET")

	router.Handle("/notes", middlewares.JWTmiddleware(http.HandlerFunc(handlers.NewNoteHandler))).Methods("POST")
	router.Handle("/note/{id}", middlewares.JWTmiddleware(http.HandlerFunc(handlers.GetNoteHandler))).Methods("GET")
	router.Handle("/note/{id}", middlewares.JWTmiddleware(http.HandlerFunc(handlers.DeleteNoteHandler))).Methods("DELETE")

	return router
}
