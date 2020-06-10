package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/bitly/go-simplejson"

	"github.com/qlcchain/go-lsobus/cmd/agent/client"
	"github.com/qlcchain/go-lsobus/cmd/agent/client/orchestra_api"
	"github.com/qlcchain/go-lsobus/cmd/agent/client/order_api"
	"github.com/qlcchain/go-lsobus/cmd/agent/models"
)

type ProductParam struct {
	BuyerAddr      string
	BuyerName      string
	SellerAddr     string
	SellerName     string
	ProductOfferID string

	SrcPort  string
	SrcLocID string
	DstPort  string
	DstLocID string

	Name      string
	Bandwidth int32
	CosName   string
	StartTime int64
	EndTime   int64

	BuyerProductID string

	ExistProductID       string
	ExistConnectionParam *models.ProtoConnectionParam
}

type ProductOrder struct {
	Param  *ProductParam
	Client *client.TypesProto

	QuoteID       string
	QuoteItemID   string
	QuoteCurrency string
	QuotePrice    float64

	InternalID string

	QuoteReq *models.ProtoOrchestraCommonRequest
	QuoteRsp *models.ProtoOrchestraCommonResponse

	CreateOrderReq    *models.ProtoCreateOrderParam
	ChangeOrderReq    *models.ProtoChangeOrderParam
	TerminateOrderReq *models.ProtoTerminateOrderParam
	CommonOrderRsp    *models.ProtoOrderID

	OrderInfoRsp *models.ProtoOrderInfo
}

func (o *ProductOrder) Init() error {
	retUrl, err := url.Parse(lsobusUrl)
	if err != nil {
		return fmt.Errorf("url parse err %s", err)
	}

	tranCfg := client.TransportConfig{
		Schemes:  []string{retUrl.Scheme},
		Host:     retUrl.Host,
		BasePath: "/",
	}
	o.Client = client.NewHTTPClientWithConfig(nil, &tranCfg)
	return nil
}

func (o *ProductOrder) CreateQuote(action string) error {
	reqDataJson, err := o.buildQuoteReqJson(action)
	if err != nil {
		return err
	}

	req := orchestra_api.NewExecCreateParams()
	req.Body = &models.ProtoOrchestraCommonRequest{}
	req.Body.Action = "ExecQuoteCreate"
	req.Body.Data = string(reqDataJson)
	rsp, err := o.Client.OrchestraAPI.ExecCreate(req)
	if err != nil {
		return err
	}

	o.QuoteRsp = rsp.GetPayload()

	err = o.parseQuoteRspJson()
	if err != nil {
		return err
	}

	fmt.Printf("Create Quote is OK, QuoteID %s, QuoteItemID %s, Price: %f/%s\n",
		o.QuoteID, o.QuoteItemID, o.QuotePrice, o.QuoteCurrency)

	return nil
}

