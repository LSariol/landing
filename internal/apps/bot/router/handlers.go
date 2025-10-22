package router

import (
	"net/http"

	"github.com/mobasity-web-landing/internal/router"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	router.RenderTemplate(w, "bot.html", nil)
}
