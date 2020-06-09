package contract

import (
	"errors"
	"strconv"
	"strings"
	"time"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"

	"github.com/qlcchain/go-lsobus/orchestra"
)

func (cs *ContractService) checkContractStatus() {
	ticker := time.NewTicker(checkContractStatusInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.chainReady {
				cs.getContractStatus()
			}
		}
	}
}

func (cs *ContractService) getContractStatus() {
	cs.orderIdOnChain.Range(func(key, value interface{}) bool {
		internalId := key.(string)
		orderInfo, err := cs.GetOrderInfoByInternalId(internalId)
		if err != nil {
			cs.logger.Error(err)
			return true
		}
		if orderInfo.ContractState == qlcSdk.DoDSettleContractStateConfirmed {
			cs.logger.Infof(" contract %s confirmed", internalId)
			cs.logger.Info(" call sonata API to place order")
			orderId, err := cs.createOrderToSonataServer(internalId, orderInfo)
			if err != nil {
				cs.logger.Error(err)
				return true
			}
			cs.logger.Infof("order place success,order id from sonata is:%s", orderId)
			cs.orderIdOnChain.Delete(internalId)
		}
		return true
	})
}

func (cs *ContractService) createOrderToSonataServer(internalId string, orderInfo *qlcSdk.DoDSettleOrderInfo) (string, error) {
	orderActivity := ""
	itemAction := ""
	if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeCreate {
		orderActivity = "install"
		itemAction = "add"
	} else if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeChange {
		orderActivity = "change"
		itemAction = "change"
	} else if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeTerminate {
		orderActivity = "disconnect"
		itemAction = "remove"
	}

	eLines := make([]*orchestra.ELineItemParams, 0)
	for _, v := range orderInfo.Connections {
		eLine := &orchestra.ELineItemParams{
			SrcPortID:     v.SrcPort,
			DstPortID:     v.DstPort,
			DstCompanyID:  v.DstCompanyName,
			DstMetroID:    v.DstCity,
			SrcLocationID: v.SrcDataCenter,
			DstLocationID: v.DstDataCenter,
			CosName:       strings.ToUpper(v.ServiceClass.String()),
			BaseItemParams: orchestra.BaseItemParams{
				BuyerProductID: v.BuyerProductId,
			},
		}

		eLine.ItemID = v.ItemId
		billingParams := &orchestra.BillingParams{}
		if len(v.Bandwidth) != 0 {
			bws := strings.Split(v.Bandwidth, " ")
			if len(bws) != 2 {
				return "", errors.New("bandwidth error")
			}
			bw, err := strconv.Atoi(bws[0])
			if err != nil {
				return "", err
			}
			eLine.Bandwidth = uint(bw)
			eLine.BwUnit = bws[1]
		}

		if v.PaymentType.String() != "null" {
			billingParams.PaymentType = strings.ToUpper(v.PaymentType.String())
		}
		if v.BillingType.String() != "null" {
			billingParams.BillingType = strings.ToUpper(v.BillingType.String())
		}
		if v.BillingUnit.String() != "null" {
			billingParams.BillingUnit = v.BillingUnit.String()
		}
		if v.BillingUnit.String() != "null" {
			billingParams.MeasureUnit = v.BillingUnit.String()
		}
		billingParams.StartTime = v.StartTime
		billingParams.EndTime = v.EndTime
		billingParams.Currency = v.Currency
		billingParams.Price = float32(v.Price)
		eLine.BillingParams = billingParams

		eLine.Name = v.ConnectionName
		eLine.QuoteID = v.QuoteId
		eLine.QuoteItemID = v.QuoteItemId
		eLine.Action = itemAction
		eLine.ProdOfferID = v.ProductOfferingId
		eLines = append(eLines, eLine)
	}
	op := &orchestra.OrderParams{
		OrderActivity: orderActivity,
		Buyer: &orchestra.PartnerParams{
			ID:   orderInfo.Buyer.Address.String(),
			Name: orderInfo.Buyer.Name,
		},
		Seller: &orchestra.PartnerParams{
			ID:   orderInfo.Seller.Address.String(),
			Name: orderInfo.Seller.Name,
		},
		ExternalID:  internalId,
		ELineItems:  eLines,
		PaymentType: eLines[0].BillingParams.PaymentType,
		BillingType: eLines[0].BillingParams.BillingType,
	}
	err := cs.orchestra.ExecOrderCreate(op)
	if err != nil {
		return "", err
	}
	orderId := op.RspOrder.ID
	orderInfo.OrderId = *orderId
	cs.orderIdFromSonata.Store(internalId, orderInfo)
	return *orderId, nil
}
