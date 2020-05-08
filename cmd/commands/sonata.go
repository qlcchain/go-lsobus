package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	sonataCmd.AddCommand(sonataPoqCmd)
	sonataCmd.AddCommand(sonataQuoteCmd)
	sonataCmd.AddCommand(sonataOrderCmd)
}

var sonataCmd = &cobra.Command{
	Use:   "sonata",
	Short: "sonata client",
	Long:  `sonata client`,
}
