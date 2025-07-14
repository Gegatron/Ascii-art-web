package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	asciiartweb "asciiartweb/asciiget"
)

type PageData struct {
	Output string
	Name   string
	Banner string
}

func main() {
	http.HandleFunc("/", handleForm)
	http.HandleFunc("/result", handleResult)
	http.ListenAndServe(":8080", nil)
}

// Handles GET request to display the form
func handleForm(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl2, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl2.Execute(w, data)
}

// Handles POST request and displays ASCII result
func handleResult(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		data.Name = r.FormValue("input")
		data.Banner = r.FormValue("select")

		bannerContent, err := os.ReadFile(data.Banner + ".txt")
		if err != nil {
			http.Error(w, "Banner file not found", http.StatusNotFound)
			return
		}
		fmt.Println(data.Name)
		lines := strings.Split(string(bannerContent), "\n")
		textLines := strings.Split(data.Name, "\r\n") // Important: user should use \n in textarea input
		data.Output = asciiartweb.AsciiPrint(lines, textLines)
			fmt.Println(data.Output)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

