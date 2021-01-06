package rpc

import (
	"context"

	"github.com/qlcchain/go-lsobus/contract"
	grpcServer "github.com/qlcchain/go-lsobus/rpc/grpc/server"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/common/event"
	"github.com/qlcchain/go-lsobus/config"
	"github.com/qlcchain/go-lsobus/log"
	chainctx "github.com/qlcchain/go-lsobus/services/context"
)

type RPC struct {
	config  *config.Config
	ctx     context.Context
	cancel  context.CancelFunc
	eb      event.EventBus
	cfgFile string
	logger  *zap.SugaredLogger
	cc      *chainctx.ServiceContext
	grpc    *grpcServer.GRPCServer
}

func NewRPC(cfgFile string, caller *contract.ContractCaller) (*RPC, error) {
	cc := chainctx.NewServiceContext(cfgFile)
	cfg, _ := cc.Config()
	ctx, cancel := context.WithCancel(context.Background())
	r := RPC{
		eb:      cc.EventBus(),
		config:  cfg,
		cfgFile: cfgFile,
		ctx:     ctx,
		cancel:  cancel,
		logger:  log.NewLogger("rpc"),
		cc:      cc,
	}
	if cfg.RPC.Enable {
		r.grpc = grpcServer.NewGRPCServer(caller)
	}
	return &r, nil
}

func (r *RPC) StopRPC() {
	r.cancel()
	if r.config.RPC.Enable {
		r.grpc.Stop()
	}
}

func (r *RPC) StartRPC() error {
	if r.config.RPC.Enable {
		err := r.grpc.Start(r.config)
		if err != nil {
			return err
		}
	}
	return nil
}
