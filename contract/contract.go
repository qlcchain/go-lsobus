package contract

import (
	"context"
	"time"

	"github.com/qlcchain/go-qlc/common/types"
	qlcchain "github.com/qlcchain/go-qlc/rpc/api"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
	rpc "github.com/qlcchain/jsonrpc2"
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/common"
	"github.com/qlcchain/go-lsobus/common/event"
	"github.com/qlcchain/go-lsobus/config"
	"github.com/qlcchain/go-lsobus/log"
	ct "github.com/qlcchain/go-lsobus/services/context"
)

const (
	checkContractInterval    = 1 * time.Minute
	connectRpcServerInterval = 5 * time.Second
)

type ContractService struct {
	cfg        *config.Config
	account    *types.Account
	logger     *zap.SugaredLogger
	client     *rpc.Client
	ctx        context.Context
	cancel     context.CancelFunc
	handlerIds map[common.TopicType]string
	eb         event.EventBus
	chainReady bool
	quit       chan bool
}

func NewContractService(cfgFile string) (*ContractService, error) {
	cc := ct.NewServiceContext(cfgFile)
	cfg, err := cc.Config()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	cs := &ContractService{
		cfg:        cfg,
		account:    cc.Account(),
		logger:     log.NewLogger("contract"),
		ctx:        ctx,
		cancel:     cancel,
		handlerIds: make(map[common.TopicType]string),
		eb:         cc.EventBus(),
		quit:       make(chan bool, 1),
	}
	return cs, nil
}

func (cs *ContractService) Init() error {
	return nil
}

func (cs *ContractService) Start() error {
	go cs.checkDoDContract()
	go cs.connectRpcServer()
	return nil
}

func (cs *ContractService) checkDoDContract() {
	ticker := time.NewTicker(checkContractInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.chainReady {
				cs.processDoDContract()
			}
		}
	}
}

func (cs *ContractService) processDoDContract() {

	/* TODO: automatically verify order and sign the settlement smartcontract between CBC and PCCWG
		1. call dod_settlement_getPendingRequest to check DoD Contract
	    2. if there is a contract that needs to be signed,take out order data
		3. call orchestra interface to verify order information
		4. call dod_settlement_getResponseBlock update contract action(confirm/reject) and sign responseBlock
		5. call process to process responseBlock
	*/

	dod := make([]*qlcchain.DoDPendingRequestRsp, 0)
	addr := cs.account.Address()
	err := cs.client.Call(&dod, "DoDSettlement_getPendingRequest", &addr)
	if err != nil || len(dod) == 0 {
		return
	}
	for _, v := range dod {
		action, err := abi.ParseDoDSettleResponseAction("confirm")
		if err != nil {
			cs.logger.Error(err)
			continue
		}

		param := &abi.DoDSettleResponseParam{
			RequestHash: v.Hash,
			Action:      action,
		}
		block := new(types.StateBlock)
		err = cs.client.Call(&block, "DoDSettlement_getCreateOrderRewardBlock", param)
		if err != nil {
			cs.logger.Error(err)
			continue
		}

		var w types.Work
		worker, _ := types.NewWorker(w, block.Root())
		block.Work = worker.NewWork()

		hash := block.GetHash()
		block.Signature = cs.account.Sign(hash)

		var h types.Hash
		err = cs.client.Call(&h, "ledger_process", &block)
		if err != nil {
			cs.logger.Error(err)
			continue
		}
	}
}

func (cs *ContractService) connectRpcServer() {
	ticker := time.NewTicker(connectRpcServerInterval)
	for {
		select {
		case <-cs.quit:
			return
		case <-ticker.C:
			if cs.cfg.ChainUrl != "" {
				if cs.client == nil {
					client, err := rpc.Dial(cs.cfg.ChainUrl)
					if err != nil || client == nil {
						continue
					} else {
						cs.client = client
						var pov qlcchain.PovStatus
						err := cs.client.Call(&pov, "pov_getPovStatus")
						if err != nil {
							continue
						} else if pov.SyncState == 2 {
							cs.chainReady = true
							cs.quit <- true
						}
					}
				} else {
					var pov qlcchain.PovStatus
					err := cs.client.Call(&pov, "pov_getPovStatus")
					if err != nil {
						continue
					} else if pov.SyncState == 2 {
						cs.chainReady = true
						cs.quit <- true
					}
				}
			}
		}
	}
}

func (cs *ContractService) GetOrderInfoByInternalId(id string) (*abi.DoDSettleOrderInfo, error) {
	orderInfo := new(abi.DoDSettleOrderInfo)
	err := cs.client.Call(&orderInfo, "DoDSettlement_getOrderInfoByInternalId", &id)
	if err != nil {
		cs.logger.Error(err)
		return nil, err
	}
	return orderInfo, nil
}

func (cs *ContractService) Stop() error {
	//this must be the first step
	cs.cancel()
	err := cs.unsubscribeEvent()
	if err != nil {
		return err
	}
	if cs.client != nil {
		_ = cs.client.Close
	}
	return nil
}

func (cs *ContractService) unsubscribeEvent() error {
	for k, v := range cs.handlerIds {
		if err := cs.eb.Unsubscribe(k, v); err != nil {
			return err
		}
	}
	return nil
}
