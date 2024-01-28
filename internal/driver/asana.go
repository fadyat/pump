package driver

import (
	"github.com/fadyat/pump/internal/api"
	"github.com/fadyat/pump/internal/model"
	"time"
)

type Asana struct {
	c *api.AsanaClient
}

func (a *Asana) Get() ([]*model.Task, error) {
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

func (a *Asana) GetByID(taskID string) (*model.Task, error) {
	taskAsana, err := a.c.GetTask(taskID)
	if err != nil {
		return nil, err
	}

	return model.FromAsanaTask(taskAsana), nil
}

func (a *Asana) Create(taskName string) error {
	return a.c.CreateTask(taskName)
}

func (a *Asana) MarkAsDone(taskID, summary string) error {
	return a.c.ChangeCompletedStatus(taskID, summary, true)
}

func (a *Asana) Reopen(taskID, summary string) error {
	return a.c.ChangeCompletedStatus(taskID, summary, false)
}

func (a *Asana) SetDueDate(taskID string, dueAt *time.Time) error {
	return a.c.SetDueDate(taskID, dueAt)
}

func (a *Asana) Update(task *model.Task) error {
	return a.c.UpdateTask(task)
}

func NewAsana(c *api.AsanaClient) Storage {
	return &Asana{c: c}
}
