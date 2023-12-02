package internal

import (
	"crypto/rand"
	"errors"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/internal/model"
	"math/big"
)

var (
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrTaskNotFound      = errors.New("task not found")
)

type Service interface {
	Get(filters ...func(task *model.Task) bool) ([]*model.Task, error)
	Create(taskName string) error
	MarkAsDone(taskName string) error
	SelectGoal(filters ...func(task *model.Task) bool) (*model.Task, error)
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

func (r *svc) Create(taskName string) error {
	return r.storage.Create(taskName)
}

func (r *svc) MarkAsDone(taskName string) error {
	return r.storage.MarkAsDone(taskName)
}

func (r *svc) SelectGoal(
	filters ...func(task *model.Task) bool,
) (*model.Task, error) {
	tasks, err := r.storage.Get(filters...)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, ErrTaskNotFound
	}

	return takeRandomTask(tasks)
}

func takeRandomTask(tasks []*model.Task) (*model.Task, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(tasks))))
	if err != nil {
		return nil, err
	}

	return tasks[bigInt.Int64()], nil
}
