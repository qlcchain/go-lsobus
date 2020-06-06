package commands

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

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
	addSonataOfferCmd(sonataCmd)
}

var sonataCmd = &cobra.Command{
	Use:   "sonata",
	Short: "sonata client",
	Long:  `sonata client`,
}

func addFlagsForOrderParams(cmd *cobra.Command) {
	cmd.Flags().Bool("fakeMode", false, "Fake mode")
	cmd.Flags().String("apiToken", "", "API token")

	// Common
	cmd.Flags().String("orderActivity", "install", "Type of order, (e.g., install, change, disconnect)")
	cmd.Flags().String("itemAction", "add", "Type of product action, (e.g., add, change, remove)")
	cmd.Flags().String("prodSpecID", "", "Production specification ID")
	cmd.Flags().String("prodOfferID", "", "Production offering ID")
	cmd.Flags().String("productID", "", "Product ID of existing service")

	cmd.Flags().Int64("startTime", 0, "Start time, (unix seconds)")
	cmd.Flags().Int64("endTime", 0, "End time, (unix seconds)")

	// UNI
	cmd.Flags().String("srcUniName", "", "Name of source port")
	cmd.Flags().String("srcSiteID", "", "Source Port geographic site ID")
	cmd.Flags().Uint("srcPortSpeed", 1000, "Source Port speed, Unit is Mbps")
	cmd.Flags().String("dstUniName", "", "Name of source port")
	cmd.Flags().String("dstSiteID", "", "Destination Port geographic site ID")
	cmd.Flags().Uint("dstPortSpeed", 1000, "Destination Port speed, Unit is Mbps")

	// Existing UNI for ELine
	cmd.Flags().String("srcUniID", "", "Source UNI ID of connection")
	cmd.Flags().UintSlice("srcVlanID", nil, "Source CE VLAN IDs of UNI")
	cmd.Flags().String("dstUniID", "", "Destination UNI ID of connection")
	cmd.Flags().UintSlice("dstVlanID", nil, "Destination CE VLAN IDs of UNI")

	cmd.Flags().String("srcLocationID", "", "Source location ID of connection")
	cmd.Flags().String("dstLocationID", "", "Destination location ID of connection")

	// ELine
	cmd.Flags().String("name", "", "Name of service")
	cmd.Flags().Uint("bandwidth", 0, "Bandwidth of connection, Unit is Mbps")
	cmd.Flags().String("cosName", "", "class of service name")
	cmd.Flags().Uint("sVlanID", 0, "Service VLAN ID of connection")

	// Price
	cmd.Flags().String("quoteID", "", "Quote ID")
	cmd.Flags().String("quoteItemID", "", "Quote Item ID")
	cmd.Flags().String("currency", "USD", "Currency, (e.g., USD, HKD, CNY)")
	cmd.Flags().Float32("price", 0, "price")
}

func addFlagsForFindParams(cmd *cobra.Command) {
	cmd.Flags().Bool("fakeMode", false, "Fake mode")
	cmd.Flags().String("apiToken", "", "API token")

	cmd.Flags().String("projectID", "", "Project ID")
	cmd.Flags().String("state", "", "Service state or status")
	cmd.Flags().String("externalID", "", "External ID")

	cmd.Flags().String("prodSpecID", "", "Production specification ID")
	cmd.Flags().String("prodOrderID", "", "Production order ID")
}