func (o *ProductOrder) buildQuoteReqJson(action string) ([]byte, error) {
	quoteParam := &models.OrderParams{}
	quoteParam.ExternalID = uuid.New().String()
	quoteParam.ProjectID = "CBC-AGENT"

	quoteParam.OrderActivity = action
	quoteParam.Buyer = &models.PartnerParams{Name: o.Param.BuyerName, ID: o.Param.BuyerAddr}
	quoteParam.Seller = &models.PartnerParams{Name: o.Param.SellerName, ID: o.Param.SellerAddr}
	quoteParam.BillingType = "DOD"
	quoteParam.PaymentType = "INVOICE"

	lineItem := &models.ELineItemParams{}
	lineItem.ItemID = uuid.New().String()
	lineItem.Action = action
	lineItem.Name = o.Param.Name
	lineItem.CosName = strings.ToUpper(o.Param.CosName)
	lineItem.Bandwidth = uint(o.Param.Bandwidth)
	lineItem.ProdOfferID = o.Param.ProductOfferID
	//lineItem.SrcPortID = o.Param.SrcPort
	//lineItem.DstPortID = o.Param.DstPort
	lineItem.SrcLocationID = o.Param.SrcLocID
	lineItem.DstLocationID = o.Param.DstLocID
	//lineItem.BuyerProductID = o.Param.BuyerProductID
	lineItem.ProductID = o.Param.ExistProductID

	lineItem.BillingParams = &models.BillingParams{}
	lineItem.BillingParams.BillingType = quoteParam.BillingType
	lineItem.BillingParams.PaymentType = quoteParam.PaymentType
	lineItem.BillingParams.StartTime = o.Param.StartTime
	lineItem.BillingParams.EndTime = o.Param.EndTime

	quoteParam.ELineItems = append(quoteParam.ELineItems, lineItem)

	/*
		quoteJson, err := simplejson.NewJson([]byte())
		if err != nil {
			return nil, err
		}

		itemJson := quoteJson.Get("quoteItem").GetIndex(0)
		descJson := itemJson.GetPath("product", "productSpecification", "describing")
		descJson.Set("bandwidth", o.Param.Bandwidth)
		descJson.Set("classOfService", o.Param.CosName)
		descJson.Set("srcLocationId", o.Param.SrcLocID)
		descJson.Set("destLocationId", o.Param.DstLocID)
		descJson.Set("startedAt", time.Unix(o.Param.StartTime, 0).Format("2020-06-04T18:42:13.000+08:00"))
		descJson.Set("terminatedAt", time.Unix(o.Param.EndTime, 0).Format("2020-06-04T18:42:13.000+08:00"))
	*/

	return json.Marshal(quoteParam)
}

func (o *ProductOrder) parseQuoteRspJson() error {
	quoteJson, err := simplejson.NewJson([]byte(o.QuoteRsp.Data))
	if err != nil {
		return err
	}

	o.QuoteID, err = quoteJson.Get("id").String()
	if err != nil {
		return err
	}

	quoteItem := quoteJson.Get("quoteItem").GetIndex(0)

	o.QuoteItemID, err = quoteItem.Get("id").String()
	if err != nil {
		return err
	}

	quotePrice := quoteItem.Get("preCalculatedPrice").Get("price")
	o.QuoteCurrency, err = quotePrice.Get("preTaxAmount").Get("unit").String()
	if err != nil {
		return err
	}
	o.QuotePrice, err = quotePrice.Get("preTaxAmount").Get("value").Float64()
	if err != nil {
		return err
	}

	return nil
}

func (o *ProductOrder) CreateNewOrder() error {
	if o.QuoteRsp == nil {
		return errors.New("quote not exist")
	}

	o.CreateOrderReq = &models.ProtoCreateOrderParam{}
	o.CreateOrderReq.Buyer = &models.ProtoUser{Name: o.Param.BuyerName, Address: o.Param.BuyerAddr}
	o.CreateOrderReq.Seller = &models.ProtoUser{Name: o.Param.SellerName, Address: o.Param.SellerAddr}

	connParam := &models.ProtoConnectionParam{}
	connParam.StaticParam = &models.ProtoConnectionStaticParam{}
	connParam.StaticParam.ProductOfferingID = o.Param.ProductOfferID

	connParam.StaticParam.BuyerProductID = o.Param.BuyerProductID

	connParam.StaticParam.SrcRegion = "KR"
	connParam.StaticParam.SrcCity = "9d242983f7a504eebb3eb478"
	connParam.StaticParam.SrcDataCenter = "5ae7e56bbbc9a8001231fa5d"
	connParam.StaticParam.SrcCompanyName = "5d02fa08a5b531000a764046"
	connParam.StaticParam.SrcPort = o.Param.SrcPort

	connParam.StaticParam.DstRegion = "JP"
	connParam.StaticParam.DstCity = "9d242983f7a504eebb3eb478"
	connParam.StaticParam.DstDataCenter = "5ae7e56bbbc9a8001231fa5d"
	connParam.StaticParam.DstCompanyName = "5d02fa08a5b531000a764046"
	connParam.StaticParam.DstPort = o.Param.DstPort

	connParam.DynamicParam = &models.ProtoConnectionDynamicParam{}
	connParam.DynamicParam.ItemID = uuid.New().String()
	connParam.DynamicParam.ConnectionName = o.Param.Name
	connParam.DynamicParam.Bandwidth = fmt.Sprintf("%d Mbps", o.Param.Bandwidth)
	connParam.DynamicParam.ServiceClass = o.Param.CosName
	connParam.DynamicParam.QuoteID = o.QuoteID
	connParam.DynamicParam.QuoteItemID = o.QuoteItemID
	connParam.DynamicParam.Currency = o.QuoteCurrency
	connParam.DynamicParam.Price = float32(o.QuotePrice)
	connParam.DynamicParam.BillingType = "DOD"
	connParam.DynamicParam.PaymentType = "invoice"
	connParam.DynamicParam.BillingUnit = "day"
	connParam.DynamicParam.StartTime = strconv.Itoa(int(o.Param.StartTime))
	connParam.DynamicParam.EndTime = strconv.Itoa(int(o.Param.EndTime))

	o.CreateOrderReq.ConnectionParam = append(o.CreateOrderReq.ConnectionParam, connParam)

	req := order_api.NewCreateOrderParams()
	req.Body = o.CreateOrderReq

	rsp, err := o.Client.OrderAPI.CreateOrder(req)
	if err != nil {
		return err
	}

	o.CommonOrderRsp = rsp.GetPayload()
	o.InternalID = o.CommonOrderRsp.InternalID

	fmt.Printf("Create New Order is OK, InternalID %s\n", o.InternalID)

	return nil
}

