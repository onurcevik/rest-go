package api

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

//GetRouter returns pointer to mux.Router after adding middleware wrapped handler functions
func GetRouter(api *API) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", LoginController{API: api}.Create).Methods("POST")
	router.HandleFunc("/register", RegisterController{API: api}.Create).Methods("POST")
	router.Handle("/", MiddlewareController{API: api}.JWTmiddleware(http.HandlerFunc(IndexController{}.Index))).Methods("GET")
	router.Handle("/notes", MiddlewareController{API: api}.JWTmiddleware(http.HandlerFunc(NoteController{API: api}.Index))).Methods("GET")
	router.Handle("/note", MiddlewareController{API: api}.JWTmiddleware(http.HandlerFunc(NoteController{API: api}.Create))).Methods("POST")
	router.Handle("/note/{id}", MiddlewareController{API: api}.JWTmiddleware(http.HandlerFunc(NoteController{API: api}.Show))).Methods("GET")
	router.Handle("/note/{id}", MiddlewareController{API: api}.JWTmiddleware(http.HandlerFunc(NoteController{API: api}.Update))).Methods("POST")
	router.Handle("/note/{id}", MiddlewareController{API: api}.JWTmiddleware(http.HandlerFunc(NoteController{API: api}.Delete))).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return router
}
