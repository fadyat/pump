package internal

import (
	"errors"
	"slices"
	"time"
)

var (
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrTaskNotFound      = errors.New("task not found")
)

type Service interface {
	Get(filters ...func(task *Task) bool) ([]*Task, error)
	Create(taskName string) error
	MarkAsDone(taskName string) error
}

type svc struct {
	tasksFile string
}

func NewSvc(
	tasksFile string,
) Service {
	return &svc{tasksFile: tasksFile}
}

func (r *svc) Get(
	filters ...func(task *Task) bool,
) ([]*Task, error) {
	var tasks []*Task
	if err := readJson(r.tasksFile, &tasks); err != nil {
		return nil, err
	}

	var filteredTasks []*Task = make([]*Task, 0)
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

func (r *svc) Create(taskName string) (err error) {
	var tasks []*Task
	if tasks, err = r.Get(); err != nil {
		return err
	}

	if slices.IndexFunc(tasks, func(task *Task) bool {
		return task.Name == taskName
	}) != -1 {
		return ErrTaskAlreadyExists
	}

	tasks = append(tasks, &Task{
		Name:      taskName,
		CreatedAt: time.Now(),
	})

	return writeJson(r.tasksFile, tasks)
}

func (r *svc) MarkAsDone(taskName string) (err error) {
	var tasks []*Task
	if tasks, err = r.Get(); err != nil {
		return err
	}

	var task *Task
	if task, err = r.findTaskByName(tasks, taskName); err != nil {
		return err
	}

	task.Done = true
	return writeJson(r.tasksFile, tasks)
}

func (r *svc) findTaskByName(
	tasks []*Task,
	taskName string,
) (*Task, error) {
	var idx = slices.IndexFunc(tasks, func(task *Task) bool {
		return task.Name == taskName
	})

	if idx == -1 {
		return nil, ErrTaskNotFound
	}

	return tasks[idx], nil
}
