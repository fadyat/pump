package internal

import (
	"crypto/rand"
	"errors"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/internal/model"
	"math/big"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type Service interface {
	Get() ([]*model.Task, error)
	Create(taskName string) error
	MarkAsDone(taskName string) error
	SelectGoal(dueAt *time.Time) (*model.Task, error)
}

type svc struct {
	storage driver.Storage
}

func NewSvc(storage driver.Storage) Service {
	return &svc{storage: storage}
}

func (r *svc) Get() ([]*model.Task, error) {
	return r.storage.Get()
}

func (r *svc) Create(taskName string) error {
	return r.storage.Create(taskName)
}

func (r *svc) MarkAsDone(taskName string) error {
	return r.storage.MarkAsDone(taskName)
}

func (r *svc) SelectGoal(dueAt *time.Time) (*model.Task, error) {
	tasks, err := r.storage.Get()
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, ErrTaskNotFound
	}

	selectedTask, err := takeRandomTask(tasks)
	if err != nil {
		return nil, err
	}

	if err = r.storage.SetDueDate(selectedTask.Name, dueAt); err != nil {
		return nil, err
	}

	return selectedTask, nil
}

func takeRandomTask(tasks []*model.Task) (*model.Task, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(tasks))))
	if err != nil {
		return nil, err
	}

	return tasks[bigInt.Int64()], nil
}
