package contract

import (
	"errors"
	"time"

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
	cs.orderIdFromSonata.Range(func(key, value interface{}) bool {
		idOnChain := key.(string)
		orderInfo := value.(*qlcSdk.DoDSettleOrderInfo)
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

func (cs *ContractService) updateOrderInfoToChain(idOnChain string, products []*Product, orderInfo *qlcSdk.DoDSettleOrderInfo) error {
	var id pkg.Hash
	_ = id.Of(idOnChain)
	ProductIds := make([]*qlcSdk.DoDSettleProductItem, 0)
	if orderInfo.OrderType == qlcSdk.DoDSettleOrderTypeCreate {
		for _, v := range products {
			cs.logger.Infof("productId is %s,buyProductID id is %s", v.productID, v.buyerProductID)
			pi := &qlcSdk.DoDSettleProductItem{
				ProductId:      v.productID,
				BuyerProductId: v.buyerProductID,
			}
			ProductIds = append(ProductIds, pi)
		}
	} else {
		for _, v := range orderInfo.Connections {
			cs.logger.Infof("productId is %s,buyerProductId id is %s", v.ProductId, v.BuyerProductId)
			pi := &qlcSdk.DoDSettleProductItem{
				ProductId:      v.ProductId,
				BuyerProductId: v.BuyerProductId,
			}
			ProductIds = append(ProductIds, pi)
		}
	}

	param := &qlcSdk.DoDSettleUpdateOrderInfoParam{
		Buyer:      cs.account.Address(),
		InternalId: id,
		OrderId:    orderInfo.OrderId,
		ProductIds: ProductIds,
		Status:     qlcSdk.DoDSettleOrderStateSuccess,
		FailReason: "",
	}

	if blk, err := cs.client.DoDSettlement.GetUpdateOrderInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return cs.account.Sign(hash), nil
	}); err != nil {
		return err
	} else {
		var w pkg.Work
		worker, _ := pkg.NewWorker(w, blk.Root())
		blk.Work = worker.NewWork()

		hash, err := cs.client.Ledger.Process(blk)
		if err != nil {
			cs.logger.Errorf("process block error: %s", err)
			return err
		}
		cs.logger.Infof("process hash %s success", hash.String())
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
