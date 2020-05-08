package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	sonataQuoteCmd.AddCommand(sonataQuoteCreateCmd)
}

var sonataQuoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "product quoting",
	Long:  `product quoting`,
}

var sonataQuoteCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create product quoting",
	Long:  `create product quoting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running", cmd.Name())
	},
}
