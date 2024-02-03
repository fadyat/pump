package driver

import (
	"errors"
	"github.com/fadyat/pump/cmd/flags"
	"github.com/fadyat/pump/internal/api"
	"github.com/fadyat/pump/internal/model"
	"time"
)

type Storage interface {
	Get(f *flags.GetFlags) ([]*model.Task, error)
	GetByID(taskID string) (*model.Task, error)
	Create(taskName string) error
	SetDueDate(taskID string, dueAt *time.Time) error
	MarkAsDone(taskID, summary string) error
	Reopen(taskID, summary string) error
	Update(task *model.Task) error
}

const (
	AsanaDriver = "asana"
)

func New(
	driverType string,
	storageOpts map[string]any,
) (Storage, error) {
	if driverType == AsanaDriver {
		return NewAsana(api.NewAsanaClient(
			storageOpts["token"].(string),
			storageOpts["project"].(string),
		)), nil
	}

	return nil, errors.New("requested driver not found, run `pump configure`")
}
