package commands

import (
	"github.com/fadyat/pump/internal"
	"github.com/spf13/cobra"
	"strings"
)

const (
	editShort = "Edit a task name and description"

	editExample = `  pump edit 1234567890`
)

type EditFlags struct {
	TaskID string
}

func NewEditFlags() *EditFlags {
	return &EditFlags{}
}

func (f *EditFlags) Override(args []string) {
	if len(args) > 0 {
		f.TaskID = args[0]
	}
}

func (f *EditFlags) Prepare() {
	f.TaskID = strings.TrimSpace(f.TaskID)
}

func (f *EditFlags) Validate() error {
	f.Prepare()

	if f.TaskID == "" {
		return ErrTaskIDRequired
	}

	return nil
}

func EditTaskV2(m *Manager) *cobra.Command {
	var (
		flags  = NewEditFlags()
		editor = internal.NewEditor("Name", "Description")
	)

	cmd := &cobra.Command{
		Use:                   "edit_v2 (-t TASK_ID)",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Short:                 editShort,
		Example:               editExample,
		Args:                  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !m.IsCommandAvailable("edit") {
				return ErrCommandNotAvailable
			}

			flags.Override(args)
			return flags.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			service := m.ServiceMaker()

			task, err := service.GetByID(flags.TaskID)
			if err != nil {
				return err
			}

			modified, err := editor.Edit([]string{task.Name, task.Description})
			if err != nil {
				return err
			}

			task.Name = modified[0]
			task.Description = modified[1]
			return service.Update(task)
		},
	}

	cmd.Flags().StringVarP(&flags.TaskID, "task-id", "t", "", "Task ID")

	return cmd
}
