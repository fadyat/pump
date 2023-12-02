package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/spf13/cobra"
)

func MarkTaskAsDone(
	config *internal.Config,
) *cobra.Command {
	var taskName string

	return &cobra.Command{
		Use:     "done [name]",
		Short:   "Mark a task as done",
		Aliases: []string{"do", "ok", "complete", "finish"},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("task name is required")
			}

			taskName = args[0]
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			driv, err := driver.New(config.Driver, config.GetDriverOpts())
			if err != nil {
				return err
			}

			var svc = internal.NewSvc(driv)
			if err := svc.MarkAsDone(taskName); err != nil {
				return err
			}

			fmt.Println("Task marked as done successfully")
			return nil
		},
		SilenceUsage: true,
	}
}
