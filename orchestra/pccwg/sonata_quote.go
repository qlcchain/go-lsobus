package pccwg

import (
	"time"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/mock"

	"github.com/go-openapi/strfmt"

	quocli "github.com/qlcchain/go-lsobus/orchestra/sonata/quote/client"
	quoapi "github.com/qlcchain/go-lsobus/orchestra/sonata/quote/client/quote"
	quomod "github.com/qlcchain/go-lsobus/orchestra/sonata/quote/models"
)

type sonataQuoteImpl struct {
	sonataBaseImpl
}

func newSonataQuoteImpl(p api.DoDSeller) *sonataQuoteImpl {
	s := &sonataQuoteImpl{}
	s.Partner = p
	s.Version = MEFAPIVersionQuote
	return s
}

func (s *sonataQuoteImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataQuoteImpl) NewHTTPClient() *quocli.APIQuoteManagement {
	httpTran := s.NewHttpTransport(quocli.DefaultBasePath)
	httpCli := quocli.New(httpTran, nil)
	return httpCli
}

func (s *sonataQuoteImpl) SendCreateRequest(orderParams *api.OrderParams) error {
	s.logger.Debugf("params: %s", s.DumpValue(orderParams))

	reqParams := s.BuildCreateParams(orderParams)

	httpCli := s.NewHTTPClient()

	s.logger.Debugf("send request, payload %s", s.DumpValue(reqParams.Quote))

	rspParams, err := httpCli.Quote.QuoteCreate(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateQuoteCreateResponse(reqParams)
	} else if err != nil {
		s.logger.Errorf("send request, error %s", err)
		s.handleResponseError(err)
		return err
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))

	orderParams.RspQuote = rspParams.GetPayload()

	return nil
}

