package orchestra

import (
	"time"

	"github.com/qlcchain/go-lsobus/mock"

	"github.com/go-openapi/strfmt"

	quocli "github.com/qlcchain/go-lsobus/sonata/quote/client"
	quoapi "github.com/qlcchain/go-lsobus/sonata/quote/client/quote"
	quomod "github.com/qlcchain/go-lsobus/sonata/quote/models"
)

type sonataQuoteImpl struct {
	sonataBaseImpl
}

func newSonataQuoteImpl(o *Orchestra) *sonataQuoteImpl {
	s := &sonataQuoteImpl{}
	s.Orch = o
	s.Version = MEFAPIVersionQuote
	return s
}

func (s *sonataQuoteImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataQuoteImpl) NewHTTPClient() *quocli.APIQuoteManagement {
	tranCfg := quocli.DefaultTransportConfig().WithHost(s.GetHost()).WithSchemes([]string{s.GetScheme()})
	httpCli := quocli.NewHTTPClientWithConfig(nil, tranCfg)
	return httpCli
}

func (s *sonataQuoteImpl) SendCreateRequest(orderParams *OrderParams) error {
	reqParams := s.BuildCreateParams(orderParams)

	httpCli := s.NewHTTPClient()

	s.logger.Debugf("send request, payload %s", s.DumpValue(reqParams.Quote))

	rspParams, err := httpCli.Quote.QuoteCreate(reqParams)
	if err != nil {
		s.logger.Errorf("send request, error %s", err)
		//return err
		rspParams = mock.SonataGenerateQuoteCreateResponse(reqParams)
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))

	orderParams.RspQuote = rspParams.GetPayload()

	return nil
}

func (s *sonataQuoteImpl) SendFindRequest(params *FindParams) error {
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
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspQuoteList = rspParams.GetPayload()
	return nil
}

func (s *sonataQuoteImpl) SendGetRequest(params *GetParams) error {
	reqParams := quoapi.NewQuoteGetParams()
	reqParams.ID = params.ID

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.Quote.QuoteGet(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspQuote = rspParams.GetPayload()

	return nil
}

func (s *sonataQuoteImpl) BuildCreateParams(orderParams *OrderParams) *quoapi.QuoteCreateParams {
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

	// Source UNI
	srcUniItem := s.BuildUNIItem(orderParams, true)
	if srcUniItem != nil {
		reqParams.Quote.QuoteItem = append(reqParams.Quote.QuoteItem, srcUniItem)
	}

	// Destination UNI
	dstUniItem := s.BuildUNIItem(orderParams, false)
	if dstUniItem != nil {
		reqParams.Quote.QuoteItem = append(reqParams.Quote.QuoteItem, dstUniItem)
	}

	// ELine
	lineItem := s.BuildELineItem(orderParams)
	if lineItem != nil {
		// Related Items
		if srcUniItem != nil {
			relItem := &quomod.QuoteItemRelationship{
				Type: quomod.RelationshipTypeRELIESON,
				ID:   srcUniItem.ID,
			}
			lineItem.QuoteItemRelationship = append(lineItem.QuoteItemRelationship, relItem)
		}

		if dstUniItem != nil {
			relItem := &quomod.QuoteItemRelationship{
				Type: quomod.RelationshipTypeRELIESON,
				ID:   dstUniItem.ID,
			}
			lineItem.QuoteItemRelationship = append(lineItem.QuoteItemRelationship, relItem)
		}

		// Related Products
		if orderParams.SrcPortID != "" {
			relType := string(quomod.RelationshipTypeRELIESON)
			relProd := &quomod.ProductRelationship{Type: &relType}
			relProd.Product = &quomod.ProductRef{}
			relProdID := orderParams.SrcPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}

		if orderParams.DstPortID != "" {
			relType := string(quomod.RelationshipTypeRELIESON)
			relProd := &quomod.ProductRelationship{Type: &relType}
			relProd.Product = &quomod.ProductRef{}
			relProdID := orderParams.DstPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}

		reqParams.Quote.QuoteItem = append(reqParams.Quote.QuoteItem, lineItem)
	}

	return reqParams
}

func (s *sonataQuoteImpl) BuildUNIItem(orderParams *OrderParams, isDirSrc bool) *quomod.QuoteItemCreate {
	var siteID string
	if isDirSrc {
		siteID = orderParams.SrcSiteID
	} else {
		siteID = orderParams.DstSiteID
	}
	if siteID == "" {
		return nil
	}

	uniItem := &quomod.QuoteItemCreate{}

	uniItemID := s.NewItemID()
	uniItem.ID = &uniItemID
	uniItem.Action = quomod.ProductActionType(orderParams.ItemAction)

	uniOfferId := MEFProductOfferingUNI
	uniItem.ProductOffering = &quomod.ProductOfferingRef{ID: &uniOfferId}

	uniItem.Product = &quomod.Product{}
	if uniItem.Action != quomod.ProductActionTypeINSTALL {
		uniItem.Product.ID = &orderParams.ProductID
	}

	// UNI Place
	if siteID != "" {
		uniPlace := &quomod.ReferencedAddress{}
		uniPlace.ReferenceID = &siteID
		uniItem.Product.SetPlace([]quomod.RelatedPlaceRefOrValue{uniPlace})
	}

	// UNI Product Specification
	uniItem.Product.ProductSpecification = &quomod.ProductSpecificationRef{}
	uniItem.Product.ProductSpecification.ID = "UNISpec"
	uniDesc := s.BuildUNIProductSpec(orderParams)
	uniItem.Product.ProductSpecification.SetDescribing(uniDesc)

	s.BuildItemTerm(uniItem, orderParams)

	return uniItem
}

func (s *sonataQuoteImpl) BuildELineItem(orderParams *OrderParams) *quomod.QuoteItemCreate {
	if orderParams.Bandwidth == 0 {
		return nil
	}

	lineItem := &quomod.QuoteItemCreate{}
	lineItem.Action = quomod.ProductActionType(orderParams.ItemAction)

	lineItemID := s.NewItemID()
	lineItem.ID = &lineItemID

	linePoVal := MEFProductOfferingELine
	lineItem.ProductOffering = &quomod.ProductOfferingRef{ID: &linePoVal}

	lineItem.Product = new(quomod.Product)
	if lineItem.Action != quomod.ProductActionTypeINSTALL {
		lineItem.Product.ID = &orderParams.ProductID
	}

	//Product Specification
	lineItem.Product.ProductSpecification = &quomod.ProductSpecificationRef{}
	lineItem.Product.ProductSpecification.ID = "UNISpec"
	lineDesc := s.BuildELineProductSpec(orderParams)
	lineItem.Product.ProductSpecification.SetDescribing(lineDesc)

	s.BuildItemTerm(lineItem, orderParams)

	return lineItem
}

func (s *sonataQuoteImpl) BuildItemTerm(item *quomod.QuoteItemCreate, orderParams *OrderParams) {
	/*
		item.RequestedQuoteItemTerm = &quomod.ItemTerm{}
		item.RequestedQuoteItemTerm.Duration = &quomod.Duration{}
		item.RequestedQuoteItemTerm.Duration.Unit = quomod.DurationUnitYEAR
		durVal := int32(99)
		item.RequestedQuoteItemTerm.Duration.Value = &durVal
	*/
}
