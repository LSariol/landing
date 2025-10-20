package router

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html", nil)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "projects.html", nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "notfound.html", nil)
}
