package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/spf13/cobra"
	"log/slog"
	"strings"
)

func MarkTaskAsDone(config *internal.Config) *cobra.Command {
	var (
		taskID  string
		summary string
		invert  bool
	)

	cmd := &cobra.Command{
		Use:     "done [task_id]",
		Short:   "Mark a task as done",
		Aliases: []string{"ok", "complete", "finish"},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("task id is required")
			}

			if config.Driver != driver.AsanaDriver && summary != "" {
				slog.Warn("summary is only supported by asana driver")
			}

			taskID = strings.TrimSpace(args[0])
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			driv, err := driver.New(config.Driver, config.GetDriverOpts())
			if err != nil {
				return err
			}

			var svc = internal.NewSvc(driv)
			var action = svc.MarkAsDone
			if invert {
				action = svc.Reopen
			}

			return action(taskID, summary)
		},
		SilenceUsage: true,
	}

	cmd.Flags().BoolVarP(&invert, "invert", "i", false, "Reopen a task instead of marking it as done")
	cmd.Flags().StringVarP(&summary, "summary", "s", "", "Summary of the task")
	return cmd
}
