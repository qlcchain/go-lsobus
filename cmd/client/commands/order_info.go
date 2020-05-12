package commands

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"google.golang.org/grpc"

	"github.com/iixlabs/virtual-lsobus/cmd/util"
)

func addGetOrderInfoByShell(parentCmd *ishell.Cmd) {
	c := &ishell.Cmd{
		Name: "info",
		Help: "order info",
		Func: func(c *ishell.Context) {
			if util.HelpText(c, nil) {
				return
			}
			err := util.CheckArgs(c, nil)
			if err != nil {
				util.Warn(err)
				return
			}
			if err := getOrderInfo(); err != nil {
				util.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(c)
}

func getOrderInfo() error {
	_, err := grpc.Dial(endpointP, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
