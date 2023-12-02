package api

import (
	"bitbucket.org/mikehouston/asana-go"
)

type AsanaClient struct {
	c                  *asana.Client
	workspace, project string
}

type Option func(*AsanaClient)

func WithProject(project string) Option {
	return func(c *AsanaClient) { c.project = project }
}

func NewAsanaClient(
	accessToken string,
	opts ...Option,
) *AsanaClient {
	c := &AsanaClient{
		c:       asana.NewClientWithAccessToken(accessToken),
		project: "personal",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
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
