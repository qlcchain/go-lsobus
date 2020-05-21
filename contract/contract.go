package contract

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/qlcchain/go-lsobus/common/util"
	"github.com/qlcchain/go-lsobus/sonata/inventory/models"

	"github.com/qlcchain/go-lsobus/orchestra"

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
	checkNeedSignContractInterval = 30 * time.Second
	checkContractStatusInterval   = 5 * time.Second
	checkOrderStatusInterval      = 5 * time.Second
	connectRpcServerInterval      = 5 * time.Second
)

type ContractService struct {
	cfg            *config.Config
	account        *types.Account
	logger         *zap.SugaredLogger
	client         *rpc.Client
	ctx            context.Context
	cancel         context.CancelFunc
	handlerIds     map[common.TopicType]string
	eb             event.EventBus
	chainReady     bool
	quit           chan bool
	orderIdOnChain *sync.Map
	orchestra      *orchestra.Orchestra
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
		cfg:            cfg,
		account:        cc.Account(),
		logger:         log.NewLogger("contract"),
		ctx:            ctx,
		cancel:         cancel,
		handlerIds:     make(map[common.TopicType]string),
		eb:             cc.EventBus(),
		quit:           make(chan bool, 1),
		orderIdOnChain: new(sync.Map),
		orchestra:      or,
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
	//	go cs.checkOrderStatus()
	return nil
}

func (cs *ContractService) checkOrderStatus() {
	ticker := time.NewTicker(checkOrderStatusInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.chainReady {
				cs.getOrderStatus()
			}
		}
	}
}

func (cs *ContractService) getOrderStatus() {
	var id []*qlcchain.DoDPendingResourceCheckInfo
	addr := cs.account.Address()
	err := cs.client.Call(&id, "DoDSettlement_getPendingResourceCheck", &addr)
	if err != nil {
		cs.logger.Error(err)
		return
	}
	for _, v := range id {
		var productActive []string
		for _, value := range v.Products {
			if !value.Active {
				gp := &orchestra.GetParams{
					ID: value.ProductId,
				}
				err := cs.orchestra.ExecInventoryGet(gp)
				if err != nil {
					cs.logger.Error(err)
					continue
				}
				if gp.RspInv.Status == models.ProductStatusActive {
					cs.logger.Infof("product %s is active", value.ProductId)
					value.Active = true
					productActive = append(productActive, value.ProductId)
				}
			}
		}
		err = cs.updateProductStatusToChain(addr, v.OrderId, productActive)
		if err != nil {
			cs.logger.Error(err)
		}
	}
	//for _, v := range id {
	//	orderReady := true
	//	for _, value := range v.Products {
	//		if !value.Active {
	//			orderReady = false
	//			break
	//		}
	//	}
	//	if orderReady {
	//		err = cs.updateOrderCompleteStatusToChain(addr, v)
	//		if err != nil {
	//			cs.logger.Error(err)
	//		}
	//	}
	//}
}

func (cs *ContractService) updateProductStatusToChain(addr types.Address, orderId string, products []string) error {
	param := &abi.DoDSettleResourceReadyParam{
		Address:   addr,
		OrderId:   orderId,
		ProductId: products,
	}

	block := new(types.StateBlock)
	err := cs.client.Call(&block, "DoDSettlement_getResourceReadyBlock", param)
	if err != nil {
		return err
	}

	var w types.Work
	worker, _ := types.NewWorker(w, block.Root())
	block.Work = worker.NewWork()

	hash := block.GetHash()
	block.Signature = cs.account.Sign(hash)

	cs.logger.Infof("block:\n%s\nhash[%s]\n", util.ToIndentString(block), block.GetHash())

	var h types.Hash
	err = cs.client.Call(&h, "ledger_process", &block)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractService) updateOrderCompleteStatusToChain(addr types.Address, id *qlcchain.DoDPendingResourceCheckInfo) error {
	var ProductId []string
	for _, v := range id.Products {
		ProductId = append(ProductId, v.ProductId)
	}
	param := &abi.DoDSettleResourceReadyParam{
		Address:   addr,
		OrderId:   id.OrderId,
		ProductId: ProductId,
	}

	block := new(types.StateBlock)
	err := cs.client.Call(&block, "DoDSettlement_getResourceReadyBlock", param)
	if err != nil {
		return err
	}

	var w types.Work
	worker, _ := types.NewWorker(w, block.Root())
	block.Work = worker.NewWork()

	hash := block.GetHash()
	block.Signature = cs.account.Sign(hash)

	cs.logger.Infof("block:\n%s\nhash[%s]\n", util.ToIndentString(block), block.GetHash())

	var h types.Hash
	err = cs.client.Call(&h, "ledger_process", &block)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractService) checkContractStatus() {
	ticker := time.NewTicker(checkContractStatusInterval)
	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			if cs.chainReady {
				cs.getContractStatus()
			}
		}
	}
}

