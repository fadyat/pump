package commands_test

import (
	"errors"
	"github.com/fadyat/pump/cmd/commands"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

type command struct {
	out       strings.Builder
	withError error
}

func (c *command) buildCommand(cmd string, args ...string) error {
	c.out.WriteString(cmd)
	c.out.WriteString(" ")
	c.out.WriteString(strings.Join(args, " "))
	return c.withError
}

func TestBrowseTaskV2(t *testing.T) {
	cases := []struct {
		name    string
		config  *internal.Config
		runner  *command
		args    []string
		wantMsg string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "browse project",
			config:  asanaConfig,
			args:    []string{},
			runner:  &command{},
			wantMsg: "open https://app.asana.com/0/asana_project/",
			wantErr: assert.NoError,
		},
		{
			name:    "browse task",
			config:  asanaConfig,
			args:    []string{"1234567890"},
			runner:  &command{},
			wantMsg: "open https://app.asana.com/0/asana_project/1234567890",
			wantErr: assert.NoError,
		},
		{
			name:    "browse task with -t flag",
			config:  asanaConfig,
			runner:  &command{},
			args:    []string{"-t", "1234567890"},
			wantMsg: "open https://app.asana.com/0/asana_project/1234567890",
			wantErr: assert.NoError,
		},
		{
			name:   "failed to run command",
			config: asanaConfig,
			runner: &command{
				withError: errors.New("command failed"),
			},
			args: []string{"1234567890"},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.Contains(t, err.Error(), "command failed")
			},
		},
		{
			name: "command not available",
			config: &internal.Config{
				Driver: "unknown",
			},
			runner: &command{},
			args:   []string{"1234567890"},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.Contains(t, err.Error(), "command not available for the current driver")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := mocks.NewIService(t)

			manager := &commands.Manager{
				Config:     tc.config,
				RunCommand: tc.runner.buildCommand,
			}
			cmd := commands.BrowseTaskV2(manager)
			cmd.SetArgs(tc.args)

			tc.wantErr(t, cmd.Execute())
			require.True(
				t,
				strings.Contains(tc.runner.out.String(), tc.wantMsg),
				"expected %q to contain %q",
				tc.runner.out.String(),
				tc.wantMsg,
			)
			service.AssertExpectations(t)
		})
	}
}
