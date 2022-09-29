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

// var taskList data.Tasks
var users data.Users
var cookieCounterTest int = 0

func getUser(receivedCookies string) *data.User {
	for _, user := range users.Users {
		if user.Cookies == receivedCookies {
			return &user
		}
	}

	return nil
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("main-cookie")
	if err != nil {
		// cookie not found
		cookie := &http.Cookie{
			Name:     "main-cookie",
			Value:    fmt.Sprintf("Cookie%d", cookieCounterTest),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	currentUser := getUser(cookie.Name)

	var Lists = struct {
		Compl   []data.Task
		UnCompl []data.Task
	}{
		Compl:   currentUser.Tasks.GetCompletedList(),
		UnCompl: currentUser.Tasks.GetUncompletedList(),
	}
	t.ExecuteTemplate(w, "index.html", Lists)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//taskList.AddTask(r.FormValue("task"))
	user.Tasks.AddTask(r.FormValue("task"))

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index, _ := strconv.Atoi(r.FormValue("token"))
	//taskList.DeleteTask(index)
	user.Tasks.DeleteTask(index)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func markCompleted(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index, _ := strconv.Atoi(r.FormValue("token"))
	//taskList.ChangeCompleness(index)
	user.Tasks.ChangeCompleness(index)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	user.Tasks.AddTask("hee hee")
	user.Tasks.AddTask("hee hee2")
	user.Tasks.ChangeCompleness(1)
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
