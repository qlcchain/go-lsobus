package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	sonataPoqCmd.AddCommand(sonataPoqCreateCmd)
}

var sonataPoqCmd = &cobra.Command{
	Use:   "poq",
	Short: "product offering qualification",
	Long:  `product offering qualification`,
}

var sonataPoqCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create product offering qualification",
	Long:  `create product offering qualification`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
