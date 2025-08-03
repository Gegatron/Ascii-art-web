package asciiartweb

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	asciiartweb "asciiartweb/functions"
)

// ResultPrint generates the ASCII art output to be displayed on the result page.
func ResultPrint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorsHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		ErrorsHandler(w, "Form parsing error", http.StatusBadRequest)
		return
	}

	Banner := r.FormValue("Select")
	DataBanner, err := os.ReadFile("banners/" + Banner + ".txt")
	if err != nil {
		ErrorsHandler(w, "Banner file not found", http.StatusNotFound)
		return
	}

	Input := r.FormValue("Input")
	if len(Input) > 1000 {
		ErrorsHandler(w, "Input to long", http.StatusBadRequest)
		return
	}

	SliceBanner := strings.Split(string(DataBanner), "\n")
	Text := strings.Split(Input, "\r\n")
	Output := asciiartweb.AsciiPrint(SliceBanner, Text)
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		ErrorsHandler(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, map[string]string{"Output": Output})
}
