package main

import (
	"html/template"
	"log"
	"net/http"
	"todoGo/data"
)

var t *template.Template
var taskList data.Tasks

func mainPage(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "main.html", taskList.Tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	taskList.AddTask(r.FormValue("task"))
	t.ExecuteTemplate(w, "main.html", taskList.Tasks)
}

func main() {
	taskList.AddTask("hee hee")
	t, _ = template.ParseGlob("*.html")

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/addTask", addTask)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
