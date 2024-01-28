package commands

import (
	"errors"
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/pkg"
	"github.com/spf13/cobra"
	"time"
)

const (
	WeekInterval = "week"
	DayInterval  = "day"
)

func timeNeeded(interval string) time.Duration {
	if interval == WeekInterval {
		return 7 * 24 * time.Hour
	}

	return 24 * time.Hour
}

func SelectTask(
	config *internal.Config,
) *cobra.Command {
	var (
		workInterval = DayInterval
		manualTaskID string
	)

	cmd := &cobra.Command{
		Use:     "select",
		Short:   "Select a task to work on",
		Aliases: []string{"todo", "do"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if workInterval != DayInterval && workInterval != WeekInterval {
				return errors.New("invalid work interval")
			}

			if len(args) > 0 {
				manualTaskID = args[0]
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			driv, err := driver.New(config.Driver, config.GetDriverOpts())
			if err != nil {
				return err
			}

			svc := internal.NewSvc(driv)
			task, err := svc.SelectGoal(
				manualTaskID,
				pkg.Ptr(pkg.Now().Add(timeNeeded(workInterval))),
			)
			if err != nil {
				return err
			}

			cmd.Println(
				fmt.Sprintf("Your task for the %s is:", workInterval),
				fmt.Sprintf("%q (%s)", task.Name, task.ID),
			)
			return nil
		},
		SilenceUsage: true,
	}

	cmd.Flags().StringVarP(&workInterval, "interval", "i", DayInterval, "work interval")
	return cmd
}
