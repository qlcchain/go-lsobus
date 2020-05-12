package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/iixlabs/virtual-lsobus/orchestra"
)

func addSonataCmd(root *cobra.Command) {
	_ = os.Setenv("SWAGGER_DEBUG", "true")
	_ = os.Setenv("DEBUG", "true")

	root.AddCommand(sonataCmd)

	addSonataSiteCmd(sonataCmd)
	addSonataPoqCmd(sonataCmd)
	addSonataQuoteCmd(sonataCmd)
	addSonataOrderCmd(sonataCmd)
	addSonataInvCmd(sonataCmd)
}

var sonataCmd = &cobra.Command{
	Use:   "sonata",
	Short: "sonata client",
	Long:  `sonata client`,
}

func addFlagsForOrderParams(cmd *cobra.Command) {
	// UNI
	//cmd.Flags().String("siteID", "", "ID of geographic site")
	//cmd.Flags().Uint("portSpeed", 1000, "Speed of port, Unit is Mbps")

	// Existing UNI for ELine
	cmd.Flags().String("srcUniID", "", "Source UNI ID of connection")
	cmd.Flags().UintSlice("srcVlanID", nil, "Source CE VLAN IDs of UNI")
	cmd.Flags().String("dstUniID", "", "Destination UNI ID of connection")
	cmd.Flags().UintSlice("dstVlanID", nil, "Destination CE VLAN IDs of UNI")

	// ELine
	cmd.Flags().Uint("bandwidth", 0, "Bandwidth of connection, Unit is Mbps")
	cmd.Flags().String("cosName", "", "class of service name")
	cmd.Flags().Uint("sVlanID", 0, "Service VLAN ID of connection")
}

func addFlagsForFindParams(cmd *cobra.Command) {
	cmd.Flags().String("projectID", "", "Project ID")
	cmd.Flags().String("state", "", "Service state or status")
	cmd.Flags().String("prodSpecID", "", "Production specification ID")
}

func addFlagsForGetParams(cmd *cobra.Command) {
	cmd.Flags().String("id", "", "ID of site/quote/order")
}

func fillOrderParamsByCmdFlags(params *orchestra.OrderParams, cmd *cobra.Command) error {
	var err error

	params.Buyer = &orchestra.Partner{ID: "", Name: "CBC"}
	params.Seller = &orchestra.Partner{ID: "", Name: "PCCW"}

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

	params.SrcPortID, err = cmd.Flags().GetString("srcUniID")
	if err != nil {
		return err
	}

	params.SrcVlanID, err = cmd.Flags().GetUintSlice("srcVlanID")
	if err != nil {
		return err
	}

	params.DstPortID, err = cmd.Flags().GetString("dstUniID")
	if err != nil {
		return err
	}

	params.DstVlanID, err = cmd.Flags().GetUintSlice("dstVlanID")
	if err != nil {
		return err
	}

	params.Bandwidth, err = cmd.Flags().GetUint("bandwidth")
	if err != nil {
		return err
	}

	params.CosName, err = cmd.Flags().GetString("cosName")
	if err != nil {
		return err
	}

	params.SVlanID, err = cmd.Flags().GetUint("sVlanID")
	if err != nil {
		return err
	}

	return nil
}

func fillFindParamsByCmdFlags(params *orchestra.FindParams, cmd *cobra.Command) error {
	var err error

	params.ProjectID, err = cmd.Flags().GetString("projectID")
	if err != nil {
		return err
	}

	params.State, err = cmd.Flags().GetString("state")
	if err != nil {
		return err
	}

	params.ProductSpecificationID, err = cmd.Flags().GetString("prodSpecID")
	if err != nil {
		return err
	}

	return nil
}

func fillGetParamsByCmdFlags(params *orchestra.GetParams, cmd *cobra.Command) error {
	var err error
	params.ID, err = cmd.Flags().GetString("id")
	if err != nil {
		return err
	}

	return nil
}

func getOrchestraInstance() (*orchestra.Orchestra, error) {
	o := orchestra.NewOrchestra()
	err := o.Init()
	if err != nil {
		return nil, err
	}

	return o, nil
}
