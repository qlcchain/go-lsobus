package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/qlcchain/go-lsobus/orchestra"
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
	// Common
	cmd.Flags().String("orderActivity", "install", "Type of order, (e.g., install, change, disconnect)")
	cmd.Flags().String("itemAction", "add", "Type of product action, (e.g., add, change, remove)")
	cmd.Flags().String("prodSpecID", "", "Production specification ID")
	cmd.Flags().String("productID", "", "Product ID of existing service")

	// UNI
	cmd.Flags().String("srcSiteID", "", "Source Port geographic site ID")
	cmd.Flags().Uint("srcPortSpeed", 1000, "Source Port speed, Unit is Mbps")
	cmd.Flags().String("dstSiteID", "", "Destination Port geographic site ID")
	cmd.Flags().Uint("dstPortSpeed", 1000, "Destination Port speed, Unit is Mbps")

	// Existing UNI for ELine
	cmd.Flags().String("srcUniID", "", "Source UNI ID of connection")
	cmd.Flags().UintSlice("srcVlanID", nil, "Source CE VLAN IDs of UNI")
	cmd.Flags().String("dstUniID", "", "Destination UNI ID of connection")
	cmd.Flags().UintSlice("dstVlanID", nil, "Destination CE VLAN IDs of UNI")

	// ELine
	cmd.Flags().Uint("bandwidth", 0, "Bandwidth of connection, Unit is Mbps")
	cmd.Flags().String("cosName", "", "class of service name")
	cmd.Flags().Uint("sVlanID", 0, "Service VLAN ID of connection")

	// Price
	cmd.Flags().String("currency", "USA", "Currency, (e.g., USA, HKD, CNY)")
	cmd.Flags().Float32("price", 0, "price")
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

	params.ProdSpecID, err = cmd.Flags().GetString("prodSpecID")
	if err != nil {
		return err
	}

	params.OrderActivity, err = cmd.Flags().GetString("orderActivity")
	if err != nil {
		return err
	}

	params.ItemAction, err = cmd.Flags().GetString("itemAction")
	if err != nil {
		return err
	}

	params.ProductID, err = cmd.Flags().GetString("productID")
	if err != nil {
		return err
	}

	params.Buyer = &orchestra.Partner{ID: "C1B2C3", Name: "CBC"}
	params.Seller = &orchestra.Partner{ID: "P1C2C3W4", Name: "PCCW"}

	params.SrcSiteID, err = cmd.Flags().GetString("srcSiteID")
	if err != nil {
		return err
	}

	params.SrcPortSpeed, err = cmd.Flags().GetUint("srcPortSpeed")
	if err != nil {
		return err
	}

	params.DstSiteID, err = cmd.Flags().GetString("dstSiteID")
	if err != nil {
		return err
	}

	params.DstPortSpeed, err = cmd.Flags().GetUint("dstPortSpeed")
	if err != nil {
		return err
	}

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

	params.BillingParams = &orchestra.BillingParams{}
	params.BillingParams.BillingType = "DOD"
	params.BillingParams.Currency, err = cmd.Flags().GetString("currency")
	if err != nil {
		return err
	}

	params.BillingParams.Price, err = cmd.Flags().GetFloat32("price")
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
	o := orchestra.NewOrchestra("")
	err := o.Init()
	if err != nil {
		return nil, err
	}

	return o, nil
}
