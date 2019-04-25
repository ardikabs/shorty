package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {}

var deleteCmd = &cobra.Command{
	Use:   "delete [shortUrl]",
	Short: "Delete shorter URL",
	Long:  "Delete shorter URL from shortener provider",
	Args:  cobra.ExactArgs(1),
	Run:   deleteHandler,
}

func deleteHandler(cmd *cobra.Command, args []string) {
	err := api.DeleteURL(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Successfully deleted '%s'", args[0])
}
