package helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//TODO getenv
var jwtsecret = os.Getenv("jwtsecret")

//AlreadyLoggedIn checks if user is logged in or not
func GetJWTUser(r *http.Request) (string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return "", err
	}
	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return []byte(jwtsecret), nil

	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}
	return "", err
}

//GenerateJWTToken Generates JWT access and refresh tokens to be sent to client side
func GenerateJWTToken(username string) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	})

	t, err := token.SignedString([]byte(jwtsecret))
	if err != nil {
		return "", err
	}
	return t, err
}
