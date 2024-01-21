package model

import (
	"bitbucket.org/mikehouston/asana-go"
	"time"
)

type Task struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	DueAt       *time.Time `json:"due_at"`
	Done        bool       `json:"done"`
}

func (t *Task) ToPrintable() []string {
	var dueAt = ""
	if t.DueAt != nil {
		dueAt = t.DueAt.Format("2006-01-02 15:04:05")
	}

	return []string{
		t.ID,
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
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Notes,
		CreatedAt:   task.CreatedAt,
		Done:        completed,
		DueAt:       task.DueAt,
	}
}
