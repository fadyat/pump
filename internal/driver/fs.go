package driver

import (
	"github.com/fadyat/pump/internal/model"
	"github.com/fadyat/pump/pkg"
	"slices"
	"time"
)

type FileStorage struct {
	file string
}

func (f *FileStorage) Get() ([]*model.Task, error) {
	var tasks []*model.Task
	if err := pkg.ReadJson(f.file, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (f *FileStorage) Create(taskName string) (err error) {
	var tasks []*model.Task
	if tasks, err = f.Get(); err != nil {
		return err
	}

	if slices.IndexFunc(tasks, func(task *model.Task) bool {
		return task.Name == taskName
	}) != -1 {
		return ErrTaskAlreadyExists
	}

	tasks = append(tasks, &model.Task{
		ID:        pkg.RandString(8),
		Name:      taskName,
		CreatedAt: pkg.Now(),
	})

	return pkg.WriteJson(f.file, tasks)
}

func (f *FileStorage) MarkAsDone(taskID string) (err error) {
	var tasks []*model.Task
	if tasks, err = f.Get(); err != nil {
		return err
	}

	var task *model.Task
	if task, err = f.findTaskByID(tasks, taskID); err != nil {
		return err
	}

	task.Done = true
	return pkg.WriteJson(f.file, tasks)
}

func (f *FileStorage) GetByID(taskID string) (*model.Task, error) {
	if tasks, err := f.Get(); err != nil {
		return nil, err
	} else {
		return f.findTaskByID(tasks, taskID)
	}
}

func (f *FileStorage) findTaskByID(tasks []*model.Task, taskID string) (*model.Task, error) {
	var idx = slices.IndexFunc(tasks, func(task *model.Task) bool {
		return task.ID == taskID
	})

	if idx == -1 {
		return nil, ErrTaskNotFound
	}

	return tasks[idx], nil
}

func (f *FileStorage) SetDueDate(taskID string, dueAt *time.Time) (err error) {
	var tasks []*model.Task
	if tasks, err = f.Get(); err != nil {
		return err
	}

	var task *model.Task
	if task, err = f.findTaskByID(tasks, taskID); err != nil {
		return err
	}

	task.DueAt = dueAt
	return pkg.WriteJson(f.file, tasks)
}

func NewFs(path string) Storage {
	return &FileStorage{file: path}
}
