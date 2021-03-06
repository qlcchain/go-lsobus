package commands

import (
	"fmt"
	"os"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
	"github.com/spf13/cobra"
)

var (
	shell       *ishell.Shell
	rootCmd     *cobra.Command
	interactive bool
)

var (
	endpointP = "http://127.0.0.1:7777"
)

func Execute(osArgs []string) {
	interactive = isInteractive(osArgs)
	if interactive {
		shell = ishell.NewWithConfig(&readline.Config{
			Prompt:      fmt.Sprintf("%c[1;0;32m%s%c[0m", 0x1B, ">> ", 0x1B),
			HistoryFile: "/tmp/readline.tmp",
			//AutoComplete:      completer,
			InterruptPrompt:   "^C",
			EOFPrompt:         "exit",
			HistorySearchFold: true,
			//FuncFilterInputRune: filterInput,
		})
		shell.Println("LSOBUS Client")
		//set common variable
		addCommands()
		// run shell
		shell.Run()
	} else {
		rootCmd = &cobra.Command{
			Use:   "lsobus",
			Short: "lsobus is a agent for MEF Sonata APIs",
			Long:  `lsobus is a agent for MEF Sonata APIs`,
			Run: func(cmd *cobra.Command, args []string) {
			},
		}
		rootCmd.PersistentFlags().StringVarP(&endpointP, "endpoint", "e", endpointP, "DoD backend address for client")
		addCommands()
		//addSonataCmd(rootCmd)
		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func isInteractive(osArgs []string) bool {
	if len(osArgs) > 1 && osArgs[1] == "-i" {
		if len(osArgs) > 3 && osArgs[2] == "--endpoint" {
			endpointP = osArgs[3]
		}
		return true
	}
	if len(osArgs) > 2 && osArgs[1] == "--endpoint" {
		endpointP = osArgs[2]
		return true
	}
	return false
}

func addCommands() {
	version()
	addOrderCmd()
	addMockOrderCmd()
}
