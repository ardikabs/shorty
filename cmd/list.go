package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List shorter URL",
	Long:  "List shorter URL from shortener provider",
	Args:  cobra.ExactArgs(0),
	RunE:  listHandler,
}

func listHandler(cmd *cobra.Command, args []string) error {
	urls, err := api.GetListURL()

	if err != nil {
		return err
	}

	if len(urls) < 1 {
		fmt.Println("No urls found")
		return nil
	}

	fmt.Printf("Found %d urls\n\n", len(urls))

	for _, url := range urls {
		fmt.Printf("URL ID: %s\n", url.ID)
		fmt.Printf("Target URL: %s\n", url.Target)
		fmt.Printf("Short URL: %s", url.ShortURL)
	}

	return nil
}
