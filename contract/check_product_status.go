package contract

import (
	"strings"
	"time"

	"github.com/qlcchain/go-lsobus/mock"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/orchestra"
	"github.com/qlcchain/go-lsobus/sonata/inventory/models"
)

func (cs *ContractService) checkProductStatus() {
	ticker := time.NewTicker(checkOrderStatusInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.chainReady {
				cs.getProductStatus()
			}
		}
	}
}

func (cs *ContractService) getProductStatus() {
	addr := cs.account.Address()
	var err error
	var id []*qlcSdk.DoDPendingResourceCheckInfo
	if cs.GetFakeMode() {
		id = mock.GetPendingResourceCheck(addr)
	} else {
		id, err = cs.client.DoDSettlement.GetPendingResourceCheck(addr)
		if id == nil {
			return
		}
	}
	for _, v := range id {
		var productActive []*qlcSdk.DoDSettleProductInfo
		for _, value := range v.Products {
			if !value.Active {
				gp := &orchestra.GetParams{
					Seller: &orchestra.PartnerParams{},
					ID:     value.ProductId,
				}
				err := cs.orchestra.ExecInventoryGet(gp)
				if err != nil {
					cs.logger.Error(err)
					continue
				}
				if strings.EqualFold(string(gp.RspInv.Status), string(models.ProductStatusActive)) {
					cs.logger.Infof("product %s is active", value.ProductId)
					value.Active = true
					productInfo := &qlcSdk.DoDSettleProductInfo{
						OrderItemId: value.OrderItemId,
						ProductId:   value.ProductId,
						Active:      true,
					}
					productActive = append(productActive, productInfo)
				}
			}
		}
		if len(productActive) != 0 {
			err = cs.updateProductStatusToChain(addr, v.OrderId, productActive)
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
					continue
				}
			}
		}
	}
}

func (cs *ContractService) updateProductStatusToChain(addr pkg.Address, orderId string, products []*qlcSdk.DoDSettleProductInfo) error {
	param := &qlcSdk.DoDSettleUpdateProductInfoParam{
		Address:     addr,
		OrderId:     orderId,
		ProductInfo: products,
	}
	blk := new(pkg.StateBlock)
	var err error
	if cs.GetFakeMode() {
		if blk, err = mock.GetUpdateProductInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return cs.account.Sign(hash), nil
		}); err != nil {
			return err
		}
	} else {
		if blk, err = cs.client.DoDSettlement.GetUpdateProductInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return cs.account.Sign(hash), nil
		}); err != nil {
			return err
		}
	}
	var w pkg.Work
	worker, _ := pkg.NewWorker(w, blk.Root())
	blk.Work = worker.NewWork()
	if !cs.GetFakeMode() {
		if err = cs.processBlockAndWaitConfirmed(blk); err != nil {
			cs.logger.Errorf("process block error: %s", err)
			return err
		} else {
			for _, v := range products {
				cs.logger.Infof("update product %s active status to chain success", v.ProductId)
			}
		}
	}
	return nil
}

func (cs *ContractService) updateOrderCompleteStatusToChain(requestHash pkg.Hash) error {
	param := &qlcSdk.DoDSettleResponseParam{
		RequestHash: requestHash,
	}
	blk := new(pkg.StateBlock)
	var err error
	if cs.GetFakeMode() {
		if blk, err = mock.GetUpdateOrderInfoRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return cs.account.Sign(hash), nil
		}); err != nil {
			return err
		}
	} else {
		if blk, err = cs.client.DoDSettlement.GetUpdateOrderInfoRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return cs.account.Sign(hash), nil
		}); err != nil {
			return err
		}
	}
	var w pkg.Work
	worker, _ := pkg.NewWorker(w, blk.Root())
	blk.Work = worker.NewWork()
	if !cs.GetFakeMode() {
		if err = cs.processBlockAndWaitConfirmed(blk); err != nil {
			cs.logger.Errorf("process block error: %s", err)
			return err
		}
	}
	return nil
}
