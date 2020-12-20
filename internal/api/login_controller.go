package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/onurcevik/rest-go/internal/model"
	"github.com/onurcevik/rest-go/pkg/helpers"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var id int
	var hash string

	name, pwd := user.Username, user.Password
	selectQuery := `SELECT id,password FROM users WHERE username=$1;` // Simdilik burada kalsin ileride db paketinin icine tasi
	if err := db.Conn.QueryRow(selectQuery, name).Scan(&id, &hash); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Could not find user in database ")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(jwtToken)
	return
}
