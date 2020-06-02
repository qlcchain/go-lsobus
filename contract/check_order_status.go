package contract

import (
	"time"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/orchestra"
	"github.com/qlcchain/go-lsobus/sonata/inventory/models"
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
	addr := cs.account.Address()
	id, err := cs.client.DoDSettlement.GetPendingResourceCheck(addr)
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

func (cs *ContractService) updateProductStatusToChain(addr pkg.Address, InternalId pkg.Hash, products []string) error {
	param := &qlcSdk.DoDSettleResourceReadyParam{
		Address:    addr,
		InternalId: InternalId,
		ProductId:  products,
	}

	if blk, err := cs.client.DoDSettlement.GetResourceReadyBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return cs.account.Sign(hash), nil
	}); err != nil {
		return err
	} else {
		var w pkg.Work
		worker, _ := pkg.NewWorker(w, blk.Root())
		blk.Work = worker.NewWork()

		_, err := cs.client.Ledger.Process(blk)
		if err != nil {
			cs.logger.Errorf("process block error: %s", err)
			return err
		}
		//		cs.logger.Infof("process hash %s success", hash.String())
	}
	return nil
}

func (cs *ContractService) updateOrderCompleteStatusToChain(requestHash pkg.Hash) error {
	param := &qlcSdk.DoDSettleResponseParam{
		RequestHash: requestHash,
	}
	if blk, err := cs.client.DoDSettlement.GetUpdateOrderInfoRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return cs.account.Sign(hash), nil
	}); err != nil {
		return err
	} else {
		var w pkg.Work
		worker, _ := pkg.NewWorker(w, blk.Root())
		blk.Work = worker.NewWork()

		_, err := cs.client.Ledger.Process(blk)
		if err != nil {
			cs.logger.Errorf("process block error: %s", err)
			return err
		}
		//	cs.logger.Infof("process hash %s success", hash.String())
	}
	return nil
}
