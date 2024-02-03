package commands_test

import (
	"errors"
	"github.com/fadyat/pump/cmd/commands"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/model"
	"github.com/fadyat/pump/mocks"
	"github.com/fadyat/pump/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

var (
	fixedTime = time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
)

func TestGetTaskV2(t *testing.T) {
	cases := []struct {
		name       string
		config     *internal.Config
		applyMocks func(service *mocks.IService)
		args       []string
		wantMsg    string
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "command not available",
			config: &internal.Config{
				Driver: "unknown",
			},
			applyMocks: func(service *mocks.IService) {},
			args:       []string{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Contains(t, err.Error(), "command not available for the current driver")
			},
		},
		{
			name:   "get tasks",
			config: asanaConfig,
			applyMocks: func(service *mocks.IService) {
				service.On("Get").
					Return([]*model.Task{{
						ID:        "1",
						Name:      "Task 1",
						CreatedAt: pkg.Ptr(fixedTime),
					}}, nil).
					Once()
			},
			args:    []string{},
			wantMsg: `1  | Task 1 | 2021-01-01 00:00:00 | `,
			wantErr: assert.NoError,
		},
		{
			name:   "no tasks found",
			config: asanaConfig,
			applyMocks: func(service *mocks.IService) {
				service.On("Get").
					Return([]*model.Task{}, nil).
					Once()
			},
			args:    []string{},
			wantMsg: "No tasks found",
			wantErr: assert.NoError,
		},
		{
			name:   "get tasks error",
			config: asanaConfig,
			applyMocks: func(service *mocks.IService) {
				service.On("Get").
					Return(nil, errors.New("service error")).
					Once()
			},
			args: []string{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Contains(t, err.Error(), "service error")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := mocks.NewIService(t)
			tc.applyMocks(service)

			manager := &commands.Manager{
				Config:       tc.config,
				ServiceMaker: func() internal.IService { return service },
			}
			cmd := commands.GetTaskV2(manager)
			out := prepareCommand(cmd, tc.args)

			tc.wantErr(t, cmd.Execute())
			require.True(t, strings.Contains(out.String(), tc.wantMsg))
			service.AssertExpectations(t)
		})
	}
}
