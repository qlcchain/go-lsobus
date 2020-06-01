package contract

import (
	"errors"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/vm/contract/abi"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

func (cs *ContractService) GetTerminateOrderBlock(param *proto.TerminateOrderParam) (string, error) {
	if cs.chainReady {
		addr := cs.account.Address().String()
		if addr == param.Buyer.Address {
			op, err := cs.convertProtoToTerminateOrderParam(param)
			if err != nil {
				return "", err
			}
			block := new(types.StateBlock)
			err = cs.client.Call(&block, "DoDSettlement_getTerminateOrderBlock", op)
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

func (cs *ContractService) convertProtoToTerminateOrderParam(param *proto.TerminateOrderParam) (*abi.DoDSettleTerminateOrderParam, error) {
	sellerAddr, _ := types.HexToAddress(param.Seller.Address)
	buyAddr, _ := types.HexToAddress(param.Buyer.Address)
	op := new(abi.DoDSettleTerminateOrderParam)
	op.Buyer = &abi.DoDSettleUser{
		Address: buyAddr,
		Name:    param.Buyer.Name,
	}
	op.Seller = &abi.DoDSettleUser{
		Address: sellerAddr,
		Name:    param.Seller.Name,
	}
	if len(param.TerminateConnectionParam) == 0 {
		return nil, errors.New("param can not be nil")
	}
	for _, v := range param.TerminateConnectionParam {
		conn := new(abi.DoDSettleChangeConnectionParam)
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
		conn = &abi.DoDSettleChangeConnectionParam{
			ProductId: v.ProductId,
			DoDSettleConnectionDynamicParam: abi.DoDSettleConnectionDynamicParam{
				OrderId:        v.DynamicParam.OrderId,
				QuoteId:        v.DynamicParam.QuoteId,
				QuoteItemId:    v.DynamicParam.QuoteItemId,
				ConnectionName: v.DynamicParam.ConnectionName,
				Currency:       v.DynamicParam.Currency,
				Bandwidth:      v.DynamicParam.Bandwidth,
				Price:          float64(v.DynamicParam.Price),
				Addition:       float64(v.DynamicParam.Addition),
				StartTime:      v.DynamicParam.StartTime,
				StartTimeStr:   v.DynamicParam.StartTimeStr,
				EndTime:        v.DynamicParam.EndTime,
				EndTimeStr:     v.DynamicParam.EndTimeStrTimeStr,
			},
		}
		op.Connections = append(op.Connections, conn)
	}
	return op, nil
}
