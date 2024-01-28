package driver

import (
	"errors"
	"github.com/fadyat/pump/internal/api"
	"github.com/fadyat/pump/internal/model"
	"time"
)

var (
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrTaskNotFound      = errors.New("task not found")
)

type Storage interface {
	Get() ([]*model.Task, error)
	GetByID(taskID string) (*model.Task, error)
	Create(taskName string) error
	SetDueDate(taskID string, dueAt *time.Time) error
	MarkAsDone(taskID, summary string) error
	Reopen(taskID, summary string) error
	Update(task *model.Task) error
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

	return nil, errors.New("requested driver not found, run `pump configure`")
}
