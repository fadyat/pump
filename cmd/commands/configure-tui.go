package commands

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/fadyat/pump/internal/driver/options"
	"github.com/spf13/cobra"
)

const (
	defaultPrompt                    = "? "
	defaultAsanaTokenPlaceholder     = "asana-token"
	defaultAsanaProjectIdPlaceholder = "asana-project-id"
)

func getPlaceholder(stored, defaultValue string) string {
	if stored == "" {
		return defaultValue
	}

	return stored
}

func placeholderIsDefaultValue(stored, defaultValue string) bool {
	return stored == defaultValue
}

func newAsanaDriverForm(opts, storedOpts *options.AsanaDriver) *huh.Form {
	var (
		accessTokenPlaceholder = getPlaceholder(storedOpts.AccessToken, defaultAsanaTokenPlaceholder)
		projectIdPlaceholder   = getPlaceholder(storedOpts.ProjectID, defaultAsanaProjectIdPlaceholder)
	)

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter Asana access key:").
				Prompt(defaultPrompt).
				Placeholder(accessTokenPlaceholder).
				Value(&opts.AccessToken).
				Validate(func(s string) error {
					if s == "" && placeholderIsDefaultValue(accessTokenPlaceholder, defaultAsanaTokenPlaceholder) {
						return errors.New("asana token is required")
					}

					return nil
				}),
			huh.NewInput().
				Title("Enter Asana Project ID:").
				Prompt(defaultPrompt).
				Placeholder(projectIdPlaceholder).
				Value(&opts.ProjectID).
				Validate(func(s string) error {
					if s == "" && placeholderIsDefaultValue(projectIdPlaceholder, defaultAsanaProjectIdPlaceholder) {
						return errors.New("asana project id is required")
					}

					return nil
				}),
		),
	)
}

func ConfigureV2() *cobra.Command {
	return &cobra.Command{
		Use:   "configure-v2",
		Short: "Configure pump via tui",
		RunE: func(cmd *cobra.Command, args []string) error {
			var driver string

			driverSelectForm := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Select driver").
						Options(
							huh.NewOption("Asana", "asana").Selected(true),
							huh.NewOption("File system", "local"),
						).
						Value(&driver),
				),
			)

			err := driverSelectForm.Run()
			if err != nil {
				return err
			}

			var (
				driverForm *huh.Form

				// todo: read from file
				storedOpts = &options.AsanaDriver{}
				opts       = &options.AsanaDriver{}
			)

			switch driver {
			case "asana":
				driverForm = newAsanaDriverForm(opts, storedOpts)
			default:
				panic(":*")
			}

			err = driverForm.Run()
			if err != nil {
				return err
			}

			opts.Merge(storedOpts)
			fmt.Println(opts.ToMap())
			return nil
		},
	}
}