func (s *sonataQuoteImpl) SendFindRequest(params *api.FindParams) error {
	reqParams := quoapi.NewQuoteFindParams()
	if params.ProjectID != "" {
		reqParams.ProjectID = &params.ProjectID
	}
	if params.ExternalID != "" {
		reqParams.ExternalID = &params.ExternalID
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

	rspParams, err := httpCli.Quote.QuoteFind(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateQuoteFindResponse(reqParams)
	} else if err != nil {
		s.logger.Error("send request,", "error:", err)
		s.handleResponseError(err)
		return err
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspQuoteList = rspParams.GetPayload()
	params.XResultCount = rspParams.XResultCount
	params.XTotalCount = rspParams.XTotalCount

	return nil
}

func (s *sonataQuoteImpl) SendGetRequest(params *api.GetParams) error {
	reqParams := quoapi.NewQuoteGetParams()
	reqParams.ID = params.ID

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.Quote.QuoteGet(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateQuoteGetResponse(reqParams)
	} else if err != nil {
		s.logger.Error("send request,", "error:", err)
		s.handleResponseError(err)
		return err
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspQuote = rspParams.GetPayload()

	return nil
}

func (s *sonataQuoteImpl) BuildCreateParams(orderParams *api.OrderParams) *quoapi.QuoteCreateParams {
	reqParams := quoapi.NewQuoteCreateParams()

	reqParams.Quote = &quomod.QuoteCreate{}

	reqParams.Quote.ExternalID = orderParams.ExternalID
	reqParams.Quote.Description = orderParams.Description
	reqParams.Quote.ProjectID = orderParams.ProjectID

	isqVal := true
	reqParams.Quote.InstantSyncQuoting = &isqVal
	reqParams.Quote.QuoteLevel = quomod.QuoteLevelFIRM
	reqParams.Quote.ExpectedFulfillmentStartDate.Scan(time.Now())
	reqParams.Quote.RequestedQuoteCompletionDate = &strfmt.DateTime{}
	reqParams.Quote.RequestedQuoteCompletionDate.Scan(time.Now().Add(time.Minute))

	// UNI
	var allUniItems []*quomod.QuoteItemCreate
	for _, uniParams := range orderParams.UNIItems {
		uniItem := s.BuildUNIItem(uniParams)
		if uniItem == nil {
			continue
		}
		reqParams.Quote.QuoteItem = append(reqParams.Quote.QuoteItem, uniItem)
		allUniItems = append(allUniItems, uniItem)
	}

	// ELine
	var allLineItems []*quomod.QuoteItemCreate
	for _, lineParams := range orderParams.ELineItems {
		lineItem := s.BuildELineItem(lineParams)
		if lineItem == nil {
			continue
		}
		reqParams.Quote.QuoteItem = append(reqParams.Quote.QuoteItem, lineItem)
		allLineItems = append(allUniItems, lineItem)

		// Related Products
		if lineParams.SrcPortID != "" {
			relType := string(quomod.RelationshipTypeRELIESON)
			relProd := &quomod.ProductRelationship{Type: &relType}
			relProd.Product = &quomod.ProductRef{}
			relProdID := lineParams.SrcPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}

		if lineParams.DstPortID != "" {
			relType := string(quomod.RelationshipTypeRELIESON)
			relProd := &quomod.ProductRelationship{Type: &relType}
			relProd.Product = &quomod.ProductRef{}
			relProdID := lineParams.DstPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}
	}

	// Related Items
	if len(allLineItems) == 1 {
		lineItem := allLineItems[0]
		for _, uniItem := range allUniItems {
			relItem := &quomod.QuoteItemRelationship{
				Type: quomod.RelationshipTypeRELIESON,
				ID:   uniItem.ID,
			}
			lineItem.QuoteItemRelationship = append(lineItem.QuoteItemRelationship, relItem)
		}
	}

	return reqParams
}

func (s *sonataQuoteImpl) BuildUNIItem(params *api.UNIItemParams) *quomod.QuoteItemCreate {
	if params.ProdSpecID != "" && params.ProdSpecID != "UNISpec" {
		return nil
	}

	uniItem := &quomod.QuoteItemCreate{}

	if params.ItemID != "" {
		uniItem.ID = &params.ItemID
	} else {
		uniItemID := s.NewItemID()
		uniItem.ID = &uniItemID
	}

	uniItem.Action = quomod.ProductActionType(params.Action)

	uniItem.ProductOffering = &quomod.ProductOfferingRef{ID: &params.ProdOfferID}

	uniItem.Product = &quomod.Product{}
	if uniItem.Action != quomod.ProductActionTypeINSTALL {
		uniItem.Product.ID = &params.ProductID
	}

	// UNI Place
	if params.SiteID != "" {
		uniPlace := &quomod.ReferencedAddress{}
		uniPlace.ReferenceID = &params.SiteID
		uniItem.Product.SetPlace([]quomod.RelatedPlaceRefOrValue{uniPlace})
	}

	// UNI Product Specification
	if uniItem.Action != quomod.ProductActionTypeDISCONNECT {
		uniItem.Product.ProductSpecification = &quomod.ProductSpecificationRef{}
		uniItem.Product.ProductSpecification.ID = "UNISpec"
		uniDesc := s.BuildUNIProductSpec(params)
		uniItem.Product.ProductSpecification.SetDescribing(uniDesc)
	}

	// Term
	/*
		if params.DurationUnit != "" && params.DurationAmount > 0 {
			uniItem.RequestedQuoteItemTerm = &quomod.ItemTerm{}
			uniItem.RequestedQuoteItemTerm.Duration = &quomod.Duration{}
			uniItem.RequestedQuoteItemTerm.Duration.Value = sonata.NewInt32(int32(params.DurationAmount))
			uniItem.RequestedQuoteItemTerm.Duration.Unit = quomod.DurationUnit(params.DurationUnit)
		}
	*/

	//Price
	/*
		itemPrice := &quomod.QuotePrice{}
		itemPrice.PriceType = quomod.PriceTypeRECURRING
		itemPrice.RecurringChargePeriod = quomod.ChargePeriodDAY
		uniItem.QuoteItemPrice = append(uniItem.QuoteItemPrice, itemPrice)
	*/

	return uniItem
}

func (s *sonataQuoteImpl) BuildELineItem(params *api.ELineItemParams) *quomod.QuoteItemCreate {
	if params.ProdSpecID != "" && params.ProdSpecID != "ELineSpec" {
		return nil
	}

	if params.Action != string(quomod.ProductActionTypeDISCONNECT) {
		if params.Bandwidth == 0 {
			return nil
		}
	}

	lineItem := &quomod.QuoteItemCreate{}
	lineItem.Action = quomod.ProductActionType(params.Action)

	if params.ItemID != "" {
		lineItem.ID = &params.ItemID
	} else {
		lineItemID := s.NewItemID()
		lineItem.ID = &lineItemID
	}

	lineItem.ProductOffering = &quomod.ProductOfferingRef{ID: &params.ProdOfferID}

	lineItem.Product = new(quomod.Product)
	lineItem.Product.ID = &params.ProductID

	//Product Specification
	lineItem.Product.ProductSpecification = &quomod.ProductSpecificationRef{}
	lineItem.Product.ProductSpecification.ID = "PCCWConnSpec"
	lineDesc := s.BuildPCCWConnProductSpec(params)
	lineItem.Product.ProductSpecification.SetDescribing(lineDesc)

	// Term
	/*
		if params.DurationUnit != "" && params.DurationAmount > 0 {
			lineItem.RequestedQuoteItemTerm = &quomod.ItemTerm{}
			lineItem.RequestedQuoteItemTerm.Duration = &quomod.Duration{}
			lineItem.RequestedQuoteItemTerm.Duration.Value = sonata.NewInt32(int32(params.DurationAmount))
			lineItem.RequestedQuoteItemTerm.Duration.Unit = quomod.DurationUnit(params.DurationUnit)
		}
	*/

	return lineItem
}

func (s *sonataQuoteImpl) handleResponseError(rspErr error) {
	switch rspErr.(type) {
	case *quoapi.QuoteCreateUnauthorized, *quoapi.QuoteCreateForbidden:
		s.ClearApiToken()
	case *quoapi.QuoteFindUnauthorized, *quoapi.QuoteFindForbidden:
		s.ClearApiToken()
	case *quoapi.QuoteGetUnauthorized, *quoapi.QuoteGetForbidden:
		s.ClearApiToken()
	}
}
