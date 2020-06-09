package grpcServer

import (
	"context"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/contract"
	"github.com/qlcchain/go-lsobus/log"
	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

type OrderApi struct {
	logger *zap.SugaredLogger
	cs     *contract.ContractService
}

func NewOrderApi(cs *contract.ContractService) *OrderApi {
	return &OrderApi{
		cs:     cs,
		logger: log.NewLogger("OrderApi"),
	}
}

func (oa *OrderApi) CreateOrder(ctx context.Context, param *proto.CreateOrderParam) (*proto.OrderId, error) {
	oa.logger.Debugf("CreateOrder name %s", param.String())
	id, err := oa.cs.GetCreateOrderBlock(param)
	if err != nil {
		oa.logger.Debugf("GetCreateOrderBlock err %s", err)
		return nil, err
	}
	return &proto.OrderId{
		InternalId: id,
	}, nil
}

func (oa *OrderApi) ChangeOrder(ctx context.Context, param *proto.ChangeOrderParam) (*proto.OrderId, error) {
	oa.logger.Debugf("ChangeOrder name %s", param.String())
	id, err := oa.cs.GetChangeOrderBlock(param)
	if err != nil {
		oa.logger.Debugf("GetChangeOrderBlock err %s", err)
		return nil, err
	}
	return &proto.OrderId{
		InternalId: id,
	}, nil
}

func (oa *OrderApi) TerminateOrder(ctx context.Context, param *proto.TerminateOrderParam) (*proto.OrderId, error) {
	oa.logger.Debugf("TerminateOrder name %s", param.String())
	id, err := oa.cs.GetTerminateOrderBlock(param)
	if err != nil {
		oa.logger.Debugf("GetTerminateOrderBlock err %s", err)
		return nil, err
	}
	return &proto.OrderId{
		InternalId: id,
	}, nil
}

func (oa *OrderApi) GetOrderInfo(ctx context.Context, id *proto.GetOrderInfoByInternalId) (*proto.OrderInfo, error) {
	oa.logger.Debugf("GetOrderInfo name %s", id.String())
	orderInfo, err := oa.cs.GetOrderInfoByInternalId(id.InternalId)
	if err != nil {
		oa.logger.Debugf("GetOrderInfoByInternalId err %s", err)
		return nil, err
	}
	info := new(proto.OrderInfo)
	info.Buyer = &proto.User{
		Address: orderInfo.Buyer.Address.String(),
		Name:    orderInfo.Buyer.Name,
	}
	info.Seller = &proto.User{
		Address: orderInfo.Seller.Address.String(),
		Name:    orderInfo.Seller.Name,
	}
	info.OrderId = orderInfo.OrderId
	info.OrderType = orderInfo.OrderType.String()
	info.OrderState = orderInfo.OrderState.String()
	info.ContractState = orderInfo.ContractState.String()
	for _, v := range orderInfo.Connections {
		conn := &proto.ConnectionParam{
			StaticParam: &proto.ConnectionStaticParam{
				ItemId:            v.ItemId,
				BuyerProductId:    v.BuyerProductId,
				ProductOfferingId: v.ProductOfferingId,
				ProductId:         v.ProductId,
				SrcCompanyName:    v.SrcCompanyName,
				SrcRegion:         v.SrcRegion,
				SrcCity:           v.SrcCity,
				SrcDataCenter:     v.DstDataCenter,
				SrcPort:           v.SrcPort,
				DstCompanyName:    v.DstCompanyName,
				DstRegion:         v.DstRegion,
				DstCity:           v.DstCity,
				DstDataCenter:     v.DstDataCenter,
				DstPort:           v.DstPort,
			},
			DynamicParam: &proto.ConnectionDynamicParam{
				OrderId:        v.OrderId,
				QuoteId:        v.QuoteId,
				QuoteItemId:    v.QuoteItemId,
				ConnectionName: v.ConnectionName,
				Bandwidth:      v.Bandwidth,
				//BillingUnit:    v.BillingUnit.String(),
				Price:    float32(v.Price),
				Addition: float32(v.Addition),
				//ServiceClass:   v.ServiceClass.String(),
				//PaymentType:    v.PaymentType.String(),
				//BillingType:    v.BillingType.String(),
				Currency:          v.Currency,
				StartTime:         v.StartTime,
				StartTimeStr:      v.StartTimeStr,
				EndTime:           v.EndTime,
				EndTimeStrTimeStr: v.EndTimeStr,
			},
		}
		if v.BillingUnit.String() != "null" {
			conn.DynamicParam.BillingUnit = v.BillingUnit.String()
		}
		if v.ServiceClass.String() != "null" {
			conn.DynamicParam.ServiceClass = v.ServiceClass.String()
		}
		if v.PaymentType.String() != "null" {
			conn.DynamicParam.PaymentType = v.PaymentType.String()
		}
		if v.BillingType.String() != "null" {
			conn.DynamicParam.BillingType = v.BillingType.String()
		}
		info.Connections = append(info.Connections, conn)
	}
	for _, v := range orderInfo.Track {
		t := &proto.OrderLifeTrack{
			ContractState: v.ContractState.String(),
			OrderState:    v.OrderState.String(),
			Reason:        v.Reason,
			Time:          v.Time,
			Hash:          v.Hash.String(),
		}
		info.Track = append(info.Track, t)
	}
	return info, nil
}
