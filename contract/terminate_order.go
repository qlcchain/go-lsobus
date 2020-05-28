package contract

import (
	"errors"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
)

func (cs *ContractService) GetTerminateOrderBlock(param *proto.TerminateOrderParam) (string, error) {
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
	if len(param.ProductId) == 0 {
		return nil, errors.New("product can not be nil")
	}
	op.ProductId = param.ProductId
	return op, nil
}
