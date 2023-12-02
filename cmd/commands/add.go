package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/spf13/cobra"
)

func AddTask(
	config *internal.Config,
) *cobra.Command {
	var taskName string

	return &cobra.Command{
		Use:   "add [name]",
		Short: "Add a new task",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("task name is required")
			}

			taskName = args[0]
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var svc = internal.NewSvc(driver.NewFs(config.TasksFile))
			if err := svc.Create(taskName); err != nil {
				return err
			}

			fmt.Println("Task created successfully")
			return nil
		},
		SilenceUsage: true,
	}
}
