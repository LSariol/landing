package server

import (
	"log"
	"net/http"
	"time"
)

func Run(srv *http.Server) {

	log.Println("Server running on", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}

func Stop(srv *http.Server) {

	srv.Close()
}

func NewServer(mux *http.ServeMux, address string) *http.Server {

	srv := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return srv
}
