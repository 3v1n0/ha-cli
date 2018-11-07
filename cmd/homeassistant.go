package cmd

import (
	"github.com/home-assistant/hassio-cli/cmd/homeassistant"
	"github.com/spf13/cobra"
)

// homeassistantCmd represents the homeassistant command
var homeassistantCmd = &cobra.Command{
	Use:     "homeassistant",
	Aliases: []string{"ha"},
	Run: func(cmd *cobra.Command, args []string) {
		homeassistant.Execute()
	},
}

func init() {
	rootCmd.AddCommand(homeassistantCmd)
}