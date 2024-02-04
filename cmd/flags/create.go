package flags

import (
	"errors"
	"strings"
)

var (
	ErrTaskNameRequired = errors.New("task name is required")
)

type CreateFlags struct {
	Name        string
	Description string
}

func (f *CreateFlags) Override(args []string) {
	if len(args) > 0 {
		f.Name = args[0]
	}
}

func (f *CreateFlags) Prepare() {
	f.Name = strings.TrimSpace(f.Name)
	f.Description = strings.TrimSpace(f.Description)
}

func (f *CreateFlags) Validate() error {
	f.Prepare()

	if f.Name == "" {
		return ErrTaskNameRequired
	}

	return nil
}

func NewCreateFlags() *CreateFlags {
	return &CreateFlags{}
}
