package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/internal/model"
	"github.com/spf13/cobra"
)

func GetAvailableTask(
	config *internal.Config,
) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get all available tasks",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			driver, err := driver.New(config.Driver, config.GetDriverOpts())
			if err != nil {
				return err
			}

			var (
				svc     = internal.NewSvc(driver)
				printer = internal.NewTablePrinter("name", "created at")
				tasks   []*model.Task
			)

			if tasks, err = svc.Get(activeTasksFilter); err != nil {
				return err
			}

			values := make([][]string, len(tasks))
			for idx, task := range tasks {
				values[idx] = task.ToPrintable()
			}

			if len(values) == 0 {
				fmt.Println("No tasks found")
				return nil
			}

			printer.Print(values)
			return nil
		},
		SilenceUsage: true,
	}
}
