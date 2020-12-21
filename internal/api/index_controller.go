package api

import (
	"encoding/json"
	"net/http"
)

type IndexController struct {
	Controller
}

// Index godoc
// @Summary Index welcomes logged in user
// @Description user index
// @ID user-index
// @Produce  json
// @Success 200 {string} string "welcome"
// @Failure 401 {string} string "Plese login or register"
// @Router / [get]
func (ic IndexController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	claims := r.Context().Value("claims").(map[string]interface{})
	username, ok := claims["username"].(string)
	if len(username) == 0 || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Please Login or Register")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Welcome " + username)
	return
}
