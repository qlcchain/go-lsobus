package contract

import (
	"errors"
	"time"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/qlcchain/go-qlc/common/types"
	abi "github.com/qlcchain/go-qlc/vm/contract/abi"
)

func (cs *ContractService) GetCreateOrderBlock(param *proto.CreateOrderParam) (string, error) {
	addr := cs.account.Address().String()
	if addr == param.Buyer.Address {
		op, err := cs.convertProtoToCreateOrderParam(param)
		if err != nil {
			return "", err
		}
		block := new(types.StateBlock)
		err = cs.client.Call(&block, "DoDSettlement_getCreateOrderBlock", op)
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

func (cs *ContractService) convertProtoToCreateOrderParam(param *proto.CreateOrderParam) (*abi.DoDSettleCreateOrderParam, error) {
	sellerAddr, _ := types.HexToAddress(param.Seller.Address)
	buyAddr, _ := types.HexToAddress(param.Buyer.Address)
	op := new(abi.DoDSettleCreateOrderParam)
	op.Buyer = &abi.DoDSettleUser{
		Address: buyAddr,
		Name:    param.Buyer.Name,
	}
	op.Seller = &abi.DoDSettleUser{
		Address: sellerAddr,
		Name:    param.Seller.Name,
	}
	for _, v := range param.ConnectionParam {
		paymentType, err := abi.ParseDoDSettlePaymentType(v.DynamicParam.PaymentType)
		if err != nil {
			return nil, err
		}

		billingType, err := abi.ParseDoDSettleBillingType(v.DynamicParam.BillingType)
		if err != nil {
			return nil, err
		}

		var billingUnit abi.DoDSettleBillingUnit
		if len(v.DynamicParam.BillingUnit) > 0 {
			billingUnit, err = abi.ParseDoDSettleBillingUnit(v.DynamicParam.BillingUnit)
			if err != nil {
				return nil, err
			}
		}

		serviceClass, err := abi.ParseDoDSettleServiceClass(v.DynamicParam.ServiceClass)
		if err != nil {
			return nil, err
		}
		var conn *abi.DoDSettleConnectionParam
		if billingType == abi.DoDSettleBillingTypePAYG {
			conn = &abi.DoDSettleConnectionParam{
				DoDSettleConnectionStaticParam: abi.DoDSettleConnectionStaticParam{
					ItemId:         v.StaticParam.ItemId,
					ProductId:      v.StaticParam.ProductId,
					SrcCompanyName: v.StaticParam.SrcCompanyName,
					SrcRegion:      v.StaticParam.SrcRegion,
					SrcCity:        v.StaticParam.SrcCity,
					SrcDataCenter:  v.StaticParam.SrcDataCenter,
					SrcPort:        v.StaticParam.SrcPort,
					DstCompanyName: v.StaticParam.DstCompanyName,
					DstRegion:      v.StaticParam.DstRegion,
					DstCity:        v.StaticParam.DstCity,
					DstDataCenter:  v.StaticParam.DstDataCenter,
					DstPort:        v.StaticParam.DstPort,
				},
				DoDSettleConnectionDynamicParam: abi.DoDSettleConnectionDynamicParam{
					OrderId:        v.DynamicParam.OrderId,
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
			conn = &abi.DoDSettleConnectionParam{
				DoDSettleConnectionStaticParam: abi.DoDSettleConnectionStaticParam{
					ItemId:         v.StaticParam.ItemId,
					ProductId:      v.StaticParam.ProductId,
					SrcCompanyName: v.StaticParam.SrcCompanyName,
					SrcRegion:      v.StaticParam.SrcRegion,
					SrcCity:        v.StaticParam.SrcCity,
					SrcDataCenter:  v.StaticParam.SrcDataCenter,
					SrcPort:        v.StaticParam.SrcPort,
					DstCompanyName: v.StaticParam.DstCompanyName,
					DstRegion:      v.StaticParam.DstRegion,
					DstCity:        v.StaticParam.DstCity,
					DstDataCenter:  v.StaticParam.DstDataCenter,
					DstPort:        v.StaticParam.DstPort,
				},
				DoDSettleConnectionDynamicParam: abi.DoDSettleConnectionDynamicParam{
					OrderId:        v.DynamicParam.OrderId,
					QuoteItemId:    v.DynamicParam.QuoteItemId,
					ConnectionName: v.DynamicParam.ConnectionName,
					Bandwidth:      v.DynamicParam.Bandwidth,
					Price:          float64(v.DynamicParam.Price),
					ServiceClass:   serviceClass,
					PaymentType:    paymentType,
					BillingType:    billingType,
					Currency:       v.DynamicParam.Currency,
					StartTime:      v.DynamicParam.StartTime,
					EndTime:        v.DynamicParam.EndTime,
				},
			}
		}
		op.Connections = append(op.Connections, conn)
	}
	return op, nil
}

func (cs *ContractService) CheckCreateOrderContractConfirmed(internalId string) {
	ticker := time.NewTicker(connectRpcServerInterval)
	for {
		select {
		case <-ticker.C:
			orderInfo := new(abi.DoDSettleOrderInfo)
			err := cs.client.Call(&orderInfo, "DoDSettlement_getOrderInfoByInternalId", &internalId)
			if err != nil {
				cs.logger.Error(err)
			}
			if orderInfo.ContractState == abi.DoDSettleContractStateConfirmed {
				cs.logger.Infof(" order %s has been sign by seller", internalId)
				return
			}
		}
	}
}
