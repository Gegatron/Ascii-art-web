package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	asciiartweb "asciiartweb/asciiget"
)


func main() {
	http.HandleFunc("/", handleForm)
	http.HandleFunc("/result", handleResult)
	http.ListenAndServe(":8080", nil)
}

// Handles GET request to display the form
func handleForm(w http.ResponseWriter, r *http.Request) {

	tmpl2, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl2.Execute(w, nil)
}

// Handles POST request and displays ASCII result
func handleResult(w http.ResponseWriter, r *http.Request) {
	

	
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		Name := r.FormValue("input")
		Banner := r.FormValue("select")

		bannerContent, err := os.ReadFile(Banner + ".txt")
		if err != nil {
			http.Error(w, "Banner file not found", http.StatusNotFound)
			return
		}
		fmt.Println(Name)
		lines := strings.Split(string(bannerContent), "\n")
		textLines := strings.Split(Name, "\r\n") // Important: user should use \n in textarea input
		output := asciiartweb.AsciiPrint(lines, textLines)
	

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]string{"Output":output})
}
