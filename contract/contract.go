package contract

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/qlcchain/go-lsobus/orchestra"

	"github.com/qlcchain/go-qlc/common/types"
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
	checkNeedSignContractInterval = 30 * time.Second
	checkContractStatusInterval   = 5 * time.Second
	checkOrderStatusInterval      = 5 * time.Second
	checkProductInterval          = 5 * time.Second
	connectRpcServerInterval      = 5 * time.Second
)

type ContractService struct {
	cfg               *config.Config
	account           *types.Account
	logger            *zap.SugaredLogger
	client            *rpc.Client
	ctx               context.Context
	cancel            context.CancelFunc
	handlerIds        map[common.TopicType]string
	eb                event.EventBus
	chainReady        bool
	quit              chan bool
	orderIdOnChain    *sync.Map
	orderIdFromSonata *sync.Map
	orchestra         *orchestra.Orchestra
}

type Product struct {
	buyerProductID string
	productID      string
}

func NewContractService(cfgFile string) (*ContractService, error) {
	cc := ct.NewServiceContext(cfgFile)
	cfg, err := cc.Config()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	or := orchestra.NewOrchestra(cfgFile)
	cs := &ContractService{
		cfg:               cfg,
		account:           cc.Account(),
		logger:            log.NewLogger("contract"),
		ctx:               ctx,
		cancel:            cancel,
		handlerIds:        make(map[common.TopicType]string),
		eb:                cc.EventBus(),
		quit:              make(chan bool, 1),
		orderIdOnChain:    new(sync.Map),
		orderIdFromSonata: new(sync.Map),
		orchestra:         or,
	}
	return cs, nil
}

func (cs *ContractService) Init() error {
	err := cs.orchestra.Init()
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractService) Start() error {
	go cs.checkDoDContract()
	go cs.connectRpcServer()
	go cs.checkContractStatus()
	go cs.checkOrderStatus()
	go cs.checkProduct()
	return nil
}

func (cs *ContractService) inventoryFind(orderId string) ([]*Product, error) {
	fp := &orchestra.FindParams{
		ProductOrderID: orderId,
	}
	err := cs.orchestra.ExecInventoryFind(fp)
	if err != nil {
		return nil, err
	}
	var productIds []*Product
	if len(fp.RspInvList) == 0 {
		return nil, errors.New("no inventory list ")
	}
	for _, v := range fp.RspInvList {
		pt := &Product{
			buyerProductID: v.BuyerProductID,
			productID:      *v.ID,
		}
		productIds = append(productIds, pt)
	}
	return productIds, nil
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
