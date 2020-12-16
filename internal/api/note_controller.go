package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/onurcevik/restful/internal/model"
)

type NoteController struct {
	Controller
	*API
}

//Index lists all notes belong to the logged in user
func (nc NoteController) Index(w http.ResponseWriter, r *http.Request) {
	db := nc.API.GetDB()
	claims := r.Context().Value("claims").(map[string]interface{})
	ownerid := int(claims["id"].(float64))
	notes, err := db.GetUserNotes(ownerid)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
	return

}

//Create creates a new note
func (nc NoteController) Create(w http.ResponseWriter, r *http.Request) {
	db := nc.API.GetDB()
	var note model.Note
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(note.Content) > 0 {
		claims := r.Context().Value("claims").(map[string]interface{})
		ownerid := claims["id"]
		if err != nil {
			//TODO log
			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = db.Insert(&note, ownerid, note.Content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(note)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Note cant be empty string.")
	return
}

//Show gets a note with a given ID
func (nc NoteController) Show(w http.ResponseWriter, r *http.Request) {
	db := nc.API.GetDB()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	i := vars["id"]

	selectQuery := `SELECT id,note FROM notes WHERE id=$1;`
	row := db.Conn.QueryRow(selectQuery, i)

	var id, content string

	switch err := row.Scan(&id, &content); err {
	case sql.ErrNoRows:
		fmt.Println("Note doesnt exist in database")
	case nil:
		w.WriteHeader(http.StatusOK)
		integerID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		var n model.Note
		n.ID = integerID
		n.Content = content
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(n)
		return
	default:
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}

//Update
func (nc NoteController) Update(w http.ResponseWriter, r *http.Request) {
	db := nc.API.GetDB()
	var note model.Note
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var uID int
	if len(note.Content) > 0 {
		vars := mux.Vars(r)
		id := vars["id"]
		uID, err = db.Update(&note, id, note.Content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Note " + string(uID) + " Updated")
	return

}

//Delete deletes a note with given ID
func (nc NoteController) Delete(w http.ResponseWriter, r *http.Request) {
	db := nc.API.GetDB()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	dID, err := db.Delete(&model.Note{}, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Note" + string(dID) + "deleted")
	return

}
