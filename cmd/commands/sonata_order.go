package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	sonataOrderCmd.AddCommand(sonataOrderCreateCmd)
}

var sonataOrderCmd = &cobra.Command{
	Use:   "order",
	Short: "product ordering",
	Long:  `product ordering`,
}

var sonataOrderCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create product ordering",
	Long:  `create product ordering`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
