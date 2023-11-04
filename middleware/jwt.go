package middleware

import (
	"fmt"
	"net/http"
)

func WithJwtAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling middleware")
		handlerFunc(w, r)
	}
}
