package orchestra

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/api"

	"github.com/qlcchain/go-lsobus/orchestra/pccwg"

	"github.com/qlcchain/go-lsobus/config"
	chainctx "github.com/qlcchain/go-lsobus/services/context"

	"github.com/qlcchain/go-lsobus/log"
)

type Sellers struct {
	logger *zap.SugaredLogger
	cfg    *config.Config

	partners sync.Map // name => partner
}

func NewSellers(cfgFile string) *Sellers {
	cc := chainctx.NewServiceContext(cfgFile)
	cfg, _ := cc.Config()

	o := &Sellers{cfg: cfg}
	o.logger = log.NewLogger("orchestra")

	return o
}

func (o *Sellers) Init() error {
	for _, partner := range o.cfg.Partners {
		if partner.SonataUrl == "" {
			return fmt.Errorf("partner %s has invalid sonata url", partner.Name)
		}
		var sellerImpl api.DoDSeller
		switch partner.Implementation {
		case "pccwg":
			sellerImpl = pccwg.NewPCCGWImpl(partner)
		case "qlc":
		default:
			return fmt.Errorf("invalid partner %s implementation, %s", partner.Name, partner.Implementation)
		}

		o.partners.Store(partner.Name, sellerImpl)
		o.logger.Debugf("add partner %s/%s success", partner.Name, partner.ID)
	}

	return nil
}

func (o *Sellers) GetPartner(name string) api.DoDSeller {
	if v, ok := o.partners.Load(name); ok {
		p := v.(api.DoDSeller)
		return p
	}

	return nil
}

func (o *Sellers) ExecPOQCreate(params *api.OrderParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecPOQCreate(params)
}

func (o *Sellers) ExecPOQFind(params *api.FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecPOQFind(params)
}

func (o *Sellers) ExecPOQGet(params *api.GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecPOQGet(params)
}

func (o *Sellers) ExecQuoteCreate(params *api.OrderParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecQuoteCreate(params)
}

func (o *Sellers) ExecQuoteFind(params *api.FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecQuoteFind(params)
}

func (o *Sellers) ExecQuoteGet(params *api.GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecQuoteGet(params)
}

func (o *Sellers) ExecOrderCreate(params *api.OrderParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOrderCreate(params)
}

func (o *Sellers) ExecOrderFind(params *api.FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOrderFind(params)
}

func (o *Sellers) ExecOrderGet(params *api.GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOrderGet(params)
}

func (o *Sellers) ExecInventoryFind(params *api.FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecInventoryFind(params)
}

func (o *Sellers) ExecInventoryGet(params *api.GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecInventoryGet(params)
}

func (o *Sellers) ExecSiteFind(params *api.FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecSiteFind(params)
}

func (o *Sellers) ExecSiteGet(params *api.GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecSiteGet(params)
}

func (o *Sellers) ExecOfferFind(params *api.FindParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOfferFind(params)
}

func (o *Sellers) ExecOfferGet(params *api.GetParams) error {
	if params.Seller == nil {
		return fmt.Errorf("invalid seller params")
	}
	p := o.GetPartner(params.Seller.Name)
	if p == nil {
		return fmt.Errorf("seller not exist")
	}

	return p.ExecOfferGet(params)
}
