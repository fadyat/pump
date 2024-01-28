package internal

import (
	"errors"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/internal/model"
	"github.com/fadyat/pump/pkg"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type Service interface {
	Get() ([]*model.Task, error)
	GetByID(taskID string) (*model.Task, error)
	Create(taskName string) error
	MarkAsDone(taskID string) error
	SelectGoal(manualTaskID string, dueAt *time.Time) (*model.Task, error)
	Reopen(taskID string) error
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

func (r *svc) GetByID(taskID string) (*model.Task, error) {
	return r.storage.GetByID(taskID)
}

func (r *svc) Create(taskName string) error {
	return r.storage.Create(taskName)
}

func (r *svc) MarkAsDone(taskID string) error {
	return r.storage.MarkAsDone(taskID)
}

func (r *svc) SelectGoal(manualTaskID string, dueAt *time.Time) (*model.Task, error) {
	var (
		selectedTask *model.Task
		err          error
	)

	if manualTaskID != "" {
		selectedTask, err = r.storage.GetByID(manualTaskID)
	} else {
		selectedTask, err = r.selectRndTask()
	}
	if err != nil {
		return nil, err
	}

	if err := r.storage.SetDueDate(selectedTask.ID, dueAt); err != nil {
		return nil, err
	}

	selectedTask.DueAt = dueAt
	return selectedTask, nil
}

func (r *svc) Reopen(taskID string) error {
	return r.storage.Reopen(taskID)
}

func (r *svc) selectRndTask() (*model.Task, error) {
	tasks, err := r.storage.Get()
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, ErrTaskNotFound
	}

	return pkg.TakeRand(tasks), nil
}
