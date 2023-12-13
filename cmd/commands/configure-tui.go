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
		),
		huh.NewGroup(
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

func newFileSystemDriverForm(opts, storedOpts *options.FileSystemDriver) *huh.Form {
	var (

		// todo: fix this to dir(configPath)/tasks.json
		tasksPathPlaceholder = "./.pump/tasks.json"
	)

	// todo: add form
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the path to the tasks file:").
				Prompt(defaultPrompt).
				Placeholder(tasksPathPlaceholder),
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
					huh.NewOption("File system", "fs"),
				).
				Value(driver),
		),
	)
}

func selectAsanaDriverOptions(config *internal.Config) (map[string]any, error) {
	var (
		storedOpts = options.AsanaDriverFromMap(config.DriverOpts)
		opts       = &options.AsanaDriver{}
	)

	if err := newAsanaDriverForm(opts, storedOpts).Run(); err != nil {
		return nil, err
	}

	return opts.Merge(storedOpts).ToMap(), nil
}

func selectFileSystemDriverOptions(config *internal.Config) (map[string]any, error) {
	var (
		storedOpts = options.FileSystemDriverFromMap(config.DriverOpts)
		opts       = &options.FileSystemDriver{}
	)

	if err := newFileSystemDriverForm(opts, storedOpts).Run(); err != nil {
		return nil, err
	}

	return opts.Merge(storedOpts).ToMap(), nil
}

func runDriverOptionsSelection(config *internal.Config, driver string) (map[string]any, error) {
	switch driver {
	case "asana":
		return selectAsanaDriverOptions(config)
	case "fs":
		return selectFileSystemDriverOptions(config)
	}

	return nil, errors.New("unsupported driver")
}

func ConfigureV2(config *internal.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "configure-v2",
		Short: "Configure pump via tui",
		RunE: func(cmd *cobra.Command, args []string) error {
			var driver string

			if err := newDriverSelectForm(&driver).Run(); err != nil {
				return err
			}

			driverOptions, err := runDriverOptionsSelection(config, driver)
			if err != nil {
				return err
			}

			return pkg.WriteJson(config.ConfigPath, internal.Config{
				Driver:     driver,
				DriverOpts: driverOptions,
			})
		},
		SilenceUsage: true,
	}
}
