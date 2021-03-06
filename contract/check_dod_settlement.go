package contract

import (
	"fmt"
	"strings"
	"time"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"

	"github.com/qlcchain/go-lsobus/cmd/util"

	"github.com/qlcchain/go-lsobus/api"
	qm "github.com/qlcchain/go-lsobus/orchestra/sonata/quote/models"
)

// Detect dod settlement contracts that require signing
func (cs *ContractCaller) checkDoDContract() {
	ticker := time.NewTicker(checkNeedSignContractInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.seller.IsChainReady() {
				cs.processDoDContract()
			}
		}
	}
}

func (cs *ContractCaller) processDoDContract() {
	addr := cs.seller.Account().Address()
	rsps, err := cs.seller.GetPendingRequest(addr)
	if err != nil || len(rsps) == 0 {
		return
	}

	for _, v := range rsps {
		//		if v.Order.OrderType == abi.DoDSettleOrderTypeCreate || v.Order.OrderType == abi.DoDSettleOrderTypeChange {
		if v.Order == nil {
			cs.logger.Error("invalid order info")
			continue
		}

		cs.logger.Infof("find a dod settlement need sign,request hash is %s", v.Hash.String())
		b := cs.verifyOrderInfoFromSonata(v.Order)
		if !b {
			continue
		}
		//		}

		action, err := qlcSdk.ParseDoDSettleResponseAction("confirm")
		if err != nil {
			cs.logger.Error(err)
			continue
		}

		param := &qlcSdk.DoDSettleResponseParam{
			RequestHash: v.Hash,
			Action:      action,
		}

		if v.Order.OrderType == qlcSdk.DoDSettleOrderTypeCreate {
			cs.logger.Infof("order type is create, %s", util.ToIndentString(param))
			key := fmt.Sprintf("%s%d", v.Hash.String(), qlcSdk.DoDSettleOrderTypeCreate)
			if !cs.cache.Has(key) {
				if _, err = cs.seller.GetCreateOrderRewardBlock(param); err != nil {
					cs.logger.Error(err)
					continue
				} else {
					cs.cache.Set(key, true)
				}
			}
		} else if v.Order.OrderType == qlcSdk.DoDSettleOrderTypeChange {
			cs.logger.Info("order type is change")
			key := fmt.Sprintf("%s%d", v.Hash.String(), qlcSdk.DoDSettleOrderTypeChange)
			if !cs.cache.Has(key) {
				if _, err = cs.seller.GetChangeOrderRewardBlock(param); err != nil {
					cs.logger.Error(err)
					continue
				} else {
					cs.cache.Set(key, true)
				}
			}
		} else if v.Order.OrderType == qlcSdk.DoDSettleOrderTypeTerminate {
			cs.logger.Info("order type is terminate")
			key := fmt.Sprintf("%s%d", v.Hash.String(), qlcSdk.DoDSettleOrderTypeTerminate)
			if !cs.cache.Has(key) {
				if _, err = cs.seller.GetTerminateOrderRewardBlock(param); err != nil {
					cs.logger.Error(err)
					continue
				} else {
					cs.cache.Set(key, true)
				}
			}
		} else {
			cs.logger.Errorf("unknown order type==%s", v.Order.OrderType.String())
			continue
		}
		//cs.logger.Debug(blk)

		//if h, err := cs.seller.Process(blk); err != nil {
		//	cs.logger.Errorf("process block error: %s", err)
		//	continue
		//} else if h != pkg.ZeroHash {
		//	cs.logger.Infof("dod settlement sign success,request hash is :%s", v.Hash.String())
		//	err = cs.readAndWriteProcessingOrder("add", "seller", v.Order.InternalId)
		//	if err != nil {
		//		cs.logger.Error(err)
		//		continue
		//	}
		//	cs.orderIdOnChainSeller.Store(v.Order.InternalId, "")
		//}
	}
}

// use quoteId call sonata api to verify order info
func (cs *ContractCaller) verifyOrderInfoFromSonata(order *qlcSdk.DoDSettleOrderInfo) bool {
	for idx := 0; idx < len(order.Connections); idx++ {
		var quote *qm.QuoteItem
		conn := order.Connections[idx]
		op := &api.GetParams{
			Seller: &api.PartnerParams{
				ID:   order.Seller.Address.String(),
				Name: order.Seller.Name,
			},
			ID: conn.QuoteId,
		}

		err := cs.seller.ExecQuoteGet(op)
		if err != nil {
			cs.logger.Error(err)
			return false
		}
		if op.RspQuote == nil {
			cs.logger.Errorf("order information verify fail, empty quote response")
			return false
		}

		//FIXME: support ordering port
		//if len(op.RspQuote.QuoteItem) != len(order.Connections) {
		//	cs.logger.Errorf("order information verify fail, item count not equal")
		//	return false
		//}

		for _, v := range op.RspQuote.QuoteItem {
			if *v.ID == conn.QuoteItemId {
				quote = v
			}
		}
		if quote != nil {
			if quote.State != qm.QuoteItemStateTypeREADY {
				cs.logger.Errorf("quote state %s error,order information verify fail", quote.State)
				return false
			}
			if !strings.EqualFold(*quote.PreCalculatedPrice.Price.PreTaxAmount.Unit, conn.Currency) || *quote.PreCalculatedPrice.Price.PreTaxAmount.Value != float32(conn.Price) {
				cs.logger.Errorf("order information verify fail, quote %f/%s, conn %f/%s", *quote.PreCalculatedPrice.Price.PreTaxAmount.Value, *quote.PreCalculatedPrice.Price.PreTaxAmount.Unit, conn.Price, conn.Currency)
				return false
			}
		} else {
			cs.logger.Errorf("order information verify fail,can not find quote item id %s", conn.QuoteItemId)
			return false
		}
	}

	cs.logger.Infof("order information verified")
	return true
}
