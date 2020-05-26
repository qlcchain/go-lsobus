package mock

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	cmnmod "github.com/qlcchain/go-lsobus/sonata/common/models"

	"github.com/qlcchain/go-lsobus/sonata"
	sitapi "github.com/qlcchain/go-lsobus/sonata/site/client/geographic_site"
	sitmod "github.com/qlcchain/go-lsobus/sonata/site/models"

	"github.com/bitly/go-simplejson"

	"github.com/go-openapi/strfmt"

	invapi "github.com/qlcchain/go-lsobus/sonata/inventory/client/product"
	invmod "github.com/qlcchain/go-lsobus/sonata/inventory/models"
	ordapi "github.com/qlcchain/go-lsobus/sonata/order/client/product_order"
	ordmod "github.com/qlcchain/go-lsobus/sonata/order/models"
	poqapi "github.com/qlcchain/go-lsobus/sonata/poq/client/product_offering_qualification"
	poqmod "github.com/qlcchain/go-lsobus/sonata/poq/models"
	quoapi "github.com/qlcchain/go-lsobus/sonata/quote/client/quote"
	quomod "github.com/qlcchain/go-lsobus/sonata/quote/models"
)

func SonataGenerateSiteFindResponse(reqParams *sitapi.GeographicSiteFindParams) *sitapi.GeographicSiteFindOK {
	rspParams := sitapi.NewGeographicSiteFindOK()

	site1 := &sitmod.GeographicSiteFindResp{}
	site1.GeographicAddress = &sitmod.GeographicAddressFindResp{}
	site1.GeographicAddress.FormattedAddress = &sitmod.FormattedAddress{}
	site1.GeographicAddress.FormattedAddress.ID = "PCCW-Addr-111"
	site1.GeographicAddress.FormattedAddress.Country = sonata.NewString("Japan")
	site1.GeographicAddress.FormattedAddress.City = sonata.NewString("Tokyo")
	site1.ID = "PCCW-Site-111"
	site1.SiteName = "DC111"
	site1.Status = sitmod.StatusExisting
	rspParams.Payload = append(rspParams.Payload, site1)

	site2 := &sitmod.GeographicSiteFindResp{}
	site2.GeographicAddress = &sitmod.GeographicAddressFindResp{}
	site2.GeographicAddress.FormattedAddress = &sitmod.FormattedAddress{}
	site2.GeographicAddress.FormattedAddress.ID = "PCCW-Addr-222"
	site2.GeographicAddress.FormattedAddress.Country = sonata.NewString("Korea")
	site2.GeographicAddress.FormattedAddress.City = sonata.NewString("Seoul")
	site2.ID = "PCCW-Site-222"
	site1.SiteName = "DC222"
	site2.Status = sitmod.StatusExisting
	rspParams.Payload = append(rspParams.Payload, site2)

	return rspParams
}

func SonataGenerateSiteGetResponse(reqParams *sitapi.GeographicSiteGetParams) *sitapi.GeographicSiteGetOK {
	rspParams := sitapi.NewGeographicSiteGetOK()
	rspParams.Payload = &sitmod.GeographicSite{}
	rspParams.Payload.ID = reqParams.SiteID
	return rspParams
}

func SonataGenerateInvFindResponse(reqParams *invapi.ProductFindParams) *invapi.ProductFindOK {
	rspParams := invapi.NewProductFindOK()

	if reqParams.ProductSpecificationID == nil || *reqParams.ProductSpecificationID == "UNISpec" {
		prodItem1 := &invmod.ProductSummary{}
		rspParams.Payload = append(rspParams.Payload, prodItem1)

		prodID1 := uuid.New().String()
		prodItem1.ID = &prodID1
		prodItem1.BuyerProductID = uuid.New().String()
		prodItem1.ProductSpecification = &invmod.ProductSpecificationSummary{}
		prodSpecID1 := "UNISpec"
		prodItem1.ProductSpecification.ID = &prodSpecID1
		prodItem1.StartDate.Scan(time.Now().Add(time.Duration(rand.Intn(48)) * time.Hour))
		prodItem1.Status = invmod.ProductStatusActive

		prodItem2 := &invmod.ProductSummary{}
		rspParams.Payload = append(rspParams.Payload, prodItem2)

		prodID2 := uuid.New().String()
		prodItem2.ID = &prodID2
		prodItem2.BuyerProductID = uuid.New().String()
		prodItem2.ProductSpecification = &invmod.ProductSpecificationSummary{}
		prodSpecID2 := "UNISpec"
		prodItem2.ProductSpecification.ID = &prodSpecID2
		prodItem2.StartDate.Scan(time.Now().Add(time.Duration(rand.Intn(48)) * time.Hour))
		prodItem2.Status = invmod.ProductStatusActive
	}

	if reqParams.ProductSpecificationID == nil || *reqParams.ProductSpecificationID == "ELineSpec" {
		prodItem1 := &invmod.ProductSummary{}
		rspParams.Payload = append(rspParams.Payload, prodItem1)

		prodID1 := uuid.New().String()
		prodItem1.ID = &prodID1
		prodItem1.BuyerProductID = uuid.New().String()
		prodItem1.ProductSpecification = &invmod.ProductSpecificationSummary{}
		prodSpecID1 := "ELineSpec"
		prodItem1.ProductSpecification.ID = &prodSpecID1
		prodItem1.StartDate.Scan(time.Now().Add(time.Duration(rand.Intn(48)) * time.Hour))
		prodItem1.Status = invmod.ProductStatusActive
	}

	rspParams.XResultCount = strconv.Itoa(len(rspParams.Payload))
	rspParams.XTotalCount = rspParams.XResultCount

	return rspParams
}

