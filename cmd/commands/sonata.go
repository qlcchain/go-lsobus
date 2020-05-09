package commands

import (
	"github.com/spf13/cobra"

	"github.com/iixlabs/virtual-lsobus/orchestra"
)

func init() {
	sonataCmd.AddCommand(sonataSiteCmd)
	sonataCmd.AddCommand(sonataPoqCmd)
	sonataCmd.AddCommand(sonataQuoteCmd)
	sonataCmd.AddCommand(sonataOrderCmd)
	sonataCmd.AddCommand(sonataInvCmd)
}

var sonataCmd = &cobra.Command{
	Use:   "sonata",
	Short: "sonata client",
	Long:  `sonata client`,
}

func addFlagsForOrderParams(cmd *cobra.Command) {
	// UNI
	//cmd.Flags().String("SiteID", "", "ID of geographic site")
	//cmd.Flags().Uint("portSpeed", 1000, "Speed of port, Unit is Mbps")

	// Existing UNI for ELine
	cmd.Flags().String("srcPortID", "", "Source port ID of connection")
	cmd.Flags().String("dstPortID", "", "Destination port ID of connection")

	// ELine
	cmd.Flags().Uint("bandwidth", 0, "Bandwidth of connection, Unit is Mbps")
	cmd.Flags().Uint("sVlanID", 0, "Service VLAN ID of connection, 1-4095")
	cmd.Flags().String("cosName", "", "class of service name")
}

func fillOrderParamsByCmdFlags(params *orchestra.OrderParams, cmd *cobra.Command) error {
	var err error
	/*
		params.SrcSiteID, err = cmd.Flags().GetString("siteID")
		if err != nil {
			return err
		}

		params.SrcPortSpeed, err = cmd.Flags().GetUint("portSpeed")
		if err != nil {
			return err
		}
	*/
	params.SrcPortID, err = cmd.Flags().GetString("srcPortID")
	if err != nil {
		return err
	}

	params.DstPortID, err = cmd.Flags().GetString("dstPortID")
	if err != nil {
		return err
	}

	params.Bandwidth, err = cmd.Flags().GetUint("bandwidth")
	if err != nil {
		return err
	}

	params.SVlanID, err = cmd.Flags().GetUint("sVlanID")
	if err != nil {
		return err
	}

	params.CosName, err = cmd.Flags().GetString("cosName")
	if err != nil {
		return err
	}

	return nil
}
