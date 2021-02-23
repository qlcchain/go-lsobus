package contract

import (
	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

func (cs *ContractCaller) GetCreateOrderBlock(param *proto.CreateOrderParam) (string, error) {
	addr := cs.seller.Account().Address().String()
	if addr == param.Buyer.Address {
		op, err := cs.convertProtoToCreateOrderParam(param)
		if err != nil {
			return "", err
		}
		blk, err := cs.seller.GetCreateOrderBlock(op)
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

func (cs *ContractCaller) convertProtoToCreateOrderParam(param *proto.CreateOrderParam) (*qlcSdk.DoDSettleCreateOrderParam, error) {
	sellerAddr, _ := pkg.HexToAddress(param.Seller.Address)
	buyAddr, _ := pkg.HexToAddress(param.Buyer.Address)
	op := new(qlcSdk.DoDSettleCreateOrderParam)
	op.Buyer = &qlcSdk.DoDSettleUser{
		Address: buyAddr,
		Name:    param.Buyer.Name,
	}
	op.Seller = &qlcSdk.DoDSettleUser{
		Address: sellerAddr,
		Name:    param.Seller.Name,
	}
	for _, v := range param.ConnectionParam {
		paymentType, err := qlcSdk.ParseDoDSettlePaymentType(v.DynamicParam.PaymentType)
		if err != nil {
			return nil, err
		}

		billingType, err := qlcSdk.ParseDoDSettleBillingType(v.DynamicParam.BillingType)
		if err != nil {
			return nil, err
		}

		var billingUnit qlcSdk.DoDSettleBillingUnit
		if len(v.DynamicParam.BillingUnit) > 0 {
			billingUnit, err = qlcSdk.ParseDoDSettleBillingUnit(v.DynamicParam.BillingUnit)
			if err != nil {
				return nil, err
			}
		}

		serviceClass, err := qlcSdk.ParseDoDSettleServiceClass(v.DynamicParam.ServiceClass)
		if err != nil {
			return nil, err
		}
		var conn *qlcSdk.DoDSettleConnectionParam
		if billingType == qlcSdk.DoDSettleBillingTypePAYG {
			conn = &qlcSdk.DoDSettleConnectionParam{
				DoDSettleConnectionStaticParam: qlcSdk.DoDSettleConnectionStaticParam{
					BuyerProductId:    v.StaticParam.BuyerProductId,
					ProductOfferingId: v.StaticParam.ProductOfferingId,
					ProductId:         v.StaticParam.ProductId,
					SrcCompanyName:    v.StaticParam.SrcCompanyName,
					SrcRegion:         v.StaticParam.SrcRegion,
					SrcCity:           v.StaticParam.SrcCity,
					SrcDataCenter:     v.StaticParam.SrcDataCenter,
					SrcPort:           v.StaticParam.SrcPort,
					DstCompanyName:    v.StaticParam.DstCompanyName,
					DstRegion:         v.StaticParam.DstRegion,
					DstCity:           v.StaticParam.DstCity,
					DstDataCenter:     v.StaticParam.DstDataCenter,
					DstPort:           v.StaticParam.DstPort,
				},
				DoDSettleConnectionDynamicParam: qlcSdk.DoDSettleConnectionDynamicParam{
					OrderId:        v.DynamicParam.OrderId,
					InternalId:     v.DynamicParam.InternalId,
					ItemId:         v.DynamicParam.ItemId,
					OrderItemId:    v.DynamicParam.OrderItemId,
					QuoteId:        v.DynamicParam.QuoteId,
					QuoteItemId:    v.DynamicParam.QuoteItemId,
					ConnectionName: v.DynamicParam.ConnectionName,
					Bandwidth:      v.DynamicParam.Bandwidth,
					BillingUnit:    billingUnit,
					Price:          float64(v.DynamicParam.Price),
					ServiceClass:   serviceClass,
					PaymentType:    paymentType,
					BillingType:    billingType,
					Currency:       v.DynamicParam.Currency,
				},
			}
		} else {
			conn = &qlcSdk.DoDSettleConnectionParam{
				DoDSettleConnectionStaticParam: qlcSdk.DoDSettleConnectionStaticParam{
					BuyerProductId:    v.StaticParam.BuyerProductId,
					ProductOfferingId: v.StaticParam.ProductOfferingId,
					ProductId:         v.StaticParam.ProductId,
					SrcCompanyName:    v.StaticParam.SrcCompanyName,
					SrcRegion:         v.StaticParam.SrcRegion,
					SrcCity:           v.StaticParam.SrcCity,
					SrcDataCenter:     v.StaticParam.SrcDataCenter,
					SrcPort:           v.StaticParam.SrcPort,
					DstCompanyName:    v.StaticParam.DstCompanyName,
					DstRegion:         v.StaticParam.DstRegion,
					DstCity:           v.StaticParam.DstCity,
					DstDataCenter:     v.StaticParam.DstDataCenter,
					DstPort:           v.StaticParam.DstPort,
				},
				DoDSettleConnectionDynamicParam: qlcSdk.DoDSettleConnectionDynamicParam{
					OrderId:        v.DynamicParam.OrderId,
					InternalId:     v.DynamicParam.InternalId,
					ItemId:         v.DynamicParam.ItemId,
					OrderItemId:    v.DynamicParam.OrderItemId,
					QuoteId:        v.DynamicParam.QuoteId,
					QuoteItemId:    v.DynamicParam.QuoteItemId,
					ConnectionName: v.DynamicParam.ConnectionName,
					Bandwidth:      v.DynamicParam.Bandwidth,
					Price:          float64(v.DynamicParam.Price),
					Addition:       float64(v.DynamicParam.Addition),
					ServiceClass:   serviceClass,
					PaymentType:    paymentType,
					BillingType:    billingType,
					Currency:       v.DynamicParam.Currency,
					StartTime:      v.DynamicParam.StartTime,
					StartTimeStr:   v.DynamicParam.StartTimeStr,
					EndTime:        v.DynamicParam.EndTime,
					EndTimeStr:     v.DynamicParam.EndTimeStrTimeStr,
				},
			}
		}
		op.Connections = append(op.Connections, conn)
	}
	return op, nil
}
