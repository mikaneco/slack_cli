/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"slk/config"

	"github.com/spf13/cobra"
)

var (
	channel string
	token   string
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set configuration options",
	Long:  `Set configuration options for the CLI tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new config object
		c := config.New()

		// Set the channel and token
		c.SetChannel(channel)
		c.SetToken(token)

		// Save the configuration
		if err := c.Save(); err != nil {
			fmt.Printf("Error saving configuration: %s", err)
			return
		}
		fmt.Printf("Setting configuration options: channel=%s, token=%s\n", channel, token)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	// Define flags for the config command
	configCmd.Flags().StringVar(&channel, "channel", "c", "The channel to send messages to")
	configCmd.Flags().StringVar(&token, "token", "t", "The Slack API token")

	// Mark flags as required
	configCmd.MarkFlagRequired("channel")
	configCmd.MarkFlagRequired("token")
}
