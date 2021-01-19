package contract

import (
	"strings"
	"time"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/api"

	"github.com/qlcchain/go-lsobus/orchestra/sonata/inventory/models"
)

func (cs *ContractCaller) checkProductStatus() {
	ticker := time.NewTicker(checkOrderStatusInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.seller.IsChainReady() {
				cs.getProductStatus()
			}
		}
	}
}

func (cs *ContractCaller) getProductStatus() {
	addr := cs.seller.Account().Address()
	var err error
	var orderInfo *qlcSdk.DoDSettleOrderInfo

	resources, err := cs.seller.GetPendingResourceCheck(addr)
	if resources == nil {
		return
	}

	for _, v := range resources {
		var productActive []*qlcSdk.DoDSettleProductInfo
		orderInfo, err = cs.seller.GetOrderInfoByInternalId(v.InternalId.String())
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
				err := cs.seller.ExecInventoryGet(gp)
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

func (cs *ContractCaller) updateProductStatusToChain(
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

	if blk, err = cs.seller.GetUpdateProductInfoBlock(param); err != nil {
		return err
	}

	var w pkg.Work
	worker, _ := pkg.NewWorker(w, blk.Root())
	blk.Work = worker.NewWork()
	if _, err := cs.seller.Process(blk); err != nil {
		cs.logger.Errorf("process block error: %s", err)
		return err
	} else {
		for _, v := range products {
			cs.logger.Infof("update product %s active status to chain success", v.ProductId)
		}
	}
	return nil
}

func (cs *ContractCaller) updateOrderCompleteStatusToChain(requestHash pkg.Hash) error {
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

	if blk, err = cs.seller.GetUpdateOrderInfoRewardBlock(param); err != nil {
		return err
	}

	if _, err := cs.seller.Process(blk); err != nil {
		cs.logger.Errorf("process block error: %s", err)
		return err
	}

	return nil
}
