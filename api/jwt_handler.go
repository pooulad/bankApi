package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/pooulad/bankApi/model"
)

func WithJwtAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-jwt-token")

		_, err := validateJwtToken(tokenString)
		if err != nil {
			WriteJson(w, http.StatusForbidden, ApiError{Error: "invalid jwt token"})
			return
		}

		handlerFunc(w, r)
	}
}

func validateJwtToken(tokenString string) (*jwt.Token, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	jwt_secret_token := os.Getenv("JWT_SECREC_TOKEN")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwt_secret_token), nil
	})

}

func createJwt(account *model.Account) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	jwt_secret_token := os.Getenv("JWT_SECREC_TOKEN")

	claims := &jwt.MapClaims{
		"ExpiresAt":     15000,
		"AccountNumber": account.Number,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwt_secret_token)
}
