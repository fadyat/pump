package main

import (
	"github.com/fadyat/pump/cmd/commands"
	"github.com/fadyat/pump/internal"
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
	pump.AddCommand(commands.GetAvailableTask(config))
	pump.AddCommand(commands.AddTask(config))
	pump.AddCommand(commands.MarkTaskAsDone(config))
	pump.AddCommand(commands.SelectTask(config))
	pump.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of Pump",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Pump version:", Version)
		},
	})

	_ = pump.Execute()
}
