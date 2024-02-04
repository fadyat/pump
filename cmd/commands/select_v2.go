package commands

import (
	"errors"
	"fmt"
	"github.com/fadyat/pump/pkg"
	"github.com/spf13/cobra"
	"slices"
	"strings"
	"time"
)

const (
	MonthInterval = "month"
	WeekInterval  = "week"
	DayInterval   = "day"
)

const (
	selectShort = "Select random task to work on"

	selectExample = `  pump select

  // for manual task selection
  pump select 1234567890
  pump select --interval month`
)

var (
	ErrInvalidWorkInterval = errors.New("invalid work interval")
)

func timeNeeded(interval string) time.Duration {
	switch interval {
	case MonthInterval:
		return 30 * 24 * time.Hour
	case WeekInterval:
		return 7 * 24 * time.Hour
	case DayInterval:
		return 24 * time.Hour
	}

	return 0
}

func getWorkIntervals() []string {
	return []string{MonthInterval, WeekInterval, DayInterval}
}

type SelectFlags struct {
	WorkInterval string
	TaskID       string
}

func NewSelectFlags() *SelectFlags {
	return &SelectFlags{
		WorkInterval: DayInterval,
	}
}

func (f *SelectFlags) Override(args []string) {
	if len(args) > 0 {
		f.TaskID = args[0]
	}
}

func (f *SelectFlags) Prepare() {
	f.TaskID = strings.TrimSpace(f.TaskID)
}

func (f *SelectFlags) Validate() error {
	f.Prepare()

	if slices.Contains(getWorkIntervals(), f.WorkInterval) {
		return ErrInvalidWorkInterval
	}

	return nil
}

func SelectTask(m *Manager) *cobra.Command {
	var flags = NewSelectFlags()

	cmd := &cobra.Command{
		Use:                   "select (-t TASK_ID -i INTERVAL)",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Short:                 selectShort,
		Example:               selectExample,
		Args:                  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !m.IsCommandAvailable("select") {
				return ErrCommandNotAvailable
			}

			flags.Override(args)
			return flags.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			task, err := m.ServiceMaker().SelectGoal(
				flags.TaskID,
				pkg.Ptr(pkg.Now().Add(timeNeeded(flags.WorkInterval))),
			)
			if err != nil {
				return err
			}

			cmd.Println(
				fmt.Sprintf("Your task for the %s is:", flags.WorkInterval),
				fmt.Sprintf("%q (%s)", task.Name, task.ID),
			)
			return nil
		},
	}

	cmd.Flags().StringVarP(&flags.TaskID, "task-id", "t", flags.TaskID, "Task ID to select")
	cmd.Flags().StringVarP(&flags.WorkInterval, "interval", "i", flags.WorkInterval, "Work interval")
	_ = cmd.RegisterFlagCompletionFunc(
		"interval",
		func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			return getWorkIntervals(), cobra.ShellCompDirectiveNoFileComp
		},
	)

	return cmd
}
