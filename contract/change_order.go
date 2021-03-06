package contract

import (
	"errors"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"

	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

func (cs *ContractCaller) GetChangeOrderBlock(param *proto.ChangeOrderParam) (string, error) {
	addr := cs.seller.Account().Address().String()
	if addr == param.Buyer.Address {
		op, err := cs.convertProtoToChangeOrderParam(param)
		if err != nil {
			return "", err
		}
		blk, err := cs.seller.GetChangeOrderBlock(op)
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

func (cs *ContractCaller) convertProtoToChangeOrderParam(param *proto.ChangeOrderParam) (*qlcSdk.DoDSettleChangeOrderParam, error) {
	sellerAddr, _ := pkg.HexToAddress(param.Seller.Address)
	buyAddr, _ := pkg.HexToAddress(param.Buyer.Address)
	op := new(qlcSdk.DoDSettleChangeOrderParam)
	op.Buyer = &qlcSdk.DoDSettleUser{
		Address: buyAddr,
		Name:    param.Buyer.Name,
	}
	op.Seller = &qlcSdk.DoDSettleUser{
		Address: sellerAddr,
		Name:    param.Seller.Name,
	}
	for _, v := range param.ChangeConnectionParam {
		conn := new(qlcSdk.DoDSettleChangeConnectionParam)
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

		if len(v.DynamicParam.ItemId) != 0 {
			conn.ItemId = v.DynamicParam.ItemId
		}

		if len(v.DynamicParam.OrderItemId) != 0 {
			conn.OrderItemId = v.DynamicParam.OrderItemId
		}

		if len(v.DynamicParam.InternalId) != 0 {
			conn.InternalId = v.DynamicParam.InternalId
		}

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