func (cs *ContractService) getContractStatus() {
	cs.orderIdOnChain.Range(func(key, value interface{}) bool {
		internalId := key.(string)
		orderInfo := new(abi.DoDSettleOrderInfo)
		err := cs.client.Call(&orderInfo, "DoDSettlement_getOrderInfoByInternalId", &internalId)
		if err != nil {
			cs.logger.Error(err)
			return true
		}
		if orderInfo.ContractState == abi.DoDSettleContractStateConfirmed {
			cs.logger.Infof(" contract %s confirmed", internalId)
			cs.logger.Info(" call sonata API to place order")
			orderId, productId, err := cs.createOrderToSonataServer(internalId, orderInfo)
			if err != nil {
				cs.logger.Error(err)
				return true
			}
			var id []string
			id = append(id, productId[0])
			cs.logger.Infof("place order success ,orderId is %s,productId is %s", orderId, id[0])
			err = cs.updateOrderInfoToChain(orderId, internalId, id)
			if err != nil {
				cs.logger.Error(err)
				return true
			}
			cs.logger.Infof("update order info to chain success ,orderId is %s,productId is %s", orderId, id[0])
			cs.orderIdOnChain.Delete(internalId)
		}
		return true
	})
}

func (cs *ContractService) createOrderToSonataServer(internalId string, orderInfo *abi.DoDSettleOrderInfo) (string, []string, error) {
	eLines := make([]*orchestra.ELineItemParams, 0)
	for _, v := range orderInfo.Connections {
		bws := strings.Split(v.Bandwidth, " ")
		if len(bws) != 2 {
			return "", nil, errors.New("bandwidth error")
		}
		bw, err := strconv.Atoi(bws[0])
		if err != nil {
			return "", nil, err
		}
		eLine := &orchestra.ELineItemParams{
			SrcPortID: v.SrcPort,
			DstPortID: v.DstPort,
			Bandwidth: uint(bw),
			BwUnit:    bws[1],
			CosName:   v.ServiceClass.String(),
			BaseItemParams: orchestra.BaseItemParams{
				BillingParams: &orchestra.BillingParams{
					BillingType: v.BillingType.String(),
					BillingUnit: v.BillingUnit.String(),
					MeasureUnit: v.PaymentType.String(),
					StartTime:   v.StartTime,
					EndTime:     v.EndTime,
					Currency:    v.Currency,
					Price:       float32(v.Price),
				},
			},
		}
		eLines = append(eLines, eLine)
	}
	op := &orchestra.OrderParams{
		Buyer: &orchestra.Partner{
			ID:   orderInfo.Buyer.Address.String(),
			Name: orderInfo.Buyer.Name,
		},
		Seller: &orchestra.Partner{
			ID:   orderInfo.Seller.Address.String(),
			Name: orderInfo.Seller.Name,
		},
		ExternalID: internalId,
		ELineItems: eLines,
	}
	err := cs.orchestra.ExecOrderCreate(op)
	if err != nil {
		return "", nil, err
	}
	orderId := op.RspOrder.ID
	fp := &orchestra.FindParams{
		ProductOrderID: *orderId,
	}
	err = cs.orchestra.ExecInventoryFind(fp)
	if err != nil {
		return "", nil, err
	}
	var productIds []string
	if len(fp.RspInvList) == 0 {
		return "", nil, errors.New("no inventory list ")
	}
	for _, v := range fp.RspInvList {
		productIds = append(productIds, *v.ID)
	}
	return *orderId, productIds, nil
}

