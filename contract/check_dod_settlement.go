package contract

import (
	"time"

	"github.com/qlcchain/go-lsobus/orchestra"
	qm "github.com/qlcchain/go-lsobus/sonata/quote/models"

	"github.com/qlcchain/go-qlc/common/types"
	qlcchain "github.com/qlcchain/go-qlc/rpc/api"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
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
	dod := make([]*qlcchain.DoDPendingRequestRsp, 0)
	addr := cs.account.Address()
	err := cs.client.Call(&dod, "DoDSettlement_getPendingRequest", &addr)
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
		action, err := abi.ParseDoDSettleResponseAction("confirm")
		if err != nil {
			cs.logger.Error(err)
			continue
		}

		param := &abi.DoDSettleResponseParam{
			RequestHash: v.Hash,
			Action:      action,
		}
		block := new(types.StateBlock)
		if v.Order.OrderType == abi.DoDSettleOrderTypeCreate {
			cs.logger.Info(" order type is create")
			err = cs.client.Call(&block, "DoDSettlement_getCreateOrderRewardBlock", param)
		} else if v.Order.OrderType == abi.DoDSettleOrderTypeChange {
			cs.logger.Info(" order type is change")
			err = cs.client.Call(&block, "DoDSettlement_getChangeOrderRewardBlock", param)
		} else if v.Order.OrderType == abi.DoDSettleOrderTypeTerminate {
			cs.logger.Info(" order type is terminate")
			err = cs.client.Call(&block, "DoDSettlement_getTerminateOrderRewardBlock", param)
		} else {
			cs.logger.Errorf("unknown order type==%s", v.Order.OrderType.String())
			continue
		}
		if err != nil {
			cs.logger.Error(err)
			continue
		}

		var w types.Work
		worker, _ := types.NewWorker(w, block.Root())
		block.Work = worker.NewWork()

		hash := block.GetHash()
		block.Signature = cs.account.Sign(hash)

		var h types.Hash
		err = cs.client.Call(&h, "ledger_process", &block)
		if err != nil {
			cs.logger.Error(err)
			continue
		}
		cs.logger.Infof("dod settlement sign success,request hash is :%s", v.Hash.String())
	}
}

// use quoteId call sonata api to verify order info
func (cs *ContractService) verifyOrderInfoFromSonata(order *abi.DoDSettleOrderInfo) bool {
	for idx := 0; idx < len(order.Connections); idx++ {
		var quote *qm.QuoteItem
		conn := order.Connections[idx]
		op := &orchestra.GetParams{
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
