package router

import (
	"net/http"

	"github.com/mobasity-web-landing/internal/router"
)

func DefineRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Un
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			router.NotFoundHandler(w, r)
			return
		}
		homeHandler(w, r)
	})

	mux.HandleFunc("/projects", projectsHandler)
	mux.HandleFunc("/roadmap", roadmapHandler)

	return mux
}
