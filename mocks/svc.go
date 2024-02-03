// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	flags "github.com/fadyat/pump/cmd/flags"

	mock "github.com/stretchr/testify/mock"

	model "github.com/fadyat/pump/internal/model"

	time "time"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// Create provides a mock function with given fields: taskName
func (_m *IService) Create(taskName string) error {
	ret := _m.Called(taskName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(taskName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *IService) Get(_a0 *flags.GetFlags) ([]*model.Task, error) {
	ret := _m.Called(_a0)

	var r0 []*model.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(*flags.GetFlags) ([]*model.Task, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*flags.GetFlags) []*model.Task); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(*flags.GetFlags) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: taskID
func (_m *IService) GetByID(taskID string) (*model.Task, error) {
	ret := _m.Called(taskID)

	var r0 *model.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Task, error)); ok {
		return rf(taskID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Task); ok {
		r0 = rf(taskID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkAsDone provides a mock function with given fields: taskID, summary
func (_m *IService) MarkAsDone(taskID string, summary string) error {
	ret := _m.Called(taskID, summary)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(taskID, summary)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reopen provides a mock function with given fields: taskID, summary
func (_m *IService) Reopen(taskID string, summary string) error {
	ret := _m.Called(taskID, summary)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(taskID, summary)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectGoal provides a mock function with given fields: manualTaskID, dueAt
func (_m *IService) SelectGoal(manualTaskID string, dueAt *time.Time) (*model.Task, error) {
	ret := _m.Called(manualTaskID, dueAt)

	var r0 *model.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *time.Time) (*model.Task, error)); ok {
		return rf(manualTaskID, dueAt)
	}
	if rf, ok := ret.Get(0).(func(string, *time.Time) *model.Task); ok {
		r0 = rf(manualTaskID, dueAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *time.Time) error); ok {
		r1 = rf(manualTaskID, dueAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: task
func (_m *IService) Update(task *model.Task) error {
	ret := _m.Called(task)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIService creates a new instance of IService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IService {
	mock := &IService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
