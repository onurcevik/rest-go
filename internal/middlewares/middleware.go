package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			claims, err := helpers.GetJWTClaims(r, jwtToken)
			if err != nil {

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("Unauthorized")
			}
			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
