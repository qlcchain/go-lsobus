package contract

import (
	"time"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/orchestra"
	qm "github.com/qlcchain/go-lsobus/sonata/quote/models"
)

// Detect dod settlement contracts that require signing
func (cs *ContractService) checkDoDContract() {
	ticker := time.NewTicker(checkNeedSignContractInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.chainReady {
				cs.processDoDContract()
			}
		}
	}
}

func (cs *ContractService) processDoDContract() {
	addr := cs.account.Address()
	dod, err := cs.client.DoDSettlement.GetPendingRequest(addr)
	if err != nil || len(dod) == 0 {
		return
	}
	for _, v := range dod {
		cs.logger.Infof("find a dod settlement need sign,request hash is %s", v.Hash.String())
		//		if v.Order.OrderType == abi.DoDSettleOrderTypeCreate || v.Order.OrderType == abi.DoDSettleOrderTypeChange {
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
		blk := new(pkg.StateBlock)
		if v.Order.OrderType == qlcSdk.DoDSettleOrderTypeCreate {
			cs.logger.Info(" order type is create")
			if blk, err = cs.client.DoDSettlement.GetCreateOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
				return cs.account.Sign(hash), nil
			}); err != nil {
				cs.logger.Error(err)
				continue
			}
		} else if v.Order.OrderType == qlcSdk.DoDSettleOrderTypeChange {
			cs.logger.Info(" order type is change")
			if blk, err = cs.client.DoDSettlement.GetChangeOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
				return cs.account.Sign(hash), nil
			}); err != nil {
				cs.logger.Error(err)
				continue
			}
		} else if v.Order.OrderType == qlcSdk.DoDSettleOrderTypeTerminate {
			cs.logger.Info(" order type is terminate")
			if blk, err = cs.client.DoDSettlement.GetTerminateOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
				return cs.account.Sign(hash), nil
			}); err != nil {
				cs.logger.Error(err)
				continue
			}
		} else {
			cs.logger.Errorf("unknown order type==%s", v.Order.OrderType.String())
			continue
		}
		var w pkg.Work
		worker, _ := pkg.NewWorker(w, blk.Root())
		blk.Work = worker.NewWork()
		_, err = cs.client.Ledger.Process(blk)
		if err != nil {
			cs.logger.Error(err)
			continue
		}
		cs.logger.Infof("dod settlement sign success,request hash is :%s", v.Hash.String())
	}
}

// use quoteId call sonata api to verify order info
func (cs *ContractService) verifyOrderInfoFromSonata(order *qlcSdk.DoDSettleOrderInfo) bool {
	for idx := 0; idx < len(order.Connections); idx++ {
		var quote *qm.QuoteItem
		conn := order.Connections[idx]
		op := &orchestra.GetParams{
			Seller: &orchestra.PartnerParams{
				ID:   order.Seller.Address.String(),
				Name: order.Seller.Name,
			},
			ID: conn.QuoteId,
		}

		err := cs.orchestra.ExecQuoteGet(op)
		if err != nil {
			cs.logger.Error(err)
			return false
		}
		if op.RspQuote == nil {
			cs.logger.Errorf("order information verify fail, empty quote response")
			return false
		}

		if len(op.RspQuote.QuoteItem) != len(order.Connections) {
			cs.logger.Errorf("order information verify fail, item count not equal")
			return false
		}

		for _, v := range op.RspQuote.QuoteItem {
			if *v.ID == conn.QuoteItemId {
				quote = v
			}
		}
		if quote != nil {
			if quote.State != qm.QuoteItemStateTypeREADY {
				cs.logger.Error("order information verify fail")
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
