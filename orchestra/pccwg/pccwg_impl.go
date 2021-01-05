package pccwg

import (
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/api"

	"github.com/qlcchain/go-lsobus/config"

	"github.com/qlcchain/go-lsobus/log"
)

type PCCWGImpl struct {
	logger *zap.SugaredLogger
	cfg    *config.PartnerCfg

	apiToken string

	sonataSiteImpl  *sonataSiteImpl
	sonataPOQImpl   *sonataPOQImpl
	sonataQuoteImpl *sonataQuoteImpl
	sonataOrderImpl *sonataOrderImpl
	sonataInvImpl   *sonataInvImpl
	sonataOfferImpl *sonataOfferImpl
}

func (p *PCCWGImpl) GetConfig() *config.PartnerCfg {
	return p.cfg
}

func NewPCCGWImpl(cfg *config.PartnerCfg) api.DoDSeller {
	p := &PCCWGImpl{cfg: cfg}
	p.logger = log.NewLogger("PCCWGImpl")

	p.sonataSiteImpl = newSonataSiteImpl(p)
	p.sonataPOQImpl = newSonataPOQImpl(p)
	p.sonataQuoteImpl = newSonataQuoteImpl(p)
	p.sonataOrderImpl = newSonataOrderImpl(p)
	p.sonataInvImpl = newSonataInvImpl(p)
	p.sonataOfferImpl = newSonataOfferImpl(p)

	return p
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
