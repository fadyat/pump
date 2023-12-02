package internal

import (
	"errors"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/internal/model"
)

var (
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrTaskNotFound      = errors.New("task not found")
)

type Service interface {
	Get(filters ...func(task *model.Task) bool) ([]*model.Task, error)
	Create(taskName string) error
	MarkAsDone(taskName string) error
}

type svc struct {
	storage driver.Storage
}

func NewSvc(storage driver.Storage) Service {
	return &svc{storage: storage}
}

func (r *svc) Get(
	filters ...func(task *model.Task) bool,
) ([]*model.Task, error) {
	return r.storage.Get(filters...)
}

func (r *svc) Create(taskName string) (err error) {
	return r.storage.Create(taskName)
}

func (r *svc) MarkAsDone(taskName string) (err error) {
	return r.storage.MarkAsDone(taskName)
}
