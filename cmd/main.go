package main

import (
	"fmt"
	"github.com/fadyat/pump/cmd/commands"
	"github.com/fadyat/pump/internal"
	"github.com/spf13/cobra"
)

var (
	// ConfigPath will be changed when building for production environment
	//  to ~/.config/pump/config.json
	ConfigPath = "./.pump/config-dev.json"
)

func main() {
	config, err := internal.NewConfig(ConfigPath)
	if err != nil {
		fmt.Println(err)
	}

	pump := &cobra.Command{
		Use:   "pump",
		Short: "Pump is a CLI application for choosing tasks to work on",
	}

	pump.AddCommand(commands.ConfigureV2(config))
	pump.AddCommand(commands.Configure(ConfigPath))
	pump.AddCommand(commands.GetAvailableTask(config))
	pump.AddCommand(commands.AddTask(config))
	pump.AddCommand(commands.MarkTaskAsDone(config))
	pump.AddCommand(commands.SelectTask(config))

	_ = pump.Execute()
}
