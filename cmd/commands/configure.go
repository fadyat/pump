package commands

import (
	"errors"
	"github.com/charmbracelet/huh"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
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

func newDriverSelectForm(d *string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select driver").
				Options(
					huh.NewOption("Asana", driver.AsanaDriver).Selected(true),
					huh.NewOption("File system", driver.FileSystemDriver),
				).
				Value(d),
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
	var opts = &options.FileSystemDriver{}

	opts.TasksFile = pkg.GetDir(config.ConfigPath) + "/tasks.json"
	return opts.ToMap(), nil
}

func runDriverOptionsSelection(config *internal.Config, d string) (map[string]any, error) {
	switch d {
	case driver.AsanaDriver:
		return selectAsanaDriverOptions(config)
	case driver.FileSystemDriver:
		return selectFileSystemDriverOptions(config)
	}

	return nil, errors.New("unsupported driver")
}

func Configure(config *internal.Config) *cobra.Command {
	var (
		backup     bool
		fromBackup bool
	)

	cmd := &cobra.Command{
		Use:     "configure",
		Short:   "Configure pump",
		Aliases: []string{"config", "conf", "cfg"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if fromBackup {
				return pkg.RestoreJson(config.ConfigPath)
			}

			var d string
			if err := newDriverSelectForm(&d).Run(); err != nil {
				return err
			}

			driverOptions, err := runDriverOptionsSelection(config, d)
			if err != nil {
				return err
			}

			if backup {
				if e := pkg.BackupJson(config.ConfigPath); e != nil {
					return e
				}
			}

			return pkg.WriteJson(config.ConfigPath, internal.Config{
				Driver:     d,
				DriverOpts: driverOptions,
			})
		},
		SilenceUsage: true,
	}

	cmd.Flags().BoolVar(&backup, "backup", true, "backup existing config file")
	cmd.Flags().BoolVar(&fromBackup, "from-backup", false, "read config from backup file")
	return cmd
}
