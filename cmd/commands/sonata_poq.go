package commands

import (
	"github.com/spf13/cobra"

	"github.com/iixlabs/virtual-lsobus/orchestra"
)

func init() {
	addFlagsForOrderParams(sonataPoqCreateCmd)
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
		var err error
		op := &orchestra.OrderParams{}
		err = fillOrderParamsByCmdFlags(op, cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		o := orchestra.NewOrchestra()
		err = o.Init()
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		err = o.ExecPOQ(op)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}
