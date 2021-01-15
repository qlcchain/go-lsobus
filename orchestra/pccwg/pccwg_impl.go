package pccwg

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/utils"

	"github.com/qlcchain/go-lsobus/mock"

	"github.com/qlcchain/go-lsobus/api"

	"github.com/qlcchain/go-lsobus/config"

	"github.com/qlcchain/go-lsobus/log"
)

type PCCWGImpl struct {
	logger          *zap.SugaredLogger
	cfg             *config.Config
	account         *pkg.Account
	apiToken        string
	client          *qlcSdk.QLCClient
	sonataSiteImpl  *sonataSiteImpl
	sonataPOQImpl   *sonataPOQImpl
	sonataQuoteImpl *sonataQuoteImpl
	sonataOrderImpl *sonataOrderImpl
	sonataInvImpl   *sonataInvImpl
	sonataOfferImpl *sonataOfferImpl
	status          atomic.Int32
}

func (p *PCCWGImpl) GetConfig() *config.Config {
	return p.cfg
}

func (p *PCCWGImpl) GetSellerConfig() *config.PartnerCfg {
	return p.cfg.Partner
}

func NewPCCGWImpl(ctx context.Context, cfg *config.Config) (api.DoDSeller, error) {
	p := &PCCWGImpl{cfg: cfg}
	p.logger = log.NewLogger("PCCWGImpl")

	p.sonataSiteImpl = newSonataSiteImpl(p)
	p.sonataPOQImpl = newSonataPOQImpl(p)
	p.sonataQuoteImpl = newSonataQuoteImpl(p)
	p.sonataOrderImpl = newSonataOrderImpl(p)
	p.sonataInvImpl = newSonataInvImpl(p)
	p.sonataOfferImpl = newSonataOfferImpl(p)

	client, err := qlcSdk.NewQLCClient(cfg.Partner.ChainUrl)
	if err != nil {
		return nil, err
	}
	p.client = client

	bytes, err := hex.DecodeString(cfg.Partner.Account)
	if err != nil {
		return nil, err
	}
	p.account = pkg.NewAccount(bytes)

	go func(ctx context.Context, client *qlcSdk.QLCClient) {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s, err := client.Pov.GetPovStatus()
				if err != nil {
					p.status.Store(0)
				} else if s.SyncState == 2 {
					p.status.Store(1)
				} else {
					time.Sleep(time.Second)
				}
			}
		}
	}(ctx, client)

	return p, nil
}

func (p *PCCWGImpl) Init() error {
	err := p.sonataSiteImpl.Init()
	if err != nil {
		return err
	}

	err = p.sonataPOQImpl.Init()
	if err != nil {
		return err
	}

	err = p.sonataQuoteImpl.Init()
	if err != nil {
		return err
	}

	err = p.sonataOrderImpl.Init()
	if err != nil {
		return err
	}

	err = p.sonataInvImpl.Init()
	if err != nil {
		return err
	}

	err = p.sonataOfferImpl.Init()
	if err != nil {
		return err
	}

	return nil
}

func (p *PCCWGImpl) ExecPOQCreate(params *api.OrderParams) error {
	return p.sonataPOQImpl.SendCreateRequest(params)
}

func (p *PCCWGImpl) ExecPOQFind(params *api.FindParams) error {
	return p.sonataPOQImpl.SendFindRequest(params)
}

func (p *PCCWGImpl) ExecPOQGet(params *api.GetParams) error {
	return p.sonataPOQImpl.SendGetRequest(params)
}

func (p *PCCWGImpl) ExecQuoteCreate(params *api.OrderParams) error {
	return p.sonataQuoteImpl.SendCreateRequest(params)
}

func (p *PCCWGImpl) ExecQuoteFind(params *api.FindParams) error {
	return p.sonataQuoteImpl.SendFindRequest(params)
}

func (p *PCCWGImpl) ExecQuoteGet(params *api.GetParams) error {
	return p.sonataQuoteImpl.SendGetRequest(params)
}

func (p *PCCWGImpl) ExecOrderCreate(params *api.OrderParams) error {
	return p.sonataOrderImpl.SendCreateRequest(params)
}

func (p *PCCWGImpl) ExecOrderFind(params *api.FindParams) error {
	return p.sonataOrderImpl.SendFindRequest(params)
}

func (p *PCCWGImpl) ExecOrderGet(params *api.GetParams) error {
	return p.sonataOrderImpl.SendGetRequest(params)
}

func (p *PCCWGImpl) ExecInventoryFind(params *api.FindParams) error {
	return p.sonataInvImpl.SendFindRequest(params)
}

func (p *PCCWGImpl) ExecInventoryGet(params *api.GetParams) error {
	return p.sonataInvImpl.SendGetRequest(params)
}

func (p *PCCWGImpl) ExecSiteFind(params *api.FindParams) error {
	return p.sonataSiteImpl.SendFindRequest(params)
}

func (p *PCCWGImpl) ExecSiteGet(params *api.GetParams) error {
	return p.sonataSiteImpl.SendGetRequest(params)
}

func (p *PCCWGImpl) ExecOfferFind(params *api.FindParams) error {
	return p.sonataOfferImpl.SendFindRequest(params)
}

func (p *PCCWGImpl) ExecOfferGet(params *api.GetParams) error {
	return p.sonataOfferImpl.SendGetRequest(params)
}

