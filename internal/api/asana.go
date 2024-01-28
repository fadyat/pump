package api

import (
	"bitbucket.org/mikehouston/asana-go"
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

func (a *AsanaClient) ChangeCompletedStatus(taskID, summary string, status bool) error {
	task := &asana.Task{ID: taskID}

	update := &asana.UpdateTaskRequest{
		TaskBase: asana.TaskBase{
			Completed: &status,
		},
	}

	if err := task.Update(a.c, update); err != nil {
		return err
	}

	if summary == "" {
		return nil
	}

	_, err := task.CreateComment(a.c, &asana.StoryBase{
		Text:     summary,
		IsPinned: false,
	})

	return err
}

func (a *AsanaClient) SetDueDate(taskID string, dueAt *time.Time) error {
	task := &asana.Task{ID: taskID}

	update := &asana.UpdateTaskRequest{
		TaskBase: asana.TaskBase{
			Name:  task.Name,
			DueAt: dueAt,
		},
		Assignee: "me",
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
