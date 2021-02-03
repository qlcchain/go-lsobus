package dod

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-openapi/swag"
	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/cmd/util"

	invmod "github.com/qlcchain/go-lsobus/orchestra/sonata/inventory/models"

	"github.com/qlcchain/go-lsobus/log"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/config"
	sw "github.com/qlcchain/go-lsobus/generated/dod"
	quomod "github.com/qlcchain/go-lsobus/orchestra/sonata/quote/models"
)

const (
	APITokenKey = "apikey"
)

type DoDImpl struct {
	client  *sw.APIClient
	cfg     *config.Config
	account *pkg.Account
	status  atomic.Int32
	logger  *zap.SugaredLogger
}

func NewDoD(ctx context.Context, cfg *config.Config) (api.DoDSeller, error) {
	cfg.Verify(func(c *config.Config) error {
		extra := c.Partner.Extra
		if extra == nil {
			return fmt.Errorf("empty extra in partner")
		}

		if v, ok := extra[APITokenKey]; !ok || v == "" {
			return errors.New("invalid API KEY")
		}
		return nil
	})

	swCfg := sw.NewConfiguration()
	swCfg.BasePath = cfg.Partner.BackEndURL
	swCfg.AddDefaultHeader("API-KEY", cfg.Partner.Extra[APITokenKey])

	client := sw.NewAPIClient(swCfg)

	data, err := hex.DecodeString(cfg.Partner.Account)
	if err != nil {
		return nil, err
	}
	s, _ := pkg.BytesToSeed(data)
	account, _ := s.Account(0)

	p := &DoDImpl{
		client:  client,
		cfg:     cfg,
		account: account,
		logger:  log.NewLogger("qlc_dod"),
	}

	p.logger.Debug("run as " + account.Address().String())

	//FIXME: verify PoV status
	p.status.Store(1)
	if resp, _, err := client.DLTPovApi.V1DltPovStatusGet(context.Background()); err == nil {
		p.logger.Debugf("pov status %d, %s", resp.SyncState, resp.SyncStateStr)
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
	panic("implement me")
}

func (d *DoDImpl) ExecQuoteFind(params *api.FindParams) error {
	panic("implement me")
}

func (d *DoDImpl) ExecQuoteGet(params *api.GetParams) error {
	if resp, _, err := d.client.QuotesApi.V1QuotesIdGet(context.Background(), params.ID); err != nil {
		return err
	} else {
		params.RspQuote = quoteItem2Quote(&resp)
		return nil
	}
}

func quoteItem2Quote(in *sw.QuoteRes) *quomod.Quote {
	var out quomod.Quote

	// mapping price
	for _, item := range in.QuoteItem {
		for _, price := range item.QuoteItemPrice {
			if price.Price.PreTaxAmount.Value != float32(0) {
				out.QuoteItem = append(out.QuoteItem, &quomod.QuoteItem{
					PreCalculatedPrice: &quomod.QuotePrice{
						Price: &quomod.Price{
							PreTaxAmount: &quomod.Money{
								Unit:  swag.String(price.Price.PreTaxAmount.Unit),
								Value: swag.Float32(price.Price.PreTaxAmount.Value),
							},
							PriceRange: nil,
						},
						PriceType:             quomod.PriceType(price.PriceType),
						RecurringChargePeriod: quomod.ChargePeriod(price.RecurringChargePeriod),
					},
					State: quomod.QuoteItemStateType(item.State),
					ID:    swag.String(item.ID),
				})
			}
		}
	}

	return &out
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
	if resp, _, err := d.client.OrdersApi.V1OrdersProductInventoryIdGet(context.Background(), params.ID); err != nil {
		return err
	} else {
		params.RspInv = &invmod.Product{
			ID:     swag.String(resp.Id),
			Status: invmod.ProductStatus(resp.Status),
		}
		return nil
	}
}

func (d *DoDImpl) ExecInventoryStatusGet(params *api.InventoryParams) error {
	if resp, _, err := d.client.OrdersApi.V1OrdersProductStatusIdGet(context.Background(), params.OrderID); err != nil {
		return nil
	} else {
		params.Status = resp.Status
	}
	//FIXME: remove
	if params.Status != "active" {
		params.Status = "active"
	}
	return nil
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
		return resp.Result, nil
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
		return resp.Result, nil
	}
}

func (d *DoDImpl) GetOrderInfoByInternalId(id string) (*qlcSdk.DoDSettleOrderInfo, error) {
	if resp, _, err := d.client.DLTOrdersInfoApi.V1DltOrderInfoByInternalIdPost(context.Background(), sw.DltOrderInfoByInternalIdReq{
		InternalId: id,
	}); err != nil {
		return nil, err
	} else {
		return resp.Result, nil
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
		d.logger.Debugf("GetUpdateOrderInfoBlock: %s", resp.TxId)
		return nil, nil
	}
}

func (d *DoDImpl) GetUpdateProductInfoBlock(param *qlcSdk.DoDSettleUpdateProductInfoParam) (*pkg.StateBlock, error) {
	var req sw.DltOrderSellerUpdateProductInfoBlockReq
	d.logger.Debug(util.ToIndentString(param))
	if err := convert(param, &req); err != nil {
		return nil, err
	}
	d.logger.Debug(util.ToIndentString(req))
	if resp, _, err := d.client.DLTOrdersSellerApi.V1DltOrderSellerUpdateProductInfoBlockPost(context.Background(), req); err != nil {
		return nil, err
	} else {
		//FIXME: generate work and sign the block
		if resp.Result != nil {
			d.logger.Debugf("GetUpdateProductInfoBlock: %s", resp.Result.TxId)
		}
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
		if resp.Result != nil {
			d.logger.Debugf("GetUpdateOrderInfoRewardBlock: %s", resp.Result.TxId)
		}
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
		d.logger.Debugf("GetChangeOrderBlock: %s", resp.TxId)
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
		d.logger.Debugf("GetCreateOrderBlock: %s", resp.TxId)
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
		d.logger.Debugf("GetTerminateOrderBlock: %APITokenKey", resp.TxId)
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
		d.logger.Debugf("GetCreateOrderRewardBlock: %s", resp.TxId)
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
		d.logger.Debugf("GetChangeOrderRewardBlock: %s", resp.TxId)
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
		d.logger.Debugf("GetTerminateOrderRewardBlock: %s", resp.TxId)
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
