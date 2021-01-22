package orchestra

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/cmd/util"

	"github.com/qlcchain/go-lsobus/orchestra/dod"

	"github.com/qlcchain/go-lsobus/api"

	"github.com/qlcchain/go-lsobus/orchestra/pccwg"

	"github.com/qlcchain/go-lsobus/config"
	chainctx "github.com/qlcchain/go-lsobus/services/context"

	"github.com/qlcchain/go-lsobus/log"
)

type Seller struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
	api.DoDSeller
}

func NewSeller(ctx context.Context, cfgFile string) (seller api.DoDSeller, err error) {
	cc := chainctx.NewServiceContext(cfgFile)
	cfg, _ := cc.Config()

	log.Root.Debug(util.ToIndentString(cfg))

	o := &Seller{cfg: cfg, logger: log.NewLogger("seller")}
	partner := cfg.Partner

	switch partner.Implementation {
	case "pccwg":
		seller, err = pccwg.NewPCCGWImpl(ctx, cfg)
		if err != nil {
			return nil, err
		}
	case "qlc":
		seller, err = dod.NewDoD(ctx, cfg)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid implementation %s", partner.Implementation)
	}

	o.DoDSeller = seller

	return o, nil
}
