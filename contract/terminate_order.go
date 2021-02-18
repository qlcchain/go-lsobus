package contract

import (
	"errors"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

func (cs *ContractCaller) GetTerminateOrderBlock(param *proto.TerminateOrderParam) (string, error) {
	addr := cs.seller.Account().Address().String()
	if addr == param.Buyer.Address {
		op, err := cs.convertProtoToTerminateOrderParam(param)
		if err != nil {
			return "", err
		}
		blk, err := cs.seller.GetTerminateOrderBlock(op)
		if err != nil {
			return "", err
		}
		hash, err := cs.seller.Process(blk)
		if err != nil {
			cs.logger.Errorf("process block error: %s", err)
			return "", err
		}
		cs.logger.Infof("process hash %s success", hash.String())

		internalId := blk.Previous.String()
		err = cs.readAndWriteProcessingOrder("add", "buyer", internalId)
		if err != nil {
			return "", err
		}
		cs.orderIdOnChainBuyer.Store(internalId, "")
		return internalId, nil
	} else {
		cs.logger.Errorf("buyer address not match,have %s,want %s", param.Buyer.Address, addr)
		return "", buyerAddrNotMatch
	}
}

func (cs *ContractCaller) convertProtoToTerminateOrderParam(param *proto.TerminateOrderParam) (*qlcSdk.DoDSettleTerminateOrderParam, error) {
	sellerAddr, _ := pkg.HexToAddress(param.Seller.Address)
	buyAddr, _ := pkg.HexToAddress(param.Buyer.Address)
	op := new(qlcSdk.DoDSettleTerminateOrderParam)
	op.Buyer = &qlcSdk.DoDSettleUser{
		Address: buyAddr,
		Name:    param.Buyer.Name,
	}
	op.Seller = &qlcSdk.DoDSettleUser{
		Address: sellerAddr,
		Name:    param.Seller.Name,
	}
	if len(param.TerminateConnectionParam) == 0 {
		return nil, errors.New("param can not be nil")
	}
	for _, v := range param.TerminateConnectionParam {
		conn := new(qlcSdk.DoDSettleChangeConnectionParam)
		if len(v.DynamicParam.PaymentType) != 0 {
			paymentType, err := qlcSdk.ParseDoDSettlePaymentType(v.DynamicParam.PaymentType)
			if err != nil {
				return nil, err
			}
			conn.PaymentType = paymentType
		}
		if len(v.DynamicParam.BillingType) != 0 {
			billingType, err := qlcSdk.ParseDoDSettleBillingType(v.DynamicParam.BillingType)
			if err != nil {
				return nil, err
			}
			conn.BillingType = billingType
		}
		var billingUnit qlcSdk.DoDSettleBillingUnit
		var err error
		if len(v.DynamicParam.BillingUnit) > 0 {
			billingUnit, err = qlcSdk.ParseDoDSettleBillingUnit(v.DynamicParam.BillingUnit)
			if err != nil {
				return nil, err
			}
			conn.BillingUnit = billingUnit
		}
		if len(v.DynamicParam.ServiceClass) > 0 {
			serviceClass, err := qlcSdk.ParseDoDSettleServiceClass(v.DynamicParam.ServiceClass)
			if err != nil {
				return nil, err
			}
			conn.ServiceClass = serviceClass
		}
		conn = &qlcSdk.DoDSettleChangeConnectionParam{
			ProductId: v.ProductId,
			DoDSettleConnectionDynamicParam: qlcSdk.DoDSettleConnectionDynamicParam{
				OrderId:        v.DynamicParam.OrderId,
				InternalId:     v.DynamicParam.InternalId,
				ItemId:         v.DynamicParam.ItemId,
				OrderItemId:    v.DynamicParam.OrderItemId,
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
