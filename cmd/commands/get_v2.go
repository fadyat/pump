package commands

import (
	"github.com/fadyat/pump/internal"
	"github.com/spf13/cobra"
)

const (
	getShort = "Get tasks based on some criteria"

	getExample = `  pump get`
)

func GetTaskV2(m *Manager) *cobra.Command {
	return &cobra.Command{
		Use:                   "get_v2",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Short:                 getShort,
		Example:               getExample,
		Args:                  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !m.IsCommandAvailable("get") {
				return ErrCommandNotAvailable
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			printer := internal.NewTablePrinter(
				cmd.OutOrStdout(),
				"id", "name", "created at", "due at",
			)

			tasks, err := m.ServiceMaker().Get()
			if err != nil {
				return err
			}

			values := internal.AsPrintable(tasks)
			if len(values) == 0 {
				cmd.Println("No tasks found")
				return nil
			}

			printer.Print(values)
			return nil
		},
	}

}
