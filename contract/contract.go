package contract

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/qlcchain/go-lsobus/mock"

	"github.com/qlcchain/go-lsobus/orchestra"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/common"
	"github.com/qlcchain/go-lsobus/common/event"
	"github.com/qlcchain/go-lsobus/config"
	"github.com/qlcchain/go-lsobus/log"
	ct "github.com/qlcchain/go-lsobus/services/context"
)

const (
	checkNeedSignContractInterval = 15 * time.Second
	checkContractStatusInterval   = 10 * time.Second
	checkOrderStatusInterval      = 10 * time.Second
	checkProductInterval          = 10 * time.Second
	connectRpcServerInterval      = 5 * time.Second
	processingOrderList           = "processingOrder.json" // a collection of processing order
)

var (
	chainNotReady     = errors.New("chain is not ready")
	buyerAddrNotMatch = errors.New("buyer address not match")
	noInventoryList   = errors.New("no inventory list ")
)

type OrderList struct {
	Role         string `json:"role"`
	ChainOrderID string `json:"chainOrderID"`
}

type ProcessingOrderList struct {
	Processing []*OrderList `json:"processing"`
}

type ContractService struct {
	cfg                  *config.Config
	account              *types.Account
	logger               *zap.SugaredLogger
	client               *qlcSdk.QLCClient
	ctx                  context.Context
	cancel               context.CancelFunc
	handlerIds           map[common.TopicType]string
	eb                   event.EventBus
	chainReady           bool
	quit                 chan bool
	orderIdOnChainSeller *sync.Map
	orderIdOnChainBuyer  *sync.Map
	orchestra            *orchestra.Orchestra
	fakeMode             bool
	mutex                *sync.Mutex
}

type Product struct {
	orderItemID string
	productID   string
}

func NewContractService(cfgFile string) (*ContractService, error) {
	cc := ct.NewServiceContext(cfgFile)
	cfg, err := cc.Config()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	or := orchestra.NewOrchestra(cfgFile)
	or.SetFakeMode(cfg.FakeMode)
	cs := &ContractService{
		cfg:                  cfg,
		account:              cc.Account(),
		logger:               log.NewLogger("contract"),
		ctx:                  ctx,
		cancel:               cancel,
		handlerIds:           make(map[common.TopicType]string),
		eb:                   cc.EventBus(),
		quit:                 make(chan bool, 1),
		orderIdOnChainSeller: new(sync.Map),
		orderIdOnChainBuyer:  new(sync.Map),
		orchestra:            or,
		mutex:                new(sync.Mutex),
	}
	return cs, nil
}

func (cs *ContractService) SetFakeMode(mode bool) {
	cs.fakeMode = mode
}

func (cs *ContractService) GetFakeMode() bool {
	return cs.fakeMode
}

func (cs *ContractService) GetAccount() *types.Account {
	return cs.account
}

func (cs *ContractService) SetAccount(account *types.Account) {
	cs.account = account
}

func (cs *ContractService) GetOrchestra() *orchestra.Orchestra {
	return cs.orchestra
}

func (cs *ContractService) Init() error {
	err := cs.orchestra.Init()
	if err != nil {
		return err
	}
	err = cs.readProcessingOrder()
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractService) readProcessingOrder() error {
	file := filepath.Join(cs.cfg.DataDir, processingOrderList)
	_, err := os.Stat(file)
	if err != nil {
		f, _ := os.Create(file)
		defer f.Close()
		return nil
	}
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if len(f) == 0 {
		return nil
	}
	var po ProcessingOrderList
	err = json.Unmarshal(f, &po)
	if err != nil {
		return err
	}
	for _, v := range po.Processing {
		if v.Role == "buyer" {
			cs.orderIdOnChainBuyer.Store(v.ChainOrderID, "")
		}
		if v.Role == "seller" {
			cs.orderIdOnChainSeller.Store(v.ChainOrderID, "")
		}
	}
	return nil
}

func (cs *ContractService) readAndWriteProcessingOrder(action, role, chainOrderID string) error {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	f, err := ioutil.ReadFile(filepath.Join(cs.cfg.DataDir, processingOrderList))
	if err != nil {
		return err
	}
	var po ProcessingOrderList
	switch action {
	case "add":
		if len(f) != 0 {
			err = json.Unmarshal(f, &po)
			if err != nil {
				return err
			}
		}
		info := &OrderList{
			Role:         role,
			ChainOrderID: chainOrderID,
		}
		po.Processing = append(po.Processing, info)
		b, err := json.Marshal(po)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(cs.cfg.DataDir, processingOrderList), b, 0600)
		if err != nil {
			return err
		}
	case "delete":
		if len(f) == 0 {
			cs.logger.Warnf("chain order id %s does not exit in the file", chainOrderID)
			return nil
		}
		err = json.Unmarshal(f, &po)
		if err != nil {
			return err
		}
		var temp ProcessingOrderList
		for _, v := range po.Processing {
			if v.ChainOrderID != chainOrderID {
				temp.Processing = append(temp.Processing, v)
			}
		}
		b, err := json.Marshal(temp)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(cs.cfg.DataDir, processingOrderList), b, 0600)
		if err != nil {
			return err
		}
	default:
		return errors.New("unKnow action")
	}
	return nil
}

func (cs *ContractService) Start() error {
	go cs.checkDoDContract()
	go cs.connectRpcServer()
	go cs.checkContractStatus()
	go cs.checkProductStatus()
	go cs.checkProduct()
	return nil
}

func (cs *ContractService) GetOrderInfoByInternalId(id string) (*qlcSdk.DoDSettleOrderInfo, error) {
	if cs.GetFakeMode() {
		return mock.GetOrderInfoByInternalId(id)
	}
	if cs.chainReady {
		orderInfo, err := cs.client.DoDSettlement.GetOrderInfoByInternalId(id)
		if err != nil {
			cs.logger.Error(err)
			return nil, err
		}
		return orderInfo, nil
	} else {
		return nil, chainNotReady
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

func (cs *ContractService) processBlockAndWaitConfirmed(block *types.StateBlock) error {
	_, err := cs.client.Ledger.Process(block)
	if err != nil {
		return fmt.Errorf("process block error: %s", err)
	}
	return cs.waitBlockConfirmed(block.GetHash())
}

func (cs *ContractService) waitBlockConfirmed(hash types.Hash) error {
	t := time.NewTimer(time.Second * 180)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			return errors.New("consensus confirmed timeout")
		default:
			confirmed, err := cs.client.Ledger.BlockConfirmedStatus(hash)
			if err != nil {
				return err
			}
			if confirmed {
				return nil
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}
}
