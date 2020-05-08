package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/iixlabs/virtual-lsobus/orchestra"
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
		fmt.Println("running", cmd.Name())
		orchestra.SendSonataPOQCreateRequest()
	},
}
