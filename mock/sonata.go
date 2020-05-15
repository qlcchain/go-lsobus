package mock

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/iixlabs/virtual-lsobus/sonata"
	sitapi "github.com/iixlabs/virtual-lsobus/sonata/site/client/geographic_site"
	sitmod "github.com/iixlabs/virtual-lsobus/sonata/site/models"

	"github.com/bitly/go-simplejson"

	"github.com/go-openapi/strfmt"

	invapi "github.com/iixlabs/virtual-lsobus/sonata/inventory/client/product"
	invmod "github.com/iixlabs/virtual-lsobus/sonata/inventory/models"
	ordapi "github.com/iixlabs/virtual-lsobus/sonata/order/client/product_order"
	ordmod "github.com/iixlabs/virtual-lsobus/sonata/order/models"
	poqapi "github.com/iixlabs/virtual-lsobus/sonata/poq/client/product_offering_qualification"
	poqmod "github.com/iixlabs/virtual-lsobus/sonata/poq/models"
	quoapi "github.com/iixlabs/virtual-lsobus/sonata/quote/client/quote"
	quomod "github.com/iixlabs/virtual-lsobus/sonata/quote/models"
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

func SonataGenerateInvFindResponse(reqParams *invapi.ProductFindParams) *invapi.ProductFindOK {
	rspParams := invapi.NewProductFindOK()

	if reqParams.ProductSpecificationID == nil || *reqParams.ProductSpecificationID == "UNISpec" {
		prodItem1 := &invmod.ProductSummary{}
		rspParams.Payload = append(rspParams.Payload, prodItem1)

		prodID1 := "PCCW-Port-111"
		prodItem1.ID = &prodID1
		prodItem1.BuyerProductID = "CBC-Port-111"
		prodItem1.ProductSpecification = &invmod.ProductSpecificationSummary{}
		prodSpecID1 := "UNISpec"
		prodItem1.ProductSpecification.ID = &prodSpecID1
		prodItem1.StartDate.Scan(time.Now().Add(time.Duration(rand.Intn(48)) * time.Hour))
		prodItem1.Status = invmod.ProductStatusActive

		prodItem2 := &invmod.ProductSummary{}
		rspParams.Payload = append(rspParams.Payload, prodItem2)

		prodID2 := "PCCW-Port-222"
		prodItem2.ID = &prodID2
		prodItem2.BuyerProductID = "CBC-Port-222"
		prodItem2.ProductSpecification = &invmod.ProductSpecificationSummary{}
		prodSpecID2 := "UNISpec"
		prodItem2.ProductSpecification.ID = &prodSpecID2
		prodItem2.StartDate.Scan(time.Now().Add(time.Duration(rand.Intn(48)) * time.Hour))
		prodItem2.Status = invmod.ProductStatusActive
	}

	if reqParams.ProductSpecificationID == nil || *reqParams.ProductSpecificationID == "ELineSpec" {
		prodItem1 := &invmod.ProductSummary{}
		rspParams.Payload = append(rspParams.Payload, prodItem1)

		prodID1 := "PCCW-Connection-111"
		prodItem1.ID = &prodID1
		prodItem1.BuyerProductID = "CBC-Connection-111"
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
	poqID := "PCCW-POQ-111"
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
	rspQuote.ID = "PCCW-Quote-111"
	rspQuote.ExpectedQuoteCompletionDate.Scan(time.Now())
	rspQuote.QuoteDate.Scan(time.Now())
	rspQuote.State = quomod.QuoteStateTypeREADY

	for _, quoteItem := range rspQuote.QuoteItem {
		quoteItem.State = quomod.QuoteItemStateTypeREADY

		itemPrice := &quomod.QuotePrice{}
		priName := "RENTAL"
		itemPrice.Name = &priName
		itemPrice.PriceType = quomod.PriceTypeRECURRING
		itemPrice.Price = &quomod.Price{}
		curUnit := "USA"
		priVal := float32(12.34)
		itemPrice.Price.PreTaxAmount = &quomod.Money{Unit: &curUnit, Value: &priVal}
		itemPrice.Price.PriceRange = &quomod.PriceRange{MaxPreTaxAmount: itemPrice.Price.PreTaxAmount, MinPreTaxAmount: itemPrice.Price.PreTaxAmount}
		itemPrice.RecurringChargePeriod = quomod.ChargePeriodDAY
		quoteItem.QuoteItemPrice = append(quoteItem.QuoteItemPrice, itemPrice)
	}

	rspParams := quoapi.NewQuoteCreateCreated()
	rspParams.Payload = rspQuote

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
	ordID := "PCCW-Order-111"
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