func SonataGenerateInvGetResponse(reqParams *invapi.ProductGetParams) *invapi.ProductGetOK {
	rspParams := invapi.NewProductGetOK()

	prodItem := &invmod.Product{}
	prodItem.ID = &reqParams.ProductID
	prodItem.Status = invmod.ProductStatusActive
	prodItem.LastUpdateDate.Scan(time.Now())
	prodItem.ProductSpecification = &invmod.ProductSpecificationRef{}
	prodItem.ProductSpecification.ID = sonata.NewString("ELineSpec")

	rspParams.Payload = prodItem

	return rspParams
}

func SonataGeneratePoqCreateResponse(reqParams *poqapi.ProductOfferingQualificationCreateParams) *poqapi.ProductOfferingQualificationCreateCreated {
	rspPoq := &poqmod.ProductOfferingQualification{}

	for i := 0; i < len(reqParams.ProductOfferingQualification.ProductOfferingQualificationItem); i++ {
		rspPoq.ProductOfferingQualificationItem = append(rspPoq.ProductOfferingQualificationItem, &poqmod.ProductOfferingQualificationItem{})
	}

	// Most fields can been filled by request parameters
	reqData, err := reqParams.ProductOfferingQualification.MarshalBinary()
	if err != nil {
		fmt.Println("reqParams MarshalBinary", err)
	}
	fmt.Println("reqData", string(reqData))

	// fixup some fields
	reqJson, err := simplejson.NewJson(reqData)
	if err != nil {
		fmt.Println("reqData NewJson", err)
	}
	reqJson.Del("requestedResponseDate")

	reqDataFixed, err := reqJson.MarshalJSON()
	if err != nil {
		fmt.Println("reqJson MarshalJSON", err)
	}

	err = rspPoq.UnmarshalBinary(reqDataFixed)
	if err != nil {
		fmt.Println("rspParams UnmarshalBinary", err)
	}

	// Response generated fields
	poqID := uuid.New().String()
	rspPoq.ID = &poqID
	rspPoq.State = poqmod.ProductOfferingQualificationStateTypeDone
	rspPoq.RequestedResponseDate.Scan(reqParams.ProductOfferingQualification.RequestedResponseDate.String() + "T00:00")
	rspPoq.ExpectedResponseDate.Scan(time.Now())
	rspPoq.EffectiveQualificationCompletionDate.Scan(time.Now())

	for _, poqItem := range rspPoq.ProductOfferingQualificationItem {
		poqItem.State = poqmod.ProductOfferingQualificationItemStateTypeDone
		poqItem.InstallationInterval = &poqmod.TimeInterval{}
		installVal := int32(5)
		poqItem.InstallationInterval.TimeUnit = poqmod.TimeUnitCalendarMinutes
		poqItem.InstallationInterval.Amount = &installVal
		poqItem.GuaranteedUntilDate.Scan(time.Now())
		poqItem.ServiceabilityConfidence = poqmod.ServiceabilityColorGreen
		//poqItem.ServiceConfidenceReason = ""
	}

	rspParams := poqapi.NewProductOfferingQualificationCreateCreated()
	rspParams.Payload = rspPoq

	return rspParams
}

