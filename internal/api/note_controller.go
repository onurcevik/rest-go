package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/onurcevik/rest-go/internal/cache"

	"github.com/gorilla/mux"
	"github.com/onurcevik/rest-go/internal/model"
)

type NoteController struct {
	Controller
	*API
}

var (
	//NoteControllerin icine tasi
	notescache cache.NotesCache = cache.NewRedisCache("localhost:6379", 0, 20)
)

//Index lists all notes belong to the logged in user
// @Summary Returns all notes for user
// @Description NoteController Index handler
// @ID note-index
// @Produce json
// @Success 200 {array} model.Note "notes"
// @Failure 500 {string} string "note index error"
// @Router /notes [get]
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
// @Summary Creates a new note for user
// @Description NoteController Create handler
// @ID note-create
// @Accept  json
// @Produce json
// @Param registerData body string true "note content"
// @Success 200 {object} model.Note "note"
// @Failure 400,500 {string} string "note create error"
// @Router /note [post]
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
		insertID, err := db.Insert(&note, ownerid, note.Content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		note.ID = insertID
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(note)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Note cant be empty string.")
	return
}

//Show get a note of a user with given ID
// @Summary shows a note for a given ID
// @Description NoteController Show handler
// @ID note-show
// @Accept  json
// @Produce json
// @Param noteID path string true "note id"
// @Success 200 {object} model.Note "note"
// @Failure 500 {string} string "note show error"
// @Router /note/{id} [get]
func (nc NoteController) Show(w http.ResponseWriter, r *http.Request) {
	db := nc.API.GetDB()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	i := vars["id"]
	var n model.Note   //
	var cn *model.Note //cache note

	cn, _ = notescache.Get(i)

	if cn == nil {
		selectQuery := `SELECT id,content FROM notes WHERE id=$1;`
		row := db.Conn.QueryRow(selectQuery, i)

		var id, content string

		switch err := row.Scan(&id, &content); err {
		case sql.ErrNoRows:
			fmt.Println("Note doesnt exist in database")
		case nil:
			w.WriteHeader(http.StatusOK)
			integerID, _ := strconv.Atoi(id)

			n.ID = integerID
			n.Content = content

		default:
			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err := notescache.Set(i, n)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(n)
		return

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cn)
	return
}

//Update content of a note of a user with given ID
// @Summary Update content of a note with a given ID
// @Description NoteController Update handler
// @ID note-update
// @Accept  json
// @Produce  json
// @Param noteID path string true "ID of the note to be updated"
// @Success 200 {object} model.Note "note"
// @Failure 500 {string} string "note update error"
// @Router /note/{id} [post]
func (nc NoteController) Update(w http.ResponseWriter, r *http.Request) {
	db := nc.API.GetDB()
	var note *model.Note
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
		uID, err = db.Update(note, id, note.Content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		note.ID = uID
		err := notescache.Set(id, *note)
		if err != nil {
			log.Fatalln(err)
		}

	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("Note " + fmt.Sprintf("Note %d deleted", uID) + " Updated")
	return

}

//Delete note of a user with given ID
// @Summary Delete the note with given ID
// @Description NoteController Delete handler
// @ID note-delete
// @Accept  json
// @Produce  json
// @Param noteID path string true "ID of the note to be delete"
// @Success 200 {object} model.Note "note"
// @Failure 500 {string} string "note delete error"
// @Router /note/{id} [delete]
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

	json.NewEncoder(w).Encode(fmt.Sprintf("Note %d deleted", dID))
	return

}
