package router

import (
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "notfound.html", nil)
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "bot.html", nil)
}