func addFlagsForGetParams(cmd *cobra.Command) {
	cmd.Flags().Bool("fakeMode", false, "Fake mode")
	cmd.Flags().String("apiToken", "", "API token")

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

	params.Buyer = &orchestra.PartnerParams{ID: "CBC", Name: "CBC"}
	params.Seller = &orchestra.PartnerParams{ID: "PCCW", Name: "PCCW"}

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

		uniItem.Name, err = cmd.Flags().GetString("srcUniName")
		if err != nil {
			return err
		}

		uniItem.PortSpeed, err = cmd.Flags().GetUint("srcPortSpeed")
		if err != nil {
			return err
		}

		uniItem.QuoteID, err = cmd.Flags().GetString("quoteID")
		if err != nil {
			return err
		}

		uniItem.QuoteItemID, err = cmd.Flags().GetString("quoteItemID")
		if err != nil {
			return err
		}

		uniItem.ProdOfferID, err = cmd.Flags().GetString("prodOfferID")
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

		uniItem.Name, err = cmd.Flags().GetString("dstUniName")
		if err != nil {
			return err
		}

		uniItem.PortSpeed, err = cmd.Flags().GetUint("dstPortSpeed")
		if err != nil {
			return err
		}

		uniItem.QuoteID, err = cmd.Flags().GetString("quoteID")
		if err != nil {
			return err
		}

		uniItem.QuoteItemID, err = cmd.Flags().GetString("quoteItemID")
		if err != nil {
			return err
		}

		uniItem.ProdOfferID, err = cmd.Flags().GetString("prodOfferID")
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

		lineItem.Name, err = cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		lineItem.SrcLocationID, err = cmd.Flags().GetString("srcLocationID")
		if err != nil {
			return err
		}

		lineItem.DstLocationID, err = cmd.Flags().GetString("dstLocationID")
		if err != nil {
			return err
		}

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

		lineItem.QuoteID, err = cmd.Flags().GetString("quoteID")
		if err != nil {
			return err
		}

		lineItem.QuoteItemID, err = cmd.Flags().GetString("quoteItemID")
		if err != nil {
			return err
		}

		lineItem.ProdOfferID, err = cmd.Flags().GetString("prodOfferID")
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

	params.BillingType = "DOM"
	//params.BillingType = "PAYG"
	params.PaymentType = "INVOICE"
	//params.PaymentType = "CREDITCARD"

	upOrdActStr := strings.ToUpper(params.OrderActivity)
	if upOrdActStr == "INSTALL" || upOrdActStr == "ADD" {
		params.ExternalID = uuid.New().String()
	}

	return nil
}

func fillBillingParamsByCmdFlags(cmd *cobra.Command) *orchestra.BillingParams {
	var err error

	params := &orchestra.BillingParams{}
	params.BillingType = "DOD"
	params.PaymentType = "invoice"

	params.Currency, err = cmd.Flags().GetString("currency")
	if err != nil {
		return nil
	}

	params.Price, err = cmd.Flags().GetFloat32("price")
	if err != nil {
		return nil
	}

	params.StartTime, err = cmd.Flags().GetInt64("startTime")
	if err != nil {
		return nil
	}
	if params.StartTime <= 0 {
		params.StartTime = time.Now().Unix()
	}

	params.EndTime, err = cmd.Flags().GetInt64("endTime")
	if err != nil {
		return nil
	}
	if params.EndTime <= params.StartTime {
		params.EndTime = params.StartTime + 24*3600
	}

	return params
}

func fillFindParamsByCmdFlags(params *orchestra.FindParams, cmd *cobra.Command) error {
	var err error

	params.Buyer = &orchestra.PartnerParams{ID: "CBC", Name: "CBC"}
	params.Seller = &orchestra.PartnerParams{ID: "PCCW", Name: "PCCW"}

	params.ProjectID, err = cmd.Flags().GetString("projectID")
	if err != nil {
		return err
	}

	params.State, err = cmd.Flags().GetString("state")
	if err != nil {
		return err
	}

	params.ExternalID, err = cmd.Flags().GetString("externalID")
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

	params.Buyer = &orchestra.PartnerParams{ID: "CBC", Name: "CBC"}
	params.Seller = &orchestra.PartnerParams{ID: "PCCW", Name: "PCCW"}

	params.ID, err = cmd.Flags().GetString("id")
	if err != nil {
		return err
	}

	if params.ID == "" {
		return errors.New("id can not be empty")
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

		apiToken, err := cmd.Flags().GetString("apiToken")
		if err == nil {
			o.SetApiToken("PCCW", apiToken)
		}
	}

	return o, nil
}
