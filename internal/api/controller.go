package api

import "net/http"

// Controller rest api interface
type Controller interface {
	Index(w http.ResponseWriter, r *http.Request)  //list all
	Show(w http.ResponseWriter, r *http.Request)   // show one
	Create(w http.ResponseWriter, r *http.Request) // create one
	Update(w http.ResponseWriter, r *http.Request) // update one
	Delete(w http.ResponseWriter, r *http.Request) // delete one
}
