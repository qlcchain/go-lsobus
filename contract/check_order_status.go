package contract

import (
	"time"

	"github.com/qlcchain/go-lsobus/orchestra"
	"github.com/qlcchain/go-lsobus/sonata/inventory/models"
	"github.com/qlcchain/go-qlc/common/types"
	qlcchain "github.com/qlcchain/go-qlc/rpc/api"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
)

func (cs *ContractService) checkOrderStatus() {
	ticker := time.NewTicker(checkOrderStatusInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.chainReady {
				cs.getOrderStatus()
			}
		}
	}
}

func (cs *ContractService) getOrderStatus() {
	var id []*qlcchain.DoDPendingResourceCheckInfo
	addr := cs.account.Address()
	err := cs.client.Call(&id, "DoDSettlement_getPendingResourceCheck", &addr)
	if err != nil {
		cs.logger.Error(err)
		return
	}
	for _, v := range id {
		var productActive []string
		for _, value := range v.Products {
			if !value.Active {
				gp := &orchestra.GetParams{
					ID: value.ProductId,
				}
				err := cs.orchestra.ExecInventoryGet(gp)
				if err != nil {
					cs.logger.Error(err)
					continue
				}
				if gp.RspInv.Status == models.ProductStatusActive {
					cs.logger.Infof("product %s is active", value.ProductId)
					value.Active = true
					productActive = append(productActive, value.ProductId)
				}
			}
		}
		if len(productActive) != 0 {
			err = cs.updateProductStatusToChain(addr, v.InternalId, productActive)
			if err != nil {
				cs.logger.Error(err)
			}
			orderReady := true
			for _, value := range v.Products {
				if !value.Active {
					orderReady = false
					break
				}
			}
			if orderReady {
				err = cs.updateOrderCompleteStatusToChain(v.SendHash)
				if err != nil {
					cs.logger.Error(err)
				}
				cs.logger.Infof("update order %s complete status to chain success", v.OrderId)
			}
		}
	}
}

func (cs *ContractService) updateProductStatusToChain(addr types.Address, InternalId types.Hash, products []string) error {
	param := &abi.DoDSettleResourceReadyParam{
		Address:    addr,
		InternalId: InternalId,
		ProductId:  products,
	}

	block := new(types.StateBlock)
	err := cs.client.Call(&block, "DoDSettlement_getResourceReadyBlock", param)
	if err != nil {
		return err
	}

	var w types.Work
	worker, _ := types.NewWorker(w, block.Root())
	block.Work = worker.NewWork()

	hash := block.GetHash()
	block.Signature = cs.account.Sign(hash)

	var h types.Hash
	err = cs.client.Call(&h, "ledger_process", &block)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractService) updateOrderCompleteStatusToChain(requestHash types.Hash) error {
	param := &abi.DoDSettleResponseParam{
		RequestHash: requestHash,
	}

	block := new(types.StateBlock)
	err := cs.client.Call(&block, "DoDSettlement_getUpdateOrderInfoRewardBlock", param)
	if err != nil {
		return err
	}

	var w types.Work
	worker, _ := types.NewWorker(w, block.Root())
	block.Work = worker.NewWork()

	hash := block.GetHash()
	block.Signature = cs.account.Sign(hash)

	h := block.GetHash()
	err = cs.client.Call(&h, "ledger_process", &block)
	if err != nil {
		return err
	}
	return nil
}
