package commands

import (
	"github.com/abiosoft/ishell"
)

func addOrderCmd() {
	if interactive {
		orderCmd := &ishell.Cmd{
			Name: "order",
			Help: "order commands",
			Func: func(c *ishell.Context) {
				c.Println(c.Cmd.HelpText())
			},
		}
		shell.AddCmd(orderCmd)
		addCreateOrderCmdByShell(orderCmd)
		addGetOrderInfoByShell(orderCmd)
	}
}
