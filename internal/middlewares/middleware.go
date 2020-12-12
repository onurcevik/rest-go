package middlewares

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/onurcevik/restful/internal/db"

	"github.com/gorilla/mux"
	"github.com/onurcevik/restful/pkg/helpers"
)

func JWTmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			claims, err := helpers.GetJWTClaims(r, jwtToken)
			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			vars := mux.Vars(r)
			resourceid, _ := strconv.Atoi(vars["id"])
			ownerid := int(claims["id"].(float64))
			if r.URL.Path != "/" && r.URL.Path != "/note" {
				if !db.IsResourceOwner(resourceid, ownerid) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("You are not the owner of this resource"))
					return
				}
			}
			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
