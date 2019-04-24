package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {}

var deleteCmd = &cobra.Command{
	Use:   "delete [url]",
	Short: "Delete shorter URL",
	Long:  "Delete shorter URL from shortener provider",
	Args:  cobra.ExactArgs(1),
	RunE:  deleteHandler,
}

func deleteHandler(cmd *cobra.Command, args []string) error {
	err := api.DeleteURL(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Successfully deleted '%s'", args[0])
	return nil
}