func SonataGeneratePoqFindResponse(reqParams *poqapi.ProductOfferingQualificationFindParams) *poqapi.ProductOfferingQualificationFindOK {
	rspParams := poqapi.NewProductOfferingQualificationFindOK()

	poq1 := &poqmod.ProductOfferingQualificationFind{}
	poq1.ID = uuid.New().String()
	poq1.State = poqmod.ProductOfferingQualificationStateTypeDone
	rspParams.Payload = append(rspParams.Payload, poq1)

	poq2 := &poqmod.ProductOfferingQualificationFind{}
	poq2.ID = uuid.New().String()
	poq2.State = poqmod.ProductOfferingQualificationStateTypeDone
	rspParams.Payload = append(rspParams.Payload, poq2)

	return rspParams
}

func SonataGeneratePoqGetResponse(reqParams *poqapi.ProductOfferingQualificationGetParams) *poqapi.ProductOfferingQualificationGetOK {
	rspParams := poqapi.NewProductOfferingQualificationGetOK()
	rspParams.Payload = &poqmod.ProductOfferingQualification{}
	rspParams.Payload.ID = &reqParams.ProductOfferingQualificationID
	rspParams.Payload.State = poqmod.ProductOfferingQualificationStateTypeDone
	poqItem1 := &poqmod.ProductOfferingQualificationItem{}
	itemId1 := "1"
	poqItem1.ID = &itemId1
	poqItem1.State = poqmod.ProductOfferingQualificationItemStateTypeDone
	rspParams.Payload.ProductOfferingQualificationItem = append(rspParams.Payload.ProductOfferingQualificationItem, poqItem1)

	return rspParams
}

func SonataGenerateQuoteCreateResponse(reqParams *quoapi.QuoteCreateParams) *quoapi.QuoteCreateCreated {
	rspQuote := &quomod.Quote{}

	for i := 0; i < len(reqParams.Quote.QuoteItem); i++ {
		quoteItem := &quomod.QuoteItem{}
		rspQuote.QuoteItem = append(rspQuote.QuoteItem, quoteItem)
	}

	// Most fields can been filled by request parameters
	reqData, err := reqParams.Quote.MarshalBinary()
	if err != nil {
		fmt.Println("reqParams MarshalBinary", err)
	}
	//fmt.Println("reqData", string(reqData))
	err = rspQuote.UnmarshalBinary(reqData)
	if err != nil {
		fmt.Println("rspParams UnmarshalBinary", err)
	}

	// Response generated fields
	rspQuote.ID = uuid.New().String()
	rspQuote.ExpectedQuoteCompletionDate.Scan(time.Now())
	rspQuote.QuoteDate.Scan(time.Now())
	rspQuote.State = quomod.QuoteStateTypeREADY

	for _, quoteItem := range rspQuote.QuoteItem {
		quoteItem.State = quomod.QuoteItemStateTypeREADY

		var uniSpec *cmnmod.UNISpec
		var lineSpec *cmnmod.ELineSpec
		if quoteItem.Product != nil && quoteItem.Product.ProductSpecification != nil {
			specDescVal := quoteItem.Product.ProductSpecification.Describing()
			specDescType := specDescVal.AtType()
			specOk := false
			if specDescType == "UNISpec" {
				uniSpec, specOk = specDescVal.(*cmnmod.UNISpec)
			} else if specDescType == "ELineSpec" {
				lineSpec, specOk = specDescVal.(*cmnmod.ELineSpec)
			}
			if !specOk {
				fmt.Println("Describing to Product Specification error")
			}
		}

		itemPrice := &quomod.QuotePrice{}
		priName := "RENTAL"
		itemPrice.Name = &priName
		itemPrice.PriceType = quomod.PriceTypeRECURRING
		itemPrice.Price = &quomod.Price{}
		curUnit := "USD"
		priVal := float32(0)
		if uniSpec != nil {
			if len(uniSpec.PhysicalLayer) > 0 {
				if strings.HasPrefix(string(uniSpec.PhysicalLayer[0]), "1000BASE") {
					priVal = 100
				} else if strings.HasPrefix(string(uniSpec.PhysicalLayer[0]), "10GBASE") {
					priVal = 200
				} else {
					priVal = 50
				}
			}
		} else if lineSpec != nil {
			bwCir := lineSpec.UNIIngressBWProfile[0].Cir
			priVal = 3 * float32(*bwCir.Amount)
		}
		itemPrice.Price.PreTaxAmount = &quomod.Money{Unit: &curUnit, Value: &priVal}
		itemPrice.Price.PriceRange = &quomod.PriceRange{MaxPreTaxAmount: itemPrice.Price.PreTaxAmount, MinPreTaxAmount: itemPrice.Price.PreTaxAmount}
		itemPrice.RecurringChargePeriod = quomod.ChargePeriodDAY
		quoteItem.QuoteItemPrice = append(quoteItem.QuoteItemPrice, itemPrice)
	}

	rspParams := quoapi.NewQuoteCreateCreated()
	rspParams.Payload = rspQuote

	return rspParams
}

