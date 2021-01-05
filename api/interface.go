package api

import "github.com/qlcchain/go-lsobus/config"

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
	GetConfig() *config.PartnerCfg
}

type DoDSeller interface {
	Authenticator
	Orchestrator
	configer
}
