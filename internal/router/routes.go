package router

import (
	"net/http"
)

func DefineRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Un
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFoundHandler(w, r)
			return
		}
		homeHandler(w, r)
	})

	mux.HandleFunc("GET /projects", projectsHandler)

	return mux
}
