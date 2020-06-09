package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command
var lsobusUrl string
var flgDebug bool

func InitCmd() {
	rootCmd = &cobra.Command{
		Use:   "cbc-agent",
		Short: "cbc-agent is a agent for go-lsobus",
		Long:  `cbc-agent is a agent for go-lsobus`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.PersistentFlags().BoolVar(&flgDebug, "debug", false, "enable debug")
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
		Run:   runConnectionCreateCmd,
	}
	connCreateCmd.Flags().String("connName", "", "connect name")
	connCreateCmd.Flags().Int32("bandwidth", 0, "bandwidth, unit is Mbps")
	connCreateCmd.Flags().String("cosName", "gold", "service class")
	connCreateCmd.Flags().Int32("duration", 1, "duration in days")
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

	connGetCmd := &cobra.Command{
		Use:   "get",
		Short: "get connection info",
		Long:  `get connection info`,
		Run:   runConnectionGetCmd,
	}
	connGetCmd.Flags().String("internalId", "", "order internal id")
	connectionCmd.AddCommand(connGetCmd)
}

func Execute(osArgs []string) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runConnectionCreateCmd(cmd *cobra.Command, args []string) {
	var err error

	pp := &ProductParam{}

	pp.SellerName = "PCCWG"
	pp.SellerAddr = "qlc_18yjtai4cwecsn3aasxx7gky6sprxdpkkcyjm9jxhynw5eq4p4ntm16shxmp"

	pp.BuyerName = "CBC"
	pp.BuyerAddr = "qlc_1gnqid9up5y998uwig44x1yfrppsdo8f9jfszgqin7pr7ixsyyae1y81w9xp"

	pp.ProductOfferID = "29f855fb-4760-4e77-877e-3318906ee4bc"

	duration, err := cmd.Flags().GetInt32("duration")
	if err != nil {
		fmt.Println(err)
		return
	}
	pp.StartTime = time.Now().Unix()
	pp.EndTime = pp.StartTime + int64(duration)*24*3600

	pp.Name, err = cmd.Flags().GetString("connName")
	if err != nil {
		fmt.Println(err)
		return
	}

	pp.Bandwidth, err = cmd.Flags().GetInt32("bandwidth")
	if err != nil {
		fmt.Println(err)
		return
	}

	pp.CosName, err = cmd.Flags().GetString("cosName")
	if err != nil {
		fmt.Println(err)
		return
	}

	pp.SrcLocID = "5ae7e56bbbc9a8001231fa5d"
	pp.SrcPort = "5d098e7e96f045000a4164fa"
	pp.DstLocID = "5ae7e56bbbc9a8001231fa5d"
	pp.DstPort = "5d269f1760e409000ad83c58"
	pp.BuyerProductID = uuid.New().String()

	po := &ProductOrder{Param: pp}
	err = po.Init()
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

	for {
		err = po.GetOrderInfo()
		if err != nil {
			fmt.Println(err)
			time.Sleep(5 * time.Second)
		}
	}
}

func runConnectionGetCmd(cmd *cobra.Command, args []string) {
	internalId, err := cmd.Flags().GetString("internalId")
	if err != nil {
		fmt.Println(err)
		return
	}

	po := &ProductOrder{Param: &ProductParam{}}
	err = po.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.GetOrderInfoByInternalId(internalId)
	if err != nil {
		fmt.Println(err)
		return
	}
}
