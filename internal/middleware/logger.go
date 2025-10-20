package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Before calling work
		fmt.Println("Logging Prework")
		// Call
		next.ServeHTTP(w, r)

		// Aftercall work
		fmt.Println("Logging postwork")
	})
}
