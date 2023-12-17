package model

import (
	"bitbucket.org/mikehouston/asana-go"
	"time"
)

type Task struct {
	Name      string
	CreatedAt *time.Time
	DueAt     *time.Time
	Done      bool
}

func (t *Task) ToPrintable() []string {
	var dueAt = ""
	if t.DueAt != nil {
		dueAt = t.DueAt.Format("2006-01-02 15:04:05")
	}

	return []string{
		t.Name,
		t.CreatedAt.Format("2006-01-02 15:04:05"),
		dueAt,
	}
}

func FromAsanaTask(task *asana.Task) *Task {
	var completed = false
	if task.Completed != nil {
		completed = *task.Completed
	}

	return &Task{
		Name:      task.Name,
		CreatedAt: task.CreatedAt,
		Done:      completed,
		DueAt:     task.DueAt,
	}
}
