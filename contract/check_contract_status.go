package contract

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/qlcchain/go-lsobus/orchestra"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
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
		orderInfo := new(abi.DoDSettleOrderInfo)
		err := cs.client.Call(&orderInfo, "DoDSettlement_getOrderInfoByInternalId", &internalId)
		if err != nil {
			cs.logger.Error(err)
			return true
		}
		if orderInfo.ContractState == abi.DoDSettleContractStateConfirmed {
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

func (cs *ContractService) createOrderToSonataServer(internalId string, orderInfo *abi.DoDSettleOrderInfo) (string, error) {
	orderActivity := ""
	itemAction := ""
	if orderInfo.OrderType == abi.DoDSettleOrderTypeCreate {
		orderActivity = "install"
		itemAction = "add"
	} else if orderInfo.OrderType == abi.DoDSettleOrderTypeChange {
		orderActivity = "change"
		itemAction = "change"
	} else if orderInfo.OrderType == abi.DoDSettleOrderTypeTerminate {
		orderActivity = "disconnect"
		itemAction = "remove"
	}

	eLines := make([]*orchestra.ELineItemParams, 0)
	for _, v := range orderInfo.Connections {
		bws := strings.Split(v.Bandwidth, " ")
		if len(bws) != 2 {
			return "", errors.New("bandwidth error")
		}
		bw, err := strconv.Atoi(bws[0])
		if err != nil {
			return "", err
		}
		eLine := &orchestra.ELineItemParams{
			SrcPortID:     v.SrcPort,
			DstPortID:     v.DstPort,
			SrcLocationID: v.SrcDataCenter,
			DstLocationID: v.DstDataCenter,
			Bandwidth:     uint(bw),
			BwUnit:        bws[1],
			CosName:       v.ServiceClass.String(),
			BaseItemParams: orchestra.BaseItemParams{
				BillingParams: &orchestra.BillingParams{
					PaymentType: v.PaymentType.String(),
					BillingType: v.BillingType.String(),
					BillingUnit: v.BillingUnit.String(),
					MeasureUnit: v.BillingUnit.String(),
					StartTime:   v.StartTime,
					EndTime:     v.EndTime,
					Currency:    v.Currency,
					Price:       float32(v.Price),
				},
				BuyerProductID: v.ItemId,
			},
		}
		eLine.Action = itemAction
		eLines = append(eLines, eLine)
	}
	op := &orchestra.OrderParams{
		OrderActivity: orderActivity,
		Buyer: &orchestra.Partner{
			ID:   orderInfo.Buyer.Address.String(),
			Name: orderInfo.Buyer.Name,
		},
		Seller: &orchestra.Partner{
			ID:   orderInfo.Seller.Address.String(),
			Name: orderInfo.Seller.Name,
		},
		ExternalID: internalId,
		ELineItems: eLines,
		//PaymentType: "",
		//BillingType: "",
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
