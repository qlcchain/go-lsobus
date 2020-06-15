package contract

import (
	"time"

	"github.com/qlcchain/go-lsobus/mock"

	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/orchestra"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
)

func (cs *ContractService) checkProduct() {
	ticker := time.NewTicker(checkProductInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.chainReady {
				cs.getProductId()
			}
		}
	}
}

func (cs *ContractService) getProductId() {
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
	for _, order := range id {
		if len(order.Products) == 0 {
			orderInfo, err = cs.GetOrderInfoByInternalId(order.InternalId.String())
			if err != nil {
				cs.logger.Error(err)
				continue
			}
			if !cs.GetFakeMode() {
				if orderInfo.ContractState != qlcSdk.DoDSettleContractStateConfirmed || orderInfo.OrderState != qlcSdk.DoDSettleOrderStateSuccess {
					cs.logger.Info("waiting for buyer place order")
					continue
				}
			}
			productIds, err := cs.inventoryFind(orderInfo.Seller.Name, orderInfo)
			if err != nil {
				if err == noInventoryList {
					cs.logger.Info(noInventoryList)
				} else {
					cs.logger.Error(err)
				}
				continue
			}

			cs.logger.Infof("get product success ,orderId is %s", orderInfo.OrderId)
			err = cs.updateProductInfoToChain(order.InternalId.String(), productIds, orderInfo)
			if err != nil {
				cs.logger.Error(err)
				continue
			}
			cs.logger.Info("update product info to chain success")
			err = cs.readAndWriteProcessingOrder("delete", "seller", order.InternalId.String())
			if err != nil {
				cs.logger.Error(err)
				continue
			}
			cs.orderIdOnChainSeller.Delete(order.InternalId.String())
		}
	}
}

func (cs *ContractService) updateProductInfoToChain(idOnChain string, productIds []*Product, orderInfo *qlcSdk.DoDSettleOrderInfo) error {
	var id pkg.Hash
	_ = id.Of(idOnChain)
	productInfos := make([]*qlcSdk.DoDSettleProductInfo, 0)
	if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeCreate {
		for _, v := range productIds {
			cs.logger.Infof("productID is %s,orderItemId id is %s", v.productID, v.orderItemID)
			pi := &qlcSdk.DoDSettleProductInfo{
				OrderItemId: v.orderItemID,
				ProductId:   v.productID,
			}
			productInfos = append(productInfos, pi)
		}

		param := &qlcSdk.DoDSettleUpdateProductInfoParam{
			Address:     cs.account.Address(),
			OrderId:     orderInfo.OrderId,
			ProductInfo: productInfos,
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
			}
		}
	}
	return nil
}

func (cs *ContractService) inventoryFind(sellName string, orderInfo *qlcSdk.DoDSettleOrderInfo) ([]*Product, error) {
	fp := &orchestra.FindParams{
		Seller:         &orchestra.PartnerParams{Name: sellName},
		ProductOrderID: orderInfo.OrderId,
	}
	err := cs.orchestra.ExecInventoryFind(fp)
	if err != nil {
		return nil, err
	}
	var productIds []*Product
	if len(fp.RspInvList) == 0 {
		return nil, noInventoryList
	}
	for _, conn := range orderInfo.Connections {
		var b bool
		for _, productSummary := range fp.RspInvList {
			for _, productOrderRef := range productSummary.ProductOrder {
				if conn.OrderItemId == *productOrderRef.OrderItemID {
					pt := &Product{
						orderItemID: *productOrderRef.OrderItemID,
						productID:   *productSummary.ID,
					}
					productIds = append(productIds, pt)
					b = true
					break
				}
				if b {
					break
				}
			}
		}
	}

	return productIds, nil
}
