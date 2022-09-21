package data

import (
	"errors"
	"time"
)

var defaultCompletedTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

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

func (t *Tasks) AddTask(taskName string) {
	t.Tasks = append(t.Tasks, Task{taskName, len(t.Tasks) + 1, false, time.Now(), defaultCompletedTime})
}

func (t *Tasks) DeleteTask(index int) error {
	index = index - 1
	if index > len(t.Tasks) {
		return errors.New("Wrong index") // todo технически такой ошибки быть не должно, потестить
	}
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)

	return nil
}

func (t *Tasks) MarkCompleted(id int) {
	t.Tasks[id-1].isDone = true
}
