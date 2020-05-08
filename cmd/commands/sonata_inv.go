package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	sonataInvCmd.AddCommand(sonataInvFindCmd)
}

var sonataInvCmd = &cobra.Command{
	Use:   "inventory",
	Short: "product inventory",
	Long:  `product inventory`,
}

var sonataInvFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find product inventory",
	Long:  `find product inventory`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
