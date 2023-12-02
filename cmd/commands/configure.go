package commands

import (
	"bufio"
	"fmt"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/pkg"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func Configure(
	configPath string,
) *cobra.Command {
	reader := bufio.NewReader(os.Stdin)

	return &cobra.Command{
		Use:   "configure",
		Short: "Configure pump",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Select a driver:")
			fmt.Println("1. Asana")
			fmt.Println("2. Filesystem")

			choice, err := readInput(reader)
			if err != nil {
				return err
			}

			var opts map[string]any
			switch choice {
			case "1", "asana":
				choice = "asana"
				opts, err = asanaOpts(reader)
			case "2", "fs":
				choice = "fs"
				opts, err = fsOpts(reader)
			default:
				return fmt.Errorf("invalid choice")
			}

			if err != nil {
				return err
			}

			return pkg.WriteJson(configPath, internal.Config{
				Driver:     choice,
				DriverOpts: opts,
			})
		},
	}
}

func readInput(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func asanaOpts(
	reader *bufio.Reader,
) (map[string]any, error) {
	fmt.Println("Enter your Asana personal access token:")
	token, err := readInput(reader)
	if err != nil {
		return nil, err
	}

	fmt.Println("Enter the project ID: (can be found in the URL)")
	project, err := readInput(reader)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"token":   token,
		"project": project,
	}, nil
}

func fsOpts(
	reader *bufio.Reader,
) (map[string]any, error) {
	fmt.Println("Enter the path to the tasks file:")
	path, err := readInput(reader)
	if err != nil {
		return nil, err
	}

	return map[string]any{"file": path}, nil
}
