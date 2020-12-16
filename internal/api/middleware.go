package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/onurcevik/restful/pkg/helpers"
)

type MiddlewareController struct {
	*API
}

func (md MiddlewareController) JWTmiddleware(next http.Handler) http.Handler {
	db := md.API.GetDB()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			claims, err := helpers.GetJWTClaims(jwtToken)
			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			vars := mux.Vars(r)
			resourceid, _ := strconv.Atoi(vars["id"])
			ownerid := int(claims["id"].(float64))
			if r.URL.Path != "/" && r.URL.Path != "/note" && r.URL.Path != "/notes" {

				if !db.IsResourceOwner(resourceid, ownerid) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("You dont have any notes with given ID"))
					return
				}
			}
			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
