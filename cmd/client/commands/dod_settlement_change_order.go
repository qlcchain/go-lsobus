package commands

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/cmd/util"

	"google.golang.org/grpc"

	pb "github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/abiosoft/ishell"
)

func addDSChangeOrderCmdByShell(parentCmd *ishell.Cmd) {
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
	productId := util.Flag{
		Name:  "productId",
		Must:  true,
		Usage: "productId (separate by comma)",
		Value: "",
	}
	quoteId := util.Flag{
		Name:  "quoteId",
		Must:  true,
		Usage: "quoteId",
		Value: "",
	}

	args := []util.Flag{buyerAddress, buyerName, sellerAddress, sellerName, billingType, bandwidth, billingUnit, price,
		startTime, endTime, productId, quoteId}
	cmd := &ishell.Cmd{
		Name:                "changeOrder",
		Help:                "create a change order request",
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
			billingTypeP := util.StringVar(c.Args, billingType)
			bandwidthP := util.StringVar(c.Args, bandwidth)
			billingUnitP := util.StringVar(c.Args, billingUnit)
			priceP := util.StringVar(c.Args, price)
			startTimeP := util.StringVar(c.Args, startTime)
			endTimeP := util.StringVar(c.Args, endTime)
			productIdP := util.StringVar(c.Args, productId)
			quoteIdP := util.StringVar(c.Args, quoteId)

			if err := DSChangeOrder(buyerAddressP, buyerNameP, sellerAddressP, sellerNameP, startTimeP, endTimeP,
				billingTypeP, bandwidthP, billingUnitP, priceP, productIdP, quoteIdP); err != nil {
				util.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(cmd)
}

func DSChangeOrder(buyerAddressP, buyerNameP, sellerAddressP, sellerNameP, startTimeP, endTimeP, billingTypeP,
	bandwidthP, billingUnitP, priceP, productIdP, quoteIdP string) error {
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

	billingType, err := qlcSdk.ParseDoDSettleBillingType(billingTypeP)
	if err != nil {
		return err
	}

	price, err := strconv.ParseFloat(priceP, 64)
	if err != nil {
		return err
	}

	var startTime, endTime int64
	var billingUnit qlcSdk.DoDSettleBillingUnit

	if billingType == qlcSdk.DoDSettleBillingTypeDOD {
		startTime, err = strconv.ParseInt(startTimeP, 10, 64)
		if err != nil {
			return err
		}

		endTime, err = strconv.ParseInt(endTimeP, 10, 64)
		if err != nil {
			return err
		}
	} else {
		billingUnit, err = qlcSdk.ParseDoDSettleBillingUnit(billingUnitP)
		if err != nil {
			return err
		}
	}

	param := &pb.ChangeOrderParam{
		Buyer: &pb.User{
			Address: acc.Address().String(),
			Name:    buyerNameP,
		},
		Seller: &pb.User{
			Address: sellerAddress.String(),
			Name:    sellerNameP,
		},
		ChangeConnectionParam: make([]*pb.ChangeConnectionParam, 0),
	}

	pids := strings.Split(productIdP, ",")

	for _, productId := range pids {
		var conn *pb.ChangeConnectionParam

		if billingType == qlcSdk.DoDSettleBillingTypePAYG {
			conn = &pb.ChangeConnectionParam{
				DynamicParam: &pb.ConnectionDynamicParam{
					Bandwidth:   bandwidthP,
					BillingUnit: billingUnit.String(),
					Price:       float32(price),
				},
			}
		} else {
			conn = &pb.ChangeConnectionParam{
				DynamicParam: &pb.ConnectionDynamicParam{
					QuoteId:     quoteIdP,
					QuoteItemId: "1",
					Bandwidth:   bandwidthP,
					StartTime:   startTime,
					EndTime:     endTime,
					Price:       float32(price),
				},
			}
		}

		conn.ProductId = productId
		param.ChangeConnectionParam = append(param.ChangeConnectionParam, conn)
	}
	c := pb.NewOrderAPIClient(cn)
	internalId, err := c.ChangeOrder(context.Background(), param)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", internalId)

	return nil
}
