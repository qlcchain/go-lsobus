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
	s.Version = MEFAPIVersionOrder
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
	reqParams.ProductOrder.OrderVersion = &s.Version

	reqParams.ProductOrder.ProjectID = orderParams.ProjectID
	if orderParams.ExternalID != "" {
		reqParams.ProductOrder.ExternalID = &orderParams.ExternalID
	}

	reqParams.ProductOrder.OrderActivity = ordmod.OrderActivityInstall
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

	// Billing
	s.BuildOrderBilling(reqParams.ProductOrder, orderParams)

	// Party
	s.BuildOrderRelatedParty(reqParams.ProductOrder, orderParams)

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
	s.BuildUNIProductSpec(uniDesc, orderParams)
	uniItem.Product.ProductSpecification.SetDescribing(uniDesc)

	// Price
	s.BuildItemPrice(uniItem, orderParams)

	// Billing
	s.BuildItemBilling(uniItem, orderParams)

	// Party
	s.BuildItemRelatedParty(uniItem, orderParams)

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
	s.BuildELineProductSpec(lineDesc, orderParams)
	lineItem.Product.ProductSpecification.SetDescribing(lineDesc)

	// Price
	s.BuildItemPrice(lineItem, orderParams)

	// Billing
	s.BuildItemBilling(lineItem, orderParams)

	// Party
	s.BuildItemRelatedParty(lineItem, orderParams)

	return lineItem
}

func (s *sonataOrderImpl) BuildItemPrice(item *ordmod.ProductOrderItemCreate, params *OrderParams) {
	// Price
	item.PricingMethod = ordmod.PricingMethodContract
	item.PricingReference = params.ContractID

	itemPrice := &ordmod.OrderItemPrice{}
	itemPrice.PriceType = ordmod.PriceTypeRecurring
	itemPrice.Price = &ordmod.Price{}
	curUnit := "USA"
	curVal := float32(12.34)
	itemPrice.Price.DutyFreeAmount = &ordmod.Money{Unit: &curUnit, Value: &curVal}
	itemPrice.Price.TaxIncludedAmount = &ordmod.Money{Unit: &curUnit, Value: &curVal}
	taxRate := float32(0)
	itemPrice.Price.TaxRate = &taxRate
	//itemPrice.Price.UnitOfMesure = ""
	item.OrderItemPrice = append(item.OrderItemPrice, itemPrice)

	itemPrice.RecurringChargePeriod = ordmod.ChargePeriodDay
}

func (s *sonataOrderImpl) BuildItemRelatedParty(item *ordmod.ProductOrderItemCreate, params *OrderParams) {
}

func (s *sonataOrderImpl) BuildItemBilling(item *ordmod.ProductOrderItemCreate, params *OrderParams) {
}

func (s *sonataOrderImpl) BuildOrderRelatedParty(order *ordmod.ProductOrderCreate, params *OrderParams) {
}

func (s *sonataOrderImpl) BuildOrderBilling(order *ordmod.ProductOrderCreate, params *OrderParams) {
	if params.Buyer != nil {
		partBuy := &ordmod.RelatedParty{}
		partBuy.Role = []string{"Buyer"}
		partBuy.ID = &params.Buyer.ID
		partBuy.Name = &params.Buyer.Name
		order.RelatedParty = append(order.RelatedParty, partBuy)
	}

	if params.Seller != nil {
		partSell := &ordmod.RelatedParty{}
		partSell.Role = []string{"Seller"}
		partSell.ID = &params.Seller.ID
		partSell.Name = &params.Seller.Name
		order.RelatedParty = append(order.RelatedParty, partSell)
	}
}
