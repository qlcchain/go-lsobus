package orchestra

import (
	"time"

	cmnmod "github.com/iixlabs/virtual-lsobus/sonata/common/models"

	"github.com/go-openapi/strfmt"

	quocli "github.com/iixlabs/virtual-lsobus/sonata/quote/client"
	quoapi "github.com/iixlabs/virtual-lsobus/sonata/quote/client/quote"
	quomod "github.com/iixlabs/virtual-lsobus/sonata/quote/models"
)

type sonataQuoteImpl struct {
	sonataBaseImpl
}

func newSonataQuoteImpl() *sonataQuoteImpl {
	s := &sonataQuoteImpl{}
	return s
}

func (s *sonataQuoteImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataQuoteImpl) SendCreateRequest(orderParams *OrderParams) error {
	reqParams := s.BuildCreateParams(orderParams)

	tranCfg := quocli.DefaultTransportConfig().WithHost("localhost").WithSchemes([]string{"http"})
	httpCli := quocli.NewHTTPClientWithConfig(nil, tranCfg)

	rspParams, err := httpCli.Quote.QuoteCreate(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())
	return nil
}

func (s *sonataQuoteImpl) BuildCreateParams(orderParams *OrderParams) *quoapi.QuoteCreateParams {
	reqParams := &quoapi.QuoteCreateParams{}

	reqParams.Quote = &quomod.QuoteCreate{}

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
	uniItem.Action = quomod.ProductActionTypeINSTALL

	uniOfferId := "LSO_Sonata_ProviderOnDemand_EthernetPort_UNI"
	uniItem.ProductOffering = &quomod.ProductOfferingRef{ID: &uniOfferId}

	uniItem.Product = &quomod.Product{}

	// UNI Place
	if siteID != "" {
		uniPlace := &quomod.ReferencedAddress{}
		uniPlace.ReferenceID = &siteID
		uniItem.Product.SetPlace([]quomod.RelatedPlaceRefOrValue{uniPlace})
	}

	// UNI Product Specification
	uniDesc := &cmnmod.UNIProductSpecification{}
	s.FillUNIProductSpec(uniDesc, orderParams)
	uniItem.Product.ProductSpecification.SetDescribing(uniDesc)

	return uniItem
}

func (s *sonataQuoteImpl) BuildELineItem(orderParams *OrderParams) *quomod.QuoteItemCreate {
	if orderParams.Bandwidth == 0 {
		return nil
	}

	lineItem := &quomod.QuoteItemCreate{}
	lineItem.Action = quomod.ProductActionTypeINSTALL

	lineItemID := s.NewItemID()
	lineItem.ID = &lineItemID

	linePoVal := "LSO_Sonata_ProviderOnDemand_EthernetConnection"
	lineItem.ProductOffering = &quomod.ProductOfferingRef{ID: &linePoVal}

	lineItem.Product = new(quomod.Product)

	//Product Specification
	lineItem.Product.ProductSpecification = &quomod.ProductSpecificationRef{}
	lineDesc := &cmnmod.ELineProductSpecification{}
	s.FillELineProductSpec(lineDesc, orderParams)
	lineItem.Product.ProductSpecification.SetDescribing(lineDesc)

	return lineItem
}
