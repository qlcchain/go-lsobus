package commands

import (
	"github.com/abiosoft/ishell"
)

func addOrderCmd() {
	if interactive {
		orderCmd := &ishell.Cmd{
			Name: "dod",
			Help: "dod settlement commands",
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

func addMockOrderCmd() {
	if interactive {
		orderCmd := &ishell.Cmd{
			Name: "order",
			Help: "mock DoD order commands",
			Func: func(c *ishell.Context) {
				c.Println(c.Cmd.HelpText())
			},
		}
		shell.AddCmd(orderCmd)
		addMockOrderCmdByShell(orderCmd)
	}
}
