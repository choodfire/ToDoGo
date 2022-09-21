package main

import (
	"ToDoGo/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var t *template.Template
var taskList data.Tasks

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(taskList.Tasks)
	t.ExecuteTemplate(w, "index.html", taskList.Tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	taskList.AddTask(r.FormValue("task"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	taskList.AddTask("hee hee")
	t, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/addTask", addTask)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
