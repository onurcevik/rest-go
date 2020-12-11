package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/onurcevik/restful/internal/model"

	"github.com/onurcevik/restful/internal/db"
	"github.com/onurcevik/restful/pkg/helpers"

	"golang.org/x/crypto/bcrypt"
)

//IndexHandler handles requests to / path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	claims, _ := r.Context().Value("claims").(jwt.MapClaims)
	username := claims["username"].(string)
	if len(username) == 0 {
		json.NewEncoder(w).Encode("Please Login or Register")
	}
	json.NewEncoder(w).Encode("Welcome " + username)
}

//RegisterHandler handles requests to /request path and allows users to request a new user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var user model.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusInternalServerError)
	}
	name, pwd := user.Username, user.Password
	if len(name) > 0 {

		var usernameExists bool
		sqlstmnt := `SELECT EXISTS(SELECT * FROM users WHERE username=$1);`
		_ = db.Conn.QueryRow(sqlstmnt, name).Scan(&usernameExists)
		if usernameExists {
			http.Error(w, "Username exists", http.StatusInternalServerError)
			return
		}
		insertQuery := `INSERT INTO users (username, password)
		VALUES ($1,$2 ) RETURNING ID;
		`

		var id int
		registerpasswd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
		_ = db.Conn.QueryRow(insertQuery, name, string(registerpasswd)).Scan(&id)

		if err != nil {
			//TODO log
			w.WriteHeader(http.StatusInternalServerError)
		}

		jwtToken, err := helpers.GenerateJWTTokenWithClaims(name, jwt.MapClaims{
			"id":       id,
			"username": name,
			"exp":      time.Now().Add(time.Minute * 5).Unix(),
		})

		if err != nil {
			//TODO log
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(jwtToken)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Name cant be empty string.")
	return
}

//LoginHandler handles requests to /login path and allows users to login with a exusting acount
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var user model.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusInternalServerError)
	}
	name, pwd := user.Username, user.Password
	selectQuery := `SELECT id,password FROM users WHERE username=$1;`
	row := db.Conn.QueryRow(selectQuery, name)

	var id int
	var hash string

	//TODO println sil
	switch err := row.Scan(&id, &hash); err {
	case sql.ErrNoRows:
		fmt.Println("User doesnt exist in database")
	case nil:
		//TODO sil
		fmt.Println(hash)
	default:
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	jwtToken, err := helpers.GenerateJWTTokenWithClaims(name, jwt.MapClaims{
		"id":       id,
		"username": name,
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
	})
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(jwtToken)

	return
}

//LogoutHandler handles requests to /logout path and logsout the current logged-in user and deletes the session from database

//ListNotesHandler lists all notes belong to the logged in user
func ListNotesHandler(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value("claims").(jwt.MapClaims)
	ownerid := claims["id"].(int)
	notes, err := db.GetUserNotes(ownerid)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
	return

}

//NewNoteHandler creates a new note
func NewNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var note model.Note
	err := json.Unmarshal(body, &note)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusInternalServerError)
	}
	if len(note.Content) > 0 {
		claims := r.Context().Value("claims").(map[string]interface{})
		ownerid := claims["id"]
		if err != nil {
			//TODO log
			w.WriteHeader(http.StatusInternalServerError)
		}
		insertQuery := `INSERT INTO notes (ownerid, note)
				VALUES ($1,$2 );`

		_, err = db.Conn.Exec(insertQuery, ownerid, note.Content)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(note)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Note cant be empty string.")
	return
}

//GetNoteHandler gets a note with a given ID
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
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
			//TODO log
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(model.Note{ID: integerID, Content: content})
		return
	default:
		//TODO log
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

//DeleteNoteHandler deletes a note with given ID
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id := vars["id"]

	deleteQuery := `DELETE  FROM notes WHERE id=$1;`

	_, err := db.Conn.Exec(deleteQuery, id)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Note deleted")
	return

}