func (o *ProductOrder) CreateChangeOrder() error {
	if o.Param.ExistConnectionParam == nil {
		return errors.New("exist product not exist")
	}

	if o.QuoteRsp == nil {
		return errors.New("quote not exist")
	}

	o.ChangeOrderReq = &models.ProtoChangeOrderParam{}
	o.ChangeOrderReq.Buyer = &models.ProtoUser{Name: o.Param.BuyerName, Address: o.Param.BuyerAddr}
	o.ChangeOrderReq.Seller = &models.ProtoUser{Name: o.Param.SellerName, Address: o.Param.SellerAddr}

	connParam := &models.ProtoChangeConnectionParam{}
	connParam.ProductID = o.Param.ExistProductID
	connParam.DynamicParam = &models.ProtoConnectionDynamicParam{}
	connParam.DynamicParam.ItemID = uuid.New().String()
	connParam.DynamicParam.ConnectionName = o.Param.Name
	connParam.DynamicParam.Bandwidth = fmt.Sprintf("%d Mbps", o.Param.Bandwidth)
	connParam.DynamicParam.ServiceClass = o.Param.CosName
	connParam.DynamicParam.QuoteID = o.QuoteID
	connParam.DynamicParam.QuoteItemID = o.QuoteItemID
	connParam.DynamicParam.Currency = o.QuoteCurrency
	connParam.DynamicParam.Price = float32(o.QuotePrice)
	connParam.DynamicParam.BillingType = "DOD"
	connParam.DynamicParam.PaymentType = "invoice"
	connParam.DynamicParam.BillingUnit = "day"
	connParam.DynamicParam.StartTime = strconv.Itoa(int(o.Param.StartTime))
	connParam.DynamicParam.EndTime = strconv.Itoa(int(o.Param.EndTime))

	o.ChangeOrderReq.ChangeConnectionParam = append(o.ChangeOrderReq.ChangeConnectionParam, connParam)

	req := order_api.NewChangeOrderParams()
	req.Body = o.ChangeOrderReq

	rsp, err := o.Client.OrderAPI.ChangeOrder(req)
	if err != nil {
		return err
	}

	o.CommonOrderRsp = rsp.GetPayload()
	o.InternalID = o.CommonOrderRsp.InternalID

	fmt.Printf("Create Change Order is OK, InternalID %s\n", o.InternalID)

	return nil
}

