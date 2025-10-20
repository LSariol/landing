package middleware

import (
	"fmt"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Before calling work
		fmt.Println("Auth Prework")
		// Call
		next.ServeHTTP(w, r)

		// Aftercall work
		fmt.Println("Auth Postwork")
	})
}
