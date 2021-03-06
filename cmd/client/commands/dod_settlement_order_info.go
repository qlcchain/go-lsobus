package commands

import (
	"context"
	"fmt"

	pb "github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/abiosoft/ishell"
	"google.golang.org/grpc"

	"github.com/qlcchain/go-lsobus/cmd/util"
)

func addGetOrderInfoByShell(parentCmd *ishell.Cmd) {
	internalId := util.Flag{
		Name:  "internalId",
		Must:  true,
		Usage: "order id on the chain",
		Value: "",
	}
	args := []util.Flag{internalId}
	c := &ishell.Cmd{
		Name:                "orderInfo",
		Help:                "order info",
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
			internalIdP := util.StringVar(c.Args, internalId)
			if err := getOrderInfo(internalIdP); err != nil {
				util.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(c)
}

func getOrderInfo(internalId string) error {
	cn, err := grpc.Dial(endpointP, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer cn.Close()
	id := &pb.GetOrderInfoParam{
		InternalId: internalId,
	}
	c := pb.NewOrderAPIClient(cn)
	orderInfo, err := c.GetOrderInfo(context.Background(), id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("orderInfo[%s]\n", util.ToIndentString(orderInfo))
	return nil
}
