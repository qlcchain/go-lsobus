package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	sonataSiteCmd.AddCommand(sonataSiteFindCmd)
}

var sonataSiteCmd = &cobra.Command{
	Use:   "site",
	Short: "geographic site",
	Long:  `geographic site`,
}

var sonataSiteFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find geographic site",
	Long:  `find geographic site`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
