package orchestra

import (
	"time"

	cmnmod "github.com/iixlabs/virtual-lsobus/sonata/common/models"

	"github.com/go-openapi/strfmt"

	ordcli "github.com/iixlabs/virtual-lsobus/sonata/order/client"
	ordapi "github.com/iixlabs/virtual-lsobus/sonata/order/client/product_order"
	ordmod "github.com/iixlabs/virtual-lsobus/sonata/order/models"
)

type sonataOrderImpl struct {
	sonataBaseImpl
}

func newSonataOrderImpl() *sonataOrderImpl {
	s := &sonataOrderImpl{}
	return s
}

func (s *sonataOrderImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataOrderImpl) SendCreateRequest(orderParams *OrderParams) error {
	reqParams := s.BuildCreateParams(orderParams)

	tranCfg := ordcli.DefaultTransportConfig().WithHost("localhost").WithSchemes([]string{"http"})
	httpCli := ordcli.NewHTTPClientWithConfig(nil, tranCfg)

	rspParams, err := httpCli.ProductOrder.ProductOrderCreate(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())

	//rspOrder := rspParams.GetPayload()

	return nil
}

func (s *sonataOrderImpl) SendFindRequest(params *FindParams) error {
	reqParams := ordapi.NewProductOrderFindParams()
	if params.ProjectID != "" {
		reqParams.ProjectID = &params.ProjectID
	}
	if params.BuyerID != "" {
		reqParams.BuyerID = &params.BuyerID
	}
	if params.State != "" {
		reqParams.State = &params.State
	}
	if params.Offset != "" {
		reqParams.Offset = &params.Offset
	}
	if params.Limit != "" {
		reqParams.Limit = &params.Limit
	}

	tranCfg := ordcli.DefaultTransportConfig().WithHost("localhost").WithSchemes([]string{"http"})
	httpCli := ordcli.NewHTTPClientWithConfig(nil, tranCfg)

	rspParams, err := httpCli.ProductOrder.ProductOrderFind(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())

	//rspOrder := rspParams.GetPayload()

	return nil
}

func (s *sonataOrderImpl) SendGetRequest(id string) error {
	reqParams := ordapi.NewProductOrderGetParams()
	reqParams.ProductOrderID = id

	tranCfg := ordcli.DefaultTransportConfig().WithHost("localhost").WithSchemes([]string{"http"})
	httpCli := ordcli.NewHTTPClientWithConfig(nil, tranCfg)

	rspParams, err := httpCli.ProductOrder.ProductOrderGet(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())

	//rspOrder := rspParams.GetPayload()

	return nil
}

func (s *sonataOrderImpl) BuildCreateParams(orderParams *OrderParams) *ordapi.ProductOrderCreateParams {
	reqParams := ordapi.NewProductOrderCreateParams()

	reqParams.ProductOrder = &ordmod.ProductOrderCreate{}

	reqParams.ProductOrder.ProjectID = ""
	reqParams.ProductOrder.ExternalID = nil

	reqParams.ProductOrder.BuyerRequestDate = &strfmt.DateTime{}
	reqParams.ProductOrder.BuyerRequestDate.Scan(time.Now())
	reqParams.ProductOrder.RequestedStartDate.Scan(time.Now().Add(time.Minute))
	reqParams.ProductOrder.RequestedCompletionDate = &strfmt.DateTime{}
	reqParams.ProductOrder.RequestedCompletionDate.Scan(time.Now().Add(time.Hour))
	reqParams.ProductOrder.DesiredResponse = ordmod.DesiredOrderResponsesConfirmationAndEngineeringDesign
	reqParams.ProductOrder.ExpeditePriority = true

	// Source UNI
	srcUniItem := s.BuildUNIItem(orderParams, true)
	if srcUniItem != nil {
		reqParams.ProductOrder.OrderItem = append(reqParams.ProductOrder.OrderItem, srcUniItem)
	}

	// Destination UNI
	dstUniItem := s.BuildUNIItem(orderParams, false)
	if dstUniItem != nil {
		reqParams.ProductOrder.OrderItem = append(reqParams.ProductOrder.OrderItem, dstUniItem)
	}

	// ELine
	lineItem := s.BuildELineItem(orderParams)
	if lineItem != nil {
		// Related Items
		if srcUniItem != nil {
			relType := string("RELIES_ON")
			relItem := &ordmod.OrderItemRelationShip{
				Type: &relType,
				ID:   srcUniItem.ID,
			}
			lineItem.OrderItemRelationship = append(lineItem.OrderItemRelationship, relItem)
		}

		if dstUniItem != nil {
			relType := string("RELIES_ON")
			relItem := &ordmod.OrderItemRelationShip{
				Type: &relType,
				ID:   dstUniItem.ID,
			}
			lineItem.OrderItemRelationship = append(lineItem.OrderItemRelationship, relItem)
		}

		// Related Products
		if orderParams.SrcPortID != "" {
			relType := string("RELIES_ON")
			relProd := &ordmod.ProductRelationship{Type: &relType}
			relProd.Product = &ordmod.ProductRef{}
			relProdID := orderParams.SrcPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}

		if orderParams.DstPortID != "" {
			relType := string("RELIES_ON")
			relProd := &ordmod.ProductRelationship{Type: &relType}
			relProd.Product = &ordmod.ProductRef{}
			relProdID := orderParams.DstPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}

		reqParams.ProductOrder.OrderItem = append(reqParams.ProductOrder.OrderItem, lineItem)
	}

	// Pricing
	reqParams.ProductOrder.PricingTerm = 12
	reqParams.ProductOrder.PricingMethod = ordmod.PricingMethodTariff

	// Billing
	reqParams.ProductOrder.BillingAccount = &ordmod.BillingAccountRef{}
	reqParams.ProductOrder.BillingAccount.BillingContact = &ordmod.Contact{}
	//reqParams.ProductOrder.BillingAccount.BillingContact.ContactName = ""

	// Party
	//reqParams.ProductOrder.RelatedParty = &ordmod.RelatedParty{}

	return reqParams
}

