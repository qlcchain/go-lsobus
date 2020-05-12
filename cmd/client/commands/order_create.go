package commands

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"google.golang.org/grpc"

	"github.com/iixlabs/virtual-lsobus/cmd/util"
)

func addCreateOrderCmdByShell(parentCmd *ishell.Cmd) {
	c := &ishell.Cmd{
		Name: "create",
		Help: "create order",
		Func: func(c *ishell.Context) {
			if util.HelpText(c, nil) {
				return
			}
			err := util.CheckArgs(c, nil)
			if err != nil {
				util.Warn(err)
				return
			}
			if err := createOrder(); err != nil {
				util.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(c)
}

func createOrder() error {
	_, err := grpc.Dial(endpointP, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
