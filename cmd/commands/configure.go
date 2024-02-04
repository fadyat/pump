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
	configureShort = "Configure Pump"

	configureExample = `  pump configure

  // backup previous config and create a new one
  pump configure --backup

  // restore from backup
  pump configure --from-backup`
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

func runDriverOptionsSelection(config *internal.Config, d string) (map[string]any, error) {
	if d == driver.AsanaDriver {
		return selectAsanaDriverOptions(config)
	}

	return nil, errors.New("unsupported driver")
}

func Configure(m *Manager) *cobra.Command {
	var (
		backup     bool
		fromBackup bool
	)

	cmd := &cobra.Command{
		Use:                   "configure",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Short:                 configureShort,
		Example:               configureExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if fromBackup {
				return pkg.RestoreJson(m.Config.ConfigPath)
			}

			var d string
			if err := newDriverSelectForm(&d).Run(); err != nil {
				return err
			}

			driverOptions, err := runDriverOptionsSelection(m.Config, d)
			if err != nil {
				return err
			}

			if backup {
				if e := pkg.BackupJson(m.Config.ConfigPath); e != nil {
					return e
				}
			}

			return pkg.WriteJson(m.Config.ConfigPath, internal.Config{
				Driver:     d,
				DriverOpts: driverOptions,
			})
		},
	}

	cmd.Flags().BoolVar(&backup, "backup", true, "backup existing config file")
	cmd.Flags().BoolVar(&fromBackup, "from-backup", false, "read config from backup file")
	return cmd
}
