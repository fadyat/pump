package commands

import (
	"github.com/fadyat/pump/cmd/flags"
	"github.com/fadyat/pump/internal"
	"github.com/spf13/cobra"
)

const (
	getShort = "Get tasks based on some criteria"

	getExample = `  pump get`
)

func GetTask(m *Manager) *cobra.Command {
	var f = flags.NewGetFlags()

	cmd := &cobra.Command{
		Use:                   "get",
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

			tasks, err := m.ServiceMaker().Get(f)
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

	cmd.Flags().BoolVarP(&f.OnlyActive, "active", "a", f.OnlyActive, "Show only active tasks")
	cmd.Flags().BoolVarP(&f.OnlyInactive, "inactive", "i", f.OnlyInactive, "Show only inactive tasks")
	cmd.MarkFlagsMutuallyExclusive("active", "inactive")

	return cmd
}
