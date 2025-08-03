package asciiartweb

import (
	"html/template"
	"net/http"
)

// Handle the incomming errors from the other handlers
func ErrorsHandler(w http.ResponseWriter, ErrorMessage string, Status int) {
	temp, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		ErrorsHandler(w, "file parsing error", http.StatusBadRequest)
	}
	w.WriteHeader(Status)
	temp.Execute(w, map[string]any{
		"ErrorMessage": ErrorMessage,
		"Status":       Status,
	})
}
