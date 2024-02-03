package driver

import (
	"github.com/fadyat/pump/cmd/flags"
	"github.com/fadyat/pump/internal/api"
	"github.com/fadyat/pump/internal/model"
	"time"
)

type Asana struct {
	c *api.AsanaClient
}

func (a *Asana) Get(f *flags.GetFlags) ([]*model.Task, error) {
	tasksAsana, err := a.c.GetTasks()
	if err != nil {
		return nil, err
	}

	var tasks = make([]*model.Task, 0, len(tasksAsana))
	for _, taskAsana := range tasksAsana {
		task := model.FromAsanaTask(taskAsana)
		if a.takeByActive(f.OnlyActive, task) {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (a *Asana) takeByActive(active bool, task *model.Task) bool {
	if active {
		return task.DueAt != nil
	}

	return true
}

func (a *Asana) GetByID(taskID string) (*model.Task, error) {
	taskAsana, err := a.c.GetTask(taskID)
	if err != nil {
		return nil, err
	}

	return model.FromAsanaTask(taskAsana), nil
}

func (a *Asana) Create(f *flags.CreateFlags) error {
	return a.c.CreateTask(f.Name, f.Description)
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
