package router

import (
	"net/http"

	"github.com/mobasity-web-landing/internal/router"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	router.RenderTemplate(w, "home.html", nil)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	router.RenderTemplate(w, "projects.html", nil)
}

func roadmapHandler(w http.ResponseWriter, r *http.Request) {
	router.RenderTemplate(w, "roadmap.html", nil)
}
