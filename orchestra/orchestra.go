package orchestra

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/config"
	chainctx "github.com/qlcchain/go-lsobus/services/context"

	"github.com/qlcchain/go-lsobus/log"
)

type Orchestra struct {
	logger   *zap.SugaredLogger
	cfg      *config.Config
	fakeMode bool

	partnerImpls sync.Map // name => partner
}

func NewOrchestra(cfgFile string) *Orchestra {
	cc := chainctx.NewServiceContext(cfgFile)
	cfg, _ := cc.Config()

	o := &Orchestra{cfg: cfg}
	o.logger = log.NewLogger("orchestra")

	return o
}

func (o *Orchestra) SetFakeMode(mode bool) {
	o.fakeMode = mode
}

func (o *Orchestra) GetFakeMode() bool {
	return o.fakeMode
}

func (o *Orchestra) SetApiToken(name string, token string) {
	p := o.GetPartnerImpl(name)
	if p != nil {
		p.SetApiToken(token)
	}
}

func (o *Orchestra) Init() error {
	for _, partCfg := range o.cfg.Partners {
		if partCfg.SonataUrl == "" {
			return fmt.Errorf("partner %s has invalid sonata url", partCfg.Name)
		}

		partImpl := NewPartnerImpl(o, partCfg)
		if partImpl == nil {
			return fmt.Errorf("partner %s new nil", partCfg.Name)
		}

		err := partImpl.Init()
		if err != nil {
			return fmt.Errorf("partner %s init err %s", partCfg.Name, err)
		}

		o.partnerImpls.Store(partCfg.Name, partImpl)
		o.logger.Debugf("add partner %s/%s success", partCfg.Name, partCfg.ID)
	}

	return nil
}

func (o *Orchestra) GetPartnerCfgByID(id string) *config.PartnerCfg {
	// just for testing, if empty use first one
	if len(o.cfg.Partners) == 1 {
		return o.cfg.Partners[0]
	}

	for _, p := range o.cfg.Partners {
		if id != "" && p.ID == id {
			return p
		}
	}

	return nil
}

func (o *Orchestra) GetPartnerCfgByName(name string) *config.PartnerCfg {
	// just for testing, if empty use first one
	if len(o.cfg.Partners) == 1 {
		return o.cfg.Partners[0]
	}

	for _, p := range o.cfg.Partners {
		if name != "" && p.Name == name {
			return p
		}
	}

	return nil
}

func (o *Orchestra) GetPartnerImpl(name string) *PartnerImpl {
	// just for testing, if empty use first one
	if name == "" {
		pc := o.GetPartnerCfgByName("")
		if pc == nil {
			return nil
		}
		name = pc.Name
	}

	if v, ok := o.partnerImpls.Load(name); ok {
		p := v.(*PartnerImpl)
		return p
	}

	return nil
}

func (o *Orchestra) ExecPOQCreate(params *OrderParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecPOQCreate(params)
}

func (o *Orchestra) ExecPOQFind(params *FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecPOQFind(params)
}

func (o *Orchestra) ExecPOQGet(params *GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecPOQGet(params)
}

func (o *Orchestra) ExecQuoteCreate(params *OrderParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecQuoteCreate(params)
}

func (o *Orchestra) ExecQuoteFind(params *FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecQuoteFind(params)
}

func (o *Orchestra) ExecQuoteGet(params *GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecQuoteGet(params)
}

func (o *Orchestra) ExecOrderCreate(params *OrderParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOrderCreate(params)
}

func (o *Orchestra) ExecOrderFind(params *FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOrderFind(params)
}

func (o *Orchestra) ExecOrderGet(params *GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOrderGet(params)
}

func (o *Orchestra) ExecInventoryFind(params *FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecInventoryFind(params)
}

func (o *Orchestra) ExecInventoryGet(params *GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecInventoryGet(params)
}

func (o *Orchestra) ExecSiteFind(params *FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecSiteFind(params)
}

func (o *Orchestra) ExecSiteGet(params *GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecSiteGet(params)
}

func (o *Orchestra) ExecOfferFind(params *FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOfferFind(params)
}

func (o *Orchestra) ExecOfferGet(params *GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartnerImpl(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOfferGet(params)
}
