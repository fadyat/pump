package driver

import (
	"github.com/fadyat/pump/internal/api"
	"github.com/fadyat/pump/internal/model"
)

type Asana struct {
	c *api.AsanaClient
}

func (a *Asana) Get(filters ...func(task *model.Task) bool) ([]*model.Task, error) {
	tasksAsana, err := a.c.GetTasks()
	if err != nil {
		return nil, err
	}

	var tasks = make([]*model.Task, 0, len(tasksAsana))
	for _, taskAsana := range tasksAsana {
		tasks = append(tasks, model.FromAsanaTask(taskAsana))
	}

	return tasks, nil
}

func (a *Asana) Create(taskName string) error {
	return a.c.CreateTask(taskName)
}

func (a *Asana) MarkAsDone(taskName string) error {
	return a.c.MarkAsDone(taskName)
}

func NewAsana(c *api.AsanaClient) Storage {
	return &Asana{c: c}
}
