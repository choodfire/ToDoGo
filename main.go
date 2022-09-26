package main

import (
	"ToDoGo/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var t *template.Template
var taskList data.Tasks

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(taskList.Tasks)
	var Lists = struct {
		Compl   []data.Task
		UnCompl []data.Task
	}{
		Compl:   taskList.GetCompletedList(),
		UnCompl: taskList.GetUncompletedList(),
	}
	t.ExecuteTemplate(w, "index.html", Lists)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	taskList.AddTask(r.FormValue("task"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	index, err := strconv.Atoi(r.FormValue("token"))
	if err != nil {
		log.Fatal(err) // todo handle somehow else
	}
	taskList.DeleteTask(index)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func markCompleted(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index, err := strconv.Atoi(r.FormValue("token"))
	if err != nil {
		log.Fatal(err) // todo handle somehow else
	}
	taskList.MarkCompleness(index)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	taskList.AddTask("hee hee")
	taskList.AddTask("hee hee2")
	taskList.MarkCompleness(1)
	t, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/addTask", addTask)
	http.HandleFunc("/deleteTask", deleteTask)
	http.HandleFunc("/markCompleted", markCompleted)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
