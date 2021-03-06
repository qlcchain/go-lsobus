package commands

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"google.golang.org/grpc"

	pb "github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/abiosoft/ishell"

	"github.com/qlcchain/go-lsobus/cmd/util"
)

func addDSCreateOrderCmdByShell(parentCmd *ishell.Cmd) {
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
		Value: "",
	}
	endTime := util.Flag{
		Name:  "endTime",
		Must:  false,
		Usage: "endTime",
		Value: "",
	}
	num := util.Flag{
		Name:  "num",
		Must:  true,
		Usage: "num",
		Value: "",
	}
	quoteId := util.Flag{
		Name:  "quoteId",
		Must:  true,
		Usage: "quoteId",
		Value: "",
	}

	args := []util.Flag{buyerAddress, buyerName, sellerAddress, sellerName, srcPort, dstPort, billingType, bandwidth,
		billingUnit, price, startTime, endTime, num, quoteId}
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
			numP := util.StringVar(c.Args, num)
			quoteIdP := util.StringVar(c.Args, quoteId)

			if err := DSCreateOrder(buyerAddressP, buyerNameP, sellerAddressP, sellerNameP, srcPortP, dstPortP,
				billingTypeP, bandwidthP, billingUnitP, priceP, startTimeP, endTimeP, numP, quoteIdP); err != nil {
				util.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(cmd)
}

func DSCreateOrder(buyerAddressP, buyerNameP, sellerAddressP, sellerNameP, srcPortP, dstPortP, billingTypeP,
	bandwidthP, billingUnitP, priceP, startTimeP, endTimeP, numP, quoteIdP string) error {
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

	acc := pkg.NewAccount(accBytes)
	if acc == nil {
		return fmt.Errorf("account format err")
	}

	sellerAddress, err := pkg.HexToAddress(sellerAddressP)
	if err != nil {
		return err
	}

	paymentType, err := qlcSdk.ParseDoDSettlePaymentType("invoice")
	if err != nil {
		return err
	}

	billingType, err := qlcSdk.ParseDoDSettleBillingType(billingTypeP)
	if err != nil {
		return err
	}

	var billingUnit qlcSdk.DoDSettleBillingUnit
	if len(billingUnitP) > 0 {
		billingUnit, err = qlcSdk.ParseDoDSettleBillingUnit(billingUnitP)
		if err != nil {
			return err
		}
	}

	price, err := strconv.ParseFloat(priceP, 64)
	if err != nil {
		return err
	}

	serviceClass, err := qlcSdk.ParseDoDSettleServiceClass("gold")
	if err != nil {
		return err
	}

	num, err := strconv.Atoi(numP)
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
		ConnectionParam: make([]*pb.ConnectionParam, 0),
	}

	var conn *pb.ConnectionParam
	for i := 0; i < num; i++ {
		quoteItemId := strconv.Itoa(1 + i)
		if billingType == qlcSdk.DoDSettleBillingTypePAYG {
			conn = &pb.ConnectionParam{
				StaticParam: &pb.ConnectionStaticParam{
					BuyerProductId: quoteItemId,
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
					ItemId:         quoteItemId,
					ConnectionName: fmt.Sprintf("connection%d", rand.Int()),
					QuoteId:        quoteIdP,
					QuoteItemId:    quoteItemId,
					Bandwidth:      bandwidthP,
					BillingUnit:    billingUnit.String(),
					Price:          float32(price),
					ServiceClass:   serviceClass.String(),
					PaymentType:    paymentType.String(),
					BillingType:    billingType.String(),
					Currency:       "USD",
				},
			}
		} else {
			startTime, err := strconv.ParseInt(startTimeP, 10, 64)
			if err != nil {
				return err
			}

			endTime, err := strconv.ParseInt(endTimeP, 10, 64)
			if err != nil {
				return err
			}

			conn = &pb.ConnectionParam{
				StaticParam: &pb.ConnectionStaticParam{
					BuyerProductId:    quoteItemId,
					ProductOfferingId: "29f855fb-4760-4e77-877e-3318906ee4bc",
					SrcCompanyName:    "CBC",
					SrcRegion:         "CHN",
					SrcCity:           "HK",
					SrcDataCenter:     "DCX",
					SrcPort:           srcPortP,
					DstCompanyName:    "CBC",
					DstRegion:         "USA",
					DstCity:           "NYC",
					DstDataCenter:     "DCY",
					DstPort:           dstPortP,
				},
				DynamicParam: &pb.ConnectionDynamicParam{
					ItemId:         quoteItemId,
					ConnectionName: fmt.Sprintf("connection%d", rand.Int()),
					QuoteId:        quoteItemId,
					QuoteItemId:    quoteItemId,
					Bandwidth:      bandwidthP,
					Price:          float32(price),
					ServiceClass:   serviceClass.String(),
					PaymentType:    paymentType.String(),
					BillingType:    billingType.String(),
					Currency:       "USD",
					StartTime:      startTime,
					EndTime:        endTime,
				},
			}
		}

		param.ConnectionParam = append(param.ConnectionParam, conn)
	}

	c := pb.NewOrderAPIClient(cn)
	internalId, err := c.CreateOrder(context.Background(), param)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", internalId)
	return nil
}
