package pccwg

import (
	"time"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/mock"

	"github.com/go-openapi/strfmt"

	ordcli "github.com/qlcchain/go-lsobus/orchestra/sonata/order/client"
	ordapi "github.com/qlcchain/go-lsobus/orchestra/sonata/order/client/product_order"
	ordmod "github.com/qlcchain/go-lsobus/orchestra/sonata/order/models"
)

type sonataOrderImpl struct {
	sonataBaseImpl
}

func newSonataOrderImpl(p api.DoDSeller) *sonataOrderImpl {
	s := &sonataOrderImpl{}
	s.Partner = p
	s.Version = MEFAPIVersionOrder
	return s
}

func (s *sonataOrderImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataOrderImpl) NewHTTPClient() *ordcli.APIProductOrderManagement {
	httpTran := s.NewHttpTransport(ordcli.DefaultBasePath)
	httpCli := ordcli.New(httpTran, nil)
	return httpCli
}

func (s *sonataOrderImpl) SendCreateRequest(orderParams *api.OrderParams) error {
	s.logger.Debugf("params: %s", s.DumpValue(orderParams))

	reqParams := s.BuildCreateParams(orderParams)

	httpCli := s.NewHTTPClient()

	s.logger.Debugf("send request, payload %s", s.DumpValue(reqParams.ProductOrder))

	rspParams, err := httpCli.ProductOrder.ProductOrderCreate(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateOrderCreateResponse(reqParams)
	} else if err != nil {
		s.logger.Errorf("send request, error %s", err)
		s.handleResponseError(err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	orderParams.RspOrder = rspParams.GetPayload()

	return nil
}

func (s *sonataOrderImpl) SendFindRequest(params *api.FindParams) error {
	reqParams := ordapi.NewProductOrderFindParams()
	if params.ProjectID != "" {
		reqParams.ProjectID = &params.ProjectID
	}
	if params.ExternalID != "" {
		reqParams.ExternalID = &params.ExternalID
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

	httpCli := s.NewHTTPClient()
	rspParams, err := httpCli.ProductOrder.ProductOrderFind(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateOrderFindResponse(reqParams)
	} else if err != nil {
		s.logger.Error("send request,", "error:", err)
		s.handleResponseError(err)
		return err
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspOrderList = rspParams.GetPayload()
	params.XResultCount = rspParams.XResultCount
	params.XTotalCount = rspParams.XTotalCount

	return nil
}

func (s *sonataOrderImpl) SendGetRequest(params *api.GetParams) error {
	reqParams := ordapi.NewProductOrderGetParams()
	reqParams.ProductOrderID = params.ID

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.ProductOrder.ProductOrderGet(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateOrderGetResponse(reqParams)
	} else if err != nil {
		s.logger.Error("send request,", "error:", err)
		s.handleResponseError(err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspOrder = rspParams.GetPayload()

	return nil
}

func (s *sonataOrderImpl) BuildCreateParams(orderParams *api.OrderParams) *ordapi.ProductOrderCreateParams {
	reqParams := ordapi.NewProductOrderCreateParams()

	reqParams.ProductOrder = &ordmod.ProductOrderCreate{}
	reqParams.ProductOrder.OrderVersion = &s.Version

	reqParams.ProductOrder.ProjectID = orderParams.ProjectID
	if orderParams.ExternalID != "" {
		reqParams.ProductOrder.ExternalID = &orderParams.ExternalID
	}

	reqParams.ProductOrder.OrderActivity = ordmod.OrderActivity(orderParams.OrderActivity)
	reqParams.ProductOrder.BuyerRequestDate = &strfmt.DateTime{}
	reqParams.ProductOrder.BuyerRequestDate.Scan(time.Now())
	reqParams.ProductOrder.RequestedStartDate.Scan(time.Now().Add(time.Minute))
	reqParams.ProductOrder.RequestedCompletionDate = &strfmt.DateTime{}
	reqParams.ProductOrder.RequestedCompletionDate.Scan(time.Now().Add(time.Hour))
	reqParams.ProductOrder.DesiredResponse = ordmod.DesiredOrderResponsesConfirmationAndEngineeringDesign
	reqParams.ProductOrder.ExpeditePriority = true

	// UNI
	var uniItemList []*ordmod.ProductOrderItemCreate
	for _, uniParams := range orderParams.UNIItems {
		uniItem := s.BuildUNIItem(uniParams)
		if uniItem == nil {
			continue
		}
		reqParams.ProductOrder.OrderItem = append(reqParams.ProductOrder.OrderItem, uniItem)
		uniItemList = append(uniItemList, uniItem)
	}

	// ELine
	var lineItemList []*ordmod.ProductOrderItemCreate
	for _, lineParams := range orderParams.ELineItems {
		lineItem := s.BuildELineItem(lineParams)
		if lineItem == nil {
			continue
		}

		// Related Products
		/*
			if lineParams.SrcPortID != "" {
				relType := string("RELIES_ON")
				relProd := &ordmod.ProductRelationship{Type: &relType}
				relProd.Product = &ordmod.ProductRef{}
				relProdID := lineParams.SrcPortID
				relProd.Product.ID = &relProdID
				lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
			}

			if lineParams.DstPortID != "" {
				relType := string("RELIES_ON")
				relProd := &ordmod.ProductRelationship{Type: &relType}
				relProd.Product = &ordmod.ProductRef{}
				relProdID := lineParams.DstPortID
				relProd.Product.ID = &relProdID
				lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
			}
		*/

		reqParams.ProductOrder.OrderItem = append(reqParams.ProductOrder.OrderItem, lineItem)
		lineItemList = append(lineItemList, lineItem)
	}

	// Related Items
	if len(lineItemList) == 1 && reqParams.ProductOrder.OrderActivity == ordmod.OrderActivityInstall {
		lineItem := lineItemList[0]
		for _, uniItem := range uniItemList {
			relType := string("RELIES_ON")
			relItem := &ordmod.OrderItemRelationShip{
				Type: &relType,
				ID:   uniItem.ID,
			}
			lineItem.OrderItemRelationship = append(lineItem.OrderItemRelationship, relItem)
		}
	}

	// Billing
	s.BuildOrderBilling(reqParams.ProductOrder, orderParams)

	// Party
	s.BuildOrderRelatedParty(reqParams.ProductOrder, orderParams)

	return reqParams
}

func (s *sonataOrderImpl) BuildUNIItem(params *api.UNIItemParams) *ordmod.ProductOrderItemCreate {
	if params.ProdSpecID != "" && params.ProdSpecID != "UNISpec" {
		return nil
	}
	uniItem := &ordmod.ProductOrderItemCreate{}

	if params.ItemID != "" {
		uniItem.ID = &params.ItemID
	} else {
		uniItemID := s.NewItemID()
		uniItem.ID = &uniItemID
	}

	uniItem.Action = ordmod.ProductActionType(params.Action)

	uniItem.ProductOffering = &ordmod.ProductOfferingRef{ID: &params.ProdOfferID}

	uniItem.Product = &ordmod.Product{}
	uniItem.Product.BuyerProductID = params.BuyerProductID
	if uniItem.Action != ordmod.ProductActionTypeAdd {
		uniItem.Product.ID = params.ProductID
	}

	// UNI Place
	if params.SiteID != "" {
		uniPlace := &ordmod.ReferencedAddress{}
		uniPlace.ReferenceID = &params.SiteID
		uniItem.Product.SetPlace([]ordmod.RelatedPlaceReforValue{uniPlace})
	}

	if uniItem.Action != ordmod.ProductActionTypeRemove {
		// Quote
		uniItem.Quote = &ordmod.QuoteRef{ID: &params.QuoteID, QuoteItem: params.QuoteItemID}

		// UNI Product Specification
		uniItem.Product.ProductSpecification = &ordmod.ProductSpecificationRef{}
		uniItem.Product.ProductSpecification.ID = "UNISpec"
		uniDesc := s.BuildUNIProductSpec(params)
		uniItem.Product.ProductSpecification.SetDescribing(uniDesc)

		// Price
		s.BuildItemPrice(uniItem, params.BillingParams)

		// Term
		//uniItem.PricingTerm = sonata.NewInt32(36)
	}

	return uniItem
}

func (s *sonataOrderImpl) BuildELineItem(params *api.ELineItemParams) *ordmod.ProductOrderItemCreate {
	if params.ProdSpecID != "" && params.ProdSpecID != "ELineSpec" {
		return nil
	}

	if params.Action != string(ordmod.ProductActionTypeRemove) {
		if params.Bandwidth == 0 {
			return nil
		}
	}

	lineItem := &ordmod.ProductOrderItemCreate{}
	lineItem.Action = ordmod.ProductActionType(params.Action)

	if params.ItemID != "" {
		lineItem.ID = &params.ItemID
	} else {
		lineItemID := s.NewItemID()
		lineItem.ID = &lineItemID
	}

	lineItem.ProductOffering = &ordmod.ProductOfferingRef{ID: &params.ProdOfferID}

	lineItem.Product = &ordmod.Product{}
	lineItem.Product.BuyerProductID = params.BuyerProductID
	if lineItem.Action != ordmod.ProductActionTypeAdd {
		lineItem.Product.ID = params.ProductID
	}

	if lineItem.Action != ordmod.ProductActionTypeRemove {
		//Product Specification
		lineItem.Product.ProductSpecification = &ordmod.ProductSpecificationRef{}
		lineItem.Product.ProductSpecification.ID = "ELineSpec"
		lineDesc := s.BuildPCCWConnProductSpec(params)
		lineItem.Product.ProductSpecification.SetDescribing(lineDesc)

		// Quote
		lineItem.Quote = &ordmod.QuoteRef{ID: &params.QuoteID, QuoteItem: params.QuoteItemID}

		// Price
		s.BuildItemPrice(lineItem, params.BillingParams)

		// Term
		//lineItem.PricingTerm = sonata.NewInt32(36)
	}

	return lineItem
}

func (s *sonataOrderImpl) BuildItemPrice(item *ordmod.ProductOrderItemCreate, params *api.BillingParams) {
	if params == nil {
		return
	}

	// Price
	//item.PricingMethod = ordmod.PricingMethodContract
	//item.PricingReference = params.ContractID

	itemPrice := &ordmod.OrderItemPrice{}
	if params.BillingType == api.BillingTypeDOD {
		itemPrice.PriceType = ordmod.PriceTypeNonRecurring
	} else if params.BillingType == api.BillingTypePAYG {
		itemPrice.PriceType = ordmod.PriceTypeRecurring
		itemPrice.RecurringChargePeriod = ordmod.ChargePeriod(params.BillingUnit)
		itemPrice.Price.UnitOfMesure = params.MeasureUnit
	} else if params.BillingType == api.BillingTypeUsage {
		itemPrice.PriceType = ordmod.PriceTypeRecurring
		itemPrice.RecurringChargePeriod = ordmod.ChargePeriod(params.BillingUnit)
		itemPrice.Price.UnitOfMesure = params.MeasureUnit
	}

	itemPrice.Price = &ordmod.Price{}
	curUnit := params.Currency
	curVal := float32(params.Price)
	itemPrice.Price.DutyFreeAmount = &ordmod.Money{Unit: &curUnit, Value: &curVal}
	itemPrice.Price.TaxIncludedAmount = &ordmod.Money{Unit: &curUnit, Value: &curVal}
	taxRate := float32(0)
	itemPrice.Price.TaxRate = &taxRate
	item.OrderItemPrice = append(item.OrderItemPrice, itemPrice)
}

func (s *sonataOrderImpl) BuildOrderRelatedParty(order *ordmod.ProductOrderCreate, params *api.OrderParams) {
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

func (s *sonataOrderImpl) BuildOrderBilling(order *ordmod.ProductOrderCreate, params *api.OrderParams) {
	order.PaymentType = params.PaymentType
	order.BillingType = params.BillingType
	if order.BillingType == "DOD" {
		order.BillingType = "DOM"
	}
}

func (s *sonataOrderImpl) handleResponseError(rspErr error) {
	switch rspErr.(type) {
	case *ordapi.ProductOrderCreateUnauthorized, *ordapi.ProductOrderCreateForbidden:
		s.ClearApiToken()
	case *ordapi.ProductOrderFindUnauthorized, *ordapi.ProductOrderFindForbidden:
		s.ClearApiToken()
	case *ordapi.ProductOrderGetUnauthorized, *ordapi.ProductOrderGetForbidden:
		s.ClearApiToken()
	}
}
