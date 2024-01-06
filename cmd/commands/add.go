package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/spf13/cobra"
	"strings"
)

func AddTask(
	config *internal.Config,
) *cobra.Command {
	var taskName string

	return &cobra.Command{
		Use:     "add [name]",
		Short:   "Add a new task",
		Aliases: []string{"new", "create"},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("task name is required")
			}

			taskName = strings.TrimSpace(args[0])
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			driv, err := driver.New(config.Driver, config.GetDriverOpts())
			if err != nil {
				return err
			}

			svc := internal.NewSvc(driv)
			if err := svc.Create(taskName); err != nil {
				return err
			}

			fmt.Println("Task created successfully")
			return nil
		},
		SilenceUsage: true,
	}
}