func (s *sonataOrderImpl) BuildUNIItem(orderParams *OrderParams, isDirSrc bool) *ordmod.ProductOrderItemCreate {
	var siteID string
	if isDirSrc {
		siteID = orderParams.SrcSiteID
	} else {
		siteID = orderParams.DstSiteID
	}
	if siteID == "" {
		return nil
	}

	uniItem := &ordmod.ProductOrderItemCreate{}

	uniItemID := s.NewItemID()
	uniItem.ID = &uniItemID
	uniItem.Action = ordmod.ProductActionTypeAdd

	uniOfferId := "LSO_Sonata_ProviderOnDemand_EthernetPort_UNI"
	uniItem.ProductOffering = &ordmod.ProductOfferingRef{ID: &uniOfferId}

	uniItem.Product = &ordmod.Product{}

	// UNI Place
	if siteID != "" {
		uniPlace := &ordmod.ReferencedAddress{}
		uniPlace.ReferenceID = &siteID
		uniItem.Product.SetPlace([]ordmod.RelatedPlaceReforValue{uniPlace})
	}

	// UNI Product Specification
	uniItem.Product.ProductSpecification = &ordmod.ProductSpecificationRef{}
	uniItem.Product.ProductSpecification.ID = "UNISpec"
	uniDesc := &cmnmod.UNIProductSpecification{}
	s.FillUNIProductSpec(uniDesc, orderParams)
	uniItem.Product.ProductSpecification.SetDescribing(uniDesc)

	// Price
	uniItem.PricingMethod = ordmod.PricingMethodTariff
	//uniItem.PricingTerm = 1

	// Party
	//uniItem.RelatedParty = &ordmod.RelatedParty{}

	return uniItem
}

func (s *sonataOrderImpl) BuildELineItem(orderParams *OrderParams) *ordmod.ProductOrderItemCreate {
	if orderParams.Bandwidth == 0 {
		return nil
	}

	lineItem := &ordmod.ProductOrderItemCreate{}
	lineItem.Action = ordmod.ProductActionTypeAdd

	lineItemID := s.NewItemID()
	lineItem.ID = &lineItemID

	linePoVal := "LSO_Sonata_ProviderOnDemand_EthernetConnection"
	lineItem.ProductOffering = &ordmod.ProductOfferingRef{ID: &linePoVal}

	lineItem.Product = &ordmod.Product{}

	//Product Specification
	lineItem.Product.ProductSpecification = &ordmod.ProductSpecificationRef{}
	lineItem.Product.ProductSpecification.ID = "ELineSpec"
	lineDesc := &cmnmod.ELineProductSpecification{}
	s.FillELineProductSpec(lineDesc, orderParams)
	lineItem.Product.ProductSpecification.SetDescribing(lineDesc)

	// Price
	lineItem.PricingMethod = ordmod.PricingMethodTariff
	//lineItem.PricingTerm = 1

	// Party
	//lineItem.RelatedParty = &ordmod.RelatedParty{}

	return lineItem
}
