package contract

import (
	"time"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
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
	cs.orderIdFromSonata.Range(func(key, value interface{}) bool {
		idOnChain := key.(string)
		orderInfo := value.(*abi.DoDSettleOrderInfo)
		productIds, err := cs.inventoryFind(orderInfo.OrderId)
		if err != nil {
			cs.logger.Error(err)
			return true
		}

		cs.logger.Infof("get product success ,orderId is %s,productId len are %v", orderInfo.OrderId, productIds)
		err = cs.updateOrderInfoToChain(idOnChain, productIds, orderInfo)
		if err != nil {
			cs.logger.Error(err)
			return true
		}
		cs.logger.Infof("update order info to chain success ,orderId is %s,productId len are %v", orderInfo.OrderId, productIds)
		cs.orderIdFromSonata.Delete(idOnChain)
		return true
	})
}

func (cs *ContractService) updateOrderInfoToChain(idOnChain string, products []*Product, orderInfo *abi.DoDSettleOrderInfo) error {
	var id types.Hash
	_ = id.Of(idOnChain)
	ProductIds := make([]*abi.DoDSettleProductItem, 0)
	if orderInfo.OrderType == abi.DoDSettleOrderTypeCreate {
		for _, v := range products {
			pi := &abi.DoDSettleProductItem{
				ProductId: v.productID,
				ItemId:    v.buyerProductID,
			}
			ProductIds = append(ProductIds, pi)
		}
	} else {
		for _, v := range orderInfo.Connections {
			pi := &abi.DoDSettleProductItem{
				ProductId: v.ProductId,
				ItemId:    v.ItemId,
			}
			ProductIds = append(ProductIds, pi)
		}
	}

	param := &abi.DoDSettleUpdateOrderInfoParam{
		Buyer:      cs.account.Address(),
		InternalId: id,
		OrderId:    orderInfo.OrderId,
		ProductIds: ProductIds,
		Status:     abi.DoDSettleOrderStateSuccess,
		FailReason: "",
	}

	block := new(types.StateBlock)
	err := cs.client.Call(&block, "DoDSettlement_getUpdateOrderInfoBlock", param)
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
