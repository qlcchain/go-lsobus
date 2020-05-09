package commands

import (
	"github.com/iixlabs/virtual-lsobus/orchestra"

	"github.com/spf13/cobra"
)

func init() {
	addFlagsForOrderParams(sonataQuoteCreateCmd)
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

		err = o.ExecQuote(op)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}
