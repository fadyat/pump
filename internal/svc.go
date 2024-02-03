package internal

import (
	"errors"
	"github.com/fadyat/pump/cmd/flags"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/internal/model"
	"github.com/fadyat/pump/pkg"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

//go:generate mockery --name=IService --output=../mocks --filename=svc.go
type IService interface {
	Get(f *flags.GetFlags) ([]*model.Task, error)
	GetByID(taskID string) (*model.Task, error)
	Create(f *flags.CreateFlags) error
	MarkAsDone(taskID, summary string) error
	SelectGoal(manualTaskID string, dueAt *time.Time) (*model.Task, error)
	Reopen(taskID, summary string) error
	Update(task *model.Task) error
}

type svc struct {
	storage driver.Storage
}

func NewSvc(storage driver.Storage) IService {
	return &svc{storage: storage}
}

func (r *svc) Get(f *flags.GetFlags) ([]*model.Task, error) {
	return r.storage.Get(f)
}

func (r *svc) GetByID(taskID string) (*model.Task, error) {
	return r.storage.GetByID(taskID)
}

func (r *svc) Create(f *flags.CreateFlags) error {
	return r.storage.Create(f)
}

func (r *svc) MarkAsDone(taskID, summary string) error {
	return r.storage.MarkAsDone(taskID, summary)
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

func (r *svc) Reopen(taskID, summary string) error {
	return r.storage.Reopen(taskID, summary)
}

func (r *svc) Update(task *model.Task) error {
	return r.storage.Update(task)
}

func (r *svc) selectRndTask() (*model.Task, error) {
	tasks, err := r.storage.Get(flags.NewGetFlags())
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, ErrTaskNotFound
	}

	return pkg.TakeRand(tasks), nil
}