func (p *PCCWGImpl) GetTerminateOrderRewardBlock(param *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	account := p.Account()
	if p.IsFake() {
		if blk, err := mock.GetTerminateOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}
	if blk, err := p.client.DoDSettlement.GetTerminateOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) GetChangeOrderRewardBlock(
	param *qlcSdk.DoDSettleResponseParam,
) (*pkg.StateBlock, error) {
	account := p.Account()
	if p.IsFake() {
		if blk, err := mock.GetChangeOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}
	if blk, err := p.client.DoDSettlement.GetChangeOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) GetCreateOrderRewardBlock(
	param *qlcSdk.DoDSettleResponseParam,
) (*pkg.StateBlock, error) {
	account := p.Account()

	if p.IsFake() {
		if blk, err := mock.GetCreateOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}
	if blk, err := p.client.DoDSettlement.GetCreateOrderRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) GetTerminateOrderBlock(
	op *qlcSdk.DoDSettleTerminateOrderParam,
) (blk *pkg.StateBlock, err error) {
	account := p.Account()
	if p.IsFake() {
		if blk, err = mock.GetTerminateOrderBlock(op, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}

	if blk, err = p.client.DoDSettlement.GetTerminateOrderBlock(op, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) GetChangeOrderBlock(param *qlcSdk.DoDSettleChangeOrderParam) (*pkg.StateBlock, error) {
	account := p.Account()
	if p.IsFake() {
		if blk, err := mock.GetChangeOrderBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}

	if blk, err := p.client.DoDSettlement.GetChangeOrderBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) GetCreateOrderBlock(param *qlcSdk.DoDSettleCreateOrderParam) (*pkg.StateBlock, error) {
	account := p.Account()
	if p.IsFake() {
		if blk, err := mock.GetCreateOrderBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}

	if blk, err := p.client.DoDSettlement.GetCreateOrderBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) GetUpdateOrderInfoRewardBlock(param *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	account := p.Account()
	if p.IsFake() {
		if blk, err := mock.GetUpdateOrderInfoRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}

	if blk, err := p.client.DoDSettlement.GetUpdateOrderInfoRewardBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) GetUpdateProductInfoBlock(param *qlcSdk.DoDSettleUpdateProductInfoParam) (*pkg.StateBlock, error) {
	account := p.Account()
	if p.IsFake() {
		if blk, err := mock.GetUpdateProductInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}

	if blk, err := p.client.DoDSettlement.GetUpdateProductInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) IsFake() bool {
	return p.GetSellerConfig().IsFake
}

func (p *PCCWGImpl) GetUpdateOrderInfoBlock(param *qlcSdk.DoDSettleUpdateOrderInfoParam) (*pkg.StateBlock, error) {
	account := p.Account()
	if p.IsFake() {
		if blk, err := mock.GetUpdateOrderInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
			return account.Sign(hash), nil
		}); err != nil {
			return nil, err
		} else {
			return blk, nil
		}
	}

	if blk, err := p.client.DoDSettlement.GetUpdateOrderInfoBlock(param, func(hash pkg.Hash) (signature pkg.Signature, err error) {
		return account.Sign(hash), nil
	}); err != nil {
		return nil, err
	} else {
		blk.Work = utils.GenerateWork(blk.Root())
		return blk, nil
	}
}

func (p *PCCWGImpl) Process(block *pkg.StateBlock) (pkg.Hash, error) {
	if !p.IsFake() {
		h, err := p.client.Ledger.Process(block)
		if err != nil {
			return pkg.ZeroHash, fmt.Errorf("process block error: %s", err)
		}
		return h, p.waitBlockConfirmed(block.GetHash())
	}
	return pkg.ZeroHash, nil
}

func (p *PCCWGImpl) Account() *pkg.Account {
	return p.account
}

func (p *PCCWGImpl) GetOrderInfoBySellerAndOrderId(
	seller pkg.Address, orderId string,
) (*qlcSdk.DoDSettleOrderInfo, error) {
	if p.IsFake() {
		return mock.GetOrderInfoByInternalId("")
	}
	if orderInfo, err := p.client.DoDSettlement.GetOrderInfoBySellerAndOrderId(seller, orderId); err != nil {
		return nil, err
	} else {
		return orderInfo, nil
	}
}

func (p *PCCWGImpl) IsChainReady() bool {
	return p.status.Load() == int32(1)
}

func (p *PCCWGImpl) GetPendingRequest(addr pkg.Address) ([]*qlcSdk.DoDPendingRequestRsp, error) {
	if p.IsFake() {
		return mock.GetPendingRequest(addr), nil
	}
	return p.client.DoDSettlement.GetPendingRequest(addr)
}

func (p *PCCWGImpl) GetPendingResourceCheck(addr pkg.Address) ([]*qlcSdk.DoDPendingResourceCheckInfo, error) {
	if p.IsFake() {
		return mock.GetPendingResourceCheckForProductId(addr), nil
	}

	return p.client.DoDSettlement.GetPendingResourceCheck(addr)
}

func (p *PCCWGImpl) GetOrderInfoByInternalId(id string) (*qlcSdk.DoDSettleOrderInfo, error) {
	if p.IsFake() {
		return mock.GetOrderInfoByInternalId(id)
	}
	orderInfo, err := p.client.DoDSettlement.GetOrderInfoByInternalId(id)
	if err != nil {
		return nil, err
	}
	return orderInfo, nil
}

func (p *PCCWGImpl) waitBlockConfirmed(hash pkg.Hash) error {
	t := time.NewTimer(time.Second * 180)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			return errors.New("consensus confirmed timeout")
		default:
			confirmed, err := p.client.Ledger.BlockConfirmedStatus(hash)
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
