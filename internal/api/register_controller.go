package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/onurcevik/restful/internal/model"
	"github.com/onurcevik/restful/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	Controller
	*API
}

func (rg RegisterController) Create(w http.ResponseWriter, r *http.Request) {
	db := rg.API.GetDB()
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	name, pwd := user.Username, user.Password
	if len(name) > 0 {

		if db.DoesUserExist(name) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Username Exists")
			return
		}

		var id int
		registerpasswd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
		id, err := db.Insert(&user, name, string(registerpasswd)) //access fields in function with Reflect package

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
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(jwtToken)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Name cant be empty string.")
	return
}