func (cs *ContractService) updateOrderInfoToChain(orderId string, internalId string, productId []string) error {
	var id types.Hash
	_ = id.Of(internalId)
	param := &abi.DoDSettleUpdateOrderInfoParam{
		Buyer:      cs.account.Address(),
		InternalId: id,
		OrderId:    orderId,
		ProductId:  productId,
		Status:     abi.DoDSettleOrderStateSuccess,
		FailReason: "",
	}

	block := new(types.StateBlock)
	err := cs.client.Call(&block, "DoDSettlement_getUpdateOrderInfoBlock", param)
	if err != nil {
		return err
	}

	var w types.Work
	worker, _ := types.NewWorker(w, block.Root())
	block.Work = worker.NewWork()

	hash := block.GetHash()
	block.Signature = cs.account.Sign(hash)

	var h types.Hash
	err = cs.client.Call(&h, "ledger_process", &block)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractService) checkDoDContract() {
	ticker := time.NewTicker(checkNeedSignContractInterval)
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
		cs.logger.Infof("find a dod settlement need sign,request hash is %s", v.Hash.String())
		b := cs.verifyOrderInfoFromSonata(v.Order)
		if !b {
			continue
		}
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
		if v.Order.OrderType == abi.DoDSettleOrderTypeCreate {
			err = cs.client.Call(&block, "DoDSettlement_getCreateOrderRewardBlock", param)
		} else if v.Order.OrderType == abi.DoDSettleOrderTypeCreate {
			err = cs.client.Call(&block, "DoDSettlement_getChangeOrderRewardBlock", param)
		} else if v.Order.OrderType == abi.DoDSettleOrderTypeCreate {
			err = cs.client.Call(&block, "DoDSettlement_getTerminateOrderRewardBlock", param)
		} else {
			cs.logger.Errorf("unknown order type==%s", v.Order.OrderType.String())
			continue
		}
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
		cs.logger.Infof("dod settlement sign success,request hash is :%s", v.Hash.String())
	}
}

func (cs *ContractService) verifyOrderInfoFromSonata(order *abi.DoDSettleOrderInfo) bool {
	op := &orchestra.OrderParams{}
	for _, conn := range order.Connections {
		bws := strings.Split(conn.Bandwidth, " ")
		if len(bws) != 2 {
			cs.logger.Error("bandwidth error")
			return false
		}
		bw, err := strconv.Atoi(bws[0])
		if err != nil {
			cs.logger.Error(err)
			return false
		}

		lineItem := &orchestra.ELineItemParams{
			Bandwidth: uint(bw),
			BwUnit:    bws[1],
			SrcPortID: conn.SrcPort,
			DstPortID: conn.DstPort,
			CosName:   conn.ServiceClass.String(),
		}
		lineItem.BillingParams = &orchestra.BillingParams{
			BillingType: conn.BillingType.String(),
			BillingUnit: conn.BillingUnit.String(),
			MeasureUnit: conn.PaymentType.String(),
			StartTime:   conn.StartTime,
			EndTime:     conn.EndTime,
			Currency:    conn.Currency,
			Price:       float32(conn.Price),
		}
		op.ELineItems = append(op.ELineItems, lineItem)
	}

	err := cs.orchestra.ExecQuoteCreate(op)
	if err != nil {
		cs.logger.Error(err)
		return false
	}
	if op.RspQuote == nil {
		cs.logger.Errorf("order information verify fail, empty quote response")
		return false
	}

	if len(op.RspQuote.QuoteItem) != len(order.Connections) {
		cs.logger.Errorf("order information verify fail, item count not equal")
		return false
	}

	for idx := 0; idx < len(order.Connections); idx++ {
		conn := order.Connections[idx]
		quote := op.RspQuote.QuoteItem[idx]
		if len(quote.QuoteItemPrice) == 0 {
			cs.logger.Errorf("order information verify fail, empty price")
			return false
		}
		quotePrice := quote.QuoteItemPrice[0]
		if quotePrice.Price == nil || quotePrice.Price.PreTaxAmount == nil {
			cs.logger.Errorf("order information verify fail, invalid price")
			return false
		}

		if *quotePrice.Price.PreTaxAmount.Unit != conn.Currency || *quotePrice.Price.PreTaxAmount.Value != float32(conn.Price) {
			cs.logger.Errorf("order information verify fail")
			return false
		}
	}

	cs.logger.Infof("order information verified")
	return true
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
