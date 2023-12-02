package driver

import (
	"errors"
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
