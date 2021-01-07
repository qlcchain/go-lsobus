package dod

import (
	"context"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/config"
	sw "github.com/qlcchain/go-lsobus/generated/dod"
	quomod "github.com/qlcchain/go-lsobus/orchestra/sonata/quote/models"
	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type DoDImpl struct {
	client *sw.APIClient
}

func NewDoD(ctx context.Context, cfg *config.Config) api.DoDSeller {
	swCfg := sw.NewConfiguration()
	swCfg.Host = cfg.Partner.SonataUrl
	client := sw.NewAPIClient(swCfg)
	return &DoDImpl{
		client: client,
	}
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
		params.RspQuote = quoteItem2Quote(resp.QuoteItem)
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
		params.RspQuote = quoteItem2Quote(resp.QuoteItem)
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
	panic("implement me")
}

func (d *DoDImpl) GetConfig() *config.Config {
	panic("implement me")
}

func (d *DoDImpl) Account() *pkg.Account {
	panic("implement me")
}

func (d *DoDImpl) IsFake() bool {
	panic("implement me")
}

func (d *DoDImpl) IsChainReady() bool {
	panic("implement me")
}

func (d *DoDImpl) GetPendingRequest(addr pkg.Address) ([]*qlcSdk.DoDPendingRequestRsp, error) {
	panic("implement me")
}

func (d *DoDImpl) GetPendingResourceCheck(addr pkg.Address) ([]*qlcSdk.DoDPendingResourceCheckInfo, error) {
	panic("implement me")
}

func (d *DoDImpl) GetOrderInfoByInternalId(id string) (*qlcSdk.DoDSettleOrderInfo, error) {
	panic("implement me")
}

func (d *DoDImpl) GetOrderInfoBySellerAndOrderId(
	seller pkg.Address, orderId string,
) (*qlcSdk.DoDSettleOrderInfo, error) {
	panic("implement me")
}

func (d *DoDImpl) GetUpdateOrderInfoBlock(op *qlcSdk.DoDSettleUpdateOrderInfoParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) GetUpdateProductInfoBlock(op *qlcSdk.DoDSettleUpdateProductInfoParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) GetUpdateOrderInfoRewardBlock(param *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) GetChangeOrderBlock(op *qlcSdk.DoDSettleChangeOrderParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) GetCreateOrderBlock(op *qlcSdk.DoDSettleCreateOrderParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) GetTerminateOrderBlock(op *qlcSdk.DoDSettleTerminateOrderParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) GetCreateOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) GetChangeOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) GetTerminateOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam) (*pkg.StateBlock, error) {
	panic("implement me")
}

func (d *DoDImpl) Process(block *pkg.StateBlock) (pkg.Hash, error) {
	panic("implement me")
}
