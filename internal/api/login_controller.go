package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/onurcevik/restful/internal/model"
	"github.com/onurcevik/restful/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	Controller
	*API
}

func (lc LoginController) Create(w http.ResponseWriter, r *http.Request) {
	db := lc.API.GetDB()
	var user model.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	name, pwd := user.Username, user.Password
	selectQuery := `SELECT id,password FROM users WHERE username=$1;` // Simdilik burada kalsin ileride db paketinin icine tasi
	row := db.Conn.QueryRow(selectQuery, name)

	var id int
	var hash string

	//TODO println sil
	switch err := row.Scan(&id, &hash); err {
	case sql.ErrNoRows:
		log.Fatalln(err)
		json.NewEncoder(w).Encode("User with that username doesnt exist")
		return
	case nil:
		fmt.Println("LOGIN ID", id)
	default:
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Wrong username and/or password.")
		return
	}

	jwtToken, err := helpers.GenerateJWTTokenWithClaims(name, jwt.MapClaims{
		"id":       id,
		"username": name,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	})
	if err != nil {
		//TODO log
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jwtToken)
	return
}
