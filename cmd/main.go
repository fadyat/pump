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
	Version = "dev"
)

func withConfig() *internal.Config {
	configPath, err := pkg.HomeDirConfig("config.json")
	if err != nil {
		log.Fatalf("failed to get home directory: %v", err)
	}

	config, err := internal.NewConfig(configPath)
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

	pump.AddCommand(commands.Configure(config))
	pump.AddCommand(commands.MarkTaskAsDone(config))
	pump.AddCommand(commands.SelectTask(config))
	pump.AddCommand(commands.BrowseTask(config))
	pump.AddCommand(commands.DescribeTask(config))
	pump.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of Pump",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Pump version:", Version)
		},
	})

	manager := commands.NewManager(config, func() internal.IService {
		// fixme: this will be fixed after switching to v2
		d, _ := driver.New(config.Driver, config.GetDriverOpts())
		return internal.NewSvc(d)
	})

	pump.AddCommand(commands.CreateTaskV2(manager))
	pump.AddCommand(commands.GetTaskV2(manager))

	_ = pump.Execute()
}
