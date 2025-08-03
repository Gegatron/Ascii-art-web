package asciiartweb

import (
	"html/template"
	"net/http"
)

// Home renders the index.html template for the root path.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorsHandler(w, "Page not found", http.StatusNotFound)
		return
	}

	tem, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorsHandler(w, "file parsing error", http.StatusBadRequest)
		return
	}
	tem.Execute(w, nil)
}
