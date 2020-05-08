package orchestra

import (
	"strconv"

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

func (s *sonataQuoteImpl) SendCreateRequest() {
	reqParams := s.BuildCreateParams()

	tranCfg := quocli.DefaultTransportConfig().WithHost("localhost").WithSchemes([]string{"http"})
	httpCli := quocli.NewHTTPClientWithConfig(nil, tranCfg)

	rspParams, err := httpCli.Quote.QuoteCreate(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())
}

func (s *sonataQuoteImpl) BuildCreateParams() *quoapi.QuoteCreateParams {
	reqParams := &quoapi.QuoteCreateParams{}

	reqParams.Quote = &quomod.QuoteCreate{}

	isqVal := true
	reqParams.Quote.InstantSyncQuoting = &isqVal
	reqParams.Quote.QuoteLevel = quomod.QuoteLevelFIRM

	itemIdSeq := int(0)

	// UNI
	itemIdSeq++
	uniItem := &quomod.QuoteItemCreate{}
	uniItemID := strconv.Itoa(itemIdSeq)
	uniItem.ID = &uniItemID
	uniItem.Action = "INSTALL"

	uniPoId := "LSO_Sonata_ProviderOnDemand_EthernetPort_UNI"
	uniItem.ProductOffering = &quomod.ProductOfferingRef{ID: &uniPoId}

	uniItem.Product = &quomod.Product{}

	// UNI Place
	uniPlace := &quomod.ReferencedAddress{}
	addrId := "1111-2222-3333"
	uniPlace.ReferenceID = &addrId
	uniItem.Product.SetPlace([]quomod.RelatedPlaceRefOrValue{uniPlace})

	// UNI Product Specification

	// ELine
	itemIdSeq++
	lineItem := new(quomod.QuoteItemCreate)
	lineItem.Action = "INSTALL"
	lineItemID := strconv.Itoa(itemIdSeq)
	lineItem.ID = &lineItemID
	linePoVal := "LSO_Sonata_ProviderOnDemand_EthernetConnection"
	lineItem.ProductOffering = &quomod.ProductOfferingRef{ID: &linePoVal}
	lineItem.Product = new(quomod.Product)

	//Product Specification

	// Related Products
	enniRelyType := string(quomod.RelationshipTypeRELIESON)
	lineItemRelOnEnni := &quomod.ProductRelationship{Type: &enniRelyType}
	lineItemRelOnEnni.Product = &quomod.ProductRef{}
	enniId := "8888"
	lineItemRelOnEnni.Product.ID = &enniId
	lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, lineItemRelOnEnni)

	lineItemRelOnUni := &quomod.QuoteItemRelationship{Type: quomod.RelationshipTypeRELIESON}
	lineItemRelOnUni.ID = &uniItemID
	lineItem.QuoteItemRelationship = append(lineItem.QuoteItemRelationship, lineItemRelOnUni)

	reqParams.Quote.QuoteItem = append(reqParams.Quote.QuoteItem, lineItem)

	// Related Parties
	/*
		partyCBC := &poqmod.RelatedParty{}
		partyCBC.Role = []string{"Buyer"}
		partyCBC.ID = "Partner111"
		cbcName := "CBC"
		partyCBC.Name = &cbcName
		cbcNumber := "12345678"
		partyCBC.Number = &cbcNumber
		poqCreate.RelatedParty = append(poqCreate.RelatedParty, partyCBC)*/

	return reqParams
}
