package main

import (
	"log"
	"net/http"

	"asciiartweb/handlers"
)

func main() {
	http.HandleFunc("/css/", asciiartweb.HandlerCss)
	http.HandleFunc("/", asciiartweb.Home)
	http.HandleFunc("/ascii-art", asciiartweb.ResultPrint)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
