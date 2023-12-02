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
			var (
				svc     = internal.NewSvc(driver.NewFs(config.TasksFile))
				printer = internal.NewTablePrinter("name", "created at")
			)

			var (
				tasks   []*model.Task
				filters = []func(task *model.Task) bool{
					func(task *model.Task) bool { return !task.Done },
				}
			)
			if tasks, err = svc.Get(filters...); err != nil {
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
