package router

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, name string, data any) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/" + name,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout.html", data)

}
