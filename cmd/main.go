package main

import (
	"github.com/fadyat/pump/cmd/commands"
	"github.com/fadyat/pump/internal"
	"github.com/fadyat/pump/internal/driver"
	"github.com/fadyat/pump/pkg"
	"github.com/spf13/cobra"
	"log"
)

var (

	// Version is changed via ldflags
	Version = "dev"
)

func getConfigPath() string {
	if Version == "dev" {
		return ".pump/config.json"
	}

	configPath, err := pkg.HomeDirConfig("config.json")
	if err != nil {
		log.Fatalf("failed to get home directory: %v", err)
	}

	return configPath
}

func withConfig() *internal.Config {
	config, err := internal.NewConfig(getConfigPath())
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	return config
}

func main() {
	config := withConfig()

	pump := &cobra.Command{
		Use:   "pump",
		Short: "Pump is a CLI application for choosing tasks to work on",
	}

	pump.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of Pump",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Pump version:", Version)
		},
	})

	manager := commands.NewManager(
		config,
		func() internal.IService {
			return internal.NewSvc(driver.New(
				config.Driver, config.GetDriverOpts(),
			))
		},
		pkg.RunCmd,
	)

	pump.AddCommand(commands.Configure(manager))
	pump.AddCommand(commands.CreateTask(manager))
	pump.AddCommand(commands.GetTask(manager))
	pump.AddCommand(commands.BrowseTask(manager))
	pump.AddCommand(commands.CompleteTask(manager))
	pump.AddCommand(commands.EditTask(manager))
	pump.AddCommand(commands.SelectTask(manager))

	_ = pump.Execute()
}
