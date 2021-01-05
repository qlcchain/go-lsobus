package contract

import (
	"strings"
	"time"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/mock"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

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
	var orderInfo *qlcSdk.DoDSettleOrderInfo
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
		orderInfo, err = cs.GetOrderInfoByInternalId(v.InternalId.String())
		if err != nil {
			cs.logger.Error(err)
			continue
		}
		for _, value := range v.Products {
			if !value.Active {
				gp := &api.GetParams{
					Seller: &api.PartnerParams{},
					ID:     value.ProductId,
				}
				err := cs.sellers.ExecInventoryGet(gp)
				if err != nil {
					cs.logger.Error(err)
					continue
				}
				if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeTerminate {
					if strings.EqualFold(string(gp.RspInv.Status), string(models.ProductStatusTerminated)) {
						cs.logger.Infof("product %s is terminated", value.ProductId)
						value.Active = true
						productInfo := &qlcSdk.DoDSettleProductInfo{
							OrderItemId: value.OrderItemId,
							ProductId:   value.ProductId,
							Active:      true,
						}
						productActive = append(productActive, productInfo)
					}
				} else {
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
		}
		if len(productActive) != 0 {
			err = cs.updateProductStatusToChain(addr, v.OrderId, productActive)
			if err != nil {
				cs.logger.Error(err)
			}
		}
		var c int
		orderReady := false
		for _, value := range v.Products {
			if value.Active {
				c++
			}
		}
		if c == len(v.Products) {
			if c != 0 {
				orderReady = true
			}
		}
		if orderReady {
			err = cs.updateOrderCompleteStatusToChain(v.SendHash)
			if err != nil {
				cs.logger.Error(err)
				continue
			}
			cs.logger.Infof("update order %s complete status to chain success", v.OrderId)
		}
	}
}

func (cs *ContractService) updateProductStatusToChain(
	addr pkg.Address, orderId string, products []*qlcSdk.DoDSettleProductInfo,
) error {
	param := &qlcSdk.DoDSettleUpdateProductInfoParam{
		Address:     addr,
		OrderId:     orderId,
		ProductInfo: products,
	}
	if cs.cfg.Privacy.Enable {
		param.PrivateFrom = cs.cfg.Privacy.From
		param.PrivateFor = cs.cfg.Privacy.For
		param.PrivateGroupID = cs.cfg.Privacy.PrivateGroupID
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
	if cs.cfg.Privacy.Enable {
		param.PrivateFrom = cs.cfg.Privacy.From
		param.PrivateFor = cs.cfg.Privacy.For
		param.PrivateGroupID = cs.cfg.Privacy.PrivateGroupID
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
