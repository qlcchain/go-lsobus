package orchestra

import (
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/config"

	"github.com/qlcchain/go-lsobus/log"
)

type PartnerImpl struct {
	logger *zap.SugaredLogger
	cfg    *config.PartnerCfg
	Orch   *Orchestra

	apiToken string

	sonataSiteImpl  *sonataSiteImpl
	sonataPOQImpl   *sonataPOQImpl
	sonataQuoteImpl *sonataQuoteImpl
	sonataOrderImpl *sonataOrderImpl
	sonataInvImpl   *sonataInvImpl
	sonataOfferImpl *sonataOfferImpl
}

func NewPartnerImpl(o *Orchestra, cfg *config.PartnerCfg) *PartnerImpl {
	p := &PartnerImpl{Orch: o, cfg: cfg}
	p.logger = log.NewLogger("partnerImpl")

	p.sonataSiteImpl = newSonataSiteImpl(p)
	p.sonataPOQImpl = newSonataPOQImpl(p)
	p.sonataQuoteImpl = newSonataQuoteImpl(p)
	p.sonataOrderImpl = newSonataOrderImpl(p)
	p.sonataInvImpl = newSonataInvImpl(p)
	p.sonataOfferImpl = newSonataOfferImpl(p)

	return p
}

func (p *PartnerImpl) Init() error {
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

func (p *PartnerImpl) GetFakeMode() bool {
	return p.Orch.GetFakeMode()
}

func (p *PartnerImpl) GetSonataUrl() string {
	return p.cfg.SonataUrl
}

func (p *PartnerImpl) ExecPOQCreate(params *OrderParams) error {
	return p.sonataPOQImpl.SendCreateRequest(params)
}

func (p *PartnerImpl) ExecPOQFind(params *FindParams) error {
	return p.sonataPOQImpl.SendFindRequest(params)
}

func (p *PartnerImpl) ExecPOQGet(params *GetParams) error {
	return p.sonataPOQImpl.SendGetRequest(params)
}

func (p *PartnerImpl) ExecQuoteCreate(params *OrderParams) error {
	return p.sonataQuoteImpl.SendCreateRequest(params)
}

func (p *PartnerImpl) ExecQuoteFind(params *FindParams) error {
	return p.sonataQuoteImpl.SendFindRequest(params)
}

func (p *PartnerImpl) ExecQuoteGet(params *GetParams) error {
	return p.sonataQuoteImpl.SendGetRequest(params)
}

func (p *PartnerImpl) ExecOrderCreate(params *OrderParams) error {
	return p.sonataOrderImpl.SendCreateRequest(params)
}

func (p *PartnerImpl) ExecOrderFind(params *FindParams) error {
	return p.sonataOrderImpl.SendFindRequest(params)
}

func (p *PartnerImpl) ExecOrderGet(params *GetParams) error {
	return p.sonataOrderImpl.SendGetRequest(params)
}

func (p *PartnerImpl) ExecInventoryFind(params *FindParams) error {
	return p.sonataInvImpl.SendFindRequest(params)
}

func (p *PartnerImpl) ExecInventoryGet(params *GetParams) error {
	return p.sonataInvImpl.SendGetRequest(params)
}

func (p *PartnerImpl) ExecSiteFind(params *FindParams) error {
	return p.sonataSiteImpl.SendFindRequest(params)
}

func (p *PartnerImpl) ExecSiteGet(params *GetParams) error {
	return p.sonataSiteImpl.SendGetRequest(params)
}

func (p *PartnerImpl) ExecOfferFind(params *FindParams) error {
	return p.sonataOfferImpl.SendFindRequest(params)
}

func (p *PartnerImpl) ExecOfferGet(params *GetParams) error {
	return p.sonataOfferImpl.SendGetRequest(params)
}
