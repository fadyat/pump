package api

import (
	"bitbucket.org/mikehouston/asana-go"
	"github.com/fadyat/pump/pkg"
	"time"
)

const (
	tasksLimit = 100
)

type AsanaClient struct {
	c       *asana.Client
	project string
}

func NewAsanaClient(
	accessToken string,
	project string,
) *AsanaClient {
	return &AsanaClient{
		c:       asana.NewClientWithAccessToken(accessToken),
		project: project,
	}
}

func (a *AsanaClient) GetTasks() ([]*asana.Task, error) {
	var (
		query = &asana.TaskQuery{
			Project:        a.project,
			CompletedSince: "now",
		}
		option = &asana.Options{
			Limit: tasksLimit,
			Fields: []string{
				"id",
				"created_at",
				"due_at",
				"completed",
				"name",
			},
		}
	)

	tasks, _, err := a.c.QueryTasks(query, option)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (a *AsanaClient) CreateTask(taskName string) error {
	var taskCreateRequest = &asana.CreateTaskRequest{
		TaskBase: asana.TaskBase{Name: taskName},
		Projects: []string{a.project},
	}

	_, err := a.c.CreateTask(taskCreateRequest)
	return err
}

func (a *AsanaClient) MarkAsDone(taskID string) error {
	task := &asana.Task{ID: taskID}

	update := &asana.UpdateTaskRequest{
		TaskBase: asana.TaskBase{
			Completed: pkg.Ptr(true),
		},
	}

	return task.Update(a.c, update)
}

func (a *AsanaClient) SetDueDate(taskID string, dueAt *time.Time) error {
	task := &asana.Task{ID: taskID}

	update := &asana.UpdateTaskRequest{
		TaskBase: asana.TaskBase{
			Name:  task.Name,
			DueAt: dueAt,
		},
	}

	return task.Update(a.c, update)
}

func (a *AsanaClient) GetTask(id string) (*asana.Task, error) {
	task := &asana.Task{ID: id}
	err := task.Fetch(a.c)
	if err != nil {
		return nil, err
	}

	return task, nil
}
