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

var users data.Users
var cookieCounterTest int = 0

func getUser(receivedCookies string) *data.User {
	// for _, user := range users.Users { - iterates over a copy of array

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Cookies.Value == receivedCookies {
			return &users.Users[i]
		}
	}

	return nil
}

func getCookie(r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("main-cookie")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "main-cookie",
			Value:    fmt.Sprintf("Cookie%d", cookieCounterTest),
			HttpOnly: true,
		}
		cookieCounterTest += 1
		users.AddNewUser(cookie)
	}

	return cookie
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(r)
	http.SetCookie(w, cookie)
	currentUser := getUser(cookie.Value)

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
	cookie := getCookie(r)
	currentUser := getUser(cookie.Value)
	currentUser.Tasks.AddTask(r.FormValue("task"))

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cookie := getCookie(r)
	currentUser := getUser(cookie.Value)
	index, _ := strconv.Atoi(r.FormValue("token"))
	currentUser.Tasks.DeleteTask(index)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func markCompleted(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cookie := getCookie(r)
	currentUser := getUser(cookie.Value)
	index, _ := strconv.Atoi(r.FormValue("token"))
	currentUser.Tasks.ChangeCompleness(index)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
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
