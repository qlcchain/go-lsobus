package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/qlcchain/go-lsobus/cmd/agent/models"

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
		Run:   runConnectionChangeCmd,
	}
	connChangeCmd.Flags().String("internalId", "", "create order internal id")
	connChangeCmd.Flags().String("sellerOrderId", "", "seller order id")
	connChangeCmd.Flags().Int32("bandwidth", 0, "bandwidth, unit is Mbps")
	connChangeCmd.Flags().Int32("duration", 1, "duration in days")
	connectionCmd.AddCommand(connChangeCmd)

	connDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete exist connection",
		Long:  `delete exist connection`,
		Run:   runConnectionDeleteCmd,
	}
	connDeleteCmd.Flags().String("internalId", "", "chain order internal id")
	connDeleteCmd.Flags().String("sellerOrderId", "", "seller order id")
	connectionCmd.AddCommand(connDeleteCmd)

	connGetCmd := &cobra.Command{
		Use:   "get",
		Short: "get connection info",
		Long:  `get connection info`,
		Run:   runConnectionGetCmd,
	}
	connGetCmd.Flags().String("internalId", "", "chain order internal id")
	connGetCmd.Flags().String("sellerOrderId", "", "seller order id")
	connectionCmd.AddCommand(connGetCmd)
}

func Execute(osArgs []string) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fillProductCommonParam(pp *ProductParam) {
	pp.SellerName = "PCCWG"
	pp.SellerAddr = "qlc_18yjtai4cwecsn3aasxx7gky6sprxdpkkcyjm9jxhynw5eq4p4ntm16shxmp"

	pp.BuyerName = "CBC"
	pp.BuyerAddr = "qlc_1gnqid9up5y998uwig44x1yfrppsdo8f9jfszgqin7pr7ixsyyae1y81w9xp"

	pp.ProductOfferID = "29f855fb-4760-4e77-877e-3318906ee4bc"

	pp.SrcLocID = "5ae7e56bbbc9a8001231fa5d"
	pp.SrcPort = "5d098e7e96f045000a4164fa"
	pp.DstLocID = "5ae7e56bbbc9a8001231fa5d"
	pp.DstPort = "5d269f1760e409000ad83c58"
}

