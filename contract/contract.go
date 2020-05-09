package contract

import (
	"context"
	"time"

	"github.com/iixlabs/virtual-lsobus/common"
	"github.com/iixlabs/virtual-lsobus/common/event"
	"github.com/iixlabs/virtual-lsobus/log"

	"github.com/iixlabs/virtual-lsobus/config"
	ct "github.com/iixlabs/virtual-lsobus/services/context"
	qlcchain "github.com/qlcchain/qlc-go-sdk"
	qlctypes "github.com/qlcchain/qlc-go-sdk/pkg/types"
	"go.uber.org/zap"
)

const (
	checkContractInterval    = 1 * time.Minute
	connectRpcServerInterval = 5 * time.Second
)

type ContractService struct {
	cfg        *config.Config
	account    *qlctypes.Account
	logger     *zap.SugaredLogger
	client     *qlcchain.QLCClient
	ctx        context.Context
	cancel     context.CancelFunc
	handlerIds map[common.TopicType]string
	eb         event.EventBus
	chainReady bool
	quit       chan bool
}

func NewInteractiveService(cfgFile string) (*ContractService, error) {
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
			cs.processDoDContract()
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
					client, err := qlcchain.NewQLCClient(cs.cfg.ChainUrl)
					if err != nil || client == nil {
						continue
					} else {
						cs.client = client
						s, err := cs.client.Pov.GetPovStatus()
						if err != nil {
							continue
						} else if s.SyncState == 2 {
							cs.chainReady = true
							cs.quit <- true
						}
					}
				} else {
					s, err := cs.client.Pov.GetPovStatus()
					if err != nil {
						continue
					} else if s.SyncState == 2 {
						cs.chainReady = true
						cs.quit <- true
					}
				}
			}
		}
	}
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
