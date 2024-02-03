package commands

import (
	"errors"
	"github.com/fadyat/pump/cmd/flags"
	"github.com/spf13/cobra"
)

var (
	ErrCommandNotAvailable = errors.New("command not available for the current driver")
)

const (
	createShort = "Create a new task"

	createExample = `  pump create -n "Task name"
  pump create "Task name"`
)

func CreateTaskV2(m *Manager) *cobra.Command {
	var f = flags.NewCreateFlags()

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

			f.Override(args)
			return f.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := m.ServiceMaker().Create(f); err != nil {
				return err
			}

			cmd.Println("Task created successfully")
			return nil
		},
	}

	cmd.Flags().StringVarP(&f.Name, "name", "n", f.Name, "Task name")
	cmd.Flags().StringVarP(&f.Description, "description", "d", f.Description, "Task description")

	return cmd
}
