package dod

import (
	"context"
	"encoding/hex"
	"encoding/json"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"
	"github.com/qlcchain/qlc-go-sdk/pkg/util"
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/log"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/config"
	sw "github.com/qlcchain/go-lsobus/generated/dod"
	quomod "github.com/qlcchain/go-lsobus/orchestra/sonata/quote/models"
)

type DoDImpl struct {
	client  *sw.APIClient
	cfg     *config.Config
	account *pkg.Account
	status  atomic.Int32
	logger  *zap.SugaredLogger
}

func NewDoD(ctx context.Context, cfg *config.Config) (api.DoDSeller, error) {
	swCfg := sw.NewConfiguration()
	swCfg.BasePath = cfg.Partner.SonataUrl
	token := cfg.Partner.APIToken
	if token != "" {
		swCfg.AddDefaultHeader("API-KEY", token)
	}
	client := sw.NewAPIClient(swCfg)

	bytes, err := hex.DecodeString(cfg.Partner.Account)
	if err != nil {
		return nil, err
	}
	account := pkg.NewAccount(bytes)
	p := &DoDImpl{
		client:  client,
		cfg:     cfg,
		account: account,
		logger:  log.NewLogger("qlc_dod"),
	}

	//FIXME: verify PoV status
	p.status.Store(1)
	if resp, _, err := client.DLTPovApi.V1DltPovStatusGet(context.Background()); err == nil {
		p.logger.Debugf("pov status %f", resp.Data.SyncState)
	}

	//go func(ctx context.Context, client *sw.APIClient) {
	//	ticker := time.NewTicker(5 * time.Second)
	//	defer ticker.Stop()
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			return
	//		case <-ticker.C:
	//			resp, _, err := client.DLTPovApi.V1DltPovStatusGet(context.Background())
	//			if err != nil {
	//				p.status.Store(0)
	//			} else if resp.Data.SyncState == 2 {
	//				p.status.Store(1)
	//			} else {
	//				time.Sleep(time.Second)
	//			}
	//		}
	//	}
	//}(ctx, client)

	return p, nil
}

func (d *DoDImpl) ExecAuthLogin(params *api.LoginParams) error {
	panic("implement me")
}

func (d *DoDImpl) RenewAPIToken() string {
	panic("implement me")
}

func (d *DoDImpl) GetAPIToken() string {
	panic("implement me")
}

func (d *DoDImpl) ClearAPIToken() {
	panic("implement me")
}

