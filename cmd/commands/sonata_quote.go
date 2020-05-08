package commands

import (
	"fmt"

	"github.com/iixlabs/virtual-lsobus/orchestra"

	"github.com/spf13/cobra"
)

func init() {
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
		o := orchestra.NewOrchestra()
		err := o.Init()
		if err != nil {
			fmt.Println(err)
		}
	},
}
