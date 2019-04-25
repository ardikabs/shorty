package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List shorter URL",
	Long:  "List shorter URL from shortener provider",
	Args:  cobra.ExactArgs(0),
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	urls, err := api.GetListURL()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(urls) < 1 {
		fmt.Println("No urls found")
		os.Exit(0)
	}

	fmt.Printf("Found %d urls\n", len(urls))

	for _, url := range urls {
		fmt.Printf("\nURL ID: %s\n", url.ID)
		fmt.Printf("Target URL: %s\n", url.Target)
		fmt.Printf("Short URL: %s\n", url.ShortURL)
	}
}
