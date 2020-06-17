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
	connCreateCmd.Flags().String("runEnv", "local", "running env, local/dev/stage")
	connCreateCmd.Flags().String("sellerAddr", "", "seller address of chain")
	connCreateCmd.Flags().String("buyerAddr", "", "buyer address of chain")
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
	connChangeCmd.Flags().String("runEnv", "local", "running env, local/dev/stage")
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
	connDeleteCmd.Flags().String("runEnv", "local", "running env, local/dev/stage")
	connDeleteCmd.Flags().String("internalId", "", "chain order internal id")
	connDeleteCmd.Flags().String("sellerOrderId", "", "seller order id")
	connectionCmd.AddCommand(connDeleteCmd)

	connGetCmd := &cobra.Command{
		Use:   "get",
		Short: "get connection info",
		Long:  `get connection info`,
		Run:   runConnectionGetCmd,
	}
	connGetCmd.Flags().String("runEnv", "local", "running env, local/dev/stage")
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

func fillProductCommonParam(cmd *cobra.Command, pp *ProductParam) error {
	var err error

	pp.RunEnv, err = cmd.Flags().GetString("runEnv")
	if pp.RunEnv == "" {
		pp.RunEnv = "local"
	}
	fmt.Println("RunEnv ", pp.RunEnv)

	if pp.RunEnv == "dev" {
		err = fillProductDevParam(cmd, pp)
	} else if pp.RunEnv == "stage" {
		err = fillProductStageParam(cmd, pp)
	} else {
		err = fillProductLocalParam(cmd, pp)
	}
	if err != nil {
		return nil
	}

	sellerAddr, _ := cmd.Flags().GetString("sellerAddr")
	if sellerAddr != "" {
		pp.SellerAddr = sellerAddr
	}

	buyerAddr, _ := cmd.Flags().GetString("buyerAddr")
	if buyerAddr != "" {
		pp.BuyerAddr = buyerAddr
	}

	return nil
}

func fillProductLocalParam(cmd *cobra.Command, pp *ProductParam) error {
	pp.SellerName = "PCCWG"
	pp.SellerAddr = "qlc_18yjtai4cwecsn3aasxx7gky6sprxdpkkcyjm9jxhynw5eq4p4ntm16shxmp"
	pp.BuyerName = "CBC"
	pp.BuyerAddr = "qlc_1gnqid9up5y998uwig44x1yfrppsdo8f9jfszgqin7pr7ixsyyae1y81w9xp"

	return nil
}

func fillProductDevParam(cmd *cobra.Command, pp *ProductParam) error {
	pp.SellerName = "PCCWG"
	pp.SellerAddr = "qlc_3bkys6wonkij7zfkti1it3aw88anbn6636x5hni4jkejc47kysa14ygkqsgh"
	pp.BuyerName = "CBC"
	pp.BuyerAddr = "qlc_1n7k9jru8teem514csgww713t6xiyya3jedcc5k6ex5zehkwj1c4yhygs1e3"

	return nil
}

func fillProductStageParam(cmd *cobra.Command, pp *ProductParam) error {
	pp.SellerName = "PCCWG"
	pp.SellerAddr = "stage-todo"
	pp.BuyerName = "CBC"
	pp.BuyerAddr = "stage-todo"

	return nil
}

func fillProductExistParam(existOrderInfo *models.ProtoOrderInfo, pp *ProductParam) error {
	if existOrderInfo.Seller != nil {
		pp.SellerAddr = existOrderInfo.Seller.Address
		pp.SellerName = existOrderInfo.Seller.Name
	}

	if existOrderInfo.Buyer != nil {
		pp.BuyerAddr = existOrderInfo.Buyer.Address
		pp.BuyerName = existOrderInfo.Buyer.Name
	}

	pp.ExistConnectionParam = existOrderInfo.Connections[0]

	pp.ExistProductID = pp.ExistConnectionParam.StaticParam.ProductID
	pp.BuyerProductID = pp.ExistConnectionParam.StaticParam.BuyerProductID
	pp.Name = pp.ExistConnectionParam.DynamicParam.ConnectionName
	pp.CosName = pp.ExistConnectionParam.DynamicParam.ServiceClass

	return nil
}

func runConnectionCreateCmd(cmd *cobra.Command, args []string) {
	var err error

	pp := &ProductParam{}

	err = fillProductCommonParam(cmd, pp)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	err = fillProductCommonParam(cmd, pp)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	err = fillProductExistParam(existOrderInfo, pp)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	err = fillProductCommonParam(cmd, pp)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	err = fillProductExistParam(existOrderInfo, pp)
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

	err := fillProductCommonParam(cmd, pp)
	if err != nil {
		fmt.Println(err)
		return
	}

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
