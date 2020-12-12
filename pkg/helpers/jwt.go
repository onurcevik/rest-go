package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

//TODO getenv
var jwtsecret = os.Getenv("jwtsecret")

func checkAuthorization() {

}

func GetJWTClaims(r *http.Request, jwtToken string) (map[string]interface{}, error) {

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(jwtsecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

//GenerateJWTTokenWithClaims Generates JWT access and refresh tokens to be sent to client side
func GenerateJWTTokenWithClaims(username string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtsecret))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return t, err
}
