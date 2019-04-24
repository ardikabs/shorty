package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	customURL, password string
	reuse               bool
)

func init() {
	submitCmd.Flags().StringVarP(&customURL, "customurl", "c", "", "custom url")
	submitCmd.Flags().StringVarP(&password, "password", "p", "", "password for shortener url")
	submitCmd.Flags().BoolVarP(&reuse, "reuse", "r", false, "reuse flag")
}

var submitCmd = &cobra.Command{
	Use:   "submit [url]",
	Short: "Submit target URL to be shorten",
	Long:  "Submit target URL to be shorten on shortener provider",
	Args:  cobra.ExactArgs(1),
	RunE:  submitHandler,
}

func submitHandler(cmd *cobra.Command, args []string) error {
	url, err := api.SubmitURL(
		args[0],
		customURL,
		password,
		reuse,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Short URL: %s", url.ShortURL)
	return nil
}
