package api

import (
	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/config"
)

type Authenticator interface {
	ExecAuthLogin(params *LoginParams) error
	RenewAPIToken() string
	GetAPIToken() string
	ClearAPIToken()
}

type Orchestrator interface {
	ExecPOQCreate(params *OrderParams) error
	ExecPOQFind(params *FindParams) error
	ExecPOQGet(params *GetParams) error
	ExecQuoteCreate(params *OrderParams) error
	ExecQuoteFind(params *FindParams) error
	ExecQuoteGet(params *GetParams) error
	ExecOrderCreate(params *OrderParams) error
	ExecOrderFind(params *FindParams) error
	ExecOrderGet(params *GetParams) error
	ExecInventoryFind(params *FindParams) error
	ExecInventoryGet(params *GetParams) error
	ExecSiteFind(params *FindParams) error
	ExecSiteGet(params *GetParams) error
	ExecOfferFind(params *FindParams) error
	ExecOfferGet(params *GetParams) error
}

type configer interface {
	GetSellerConfig() *config.PartnerCfg
	GetConfig() *config.Config
	Account() *types.Account
	IsFake() bool
}

type contractor interface {
	IsChainReady() bool
	GetPendingRequest(addr types.Address) ([]*qlcSdk.DoDPendingRequestRsp, error)
	GetPendingResourceCheck(addr types.Address) ([]*qlcSdk.DoDPendingResourceCheckInfo, error)
	GetOrderInfoByInternalId(id string) (*qlcSdk.DoDSettleOrderInfo, error)
	GetOrderInfoBySellerAndOrderId(seller types.Address, orderId string) (*qlcSdk.DoDSettleOrderInfo, error)

	// QLC SDK
	GetUpdateOrderInfoBlock(op *qlcSdk.DoDSettleUpdateOrderInfoParam) (*types.StateBlock, error)
	GetUpdateProductInfoBlock(op *qlcSdk.DoDSettleUpdateProductInfoParam) (*types.StateBlock, error)
	GetUpdateOrderInfoRewardBlock(param *qlcSdk.DoDSettleResponseParam) (*types.StateBlock, error)
	GetChangeOrderBlock(op *qlcSdk.DoDSettleChangeOrderParam) (*types.StateBlock, error)
	GetCreateOrderBlock(op *qlcSdk.DoDSettleCreateOrderParam) (*types.StateBlock, error)
	GetTerminateOrderBlock(op *qlcSdk.DoDSettleTerminateOrderParam) (*types.StateBlock, error)
	GetCreateOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam) (*types.StateBlock, error)
	GetChangeOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam) (*types.StateBlock, error)
	GetTerminateOrderRewardBlock(op *qlcSdk.DoDSettleResponseParam) (*types.StateBlock, error)
	Process(block *types.StateBlock) (types.Hash, error)
}

type DoDSeller interface {
	Authenticator
	Orchestrator
	configer
	contractor
}
