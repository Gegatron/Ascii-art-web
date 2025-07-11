package main

import (
	"net/http"
	"os"
	"strings"
	"text/template"

	asciiartweb "asciiartweb/asciiget"
)

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", aa)
	http.HandleFunc("/submit", Asciirprint)
	http.ListenAndServe(":8080", nil)
}

func aa(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return
	}
	//	new := r.URL.Query().Get("input")

	tem.Execute(w, nil)
}

func Asciirprint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form parsing error", http.StatusBadRequest)
		return
	}

	banner := r.FormValue("select")
	bannertxt, err := os.ReadFile(banner + ".txt")
	if err != nil {
		http.Error(w, "Banner file not found", http.StatusNotFound)
		return
	}

	name := r.FormValue("input")
	n1 := strings.Split(string(bannertxt), "\n")
	n2 := strings.Split(name, "\\n")
	outpout := asciiartweb.AsciiPrint(n1, n2)

	tmpl, err := template.ParseFiles("templates/new.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, map[string]string{"Name": outpout})
}
