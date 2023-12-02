package driver

import (
	"errors"
	"github.com/fadyat/pump/internal/api"
	"github.com/fadyat/pump/internal/model"
)

var (
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrTaskNotFound      = errors.New("task not found")
)

type Storage interface {
	Get(filters ...func(task *model.Task) bool) ([]*model.Task, error)
	Create(taskName string) error
	MarkAsDone(taskName string) error
}

func New(
	driverType string,
	storageOpts map[string]any,
) (Storage, error) {
	switch driverType {
	case "asana":
		return NewAsana(api.NewAsanaClient(
			storageOpts["token"].(string),
			storageOpts["project"].(string),
		)), nil
	case "fs":
		return NewFs(storageOpts["file"].(string)), nil
	}

	return nil, errors.New("unknown driver type")
}
