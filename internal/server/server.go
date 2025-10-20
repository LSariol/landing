package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mobasity-web-landing/internal/router"
)

func Run() {
	srv := NewServer()

	log.Println("Server running on", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}

func Stop() {

}

func NewServer() *http.Server {

	mux := router.DefineRoutes()
	addr := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return srv
}
