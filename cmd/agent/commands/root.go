package commands

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"time"
)

var rootCmd *cobra.Command
var lsobusUrl string

func InitCmd() {
	rootCmd = &cobra.Command{
		Use:   "cbc-agent",
		Short: "cbc-agent is a agent for go-lsobus",
		Long:  `cbc-agent is a agent for go-lsobus`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.PersistentFlags().StringVar(&lsobusUrl, "lsobusUrl", "http://127.0.0.1:9998", "http url of go-lsobus")

	connectionCmd := &cobra.Command{
		Use:   "connection",
		Short: "connection",
		Long:  `connection`,
	}
	rootCmd.AddCommand(connectionCmd)

	connCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "create new connection",
		Long:  `create new connection`,
		Run: func(cmd *cobra.Command, args []string) {
			pp := &ProductParam{}
			pp.BuyerName = "CBC"
			pp.BuyerAddr = "PCCWG"
			pp.SellerName = "PCCWG"
			pp.SellerAddr = "PCCWG"
			pp.ProductOfferID = "29f855fb-4760-4e77-877e-3318906ee4bc"

			pp.StartTime = time.Now().Unix()
			pp.EndTime = pp.StartTime + 5 * 24 * 3600
			pp.Bandwidth = 100
			pp.CosName = "gold"
			pp.SrcLocID = "5ae7e56bbbc9a8001231fa5d"
			pp.SrcPort = "5d098e7e96f045000a4164fa"
			pp.DstLocID = "5ae7e56bbbc9a8001231fa5d"
			pp.DstPort = "5d269f1760e409000ad83c58"
			pp.Name = "cbc-customer-1-line-1"

			po := &ProductOrder{Param: pp}
			err := po.Init()
			if err != nil {
				fmt.Println(err)
				return
			}

			err = po.CreateQuote()
			if err != nil {
				fmt.Println(err)
				return
			}

			err = po.CreateNewOrder()
			if err != nil {
				fmt.Println(err)
				return
			}

			err = po.GetOrderInfo()
			if err != nil {
				fmt.Println(err)
				return
			}
		},
	}
	connectionCmd.AddCommand(connCreateCmd)

	connChangeCmd := &cobra.Command{
		Use:   "change",
		Short: "change exist connection",
		Long:  `change exist connection`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	connectionCmd.AddCommand(connChangeCmd)

	connDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete exist connection",
		Long:  `delete exist connection`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	connectionCmd.AddCommand(connDeleteCmd)
}

func Execute(osArgs []string) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