func (d *DoDImpl) ExecPOQCreate(params *api.OrderParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecPOQFind(params *api.FindParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecPOQGet(params *api.GetParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecQuoteCreate(params *api.OrderParams) error {
	if resp, _, err := d.client.QuotesApi.V1QuotesPost(context.Background(), sw.QuoteCreateReq{
		ProjectId:                    "",
		QuoteLevel:                   "",
		InstantSyncQuoting:           false,
		Description:                  "",
		RequestedQuoteCompletionDate: "",
		ExpectedFulfillmentStartDate: "",
		Agreement:                    nil,
		RelatedParty:                 nil,
		QuoteItem:                    nil,
	}); err != nil {
		return err
	} else {
		params.RspQuote = quoteItem2Quote(resp.Data.QuoteItem)
		return nil
	}
}

func (d *DoDImpl) ExecQuoteFind(params *api.FindParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecQuoteGet(params *api.GetParams) error {
	if resp, _, err := d.client.QuotesApi.V1QuotesIdGet(context.Background(), params.ID); err != nil {
		return err
	} else {
		d.logger.Debug(util.ToString(resp.Data))
		params.RspQuote = quoteItem2Quote(resp.Data.QuoteItem)
		d.logger.Debug(util.ToString(&quomod.Quote{}))
		return nil
	}
}

func quoteItem2Quote([]sw.QuoteQuoteItemSonataModel) *quomod.Quote {
	return nil
}

func (d *DoDImpl) ExecOrderCreate(params *api.OrderParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecOrderFind(params *api.FindParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecOrderGet(params *api.GetParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecInventoryFind(params *api.FindParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecInventoryGet(params *api.GetParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecSiteFind(params *api.FindParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecSiteGet(params *api.GetParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecOfferFind(params *api.FindParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecOfferGet(params *api.GetParams) error {
	panic("implement me")
}

func (d *DoDImpl) GetSellerConfig() *config.PartnerCfg {
	return d.cfg.Partner
}

func (d *DoDImpl) GetConfig() *config.Config {
	return d.cfg
}

func (d *DoDImpl) Account() *pkg.Account {
	return d.account
}

func (d *DoDImpl) IsFake() bool {
	return d.cfg.Partner.IsFake
}

func (d *DoDImpl) IsChainReady() bool {
	return d.status.Load() == int32(1)
}

func (d *DoDImpl) GetPendingRequest(addr pkg.Address) ([]*qlcSdk.DoDPendingRequestRsp, error) {
	if resp, _, err := d.client.DLTOrdersSellerApi.V1DltOrderSellerPendingRequestPost(
		context.Background(),
		sw.DltOrderSellerPendingRequestReq{
			QlcAddressSeller: addr.String(),
		},
	); err != nil {
		return nil, err
	} else {
		var pending qlcSdk.DoDPendingRequestRsp
		_ = convert(resp.Data, &pending)
		return []*qlcSdk.DoDPendingRequestRsp{&pending}, nil
	}
}

func (d *DoDImpl) GetPendingResourceCheck(addr pkg.Address) ([]*qlcSdk.DoDPendingResourceCheckInfo, error) {
	if resp, _, err := d.client.DLTOrdersSellerApi.V1DltOrderSellerPendingResourceCheckPost(
		context.Background(), sw.DltOrderSellerPendingResourceCheckReq{
			QlcAddressSeller: addr.String(),
		},
	); err != nil {
		return nil, err
	} else {
		var info []*qlcSdk.DoDPendingResourceCheckInfo
		_ = convert(resp, &info)
		return info, nil
	}
}

func (d *DoDImpl) GetOrderInfoByInternalId(id string) (*qlcSdk.DoDSettleOrderInfo, error) {
	if resp, _, err := d.client.DLTOrdersInfoApi.V1DltOrderInfoByInternalIdPost(context.Background(), sw.DltOrderInfoByInternalIdReq{
		InternalId: id,
	}); err != nil {
		return nil, err
	} else {
		var info qlcSdk.DoDSettleOrderInfo
		_ = convert(resp, &info)
		return &info, nil
	}
}

func (d *DoDImpl) GetOrderInfoBySellerAndOrderId(
	seller pkg.Address, orderId string,
) (*qlcSdk.DoDSettleOrderInfo, error) {
	if resp, _, err := d.client.DLTOrdersInfoApi.V1DltOrderInfoBySellerAndOrderIdPost(context.Background(), sw.DltOrderInfoBySellerAndOrderIdReq{
		QlcAddressSeller: seller.String(),
		OrderId:          orderId,
	}); err != nil {
		return nil, err
	} else {
		var info qlcSdk.DoDSettleOrderInfo
		_ = convert(resp, &info)
		return &info, nil
	}
}

func (d *DoDImpl) GetUpdateOrderInfoBlock(param *qlcSdk.DoDSettleUpdateOrderInfoParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderBuyerUpdateOrderInfoBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersBuyerApi.V1DltOrderBuyerUpdateOrderInfoBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetUpdateOrderInfoBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetUpdateProductInfoBlock(param *qlcSdk.DoDSettleUpdateProductInfoParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderSellerUpdateProductInfoBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersSellerApi.V1DltOrderSellerUpdateProductInfoBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetUpdateProductInfoBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetUpdateOrderInfoRewardBlock(param *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderSellerUpdateOrderInfoRewardBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersSellerApi.V1DltOrderSellerUpdateOrderInfoRewardBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetUpdateOrderInfoRewardBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetChangeOrderBlock(param *qlcSdk.DoDSettleChangeOrderParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderBuyerChangeOrderBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersBuyerApi.V1DltOrderBuyerChangeOrderBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetChangeOrderBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetCreateOrderBlock(param *qlcSdk.DoDSettleCreateOrderParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderBuyerChangeOrderBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersBuyerApi.V1DltOrderBuyerChangeOrderBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetCreateOrderBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetTerminateOrderBlock(param *qlcSdk.DoDSettleTerminateOrderParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderBuyerTerminateOrderBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersBuyerApi.V1DltOrderBuyerTerminateOrderBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetTerminateOrderBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetCreateOrderRewardBlock(param *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderSellerCreateOrderRewardBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersSellerApi.V1DltOrderSellerCreateOrderRewardBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetCreateOrderRewardBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetChangeOrderRewardBlock(param *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderSellerChangeOrderRewardBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersSellerApi.V1DltOrderSellerChangeOrderRewardBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetChangeOrderRewardBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetTerminateOrderRewardBlock(param *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderSellerTerminateOrderRewardBlockReq
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	if resp, _, err := d.client.DLTOrdersSellerApi.V1DltOrderSellerTerminateOrderRewardBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		d.logger.Debugf("GetTerminateOrderRewardBlock: %s", resp.Data.TxId)
		return nil, nil
	}
}

//FIXME: A new interface is required to calculate the PoW locally
func (d *DoDImpl) Process(blk *pkg.StateBlock) (pkg.Hash, error) {
	//var req sw.DltLedgerProcessReq
	//_ = convert(blk, &req)
	//if resp, _, err := d.client.DLTLedgerApi.V1DltLedgerProcessPost(context.Background(), req); err == nil {
	//	return pkg.ZeroHash, err
	//} else {
	//	return pkg.NewHash(resp.Data.Hash)
	//}
	return pkg.ZeroHash, nil
}

func convert(from, to interface{}) error {
	if data, err := json.Marshal(from); err != nil {
		return err
	} else {
		if err := json.Unmarshal(data, to); err != nil {
			return err
		} else {
			return nil
		}
	}
}
