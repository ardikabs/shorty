package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ardikabs/shorty/kutt"
	"github.com/spf13/cobra"
)

var api kutt.API

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(deleteCmd)

	apiToken := os.Getenv("KUTT_TOKEN")
	if apiToken == "" {
		fmt.Println("Kutt API Token are not set. 'KUTT_TOKEN'")
		os.Exit(1)
	}

	api = kutt.API{
		BaseURL: &url.URL{
			Scheme: "https",
			Host:   "kutt.it",
		},
		APIToken:     apiToken,
		CustomDomain: os.Getenv("KUTT_CUSTOM_DOMAIN"),
	}
}

var rootCmd = &cobra.Command{
	Use:   "shorty",
	Short: "Shorty CLI. Use with calm",
}

// Execute will running CLI application and return nothing
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
