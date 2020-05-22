package commands

import (
	"github.com/abiosoft/ishell"
)

func addOrderCmd() {
	if interactive {
		orderCmd := &ishell.Cmd{
			Name: "dod_settlement",
			Help: "dod_settlement commands",
			Func: func(c *ishell.Context) {
				c.Println(c.Cmd.HelpText())
			},
		}
		shell.AddCmd(orderCmd)
		addDSCreateOrderCmdByShell(orderCmd)
		addDSChangeOrderCmdByShell(orderCmd)
		addGetOrderInfoByShell(orderCmd)
		addDSTerminateOrderCmdByShell(orderCmd)
	}
}
