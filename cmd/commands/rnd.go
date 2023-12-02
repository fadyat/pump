package commands

import (
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/internal/model"
	"github.com/spf13/cobra"
)

var (
	activeTasksFilter = func(task *model.Task) bool {
		return !task.Done
	}
)

func SelectTask(
	config *internal.Config,
) *cobra.Command {
	return &cobra.Command{
		Use:   "select",
		Short: "Select a task to work on",
		RunE: func(cmd *cobra.Command, args []string) error {
			driv, err := driver.New(config.Driver, config.GetDriverOpts())
			if err != nil {
				return err
			}

			svc := internal.NewSvc(driv)
			task, err := svc.SelectGoal(activeTasksFilter)
			if err != nil {
				return err
			}

			printer := internal.NewTablePrinter("name", "created at")
			printer.Print([][]string{task.ToPrintable()})
			return nil
		},
		SilenceUsage: true,
	}
}
