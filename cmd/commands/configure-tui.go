package commands

import (
	"errors"
	"github.com/charmbracelet/huh"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver/options"
	"github.com/fadyat/pump/pkg"
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

func newDriverSelectForm(driver *string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select driver").
				Options(
					huh.NewOption("Asana", "asana").Selected(true),
					huh.NewOption("File system", "local"),
				).
				Value(driver),
		),
	)
}

func ConfigureV2(config *internal.Config, configPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "configure-v2",
		Short: "Configure pump via tui",
		RunE: func(cmd *cobra.Command, args []string) error {
			var driver string

			if err := newDriverSelectForm(&driver).Run(); err != nil {
				return err
			}

			var (
				driverForm *huh.Form
				storedOpts = options.AsanaDriverFromMap(config.DriverOpts)
				opts       = &options.AsanaDriver{}
			)

			switch driver {
			case "asana":
				driverForm = newAsanaDriverForm(opts, storedOpts)
			default:
				panic(":*")
			}

			if err := driverForm.Run(); err != nil {
				return err
			}

			return pkg.WriteJson(configPath, internal.Config{
				Driver:     driver,
				DriverOpts: opts.Merge(storedOpts).ToMap(),
			})
		},
		SilenceUsage: true,
	}
}
