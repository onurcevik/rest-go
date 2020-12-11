package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/onurcevik/restful/internal/db"

	"github.com/gorilla/mux"
	"github.com/onurcevik/restful/pkg/helpers"
)

func JWTmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO SIL
		fmt.Println("Authoriz:", r.Header.Get("Authorization"))
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		//TODO SIL
		fmt.Println("Header", authHeader)
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			claims, err := helpers.GetJWTClaims(r, jwtToken)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Println(err)
			}
			vars := mux.Vars(r)
			resourceid, _ := strconv.Atoi(vars["id"])

			ownerid := int(claims["id"].(float64))

			if db.IsResourceOwner(resourceid, ownerid) {
				fmt.Println("owner")
			}

			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
