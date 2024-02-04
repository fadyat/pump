package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver/options"
	"github.com/spf13/cobra"
	"runtime"
	"strings"
)

const (
	browseShort = "Browse project/specified task"

	browseExample = `  pump browse 1234567890
  pump browse -t 1234567890
  pump browse`
)

type BrowseFlags struct {
	TaskID string
}

func NewBrowseFlags() *BrowseFlags {
	return &BrowseFlags{}
}

func (f *BrowseFlags) Override(args []string) {
	if len(args) > 0 {
		f.TaskID = args[0]
	}
}

func (f *BrowseFlags) Prepare() {
	f.TaskID = strings.TrimSpace(f.TaskID)
}

func (f *BrowseFlags) Validate() error {
	f.Prepare()

	// empty TaskID is valid; it means we want to
	// browse the whole project instead of a specific task
	return nil
}

func BrowseTask(m *Manager) *cobra.Command {
	var flags = NewBrowseFlags()

	cmd := &cobra.Command{
		Use:                   "browse (-t TASK_ID)",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Short:                 browseShort,
		Example:               browseExample,
		Args:                  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !m.IsCommandAvailable("browse") {
				return ErrCommandNotAvailable
			}

			flags.Override(args)
			return flags.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			url := buildURL(m.Config, flags.TaskID)
			return m.RunCommand(getBrowserCommand(), url)
		},
	}

	cmd.Flags().StringVarP(&flags.TaskID, "task-id", "t", flags.TaskID, "Task ID to browse")
	return cmd
}

func buildURL(config *internal.Config, taskID string) string {
	var opts = options.AsanaDriverFromMap(config.GetDriverOpts())
	return fmt.Sprintf("https://app.asana.com/0/%s/%s", opts.ProjectID, taskID)
}

func getBrowserCommand() string {
	switch runtime.GOOS {
	case "linux":
		return "xdg-open"
	case "windows":
		return "rundll32"
	default:
		return "open"
	}
}
