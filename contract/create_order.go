package contract

import (
	"errors"
	"strconv"
	"time"

	"github.com/iixlabs/virtual-lsobus/rpc/grpc/proto"

	"github.com/qlcchain/go-qlc/common/types"
	abi "github.com/qlcchain/go-qlc/vm/contract/abi"
)

func (cs *ContractService) GetCreateOrderBlock(param *proto.CreateOrderParam) (string, error) {
	/* TODO: Generate a block to create an order
		1. call dod_settlement_getCreateOrderBlock to creat an order,it will return an internal id
	    2. sign orderBlock and process it to the chain
		3. periodically check whether this order has been signed and confirmed through internal id
		4. if order has been signed and confirmed,call orchestra interface to order at the sonata service,
	       will return an external order id
		5. call dod_settlement_getUpdateOrderInfoBlock to update real orderId to qlc chain
		6. call orchestra interface to periodically check whether the resource of this order has been ready?
		7. if resource is ready,call dod_settlement_getResourceReadyBlock periodically check whether the resource of this order has been ready?
	*/
	addr := cs.account.Address().String()
	if addr == param.Buyer.Address {
		op, err := cs.convertProtoToOrderParam(param)
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
		return internalId, nil
	} else {
		cs.logger.Errorf("buyer address not match,have %s,want %s", param.Buyer.Address, addr)
	}
	//	go cs.CheckCreateOrderContractConfirmed(block.Previous.String())
	return "", errors.New("buyer address not match")
}

func (cs *ContractService) convertProtoToOrderParam(param *proto.CreateOrderParam) (*abi.DoDSettleCreateOrderParam, error) {
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
	for _, v := range param.Cps {
		paymentType, err := abi.ParseDoDSettlePaymentType(v.DynamicParam.PaymentType)
		if err != nil {
			return nil, err
		}

		billingType, err := abi.ParseDoDSettleBillingType(v.DynamicParam.BillingType)
		if err != nil {
			return nil, err
		}

		billingUnit, err := abi.ParseDoDSettleBillingUnit(v.DynamicParam.BillingUnit)
		if err != nil {
			return nil, err
		}

		price, err := strconv.ParseFloat(v.DynamicParam.Price, 64)
		if err != nil {
			return nil, err
		}

		serviceClass, err := abi.ParseDoDSettleServiceClass(v.DynamicParam.ServiceClass)
		if err != nil {
			return nil, err
		}
		conn := &abi.DoDSettleConnectionParam{
			DoDSettleConnectionStaticParam: abi.DoDSettleConnectionStaticParam{
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
				ConnectionName: v.DynamicParam.ConnectionName,
				Bandwidth:      v.DynamicParam.Bandwidth,
				BillingUnit:    billingUnit,
				Price:          price,
				ServiceClass:   serviceClass,
				PaymentType:    paymentType,
				BillingType:    billingType,
				Currency:       v.DynamicParam.ConnectionName,
				StartTime:      v.DynamicParam.StartTime,
				EndTime:        v.DynamicParam.EndTime,
			},
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

func (cs *ContractService) CheckCreateOrderResourceReady(externalId string) bool {
	return true
}
