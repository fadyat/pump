package commands

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

var (
	ErrTaskIDRequired = errors.New("task id is required")
)

const (
	completeShort = "Complete a task"

	completeExample = `  pump complete 1234567890
  pump complete 1234567890 "this is a summary"
  pump complete 1234567890 "this is a summary" --reopen`
)

type CompleteFlags struct {
	TaskID  string
	Summary string
	Reopen  bool
}

func NewCompleteFlags() *CompleteFlags {
	return &CompleteFlags{
		Reopen: false,
	}
}

func (f *CompleteFlags) Override(args []string) {
	if len(args) > 0 {
		f.TaskID = args[0]
	}

	if len(args) > 1 {
		f.Summary = args[1]
	}
}

func (f *CompleteFlags) Prepare() {
	f.TaskID = strings.TrimSpace(f.TaskID)
	f.Summary = strings.TrimSpace(f.Summary)
}

func (f *CompleteFlags) Validate() error {
	f.Prepare()

	if f.TaskID == "" {
		return ErrTaskIDRequired
	}

	return nil
}

func CompleteTask(m *Manager) *cobra.Command {
	var flags = NewCompleteFlags()

	cmd := &cobra.Command{
		Use:                   "complete (-t TASK_ID [SUMMARY])",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Short:                 completeShort,
		Example:               completeExample,
		Args:                  cobra.MaximumNArgs(2),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !m.IsCommandAvailable("complete") {
				return ErrCommandNotAvailable
			}

			flags.Override(args)
			return flags.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			svc := m.ServiceMaker()
			if flags.Reopen {
				return svc.Reopen(flags.TaskID, flags.Summary)
			}

			return svc.MarkAsDone(flags.TaskID, flags.Summary)
		},
	}

	cmd.Flags().StringVarP(&flags.TaskID, "task-id", "t", flags.TaskID, "Task ID to complete")
	cmd.Flags().StringVarP(&flags.Summary, "summary", "s", flags.Summary, "Summary of the completed task")
	cmd.Flags().BoolVarP(&flags.Reopen, "reopen", "r", flags.Reopen, "Reopen a completed task")

	return cmd
}
