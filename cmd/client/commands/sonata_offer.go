package commands

import (
	"github.com/spf13/cobra"

	"github.com/qlcchain/go-lsobus/orchestra"
)

func addSonataOfferCmd(parent *cobra.Command) {
	parent.AddCommand(sonataOfferCmd)

	addFlagsForFindParams(sonataOfferFindCmd)
	sonataOfferCmd.AddCommand(sonataOfferFindCmd)

	addFlagsForGetParams(sonataOfferGetCmd)
	sonataOfferCmd.AddCommand(sonataOfferGetCmd)
}

var sonataOfferCmd = &cobra.Command{
	Use:   "offer",
	Short: "product offering",
	Long:  `product offering`,
}

var sonataOfferFindCmd = &cobra.Command{
	Use:   "find",
	Short: "retrieve product offering list",
	Long:  `retrieve product offering list`,
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

		err = o.ExecOfferFind(nil)
		if err != nil {
			cmd.PrintErrln(err)
		}
	},
}

var sonataOfferGetCmd = &cobra.Command{
	Use:   "get",
	Short: "retrieve product offering item",
	Long:  `retrieve product offering item`,
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

		err = o.ExecOfferGet(params)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}
