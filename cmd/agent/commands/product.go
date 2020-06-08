package commands

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/qlcchain/go-lsobus/cmd/agent/client"
	"github.com/qlcchain/go-lsobus/cmd/agent/client/orchestra_api"
	"github.com/qlcchain/go-lsobus/cmd/agent/client/order_api"
	"github.com/qlcchain/go-lsobus/cmd/agent/models"
	"net/url"
	"strings"
)

type ProductParam struct {
	BuyerAddr string
	BuyerName string
	SellerAddr string
	SellerName string
	ProductOfferID string

	SrcPort string
	SrcLocID string
	DstPort string
	DstLocID string

	Name string
	Bandwidth int32
	CosName string
	StartTime int64
	EndTime int64
}

type ProductOrder struct {
	Param *ProductParam
	Client *client.TypesProto

	QuoteID     string
	QuoteItemID string
	InternalID  string

	QuoteReq *models.ProtoOrchestraCommonRequest
	QuoteRsp *models.ProtoOrchestraCommonResponse

	OrderReq *models.ProtoCreateOrderParam
	OrderRsp *models.ProtoOrderID

	OrderInfoRsp *models.ProtoOrderInfo
}

func (o *ProductOrder) Init() error {
	retUrl, err := url.Parse(lsobusUrl)
	if err != nil {
		return fmt.Errorf("url parse err %s", err)
	}

	tranCfg := client.TransportConfig{
		Schemes: []string{retUrl.Scheme},
		Host: retUrl.Host,
		BasePath: "/",
	}
	o.Client = client.NewHTTPClientWithConfig(nil, &tranCfg)
	return nil
}

func (o *ProductOrder) CreateQuote() error {
	reqDataJson, err:= o.buildQuoteReqJson()
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

	fmt.Printf("Create Quote is OK, QuoteID %s, QuoteItemID %s\n", o.QuoteID, o.QuoteItemID)

	return nil
}

func (o *ProductOrder) buildQuoteReqJson() ([]byte, error) {
	quoteParam := &models.OrderParams{}
	quoteParam.OrderActivity = "INSTALL"
	quoteParam.Buyer = &models.PartnerParams{Name: o.Param.BuyerName, ID: o.Param.BuyerAddr}
	quoteParam.Seller = &models.PartnerParams{Name: o.Param.SellerName, ID: o.Param.SellerAddr}

	lineItem := &models.ELineItemParams{}
	lineItem.ItemID = "1"
	lineItem.Action = "INSTALL"
	lineItem.Name = o.Param.Name
	lineItem.CosName = strings.ToUpper(o.Param.CosName)
	lineItem.Bandwidth = uint(o.Param.Bandwidth)
	lineItem.ProdOfferID = o.Param.ProductOfferID
	//lineItem.SrcPortID = o.Param.SrcPort
	//lineItem.DstPortID = o.Param.DstPort
	lineItem.SrcLocationID = o.Param.SrcLocID
	lineItem.DstLocationID = o.Param.DstLocID
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

	o.QuoteItemID, err = quoteJson.Get("quoteItem").GetIndex(0).Get("id").String()
	if err != nil {
		return err
	}

	return nil
}

func (o *ProductOrder) CreateNewOrder() error {
	o.OrderReq = &models.ProtoCreateOrderParam{}
	o.OrderReq.Buyer = &models.ProtoUser{Name:o.Param.BuyerName, Address:o.Param.BuyerAddr}
	o.OrderReq.Seller = &models.ProtoUser{Name:o.Param.SellerName, Address:o.Param.SellerAddr}

	connParam := &models.ProtoConnectionParam{}
	connParam.StaticParam = &models.ProtoConnectionStaticParam{}
	connParam.StaticParam.ProductOfferingID = o.Param.ProductOfferID
	connParam.StaticParam.ItemID = "1"
	connParam.StaticParam.SrcPort = o.Param.SrcPort
	connParam.StaticParam.DstPort = o.Param.DstPort

	connParam.DynamicParam = &models.ProtoConnectionDynamicParam{}
	connParam.DynamicParam.QuoteID = o.QuoteID
	connParam.DynamicParam.QuoteItemID = o.QuoteItemID
	connParam.DynamicParam.Bandwidth = fmt.Sprintf("%d Mbps", o.Param.Bandwidth)
	connParam.DynamicParam.ServiceClass = o.Param.CosName

	o.OrderReq.ConnectionParam = append(o.OrderReq.ConnectionParam, connParam)

	req := order_api.NewCreateOrderParams()
	req.Body = o.OrderReq

	rsp, err := o.Client.OrderAPI.CreateOrder(req)
	if err != nil {
		return err
	}

	o.OrderRsp = rsp.GetPayload()
	o.InternalID = o.OrderRsp.InternalID

	fmt.Printf("Create New Order is OK, InternalID %s\n", o.InternalID)

	return nil
}

func (o *ProductOrder) GetOrderInfo() error {
	req := order_api.NewGetOrderInfoParams()
	req.InternalID = &(o.OrderRsp.InternalID)

	rsp, err := o.Client.OrderAPI.GetOrderInfo(req)
	if err != nil {
		return err
	}

	o.OrderInfoRsp = rsp.GetPayload()

	fmt.Printf("Get Order Info is OK, OrderID %s, OrderState %s\n", o.OrderInfoRsp.OrderID, o.OrderInfoRsp.OrderState)

	return nil
}
