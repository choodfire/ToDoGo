package data

import (
	"net/http"
	"time"
)

var defaultCompletedTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
var taskID = 0

type Task struct {
	Title         string
	Id            int
	isDone        bool
	timeCreated   time.Time
	timeCompleted time.Time
}

type Tasks struct {
	Tasks []Task
}

type User struct {
	Cookies *http.Cookie
	//UserId  int
	Tasks
}

type Users struct {
	Users []User
}

func (u *Users) AddNewUser(cookie *http.Cookie) {
	u.Users = append(u.Users, User{Cookies: cookie, Tasks: Tasks{}})
}

func (t *Tasks) AddTask(taskName string) {
	t.Tasks = append(t.Tasks, Task{taskName, taskID, true, time.Now(), defaultCompletedTime})
	taskID += 1
}

func (t *Tasks) DeleteTask(index int) {
	t.Tasks[index] = Task{
		Title:         "nothing",
		Id:            -1, // -1 means deleted
		isDone:        false,
		timeCreated:   defaultCompletedTime,
		timeCompleted: defaultCompletedTime,
	}
}

func (t *Tasks) ChangeCompleness(index int) {
	t.Tasks[index].isDone = !t.Tasks[index].isDone
}

func (t Tasks) GetCompletedList() []Task {
	res := []Task{}
	for _, task := range t.Tasks {
		if task.isDone == true && task.Id != -1 {
			res = append(res, task)
		}
	}

	return res
}

func (t Tasks) GetUncompletedList() []Task {
	res := []Task{}
	for _, task := range t.Tasks {
		if task.isDone == false && task.Id != -1 {
			res = append(res, task)
		}
	}

	return res
}
