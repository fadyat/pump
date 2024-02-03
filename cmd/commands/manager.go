package commands

import (
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"slices"
)

var (
	availableCommands = map[string][]string{
		driver.AsanaDriver: {
			"create", "get", "browse",
			"select", "edit", "configure",
			"complete",
		},
	}
)

type Manager struct {
	Config *internal.Config

	// ServiceMaker is a function that returns a new instance of the service.
	//
	// Passing it as a function, because it's not always possible to create a
	// service instance, because some commands require additional configuration.
	//
	// For example, `pump configure` command require driver to be set, but if
	// you will try to create a service instance, it will fail, because driver
	// is not set yet.
	ServiceMaker func() internal.IService

	// RunCommand is a function that runs a command with the given arguments.
	RunCommand func(cmd string, args ...string) error
}

func NewManager(
	config *internal.Config,
	serviceFn func() internal.IService,
	runCmdFn func(cmd string, args ...string) error,
) *Manager {
	return &Manager{
		Config:       config,
		ServiceMaker: serviceFn,
		RunCommand:   runCmdFn,
	}
}

func (m *Manager) IsCommandAvailable(cmd string) bool {
	commands, supported := availableCommands[m.Config.Driver]
	if !supported {
		return false
	}

	return slices.Contains(commands, cmd)
}
