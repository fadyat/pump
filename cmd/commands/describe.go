package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/internal/model"
	"github.com/spf13/cobra"
)

// todo: reuse this function in other commands
func newService(cfg *internal.Config) (internal.IService, error) {
	d, err := driver.New(cfg.Driver, cfg.GetDriverOpts())
	if err != nil {
		return nil, err
	}

	return internal.NewSvc(d), nil
}

func autoCompleteArgs[T any](args []T, take func(arg T) string) []string {
	values := make([]string, len(args))
	for idx, arg := range args {
		values[idx] = take(arg)
	}

	return values
}

func DescribeTask(cfg *internal.Config) *cobra.Command {
	var taskID string

	return &cobra.Command{
		Use:     "describe [task_id]",
		Short:   "Describe specified task",
		Aliases: []string{"edit"},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			svc, err := newService(cfg)
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}

			tasks, err := svc.Get()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}

			idTaker := func(task *model.Task) string { return task.ID }
			return autoCompleteArgs(tasks, idTaker), cobra.ShellCompDirectiveNoFileComp
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				taskID = args[0]
			}

			if taskID == "" {
				return fmt.Errorf("task id is required")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			svc, err := newService(cfg)
			if err != nil {
				return err
			}

			editor := internal.NewEditor("Name", "Description")
			task, err := svc.GetByID(taskID)
			if err != nil {
				return err
			}

			modified, err := editor.Edit([]string{task.Name, task.Description})
			if err != nil {
				return err
			}

			task.Name = modified[0]
			task.Description = modified[1]

			return svc.Update(task)
		},
		SilenceUsage: true,
	}
}
