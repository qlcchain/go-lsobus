package orchestra

import (
	"github.com/qlcchain/go-lsobus/config"
	chainctx "github.com/qlcchain/go-lsobus/services/context"
)

type Orchestra struct {
	cfg      *config.Config
	fakeMode bool

	sonataSiteImpl  *sonataSiteImpl
	sonataPOQImpl   *sonataPOQImpl
	sonataQuoteImpl *sonataQuoteImpl
	sonataOrderImpl *sonataOrderImpl
	sonataInvImpl   *sonataInvImpl
	sonataOfferImpl *sonataOfferImpl
}

func NewOrchestra(cfgFile string) *Orchestra {
	cc := chainctx.NewServiceContext(cfgFile)
	cfg, _ := cc.Config()

	o := &Orchestra{cfg: cfg}
	o.sonataSiteImpl = newSonataSiteImpl(o)
	o.sonataPOQImpl = newSonataPOQImpl(o)
	o.sonataQuoteImpl = newSonataQuoteImpl(o)
	o.sonataOrderImpl = newSonataOrderImpl(o)
	o.sonataInvImpl = newSonataInvImpl(o)
	o.sonataOfferImpl = newSonataOfferImpl(o)

	return o
}

func (o *Orchestra) SetFakeMode(mode bool) {
	o.fakeMode = mode
}

func (o *Orchestra) GetFakeMode() bool {
	return o.fakeMode
}

func (o *Orchestra) Init() error {
	err := o.sonataSiteImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataPOQImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataQuoteImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataOrderImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataInvImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataOfferImpl.Init()
	if err != nil {
		return err
	}

	return nil
}

func (o *Orchestra) GetSonataUrl(id string) string {
	if len(o.cfg.Partners) == 1 {
		return o.cfg.Partners[0].SonataUrl
	}

	for _, p := range o.cfg.Partners {
		if p.ID == id {
			return p.SonataUrl
		}
	}

	return "http://127.0.0.1:8080"
}

func (o *Orchestra) ExecPOQCreate(params *OrderParams) error {
	return o.sonataPOQImpl.SendCreateRequest(params)
}

func (o *Orchestra) ExecPOQFind(params *FindParams) error {
	return o.sonataPOQImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecPOQGet(params *GetParams) error {
	return o.sonataPOQImpl.SendGetRequest(params)
}

func (o *Orchestra) ExecQuoteCreate(params *OrderParams) error {
	return o.sonataQuoteImpl.SendCreateRequest(params)
}

func (o *Orchestra) ExecQuoteFind(params *FindParams) error {
	return o.sonataQuoteImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecQuoteGet(params *GetParams) error {
	return o.sonataQuoteImpl.SendGetRequest(params)
}

func (o *Orchestra) ExecOrderCreate(params *OrderParams) error {
	return o.sonataOrderImpl.SendCreateRequest(params)
}

func (o *Orchestra) ExecOrderFind(params *FindParams) error {
	return o.sonataOrderImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecOrderGet(params *GetParams) error {
	return o.sonataOrderImpl.SendGetRequest(params)
}

func (o *Orchestra) ExecInventoryFind(params *FindParams) error {
	return o.sonataInvImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecInventoryGet(params *GetParams) error {
	return o.sonataInvImpl.SendGetRequest(params)
}

func (o *Orchestra) ExecSiteFind(params *FindParams) error {
	return o.sonataSiteImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecSiteGet(params *GetParams) error {
	return o.sonataSiteImpl.SendGetRequest(params)
}

func (o *Orchestra) ExecOfferFind(params *FindParams) error {
	return o.sonataOfferImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecOfferGet(params *GetParams) error {
	return o.sonataOfferImpl.SendGetRequest(params)
}
