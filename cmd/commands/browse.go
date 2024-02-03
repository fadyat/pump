package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver/options"
	"github.com/fadyat/pump/pkg"
	"github.com/spf13/cobra"
	"runtime"
)

func BrowseTask(cfg *internal.Config) *cobra.Command {
	var taskID string

	return &cobra.Command{
		Use:     "browse",
		Short:   "Browse project/specified task",
		Aliases: []string{"info", "open", "view"},
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
			var opts = options.AsanaDriverFromMap(cfg.GetDriverOpts())
			asanaURL := fmt.Sprintf("https://app.asana.com/0/%s/%s", opts.ProjectID, taskID)
			return pkg.RunCmd(getBrowserCommand(), asanaURL)
		},
		SilenceUsage: true,
	}
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
