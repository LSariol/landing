package middleware

import "net/http"

func FullWrap(path http.Handler) http.Handler {
	return Logger(Authenticate(path))
}

func WrapAuth(path http.Handler) http.Handler {
	return Authenticate(path)
}

func WrapLogger(path http.Handler) http.Handler {
	return Logger(path)
}