func runConnectionCreateCmd(cmd *cobra.Command, args []string) {
	var err error

	pp := &ProductParam{}

	fillProductCommonParam(pp)

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

	pp.BuyerProductID = uuid.New().String()

	po := &ProductOrder{Param: pp}
	err = po.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CreateQuote("INSTALL")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CreateNewOrder()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CheckOrderStatus()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func runConnectionChangeCmd(cmd *cobra.Command, args []string) {
	var err error

	pp := &ProductParam{}

	fillProductCommonParam(pp)

	// find existing product by create order internal id
	internalId, err := cmd.Flags().GetString("internalId")
	if err != nil {
		fmt.Println(err)
		return
	}
	// find existing product by seller order id
	sellerOrderId, err := cmd.Flags().GetString("sellerOrderId")
	if err != nil {
		fmt.Println(err)
		return
	}

	existPo := &ProductOrder{}
	err = existPo.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	var existOrderInfo *models.ProtoOrderInfo

	if internalId != "" {
		existOrderInfo, err = existPo.GetOrderInfoByInternalId(internalId)
		if err != nil {
			fmt.Println("failed to GetOrderInfoByInternalId ", err)
			return
		}
	} else if sellerOrderId != "" {
		existOrderInfo, err = existPo.GetOrderInfoBySellerAndOrderId(pp.SellerAddr, sellerOrderId)
		if err != nil {
			fmt.Println("failed to GetOrderInfoBySellerAndOrderId ", err)
			return
		}
	} else {
		fmt.Println("invalid internalId/sellerOrderId param")
		return
	}

	if len(existOrderInfo.Connections) == 0 {
		fmt.Println("exist connection is empty")
		return
	}
	pp.ExistConnectionParam = existOrderInfo.Connections[0]

	pp.ExistProductID = pp.ExistConnectionParam.StaticParam.ProductID
	pp.BuyerProductID = pp.ExistConnectionParam.StaticParam.BuyerProductID
	pp.Name = pp.ExistConnectionParam.DynamicParam.ConnectionName
	pp.CosName = pp.ExistConnectionParam.DynamicParam.ServiceClass

	duration, err := cmd.Flags().GetInt32("duration")
	if err != nil {
		fmt.Println(err)
		return
	}
	existSTimeInt, err := strconv.Atoi(pp.ExistConnectionParam.DynamicParam.StartTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	existETimeInt, err := strconv.Atoi(pp.ExistConnectionParam.DynamicParam.EndTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	pp.StartTime = int64(existSTimeInt)
	if duration > 0 {
		pp.EndTime = pp.StartTime + int64(duration)*24*3600
	} else {
		pp.EndTime = int64(existETimeInt)
	}

	pp.Bandwidth, err = cmd.Flags().GetInt32("bandwidth")
	if err != nil {
		fmt.Println(err)
		return
	}

	po := &ProductOrder{Param: pp}
	err = po.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CreateQuote("CHANGE")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CreateChangeOrder()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CheckOrderStatus()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func runConnectionDeleteCmd(cmd *cobra.Command, args []string) {
	var err error

	pp := &ProductParam{}

	fillProductCommonParam(pp)

	// find existing product by create order internal id
	internalId, err := cmd.Flags().GetString("internalId")
	if err != nil {
		fmt.Println(err)
		return
	}
	// find existing product by seller order id
	sellerOrderId, err := cmd.Flags().GetString("sellerOrderId")
	if err != nil {
		fmt.Println(err)
		return
	}

	var existOrderInfo *models.ProtoOrderInfo

	existPo := &ProductOrder{}
	err = existPo.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	if internalId != "" {
		existOrderInfo, err = existPo.GetOrderInfoByInternalId(internalId)
		if err != nil {
			fmt.Println("failed to GetOrderInfoByInternalId ", err)
			return
		}
	} else if sellerOrderId != "" {
		existOrderInfo, err = existPo.GetOrderInfoBySellerAndOrderId(pp.SellerAddr, sellerOrderId)
		if err != nil {
			fmt.Println("failed to GetOrderInfoBySellerAndOrderId ", err)
			return
		}
	} else {
		fmt.Println("invalid internalId/sellerOrderId param")
		return
	}

	if len(existOrderInfo.Connections) == 0 {
		fmt.Println("exist connection is empty")
		return
	}
	pp.ExistConnectionParam = existOrderInfo.Connections[0]

	pp.ExistProductID = pp.ExistConnectionParam.StaticParam.ProductID
	pp.BuyerProductID = pp.ExistConnectionParam.StaticParam.BuyerProductID
	pp.Name = pp.ExistConnectionParam.DynamicParam.ConnectionName
	pp.CosName = pp.ExistConnectionParam.DynamicParam.ServiceClass

	po := &ProductOrder{Param: pp}
	err = po.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CreateQuote("DISCONNECT")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CreateTerminateOrder()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = po.CheckOrderStatus()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func runConnectionGetCmd(cmd *cobra.Command, args []string) {
	pp := &ProductParam{}
	fillProductCommonParam(pp)

	internalId, err := cmd.Flags().GetString("internalId")
	if err != nil {
		fmt.Println(err)
		return
	}

	sellerOrderId, err := cmd.Flags().GetString("sellerOrderId")
	if err != nil {
		fmt.Println(err)
		return
	}

	po := &ProductOrder{Param: pp}
	err = po.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	var orderInfo *models.ProtoOrderInfo

	if internalId != "" {
		orderInfo, err = po.GetOrderInfoByInternalId(internalId)
		if err != nil {
			fmt.Println("failed to GetOrderInfoByInternalId ", err)
			return
		}
	} else if sellerOrderId != "" {
		orderInfo, err = po.GetOrderInfoBySellerAndOrderId(pp.SellerAddr, sellerOrderId)
		if err != nil {
			fmt.Println("failed to GetOrderInfoBySellerAndOrderId ", err)
			return
		}
	} else {
		fmt.Println("invalid internalId/sellerOrderId param")
		return
	}

	infoData, err := json.MarshalIndent(orderInfo, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Get Order Info is OK, %s\n", string(infoData))
}
