package contract

import (
	"errors"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
)

func (cs *ContractService) GetChangeOrderBlock(param *proto.ChangeOrderParam) (string, error) {
	addr := cs.account.Address().String()
	if addr == param.Buyer.Address {
		op, err := cs.convertProtoToChangeOrderParam(param)
		if err != nil {
			return "", err
		}
		block := new(types.StateBlock)
		err = cs.client.Call(&block, "DoDSettlement_getChangeOrderBlock", op)
		if err != nil {
			return "", err
		}

		var w types.Work
		worker, _ := types.NewWorker(w, block.Root())
		block.Work = worker.NewWork()

		hash := block.GetHash()
		block.Signature = cs.account.Sign(hash)
		var h types.Hash
		err = cs.client.Call(&h, "ledger_process", &block)
		if err != nil {
			return "", err
		}
		cs.logger.Infof("process hash %s success", h.String())
		internalId := block.Previous.String()
		cs.orderIdOnChain.Store(internalId, "")
		return internalId, nil
	} else {
		cs.logger.Errorf("buyer address not match,have %s,want %s", param.Buyer.Address, addr)
	}
	return "", errors.New("buyer address not match")
}

func (cs *ContractService) convertProtoToChangeOrderParam(param *proto.ChangeOrderParam) (*abi.DoDSettleChangeOrderParam, error) {
	sellerAddr, _ := types.HexToAddress(param.Seller.Address)
	buyAddr, _ := types.HexToAddress(param.Buyer.Address)
	op := new(abi.DoDSettleChangeOrderParam)
	op.Buyer = &abi.DoDSettleUser{
		Address: buyAddr,
		Name:    param.Buyer.Name,
	}
	op.Seller = &abi.DoDSettleUser{
		Address: sellerAddr,
		Name:    param.Seller.Name,
	}
	op.QuoteId = param.QuoteId
	for _, v := range param.ChangeConnectionParam {
		conn := new(abi.DoDSettleChangeConnectionParam)
		if len(v.ProductId) == 0 {
			return nil, errors.New("product id can not be nil")
		}
		conn.ProductId = v.ProductId
		if len(v.QuoteItemId) != 0 {
			conn.QuoteItemId = v.QuoteItemId
		}

		if len(v.DynamicParam.PaymentType) != 0 {
			paymentType, err := abi.ParseDoDSettlePaymentType(v.DynamicParam.PaymentType)
			if err != nil {
				return nil, err
			}
			conn.PaymentType = paymentType
		}
		if len(v.DynamicParam.BillingType) != 0 {
			billingType, err := abi.ParseDoDSettleBillingType(v.DynamicParam.BillingType)
			if err != nil {
				return nil, err
			}
			conn.BillingType = billingType
		}
		var billingUnit abi.DoDSettleBillingUnit
		var err error
		if len(v.DynamicParam.BillingUnit) > 0 {
			billingUnit, err = abi.ParseDoDSettleBillingUnit(v.DynamicParam.BillingUnit)
			if err != nil {
				return nil, err
			}
			conn.BillingUnit = billingUnit
		}
		if len(v.DynamicParam.ServiceClass) > 0 {
			serviceClass, err := abi.ParseDoDSettleServiceClass(v.DynamicParam.ServiceClass)
			if err != nil {
				return nil, err
			}
			conn.ServiceClass = serviceClass
		}
		if len(v.DynamicParam.Currency) != 0 {
			conn.Currency = v.DynamicParam.Currency
		}
		if len(v.DynamicParam.Bandwidth) != 0 {
			conn.Bandwidth = v.DynamicParam.Bandwidth
		}
		if len(v.DynamicParam.ConnectionName) != 0 {
			conn.ConnectionName = v.DynamicParam.ConnectionName
		}
		if v.DynamicParam.StartTime != 0 {
			conn.StartTime = v.DynamicParam.StartTime
		}
		if v.DynamicParam.EndTime != 0 {
			conn.EndTime = v.DynamicParam.EndTime
		}
		if v.DynamicParam.Price != 0 {
			conn.Price = float64(v.DynamicParam.Price)
		}
		op.Connections = append(op.Connections, conn)
	}
	return op, nil
}
