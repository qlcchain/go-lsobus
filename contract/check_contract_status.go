package contract

import (
	"errors"
	"strconv"
	"strings"
	"time"

	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/api"

	"github.com/qlcchain/go-lsobus/mock"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
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
	cs.orderIdOnChainBuyer.Range(func(key, value interface{}) bool {
		internalId := key.(string)
		orderInfo, err := cs.GetOrderInfoByInternalId(internalId)
		if err != nil {
			cs.logger.Error(err)
			return true
		}
		if orderInfo.ContractState == qlcSdk.DoDSettleContractStateConfirmed {
			cs.logger.Infof(" contract %s confirmed", internalId)
			cs.logger.Info(" check this order have already exit on sonata server")
			var b bool
			var orderId string
			orderId, b = cs.checkOrderAlreadyExitOnSonataServer(internalId)
			if !b {
				cs.logger.Info(" call sonata API to place order")
				orderId, err = cs.createOrderToSonataServer(internalId, orderInfo)
				if err != nil {
					cs.logger.Error(err)
					return true
				}
				cs.logger.Infof("order place success,order id from sonata is:%s", orderId)
			} else {
				cs.logger.Infof("order %s already placed,just need update to the chain", orderId)
			}
			orderInfo.OrderId = orderId
			err = cs.updateOrderInfoToChain(internalId, orderInfo)
			if err != nil {
				cs.logger.Error(err)
				return true
			}
			err = cs.readAndWriteProcessingOrder("delete", "buyer", internalId)
			if err != nil {
				cs.logger.Error(err)
				return true
			}
			cs.orderIdOnChainBuyer.Delete(internalId)
		}
		return true
	})
}

func (cs *ContractService) checkOrderAlreadyExitOnSonataServer(externalID string) (string, bool) {
	fp := &api.FindParams{
		ExternalID: externalID,
	}
	err := cs.sellers.ExecOrderFind(fp)
	if err != nil {
		return "", false
	}
	if len(fp.RspOrderList) == 0 {
		return "", false
	}
	if len(fp.RspOrderList) > 1 {
		cs.logger.Warnf("rsp order list should be only one ,but have %d,and the order id are:", len(fp.RspOrderList))
		for _, v := range fp.RspOrderList {
			cs.logger.Warn(*v.ID)
		}
	}
	return *fp.RspOrderList[0].ID, true
}

func (cs *ContractService) createOrderToSonataServer(
	internalId string, orderInfo *qlcSdk.DoDSettleOrderInfo,
) (string, error) {
	orderActivity := ""
	itemAction := ""
	if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeCreate {
		orderActivity = "INSTALL"
		itemAction = "INSTALL"
	} else if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeChange {
		orderActivity = "CHANGE"
		itemAction = "CHANGE"
	} else if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeTerminate {
		orderActivity = "DISCONNECT"
		itemAction = "DISCONNECT"
	}

	eLines := make([]*api.ELineItemParams, 0)
	for _, v := range orderInfo.Connections {
		eLine := &api.ELineItemParams{
			SrcPortID:     v.SrcPort,
			DstPortID:     v.DstPort,
			DstCompanyID:  v.DstCompanyName,
			DstMetroID:    v.DstCity,
			SrcLocationID: v.SrcDataCenter,
			DstLocationID: v.DstDataCenter,
			CosName:       strings.ToUpper(v.ServiceClass.String()),
			BaseItemParams: api.BaseItemParams{
				BuyerProductID: v.BuyerProductId,
			},
		}

		eLine.ItemID = v.ItemId
		billingParams := &api.BillingParams{}
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
		eLine.ProductID = v.ProductId
		eLines = append(eLines, eLine)
	}
	op := &api.OrderParams{
		OrderActivity: orderActivity,
		Buyer: &api.PartnerParams{
			ID:   orderInfo.Buyer.Address.String(),
			Name: orderInfo.Buyer.Name,
		},
		Seller: &api.PartnerParams{
			ID:   orderInfo.Seller.Address.String(),
			Name: orderInfo.Seller.Name,
		},
		ExternalID:  internalId,
		ELineItems:  eLines,
		PaymentType: eLines[0].BillingParams.PaymentType,
		BillingType: eLines[0].BillingParams.BillingType,
	}
	err := cs.sellers.ExecOrderCreate(op)
	if err != nil {
		return "", err
	}
	orderId := op.RspOrder.ID
	for _, v := range op.RspOrder.OrderItem {
		for _, value := range orderInfo.Connections {
			if value.ItemId == v.ExternalID {
				value.OrderItemId = *v.ID
				break
			}
		}
	}
	return *orderId, nil
}

func (cs *ContractService) updateOrderInfoToChain(idOnChain string, orderInfo *qlcSdk.DoDSettleOrderInfo) error {
	var id pkg.Hash
	_ = id.Of(idOnChain)
	orderItemIds := make([]*qlcSdk.DoDSettleOrderItem, 0)
	for _, v := range orderInfo.Connections {
		cs.logger.Infof("itemId is %s,orderItemId id is %s", v.ItemId, v.OrderItemId)
		pi := &qlcSdk.DoDSettleOrderItem{
			ItemId:      v.ItemId,
			OrderItemId: v.OrderItemId,
		}
		orderItemIds = append(orderItemIds, pi)
	}

	param := &qlcSdk.DoDSettleUpdateOrderInfoParam{
		Buyer:       cs.account.Address(),
		InternalId:  id,
		OrderId:     orderInfo.OrderId,
		OrderItemId: orderItemIds,
		Status:      qlcSdk.DoDSettleOrderStateSuccess,
		FailReason:  "",
	}
	if cs.cfg.Privacy.Enable {
		param.PrivateFrom = cs.cfg.Privacy.From
		param.PrivateFor = cs.cfg.Privacy.For
		param.PrivateGroupID = cs.cfg.Privacy.PrivateGroupID
	}
	blk := new(pkg.StateBlock)
	var err error
	if cs.GetFakeMode() {
		if blk, err = mock.GetUpdateOrderInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return cs.account.Sign(hash), nil
		}); err != nil {
			return err
		}
	} else {
		if blk, err = cs.client.DoDSettlement.GetUpdateOrderInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return cs.account.Sign(hash), nil
		}); err != nil {
			return err
		}
	}
	var w pkg.Work
	worker, _ := pkg.NewWorker(w, blk.Root())
	blk.Work = worker.NewWork()
	if !cs.GetFakeMode() {
		hash, err := cs.client.Ledger.Process(blk)
		if err != nil {
			cs.logger.Errorf("process block error: %s", err)
			return err
		}
		cs.logger.Infof("process hash %s success", hash.String())
	}
	return nil
}
