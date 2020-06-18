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
	RunEnv string

	BuyerAddr  string
	BuyerName  string
	SellerAddr string
	SellerName string

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
	if o.Param.RunEnv == "dev" {
		lineItem.ProdOfferID = "29f855fb-4760-4e77-877e-3318906ee4bc"
		lineItem.SrcLocationID = "5ae7e56bbbc9a8001231fa5d"
		lineItem.DstLocationID = "5ae7e56bbbc9a8001231fa5d"
	} else if o.Param.RunEnv == "stage" {
		lineItem.ProdOfferID = "29f855fb-4760-4e77-877e-3318906ee4bc"
		lineItem.SrcLocationID = "5db662305d545c000bc68aaf"
		lineItem.DstLocationID = "5db662305d545c000bc68aaf"
	} else {
		lineItem.ProdOfferID = "29f855fb-4760-4e77-877e-3318906ee4bc"
		lineItem.SrcLocationID = "5ae7e56bbbc9a8001231fa5d"
		lineItem.DstLocationID = "5ae7e56bbbc9a8001231fa5d"
	}
	lineItem.ProductID = o.Param.ExistProductID

	lineItem.BillingParams = &models.BillingParams{}
	lineItem.BillingParams.BillingType = quoteParam.BillingType
	lineItem.BillingParams.PaymentType = quoteParam.PaymentType
	lineItem.BillingParams.StartTime = o.Param.StartTime
	lineItem.BillingParams.EndTime = o.Param.EndTime

	quoteParam.ELineItems = append(quoteParam.ELineItems, lineItem)

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
	connParam.StaticParam.BuyerProductID = o.Param.BuyerProductID
	if o.Param.RunEnv == "dev" {
		connParam.StaticParam.ProductOfferingID = "29f855fb-4760-4e77-877e-3318906ee4bc"

		connParam.StaticParam.SrcRegion = "KR"
		connParam.StaticParam.SrcCity = "9d242983f7a504eebb3eb478"
		connParam.StaticParam.SrcDataCenter = "5ae7e56bbbc9a8001231fa5d"
		connParam.StaticParam.SrcCompanyName = "5d02fa08a5b531000a764046"
		connParam.StaticParam.SrcPort = "5d098e7e96f045000a4164fa"

		connParam.StaticParam.DstRegion = "JP"
		connParam.StaticParam.DstCity = "9d242983f7a504eebb3eb478"
		connParam.StaticParam.DstDataCenter = "5ae7e56bbbc9a8001231fa5d"
		connParam.StaticParam.DstCompanyName = "5d02fa08a5b531000a764046"
		connParam.StaticParam.DstPort = "5d269f1760e409000ad83c58"
	} else if o.Param.RunEnv == "stage" {
		connParam.StaticParam.ProductOfferingID = "29f855fb-4760-4e77-877e-3318906ee4bc"

		connParam.StaticParam.SrcRegion = "co"
		connParam.StaticParam.SrcCity = "5d7593c0a266e1000afd5335"
		connParam.StaticParam.SrcDataCenter = "5db662305d545c000bc68aaf"
		connParam.StaticParam.SrcCompanyName = "5eeacca2a73b53001545aa60"
		connParam.StaticParam.SrcPort = "5eeacf65deb2d40014f2c2b5"

		connParam.StaticParam.DstRegion = "co"
		connParam.StaticParam.DstCity = "5d7593c0a266e1000afd5335"
		connParam.StaticParam.DstDataCenter = "5db662305d545c000bc68aaf"
		connParam.StaticParam.DstCompanyName = "5eeacca2a73b53001545aa60"
		connParam.StaticParam.DstPort = "5eead050def54c001413f04b"
	} else {
		connParam.StaticParam.ProductOfferingID = "29f855fb-4760-4e77-877e-3318906ee4bc"

		connParam.StaticParam.SrcRegion = "KR"
		connParam.StaticParam.SrcCity = "9d242983f7a504eebb3eb478"
		connParam.StaticParam.SrcDataCenter = "5ae7e56bbbc9a8001231fa5d"
		connParam.StaticParam.SrcCompanyName = "5d02fa08a5b531000a764046"
		connParam.StaticParam.SrcPort = "5d098e7e96f045000a4164fa"

		connParam.StaticParam.DstRegion = "JP"
		connParam.StaticParam.DstCity = "9d242983f7a504eebb3eb478"
		connParam.StaticParam.DstDataCenter = "5ae7e56bbbc9a8001231fa5d"
		connParam.StaticParam.DstCompanyName = "5d02fa08a5b531000a764046"
		connParam.StaticParam.DstPort = "5d269f1760e409000ad83c58"
	}

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

func (o *ProductOrder) GetOrderInfoBySellerAndOrderId(seller string, orderId string) (*models.ProtoOrderInfo, error) {
	req := order_api.NewGetOrderInfoParams()
	req.SellerAddress = &seller
	req.OrderID = &orderId

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
