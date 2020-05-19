package commands

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"

	"github.com/qlcchain/go-lsobus/cmd/util"
	version2 "github.com/qlcchain/go-lsobus/services/version"
)

func version() {
	if interactive {
		c := &ishell.Cmd{
			Name: "version",
			Help: "show version info for client",
			Func: func(c *ishell.Context) {
				if util.HelpText(c, nil) {
					return
				}
				if err := util.CheckArgs(c, nil); err != nil {
					util.Warn(err)
					return
				}
				versionInfo()
			},
		}
		shell.AddCmd(c)
	} else {
		var versionCmd = &cobra.Command{
			Use:   "version",
			Short: "show version info",
			Run: func(cmd *cobra.Command, args []string) {
				versionInfo()
			},
		}
		rootCmd.AddCommand(versionCmd)
	}
}

func versionInfo() {
	v := version2.VersionString()
	if interactive {
		util.Info(v)
	} else {
		fmt.Println(v)
	}
}
