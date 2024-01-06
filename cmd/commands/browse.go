package commands

import (
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/pkg"
	"github.com/spf13/cobra"
	"runtime"
)

func BrowseTask(cfg *internal.Config) *cobra.Command {
	var taskID string

	return &cobra.Command{
		Use:     "browse",
		Short:   "Browse specific task info in the browser",
		Aliases: []string{"info", "open", "view"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if cfg.Driver != "asana" {
				return fmt.Errorf("asana driver is only supported")
			}

			if len(args) == 0 {
				return fmt.Errorf("task id is required")
			}

			taskID = args[0]
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var asanaURL = fmt.Sprintf("https://app.asana.com/0/0/%s", taskID)
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
