package api

import (
	"bitbucket.org/mikehouston/asana-go"
	"errors"
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
			Limit: 100,
			Fields: []string{
				"created_at",
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

func (a *AsanaClient) MarkAsDone(taskName string) error {
	// fixme: think about marking as done via ID, fetching all tasks it's too much
	task, err := a.getTaskByName(taskName)
	if err != nil {
		return err
	}

	taskUpdateRequest := &asana.UpdateTaskRequest{
		TaskBase: asana.TaskBase{
			Name:      task.Name,
			Completed: ptr(true),
		},
	}

	return task.Update(a.c, taskUpdateRequest)
}

func (a *AsanaClient) getTaskByName(taskName string) (*asana.Task, error) {
	tasks, err := a.GetTasks()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.Name == taskName {
			return task, nil
		}
	}

	return nil, errors.New("task not found")
}

func ptr[T any](v T) *T {
	return &v
}
