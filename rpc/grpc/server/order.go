package grpcServer

import (
	"context"

	"github.com/qlcchain/go-lsobus/contract"
	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

type OrderApi struct {
	cs *contract.ContractService
}

func NewOrderApi(cs *contract.ContractService) *OrderApi {
	return &OrderApi{
		cs: cs,
	}
}

func (oa *OrderApi) CreateOrder(ctx context.Context, param *proto.CreateOrderParam) (*proto.OrderRsp, error) {
	id, err := oa.cs.GetCreateOrderBlock(param)
	if err != nil {
		return nil, err
	}
	return &proto.OrderRsp{
		InternalId: id,
	}, nil
}

func (oa *OrderApi) ChangeOrder(ctx context.Context, param *proto.ChangeOrderParam) (*proto.OrderRsp, error) {
	id, err := oa.cs.GetChangeOrderBlock(param)
	if err != nil {
		return nil, err
	}
	return &proto.OrderRsp{
		InternalId: id,
	}, nil
}

func (oa *OrderApi) TerminateOrder(ctx context.Context, param *proto.TerminateOrderParam) (*proto.OrderRsp, error) {
	id, err := oa.cs.GetTerminateOrderBlock(param)
	if err != nil {
		return nil, err
	}
	return &proto.OrderRsp{
		InternalId: id,
	}, nil
}

func (oa *OrderApi) OrderInfo(ctx context.Context, id *proto.GetOrderInfoByInternalId) (*proto.OrderInfoRsp, error) {
	orderInfo, err := oa.cs.GetOrderInfoByInternalId(id.InternalId)
	if err != nil {
		return nil, err
	}
	info := new(proto.OrderInfoRsp)
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
				ProductId:      v.ProductId,
				SrcCompanyName: v.SrcCompanyName,
				SrcRegion:      v.SrcRegion,
				SrcCity:        v.SrcCity,
				SrcDataCenter:  v.DstDataCenter,
				SrcPort:        v.SrcPort,
				DstCompanyName: v.DstCompanyName,
				DstRegion:      v.DstRegion,
				DstCity:        v.DstCity,
				DstDataCenter:  v.DstDataCenter,
				DstPort:        v.DstPort,
			},
			DynamicParam: &proto.ConnectionDynamicParam{
				ConnectionName: v.ConnectionName,
				Bandwidth:      v.Bandwidth,
				BillingUnit:    v.BillingUnit.String(),
				Price:          float32(v.Price),
				ServiceClass:   v.ServiceClass.String(),
				PaymentType:    v.PaymentType.String(),
				BillingType:    v.BillingType.String(),
				Currency:       v.Currency,
				StartTime:      v.StartTime,
				EndTime:        v.EndTime,
			},
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