func (o *ProductOrder) CreateTerminateOrder() error {
	if o.Param.ExistConnectionParam == nil {
		return errors.New("exist product not exist")
	}

	if o.QuoteRsp == nil {
		return errors.New("quote not exist")
	}

	o.TerminateOrderReq = &models.ProtoTerminateOrderParam{}
	o.TerminateOrderReq.Buyer = &models.ProtoUser{Name: o.Param.BuyerName, Address: o.Param.BuyerAddr}
	o.TerminateOrderReq.Seller = &models.ProtoUser{Name: o.Param.SellerName, Address: o.Param.SellerAddr}

	connParam := &models.ProtoTerminateConnectionParam{}
	connParam.ProductID = o.Param.ExistProductID
	connParam.DynamicParam = &models.ProtoConnectionDynamicParam{}
	connParam.DynamicParam.ItemID = uuid.New().String()
	connParam.DynamicParam.ConnectionName = o.Param.Name
	connParam.DynamicParam.Bandwidth = fmt.Sprintf("%d Mbps", o.Param.Bandwidth)
	connParam.DynamicParam.ServiceClass = o.Param.CosName
	connParam.DynamicParam.QuoteID = o.QuoteID
	connParam.DynamicParam.QuoteItemID = o.QuoteItemID
	connParam.DynamicParam.Currency = o.QuoteCurrency
	connParam.DynamicParam.Price = float32(o.QuotePrice)
	connParam.DynamicParam.BillingType = "DOD"
	connParam.DynamicParam.PaymentType = "invoice"
	connParam.DynamicParam.BillingUnit = "day"
	connParam.DynamicParam.StartTime = strconv.Itoa(int(o.Param.StartTime))
	connParam.DynamicParam.EndTime = strconv.Itoa(int(o.Param.EndTime))

	o.TerminateOrderReq.TerminateConnectionParam = append(o.TerminateOrderReq.TerminateConnectionParam, connParam)

	req := order_api.NewTerminateOrderParams()
	req.Body = o.TerminateOrderReq

	rsp, err := o.Client.OrderAPI.TerminateOrder(req)
	if err != nil {
		return err
	}

	o.CommonOrderRsp = rsp.GetPayload()
	o.InternalID = o.CommonOrderRsp.InternalID

	fmt.Printf("Create Terminate Order is OK, InternalID %s\n", o.InternalID)

	return nil
}

func (o *ProductOrder) GetOrderInfo() error {
	req := order_api.NewGetOrderInfoParams()
	req.InternalID = &(o.CommonOrderRsp.InternalID)

	rsp, err := o.Client.OrderAPI.GetOrderInfo(req)
	if err != nil {
		return err
	}

	o.OrderInfoRsp = rsp.GetPayload()

	fmt.Printf("Get Order Info is OK, ContractState %s, OrderID %s, OrderState %s\n",
		o.OrderInfoRsp.ContractState, o.OrderInfoRsp.OrderID, o.OrderInfoRsp.OrderState)

	return nil
}

func (o *ProductOrder) GetOrderInfoByInternalId(internalId string) (*models.ProtoOrderInfo, error) {
	req := order_api.NewGetOrderInfoParams()
	req.InternalID = &internalId

	rsp, err := o.Client.OrderAPI.GetOrderInfo(req)
	if err != nil {
		return nil, err
	}
	orderInfo := rsp.GetPayload()

	return orderInfo, nil
}

func (o *ProductOrder) CheckOrderStatus() error {
	var err error
	var lastContractState string
	var lastOrderState string

	for {
		time.Sleep(10 * time.Second)

		err = o.GetOrderInfo()
		if err != nil {
			fmt.Printf("wait to get order info, err %s\n", err)
			continue
		}
		if o.OrderInfoRsp == nil {
			fmt.Printf("OrderInfo not exist\n")
			continue
		}

		if o.OrderInfoRsp.ContractState != lastContractState {
			fmt.Printf("Update ContractState %s\n", o.OrderInfoRsp.ContractState)
			lastContractState = o.OrderInfoRsp.ContractState
		}

		if o.OrderInfoRsp.OrderState != lastOrderState {
			fmt.Printf("Update OrderState %s\n", o.OrderInfoRsp.OrderState)
			lastOrderState = o.OrderInfoRsp.OrderState
		}

		if o.OrderInfoRsp.OrderState == "complete" {
			break
		}
	}

	return nil
}
