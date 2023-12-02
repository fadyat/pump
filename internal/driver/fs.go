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

func (f *FileStorage) Get(filters ...func(task *model.Task) bool) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := pkg.ReadJson(f.file, &tasks); err != nil {
		return nil, err
	}

	var filteredTasks = make([]*model.Task, 0)
	for _, filter := range filters {
		for _, task := range tasks {
			if filter(task) {
				filteredTasks = append(filteredTasks, task)
			}
		}

		tasks = filteredTasks
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
		Name:      taskName,
		CreatedAt: ptr(time.Now()),
	})

	return pkg.WriteJson(f.file, tasks)
}

func (f *FileStorage) MarkAsDone(taskName string) (err error) {
	var tasks []*model.Task
	if tasks, err = f.Get(); err != nil {
		return err
	}

	var task *model.Task
	if task, err = f.findTaskByName(tasks, taskName); err != nil {
		return err
	}

	task.Done = true
	return pkg.WriteJson(f.file, tasks)
}

func (f *FileStorage) findTaskByName(tasks []*model.Task, taskName string) (*model.Task, error) {
	var idx = slices.IndexFunc(tasks, func(task *model.Task) bool {
		return task.Name == taskName
	})

	if idx == -1 {
		return nil, ErrTaskNotFound
	}

	return tasks[idx], nil
}

func NewFs(path string) Storage {
	return &FileStorage{file: path}
}

func ptr[T any](v T) *T {
	return &v
}
