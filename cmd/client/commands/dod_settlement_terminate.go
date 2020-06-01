package commands

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	pb "github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/abiosoft/ishell"
	"github.com/qlcchain/go-qlc/cmd/util"
	"github.com/qlcchain/go-qlc/common/types"
)

func addDSTerminateOrderCmdByShell(parentCmd *ishell.Cmd) {
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
	productId := util.Flag{
		Name:  "productId",
		Must:  true,
		Usage: "productId (separate by comma)",
		Value: "",
	}
	price := util.Flag{
		Name:  "price",
		Must:  true,
		Usage: "price",
		Value: "",
	}

	args := []util.Flag{buyerAddress, buyerName, sellerAddress, sellerName, productId, price}
	cmd := &ishell.Cmd{
		Name:                "terminateOrder",
		Help:                "create a terminate order request",
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
			productIdP := util.StringVar(c.Args, productId)
			priceP := util.StringVar(c.Args, price)

			if err := DSTerminateOrder(buyerAddressP, buyerNameP, sellerAddressP, sellerNameP, productIdP, priceP); err != nil {
				util.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(cmd)
}

func DSTerminateOrder(buyerAddressP, buyerNameP, sellerAddressP, sellerNameP, productIdP string, priceP string) error {
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
	price, err := strconv.ParseFloat(priceP, 64)
	if err != nil {
		return err
	}
	param := &pb.TerminateOrderParam{
		Buyer: &pb.User{
			Address: acc.Address().String(),
			Name:    buyerNameP,
		},
		Seller: &pb.User{
			Address: sellerAddress.String(),
			Name:    sellerNameP,
		},
		TerminateConnectionParam: make([]*pb.TerminateConnectionParam, 0),
	}
	pids := strings.Split(productIdP, ",")
	for _, productId := range pids {
		var conn *pb.TerminateConnectionParam

		conn = &pb.TerminateConnectionParam{
			DynamicParam: &pb.ConnectionDynamicParam{
				Price:       float32(price),
				Currency:    "USD",
				QuoteId:     "1",
				QuoteItemId: "1",
			},
		}
		conn.ProductId = productId
		param.TerminateConnectionParam = append(param.TerminateConnectionParam, conn)
	}
	c := pb.NewOrderAPIClient(cn)
	internalId, err := c.TerminateOrder(context.Background(), param)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", internalId)

	return nil
}
