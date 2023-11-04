package middleware

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
)

func WithJwtAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling middleware")
		handlerFunc(w, r)
	}
}

func validateJwtToken(tokenString string) (*jwt.Token, error) {
	
}
