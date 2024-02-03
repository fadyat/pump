package commands

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

var (
	ErrTaskNameRequired    = errors.New("task name is required")
	ErrCommandNotAvailable = errors.New("command not available for the current driver")
)

const (
	createShort = "Create a new task"

	createExample = `  pump create -n "Task name"
  pump create "Task name"`
)

type CreateFlags struct {
	Name string
}

func (f *CreateFlags) Override(args []string) {
	if len(args) > 0 {
		f.Name = args[0]
	}
}

func (f *CreateFlags) Prepare() {
	f.Name = strings.TrimSpace(f.Name)
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

func CreateTaskV2(m *Manager) *cobra.Command {
	var flags = NewCreateFlags()

	cmd := &cobra.Command{
		Use:                   "create_v2 (-n NAME)",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Short:                 createShort,
		Example:               createExample,
		Args:                  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !m.IsCommandAvailable("create") {
				return ErrCommandNotAvailable
			}

			flags.Override(args)
			return flags.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := m.ServiceMaker().Create(flags.Name); err != nil {
				return err
			}

			cmd.Println("Task created successfully")
			return nil
		},
	}

	cmd.Flags().StringVarP(&flags.Name, "name", "n", flags.Name, "Task name")

	return cmd
}
