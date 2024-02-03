package commands

import (
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
)

var (
	availableCommands = map[string]map[string]bool{
		driver.AsanaDriver: {
			"create": true,
			"get":    true,
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
}

func NewManager(
	config *internal.Config,
	serviceFn func() internal.IService,
) *Manager {
	return &Manager{
		Config:       config,
		ServiceMaker: serviceFn,
	}
}

func (m *Manager) IsCommandAvailable(cmd string) bool {
	commands, supported := availableCommands[m.Config.Driver]
	if !supported {
		return false
	}

	_, supported = commands[cmd]
	return supported
}
