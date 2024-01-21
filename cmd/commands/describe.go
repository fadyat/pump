package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/spf13/cobra"
)

func DescribeTask(cfg *internal.Config) *cobra.Command {
	var taskID string

	return &cobra.Command{
		Use:   "describe",
		Short: "Describe specified task",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if cfg.Driver != "asana" {
				return fmt.Errorf("asana driver is only supported")
			}

			if len(args) > 0 {
				taskID = args[0]
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			driv, err := driver.New(cfg.Driver, cfg.GetDriverOpts())
			if err != nil {
				return err
			}

			var (
				svc    = internal.NewSvc(driv)
				editor = internal.NewEditor("Name", "Description")
			)

			task, err := svc.GetByID(taskID)
			if err != nil {
				return err
			}

			if err := editor.Edit([]string{task.Name, task.Description}); err != nil {
				return err
			}

			return nil
		},
		SilenceUsage: true,
	}
}
