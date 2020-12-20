package api

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("superSecret")

// GenerateJWT generates JWT token
func GenerateJWT()  (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("failed to generate JWT: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
