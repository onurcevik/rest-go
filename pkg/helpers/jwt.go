package helpers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

//TODO getenv
var jwtsecret = os.Getenv("jwtsecret")

func GetJWTClaims(r *http.Request, jwtToken string) (map[string]interface{}, error) {

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return []byte(jwtsecret), nil

	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//username := claims["username"].(string)
		return claims, nil
	}
	return nil, err
}

//GenerateJWTTokenWithClaims Generates JWT access and refresh tokens to be sent to client side
func GenerateJWTTokenWithClaims(username string, claims jwt.MapClaims) (string, error) {
	// Create token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"username": username,
	// 	"exp":      time.Now().Add(time.Minute * 5).Unix(),
	// })
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtsecret))
	if err != nil {
		return "", err
	}
	return t, err
}
