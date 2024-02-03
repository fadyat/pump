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

const (
	AsanaDriver      = "asana"
	FileSystemDriver = "fs"
)

func New(
	driverType string,
	storageOpts map[string]any,
) (Storage, error) {
	switch driverType {
	case AsanaDriver:
		return NewAsana(api.NewAsanaClient(
			storageOpts["token"].(string),
			storageOpts["project"].(string),
		)), nil
	case FileSystemDriver:
		return NewFs(storageOpts["file"].(string)), nil
	}

	return nil, errors.New("requested driver not found, run `pump configure`")
}
