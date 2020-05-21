package commands

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/abiosoft/ishell"
	"github.com/qlcchain/go-qlc/common/types"

	"github.com/qlcchain/go-lsobus/cmd/util"
)

func addCreateOrderCmdByShell(parentCmd *ishell.Cmd) {
	buyerAddress := util.Flag{
		Name:  "buyerAddress",
		Must:  true,
		Usage: "buyer's address hex string",
		Value: "",
	}
	buyerName := util.Flag{
		Name:  "buyerName",
		Must:  true,
		Usage: "buyer's name",
		Value: "",
	}
	sellerAddress := util.Flag{
		Name:  "sellerAddress",
		Must:  true,
		Usage: "seller's address",
		Value: "",
	}
	sellerName := util.Flag{
		Name:  "sellerName",
		Must:  true,
		Usage: "seller's name",
		Value: "",
	}
	srcPort := util.Flag{
		Name:  "srcPort",
		Must:  true,
		Usage: "source port",
		Value: "",
	}
	dstPort := util.Flag{
		Name:  "dstPort",
		Must:  true,
		Usage: "destination port",
		Value: "",
	}
	billingType := util.Flag{
		Name:  "billingType",
		Must:  true,
		Usage: "billing type (PAYG/DOD)",
		Value: "",
	}
	bandwidth := util.Flag{
		Name:  "bandwidth",
		Must:  true,
		Usage: "connection bandwidth (10 Mbps)",
		Value: "",
	}
	billingUnit := util.Flag{
		Name:  "billingUnit",
		Must:  false,
		Usage: "billing unit (year/month/week/day/hour/minute/second)",
		Value: "",
	}
	price := util.Flag{
		Name:  "price",
		Must:  true,
		Usage: "price",
		Value: "",
	}
	startTime := util.Flag{
		Name:  "startTime",
		Must:  false,
		Usage: "startTime",
		Value: "0",
	}
	endTime := util.Flag{
		Name:  "endTime",
		Must:  false,
		Usage: "endTime",
		Value: "0",
	}
	args := []util.Flag{buyerAddress, buyerName, sellerAddress, sellerName, srcPort, dstPort, billingType, bandwidth, billingUnit, price, startTime, endTime}
	cmd := &ishell.Cmd{
		Name:                "createOrder",
		Help:                "create a order request",
		CompleterWithPrefix: util.OptsCompleter(args),
		Func: func(c *ishell.Context) {
			if util.HelpText(c, args) {
				return
			}
			err := util.CheckArgs(c, args)
			if err != nil {
				util.Warn(err)
				return
			}

			buyerAddressP := util.StringVar(c.Args, buyerAddress)
			buyerNameP := util.StringVar(c.Args, buyerName)
			sellerAddressP := util.StringVar(c.Args, sellerAddress)
			sellerNameP := util.StringVar(c.Args, sellerName)
			srcPortP := util.StringVar(c.Args, srcPort)
			dstPortP := util.StringVar(c.Args, dstPort)
			billingTypeP := util.StringVar(c.Args, billingType)
			bandwidthP := util.StringVar(c.Args, bandwidth)
			billingUnitP := util.StringVar(c.Args, billingUnit)
			priceP := util.StringVar(c.Args, price)
			startTimeP := util.StringVar(c.Args, startTime)
			endTimeP := util.StringVar(c.Args, endTime)

			if err := DSCreateOrder(buyerAddressP, buyerNameP, sellerAddressP, sellerNameP, srcPortP, dstPortP,
				billingTypeP, bandwidthP, billingUnitP, priceP, startTimeP, endTimeP); err != nil {
				util.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(cmd)
}

func DSCreateOrder(buyerAddressP, buyerNameP, sellerAddressP, sellerNameP, srcPortP, dstPortP, billingTypeP,
	bandwidthP, billingUnitP, priceP, startTimeP, endTimeP string) error {
	cn, err := grpc.Dial(endpointP, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer cn.Close()
	accBytes, err := hex.DecodeString(buyerAddressP)
	if err != nil {
		return err
	}

	acc := types.NewAccount(accBytes)
	if acc == nil {
		return fmt.Errorf("account format err")
	}

	sellerAddress, err := types.HexToAddress(sellerAddressP)
	if err != nil {
		return err
	}
	s, err := strconv.ParseInt(startTimeP, 10, 64)
	if err != nil {
		return err
	}
	e, err := strconv.ParseInt(endTimeP, 10, 64)
	if err != nil {
		return err
	}
	price, err := strconv.ParseFloat(priceP, 64)
	if err != nil {
		return err
	}
	param := &pb.CreateOrderParam{
		Buyer: &pb.User{
			Address: acc.Address().String(),
			Name:    buyerNameP,
		},
		Seller: &pb.User{
			Address: sellerAddress.String(),
			Name:    sellerNameP,
		},
		Cps: make([]*pb.ConnectionParam, 0),
	}

	conn := &pb.ConnectionParam{
		StaticParam: &pb.ConnectionStaticParam{
			SrcCompanyName: "CBC",
			SrcRegion:      "CHN",
			SrcCity:        "HK",
			SrcDataCenter:  "DCX",
			SrcPort:        srcPortP,
			DstCompanyName: "CBC",
			DstRegion:      "USA",
			DstCity:        "NYC",
			DstDataCenter:  "DCY",
			DstPort:        dstPortP,
		},
		DynamicParam: &pb.ConnectionDynamicParam{
			ConnectionName: fmt.Sprintf("connection%d", rand.Int()),
			Bandwidth:      bandwidthP,
			BillingUnit:    billingUnitP,
			Price:          float32(price),
			ServiceClass:   "gold",
			PaymentType:    "invoice",
			BillingType:    billingTypeP,
			Currency:       "USD",
			StartTime:      s,
			EndTime:        e,
		},
	}
	param.Cps = append(param.Cps, conn)

	c := pb.NewOrderAPIClient(cn)
	internalId, err := c.CreateOrder(context.Background(), param)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", internalId)

	return nil
}
