package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/pooulad/bankApi/database"
	"github.com/pooulad/bankApi/model"
	"github.com/pooulad/bankApi/util"
)

func withJwtAuth(handlerFunc http.HandlerFunc, s database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-jwt-token")

		token, err := validateJwtToken(tokenString)
		if err != nil {
			perimssionDeniedWriter(w)
			return
		}
		if !token.Valid {
			perimssionDeniedWriter(w)
			return
		}

		userId, err := util.GetAccountParameterId(r)
		if err != nil {
			perimssionDeniedWriter(w)
			return
		}

		account, err := s.GetAccountById(userId)
		if err != nil {
			perimssionDeniedWriter(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if account.Number != int64(claims["AccountNumber"].(float64)) {
			perimssionDeniedWriter(w)
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
	return token.SignedString([]byte(jwt_secret_token))
}
