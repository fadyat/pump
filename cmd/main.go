package main

import (
	"github.com/fadyat/pump/cmd/commands"
	"github.com/fadyat/pump/internal"
	"github.com/spf13/cobra"
)

func main() {
	config := internal.NewConfig()
	pump := &cobra.Command{
		Use:   "pump",
		Short: "Pump is a CLI application for choosing tasks to work on",
	}

	pump.AddCommand(commands.GetAvailableTask(config))
	pump.AddCommand(commands.AddTask(config))
	pump.AddCommand(commands.MarkTaskAsDone(config))

	pump.Execute()
}
