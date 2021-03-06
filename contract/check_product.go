package contract

import (
	"time"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"
)

func (cs *ContractCaller) checkProduct() {
	ticker := time.NewTicker(checkProductInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.seller.IsChainReady() {
				cs.getProductId()
			}
		}
	}
}

func (cs *ContractCaller) getProductId() {
	addr := cs.seller.Account().Address()
	resources, err := cs.seller.GetPendingResourceCheck(addr)
	if resources == nil || err != nil {
		return
	}
	for _, order := range resources {
		if len(order.Products) == 0 {
			orderInfo, err := cs.seller.GetOrderInfoByInternalId(order.InternalId.String())
			if err != nil {
				cs.logger.Error(err)
				continue
			}
			if !cs.seller.IsFake() {
				if orderInfo.ContractState != qlcSdk.DoDSettleContractStateConfirmed || orderInfo.OrderState != qlcSdk.DoDSettleOrderStateSuccess {
					cs.logger.Info("waiting for buyer update the order")
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
			if cs.seller.IsFake() {
				cs.orderIdOnChainSeller.Range(func(key, value interface{}) bool {
					id := key.(string)
					cs.orderIdOnChainSeller.Delete(id)
					return true
				})
			} else {
				cs.orderIdOnChainSeller.Delete(order.InternalId.String())
			}
		}
	}
}

func (cs *ContractCaller) updateProductInfoToChain(
	idOnChain string, productIds []*Product, orderInfo *qlcSdk.DoDSettleOrderInfo,
) error {
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

		account := cs.seller.Account()
		param := &qlcSdk.DoDSettleUpdateProductInfoParam{
			Address:     account.Address(),
			OrderId:     orderInfo.OrderId,
			ProductInfo: productInfos,
		}
		blk := new(pkg.StateBlock)
		var err error

		if blk, err = cs.seller.GetUpdateProductInfoBlock(param); err != nil {
			return err
		}
		if _, err = cs.seller.Process(blk); err != nil {
			cs.logger.Errorf("process block error: %s", err)
			return err
		}
	}
	return nil
}

func (cs *ContractCaller) inventoryFind(sellName string, orderInfo *qlcSdk.DoDSettleOrderInfo) ([]*Product, error) {
	var productIds []*Product
	for _, conn := range orderInfo.Connections {
		pt := &Product{
			orderItemID: conn.OrderItemId,
			productID:   conn.ItemId,
		}
		productIds = append(productIds, pt)
	}

	return productIds, nil
}
