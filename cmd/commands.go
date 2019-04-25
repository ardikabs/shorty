package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ardikabs/shorty/kutt"
	"github.com/spf13/cobra"
	"github.com/subosito/gotenv"
)

var api kutt.API

func init() {
	gotenv.Load()

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(deleteCmd)

	apiToken := os.Getenv("KUTT_TOKEN")
	customDomain := os.Getenv("KUTT_CUSTOM_DOMAIN")
	timeout, err := strconv.Atoi(os.Getenv("KUTT_TIMEOUT"))
	if err != nil {
		timeout = 5
	}

	api = kutt.API{
		BaseURL:      "https://kutt.it",
		Timeout:      time.Duration(timeout) * time.Second,
		Token:        apiToken,
		CustomDomain: customDomain,
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
