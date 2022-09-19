package data

import "time"

var defaultCompletedTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

type Task struct {
	Title         string
	id            int
	isDone        bool
	timeCreated   time.Time
	timeCompleted time.Time
}

type Tasks struct {
	Tasks []Task
}

func (t *Tasks) AddTask(taskName string) {
	t.Tasks = append(t.Tasks, Task{taskName, len(t.Tasks) + 1, false, time.Now(), defaultCompletedTime})
}

func (t *Tasks) MarkCompleted(id int) {
	t.Tasks[id-1].isDone = true
}
