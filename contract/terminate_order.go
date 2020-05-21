package contract

import (
	"errors"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
)

func (cs *ContractService) GetTerminateOrderBlock(param *proto.TerminateOrderParam) (string, error) {
	/* TODO: Genearate a block to terminate an order
		1. call dod_settlement_getTerminateOrderBlock to terminate an order,need order's id generated before
	    2. sign orderBlock and process it to the chain
		3. periodically check whether this order has been signed and confirmed through internal id
		4. if order has been signed and confirmed,call orchestra interface to order at the sonata service
		5. call orchestra interface to periodically check whether the resource of this order has been ready?
		6. if resource is ready,call dod_settlement_getResourceReadyBlock periodically check whether the resource of this order has been ready?
	*/
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
		return internalId, nil
	} else {
		cs.logger.Errorf("buyer address not match,have %s,want %s", param.Buyer.Address, addr)
	}
	return "", errors.New("buyer address not match")
}

func (cs *ContractService) CheckTerminateOrderContractSignStatus(internalId string) bool {
	return true
}

func (cs *ContractService) CheckTerminateOrderResourceReady(externalId string) bool {
	return true
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
	op.ProductId = param.ProductId
	return op, nil
}
