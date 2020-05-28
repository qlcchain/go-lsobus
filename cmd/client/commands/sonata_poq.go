package commands

import (
	"github.com/spf13/cobra"

	"github.com/qlcchain/go-lsobus/orchestra"
)

func addSonataPoqCmd(parent *cobra.Command) {
	parent.AddCommand(sonataPoqCmd)

	addFlagsForOrderParams(sonataPoqCreateCmd)
	sonataPoqCmd.AddCommand(sonataPoqCreateCmd)

	addFlagsForFindParams(sonataPoqFindCmd)
	sonataPoqCmd.AddCommand(sonataPoqFindCmd)
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

		o, err := getOrchestraInstance(cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		err = o.ExecPOQCreate(op)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}

var sonataPoqFindCmd = &cobra.Command{
	Use:   "find",
	Short: "retrieve offering qualification list",
	Long:  `retrieve offering qualification list`,
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

		err = o.ExecPOQFind(params)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	},
}
