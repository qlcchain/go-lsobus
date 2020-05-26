package commands

import (
	"github.com/qlcchain/go-lsobus/orchestra"

	"github.com/spf13/cobra"
)

func addSonataQuoteCmd(parent *cobra.Command) {
	parent.AddCommand(sonataQuoteCmd)

	addFlagsForOrderParams(sonataQuoteCreateCmd)
	sonataQuoteCmd.AddCommand(sonataQuoteCreateCmd)

	addFlagsForFindParams(sonataQuoteFindCmd)
	sonataQuoteCmd.AddCommand(sonataQuoteFindCmd)

	addFlagsForGetParams(sonataQuoteGetCmd)
	sonataQuoteCmd.AddCommand(sonataQuoteGetCmd)
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

		o, err := getOrchestraInstance(cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		err = o.ExecQuoteCreate(op)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}

var sonataQuoteFindCmd = &cobra.Command{
	Use:   "find",
	Short: "retrieve product quoting list",
	Long:  `retrieve product quoting list`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		params := &orchestra.FindParams{}
		err = fillFindParamsByCmdFlags(params, cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		o, err := getOrchestraInstance(cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		err = o.ExecQuoteFind(params)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}

var sonataQuoteGetCmd = &cobra.Command{
	Use:   "get",
	Short: "retrieve product quoting item",
	Long:  `retrieve product quoting item`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		params := &orchestra.GetParams{}
		err = fillGetParamsByCmdFlags(params, cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		o, err := getOrchestraInstance(cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		err = o.ExecQuoteGet(params)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}
