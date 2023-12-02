package driver

import (
	"encoding/json"
	"errors"
	"github.com/fadyat/pump/internal/model"
	"os"
	"path/filepath"
	"slices"
	"time"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

type FileStorage struct {
	file string
}

func (f *FileStorage) Get(filters ...func(task *model.Task) bool) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := readJson(f.file, &tasks); err != nil {
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

	return writeJson(f.file, tasks)
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
	return writeJson(f.file, tasks)
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

func readJson(path string, v interface{}) error {
	content, err := readFileBytes(path)

	switch {
	case err == nil:
		return json.Unmarshal(content, v)
	case errors.Is(err, ErrFileNotFound):
		return nil
	default:
		return err
	}
}

func readFileBytes(path string) ([]byte, error) {
	b, err := os.ReadFile(path)

	switch {
	case err == nil:
		return b, nil
	case os.IsNotExist(err):
		return nil, ErrFileNotFound
	default:
		return nil, err
	}
}

func writeJson(path string, v interface{}) error {
	var data, err = json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	var dir = filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func ptr[T any](v T) *T {
	return &v
}
