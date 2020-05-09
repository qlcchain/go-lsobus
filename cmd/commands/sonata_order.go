package commands

import (
	"github.com/spf13/cobra"

	"github.com/iixlabs/virtual-lsobus/orchestra"
)

func init() {
	addFlagsForOrderParams(sonataOrderCreateCmd)
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

		err = o.ExecOrder(op)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}
