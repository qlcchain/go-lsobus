package contract

import (
	"errors"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/vm/contract/abi"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

func (cs *ContractService) GetChangeOrderBlock(param *proto.ChangeOrderParam) (string, error) {
	if cs.chainReady {
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
	} else {
		return "", errors.New("chain is not ready")
	}
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
	for _, v := range param.ChangeConnectionParam {
		conn := new(abi.DoDSettleChangeConnectionParam)
		if len(v.DynamicParam.QuoteId) == 0 {
			return nil, errors.New("quote id can not be nil")
		}
		conn.QuoteId = v.DynamicParam.QuoteId
		if len(v.ProductId) == 0 {
			return nil, errors.New("product id can not be nil")
		}
		conn.ProductId = v.ProductId
		if len(v.DynamicParam.QuoteItemId) != 0 {
			conn.QuoteItemId = v.DynamicParam.QuoteItemId
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
		if len(v.DynamicParam.StartTimeStr) != 0 {
			conn.StartTimeStr = v.DynamicParam.StartTimeStr
		}
		if v.DynamicParam.EndTime != 0 {
			conn.EndTime = v.DynamicParam.EndTime
		}
		if len(v.DynamicParam.EndTimeStrTimeStr) != 0 {
			conn.EndTimeStr = v.DynamicParam.EndTimeStrTimeStr
		}
		if v.DynamicParam.Price != 0 {
			conn.Price = float64(v.DynamicParam.Price)
		}
		if v.DynamicParam.Addition != 0 {
			conn.Addition = float64(v.DynamicParam.Addition)
		}
		op.Connections = append(op.Connections, conn)
	}
	return op, nil
}
