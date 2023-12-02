package model

import "time"

type Task struct {
	Name      string
	CreatedAt time.Time
	Done      bool
}

func (t *Task) ToPrintable() []string {
	return []string{
		t.Name,
		t.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
