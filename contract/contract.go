package contract

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/orchestra"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/config"
	"github.com/qlcchain/go-lsobus/log"
	ct "github.com/qlcchain/go-lsobus/services/context"
)

const (
	checkNeedSignContractInterval = 10 * time.Second
	checkContractStatusInterval   = 8 * time.Second
	checkOrderStatusInterval      = 10 * time.Second
	checkProductInterval          = 8 * time.Second
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

type ContractCaller struct {
	cfg                  *config.Config
	logger               *zap.SugaredLogger
	ctx                  context.Context
	cancel               context.CancelFunc
	orderIdOnChainSeller *sync.Map
	orderIdOnChainBuyer  *sync.Map
	seller               api.DoDSeller
	mutex                *sync.Mutex
}

type Product struct {
	orderItemID string
	productID   string
}

func NewContractService(cfgFile string) (*ContractCaller, error) {
	cc := ct.NewServiceContext(cfgFile)
	cfg, err := cc.Config()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	seller, err := orchestra.NewSeller(ctx, cfgFile)
	if err != nil {
		defer cancel()
		return nil, err
	}
	cs := &ContractCaller{
		cfg:                  cfg,
		logger:               log.NewLogger("contract"),
		ctx:                  ctx,
		cancel:               cancel,
		orderIdOnChainSeller: new(sync.Map),
		orderIdOnChainBuyer:  new(sync.Map),
		seller:               seller,
		mutex:                new(sync.Mutex),
	}
	return cs, nil
}

func (cs *ContractCaller) GetOrchestra() api.DoDSeller {
	return cs.seller
}

func (cs *ContractCaller) Init() error {
	err := cs.readProcessingOrder()
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractCaller) readProcessingOrder() error {
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

func (cs *ContractCaller) readAndWriteProcessingOrder(action, role, chainOrderID string) error {
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

func (cs *ContractCaller) Start() error {
	go cs.checkDoDContract()
	go cs.checkContractStatus()
	go cs.checkProductStatus()
	go cs.checkProduct()
	return nil
}

func (cs *ContractCaller) Stop() error {
	//this must be the first step
	cs.cancel()
	return nil
}
