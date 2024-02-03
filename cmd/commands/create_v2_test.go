package commands_test

import (
	"bytes"
	"errors"
	"github.com/fadyat/pump/cmd/commands"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var (
	asanaConfig = &internal.Config{
		Driver: "asana",
		DriverOpts: map[string]any{
			"token":   "asana_token",
			"project": "asana_project",
		},
	}
)

func prepareCommand(cmd *cobra.Command, args []string) *bytes.Buffer {
	out := bytes.NewBufferString("")
	cmd.SetOut(out)
	cmd.SetArgs(args)

	return out
}

func TestCreateTask_V2(t *testing.T) {
	cases := []struct {
		name       string
		config     *internal.Config
		applyMocks func(service *mocks.IService)
		args       []string
		wantMsg    string
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name:   "create via name flag",
			config: asanaConfig,
			applyMocks: func(service *mocks.IService) {
				service.On("Create", "Task name").
					Return(nil).
					Once()
			},
			args:    []string{"-n", "Task name"},
			wantMsg: "Task created successfully",
			wantErr: assert.NoError,
		},
		{
			name:   "create via positional argument",
			config: asanaConfig,
			applyMocks: func(service *mocks.IService) {
				service.On("Create", "Task name").
					Return(nil).
					Once()
			},
			args:    []string{"Task name"},
			wantMsg: "Task created successfully",
			wantErr: assert.NoError,
		},
		{
			name:       "create with no arguments",
			config:     asanaConfig,
			applyMocks: func(service *mocks.IService) {},
			args:       []string{""},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Contains(t, err.Error(), "task name is required")
			},
		},
		{
			name:       "create with empty name",
			config:     asanaConfig,
			applyMocks: func(service *mocks.IService) {},
			args:       []string{"-n", ""},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Contains(t, err.Error(), "task name is required")
			},
		},
		{
			name:       "create too many arguments",
			config:     asanaConfig,
			applyMocks: func(service *mocks.IService) {},
			args:       []string{"Task name", "extra"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Contains(t, err.Error(), "accepts at most 1 arg(s)")
			},
		},
		{
			name:       "command not available",
			config:     &internal.Config{Driver: "unknown"},
			applyMocks: func(service *mocks.IService) {},
			args:       []string{"Task name"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Contains(t, err.Error(), "command not available for the current driver")
			},
		},
		{
			name:   "create with service error",
			config: asanaConfig,
			applyMocks: func(service *mocks.IService) {
				service.On("Create", "Task name").
					Return(errors.New("service error")).
					Once()
			},
			args: []string{"Task name"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Contains(t, err.Error(), "service error")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := mocks.NewIService(t)
			tc.applyMocks(service)

			manager := commands.NewManager(tc.config, func() internal.IService {
				return service
			})
			cmd := commands.CreateTaskV2(manager)
			out := prepareCommand(cmd, tc.args)

			tc.wantErr(t, cmd.Execute())
			require.True(t, strings.Contains(out.String(), tc.wantMsg))
			service.AssertExpectations(t)
		})
	}
}
