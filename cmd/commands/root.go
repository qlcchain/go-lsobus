package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sonataCmd)
}

var rootCmd = &cobra.Command{
	Use:   "virtual-lsobus",
	Short: "virtual lsobus is a agent for MEF Sonata APIs",
	Long:  `virtual lsobus is a agent for MEF Sonata APIs`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
