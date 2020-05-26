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
	cmd.Flags().Bool("fakeMode", false, "Fake mode")

	// Common
	cmd.Flags().String("orderActivity", "install", "Type of order, (e.g., install, change, disconnect)")
	cmd.Flags().String("itemAction", "add", "Type of product action, (e.g., add, change, remove)")
	cmd.Flags().String("prodSpecID", "", "Production specification ID")
	cmd.Flags().String("productID", "", "Product ID of existing service")

	cmd.Flags().String("durationUnit", "", "Duration unit, (e.g., YEAR, MONTH, DAY, HOUR)")
	cmd.Flags().Uint("durationAmount", 0, "Duration amount")

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
	cmd.Flags().String("currency", "USD", "Currency, (e.g., USD, HKD, CNY)")
	cmd.Flags().Float32("price", 0, "price")
}

func addFlagsForFindParams(cmd *cobra.Command) {
	cmd.Flags().Bool("fakeMode", false, "Fake mode")

	cmd.Flags().String("projectID", "", "Project ID")
	cmd.Flags().String("state", "", "Service state or status")
	cmd.Flags().String("prodSpecID", "", "Production specification ID")
	cmd.Flags().String("prodOrderID", "", "Production order ID")
}

func addFlagsForGetParams(cmd *cobra.Command) {
	cmd.Flags().Bool("fakeMode", false, "Fake mode")

	cmd.Flags().String("id", "", "ID of site/quote/order")
}

func fillOrderParamsByCmdFlags(params *orchestra.OrderParams, cmd *cobra.Command) error {
	var err error

	prodSpecID, err := cmd.Flags().GetString("prodSpecID")
	if err != nil {
		return err
	}

	params.OrderActivity, err = cmd.Flags().GetString("orderActivity")
	if err != nil {
		return err
	}

	params.Buyer = &orchestra.Partner{ID: "C1B2C3", Name: "CBC"}
	params.Seller = &orchestra.Partner{ID: "P1C2C3W4", Name: "PCCW"}

	itemAction, err := cmd.Flags().GetString("itemAction")
	if err != nil {
		return err
	}

	productID, err := cmd.Flags().GetString("productID")
	if err != nil {
		return err
	}

	// Create or Update UNI (Source)
	srcSiteID, err := cmd.Flags().GetString("srcSiteID")
	if err != nil {
		return err
	}
	if srcSiteID != "" {
		uniItem := &orchestra.UNIItemParams{}
		uniItem.Action = itemAction
		uniItem.ProductID = productID
		uniItem.SiteID = srcSiteID
		uniItem.PortSpeed, err = cmd.Flags().GetUint("srcPortSpeed")
		if err != nil {
			return err
		}

		uniItem.DurationUnit, err = cmd.Flags().GetString("durationUnit")
		if err != nil {
			return err
		}
		uniItem.DurationAmount, err = cmd.Flags().GetUint("durationAmount")
		if err != nil {
			return err
		}

		uniItem.BillingParams = fillBillingParamsByCmdFlags(cmd)

		params.UNIItems = append(params.UNIItems, uniItem)
	}

	// Create or Update UNI (Destination)
	dstSiteID, err := cmd.Flags().GetString("dstSiteID")
	if err != nil {
		return err
	}
	if dstSiteID != "" {
		uniItem := &orchestra.UNIItemParams{}
		uniItem.Action = itemAction
		uniItem.ProductID = productID
		uniItem.SiteID = dstSiteID
		uniItem.PortSpeed, err = cmd.Flags().GetUint("dstPortSpeed")
		if err != nil {
			return err
		}

		uniItem.DurationUnit, err = cmd.Flags().GetString("durationUnit")
		if err != nil {
			return err
		}
		uniItem.DurationAmount, err = cmd.Flags().GetUint("durationAmount")
		if err != nil {
			return err
		}

		uniItem.BillingParams = fillBillingParamsByCmdFlags(cmd)

		params.UNIItems = append(params.UNIItems, uniItem)
	}

	// Disconnect UNI
	if len(params.UNIItems) == 0 && prodSpecID == "UNISpec" {
		uniItem := &orchestra.UNIItemParams{}
		uniItem.Action = itemAction
		uniItem.ProductID = productID
		params.UNIItems = append(params.UNIItems, uniItem)
	}

	// Create or Update ELine
	lineBw, err := cmd.Flags().GetUint("bandwidth")
	if err != nil {
		return err
	}
	if lineBw > 0 {
		lineItem := &orchestra.ELineItemParams{}
		lineItem.Action = itemAction
		lineItem.ProductID = productID

		lineItem.SrcPortID, err = cmd.Flags().GetString("srcUniID")
		if err != nil {
			return err
		}

		lineItem.DstPortID, err = cmd.Flags().GetString("dstUniID")
		if err != nil {
			return err
		}

		lineItem.Bandwidth, err = cmd.Flags().GetUint("bandwidth")
		if err != nil {
			return err
		}
		lineItem.BwUnit = "Mbps"

		lineItem.CosName, err = cmd.Flags().GetString("cosName")
		if err != nil {
			return err
		}

		lineItem.SVlanID, err = cmd.Flags().GetUint("sVlanID")
		if err != nil {
			return err
		}

		lineItem.DurationUnit, err = cmd.Flags().GetString("durationUnit")
		if err != nil {
			return err
		}
		lineItem.DurationAmount, err = cmd.Flags().GetUint("durationAmount")
		if err != nil {
			return err
		}

		lineItem.BillingParams = fillBillingParamsByCmdFlags(cmd)

		params.ELineItems = append(params.ELineItems, lineItem)
	}

	// Disconnect ELine
	if len(params.ELineItems) == 0 && prodSpecID == "ELineSpec" {
		lineItem := &orchestra.ELineItemParams{}
		lineItem.Action = itemAction
		lineItem.ProductID = productID
		params.ELineItems = append(params.ELineItems, lineItem)
	}

	return nil
}

func fillBillingParamsByCmdFlags(cmd *cobra.Command) *orchestra.BillingParams {
	currency, err := cmd.Flags().GetString("currency")
	if err != nil {
		return nil
	}

	price, err := cmd.Flags().GetFloat32("price")
	if err != nil {
		return nil
	}

	if currency != "" && price > 0 {
		params := &orchestra.BillingParams{}
		params.BillingType = "DOD"
		params.Price = price
		params.Currency = currency

		return params
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

	params.ProductOrderID, err = cmd.Flags().GetString("prodOrderID")
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

func getOrchestraInstance(cmd *cobra.Command) (*orchestra.Orchestra, error) {
	o := orchestra.NewOrchestra("")
	err := o.Init()
	if err != nil {
		return nil, err
	}

	if cmd != nil {
		fakeMode, err := cmd.Flags().GetBool("fakeMode")
		if err == nil {
			o.SetFakeMode(fakeMode)
		}
	}

	return o, nil
}
