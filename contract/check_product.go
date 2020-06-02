package contract

import (
	"errors"
	"time"

	"github.com/qlcchain/go-lsobus/orchestra"

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
		productIds, err := cs.inventoryFind(orderInfo.Seller.Name, orderInfo.OrderId)
		if err != nil {
			cs.logger.Error(err)
			return true
		}

		cs.logger.Infof("get product success ,orderId is %s", orderInfo.OrderId)

		err = cs.updateOrderInfoToChain(idOnChain, productIds, orderInfo)
		if err != nil {
			cs.logger.Error(err)
			return true
		}
		cs.logger.Infof("update order info to chain success ,orderId is %s", orderInfo.OrderId)
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
			cs.logger.Infof("productId is %s,item id is %s", v.productID, v.buyerProductID)
			pi := &abi.DoDSettleProductItem{
				ProductId:      v.productID,
				BuyerProductId: v.buyerProductID,
			}
			ProductIds = append(ProductIds, pi)
		}
	} else {
		for _, v := range orderInfo.Connections {
			cs.logger.Infof("productId is %s,buyerProductId id is %s", v.ProductId, v.BuyerProductId)
			pi := &abi.DoDSettleProductItem{
				ProductId:      v.ProductId,
				BuyerProductId: v.BuyerProductId,
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

func (cs *ContractService) inventoryFind(sellName, orderId string) ([]*Product, error) {
	fp := &orchestra.FindParams{
		Seller:         &orchestra.PartnerParams{Name: sellName},
		ProductOrderID: orderId,
	}
	err := cs.orchestra.ExecInventoryFind(fp)
	if err != nil {
		return nil, err
	}
	var productIds []*Product
	if len(fp.RspInvList) == 0 {
		return nil, errors.New("no inventory list ")
	}
	for _, v := range fp.RspInvList {
		pt := &Product{
			buyerProductID: v.BuyerProductID,
			productID:      *v.ID,
		}
		productIds = append(productIds, pt)
	}
	return productIds, nil
}
