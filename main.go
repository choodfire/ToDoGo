package main

import (
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

func mainPage(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "main.html", nil)
}

func main() {
	t, _ = template.ParseGlob("*.html")

	http.HandleFunc("/", mainPage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
