package main

import (
	"fmt"
	"github.com/fadyat/pump/cmd/commands"
	"github.com/fadyat/pump/internal"
	"github.com/spf13/cobra"
)

var (
	devConfigPath  string = "./.pump/config.json"
	prodConfigPath string = "~/.config/pump/config.json"
)

func main() {
	// todo: move to another file, because config command can't run
	config, err := internal.NewConfig(devConfigPath)
	if err != nil {
		fmt.Println(err)
	}

	pump := &cobra.Command{
		Use:   "pump",
		Short: "Pump is a CLI application for choosing tasks to work on",
	}

	pump.AddCommand(commands.Configure(devConfigPath))
	pump.AddCommand(commands.GetAvailableTask(config))
	pump.AddCommand(commands.AddTask(config))
	pump.AddCommand(commands.MarkTaskAsDone(config))

	pump.Execute()
}