func SonataGenerateQuoteFindResponse(reqParams *quoapi.QuoteFindParams) *quoapi.QuoteFindOK {
	rspParams := quoapi.NewQuoteFindOK()

	quote1 := &quomod.QuoteFind{}
	quote1.ID = uuid.New().String()
	quote1.State = quomod.QuoteStateTypeREADY
	rspParams.Payload = append(rspParams.Payload, quote1)

	quote2 := &quomod.QuoteFind{}
	quote2.ID = uuid.New().String()
	quote2.State = quomod.QuoteStateTypeREADY
	rspParams.Payload = append(rspParams.Payload, quote2)

	return rspParams
}

func SonataGenerateQuoteGetResponse(reqParams *quoapi.QuoteGetParams) *quoapi.QuoteGetOK {
	rspParams := quoapi.NewQuoteGetOK()

	quote := &quomod.Quote{}
	quote.ID = reqParams.ID
	quote.InstantSyncQuoting = true
	quote.State = quomod.QuoteStateTypeREADY
	lineItem := &quomod.QuoteItem{}
	lineId := "1"
	lineItem.ID = &lineId
	lineItem.State = quomod.QuoteItemStateTypeREADY
	quote.QuoteItem = append(quote.QuoteItem, lineItem)
	quote.QuoteDate.Scan(time.Now())

	rspParams.Payload = quote

	return rspParams
}

func SonataGenerateOrderCreateResponse(reqParams *ordapi.ProductOrderCreateParams) *ordapi.ProductOrderCreateCreated {
	rspOrder := &ordmod.ProductOrder{}

	for i := 0; i < len(reqParams.ProductOrder.OrderItem); i++ {
		rspOrder.OrderItem = append(rspOrder.OrderItem, &ordmod.OrderItem{})
	}

	// Most fields can been filled by request parameters
	reqData, err := reqParams.ProductOrder.MarshalBinary()
	if err != nil {
		fmt.Println("reqParams MarshalBinary", err)
	}
	//fmt.Println("reqData", string(reqData))
	err = rspOrder.UnmarshalBinary(reqData)
	if err != nil {
		fmt.Println("rspParams UnmarshalBinary", err)
	}

	// Response generated fields
	ordID := uuid.New().String()
	rspOrder.ID = &ordID
	rspOrder.State = ordmod.ProductOrderStateTypeCompleted
	rspOrder.OrderDate = &strfmt.DateTime{}
	rspOrder.OrderDate.Scan(time.Now())
	rspOrder.CompletionDate.Scan(time.Now())

	for _, orderItem := range rspOrder.OrderItem {
		orderItem.State = ordmod.ProductOrderItemStateTypeCompleted
	}

	rspParams := ordapi.NewProductOrderCreateCreated()
	rspParams.Payload = rspOrder

	return rspParams
}

func SonataGenerateOrderFindResponse(reqParams *ordapi.ProductOrderFindParams) *ordapi.ProductOrderFindOK {
	rspParams := ordapi.NewProductOrderFindOK()

	order1 := &ordmod.ProductOrderSummary{}
	id1 := uuid.New().String()
	order1.ID = &id1
	order1.State = ordmod.ProductOrderStateTypeCompleted
	rspParams.Payload = append(rspParams.Payload, order1)

	order2 := &ordmod.ProductOrderSummary{}
	id2 := uuid.New().String()
	order2.ID = &id2
	order2.State = ordmod.ProductOrderStateTypeCompleted
	rspParams.Payload = append(rspParams.Payload, order2)

	return rspParams
}

func SonataGenerateOrderGetResponse(reqParams *ordapi.ProductOrderGetParams) *ordapi.ProductOrderGetOK {
	rspParams := ordapi.NewProductOrderGetOK()

	order := &ordmod.ProductOrder{}
	order.ID = &reqParams.ProductOrderID
	order.State = ordmod.ProductOrderStateTypeCompleted
	lineItem := &ordmod.OrderItem{}
	lineId := "1"
	lineItem.ID = &lineId
	lineItem.State = ordmod.ProductOrderItemStateTypeCompleted
	order.OrderItem = append(order.OrderItem, lineItem)

	rspParams.Payload = order

	return rspParams
}
