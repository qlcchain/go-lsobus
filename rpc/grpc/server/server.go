package grpcServer

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/qlcchain/go-lsobus/contract"

	"github.com/qlcchain/go-lsobus/config"

	pb "github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/qlcchain/go-lsobus/log"
)

type GRPCServer struct {
	rpc    *grpc.Server
	caller *contract.ContractCaller
	logger *zap.SugaredLogger
}

func NewGRPCServer(caller *contract.ContractCaller) *GRPCServer {
	gRpcServer := grpc.NewServer()
	r := &GRPCServer{
		rpc:    gRpcServer,
		caller: caller,
		logger: log.NewLogger("rpc"),
	}
	return r
}

func (g *GRPCServer) Start(cfg *config.Config) error {

	network, address, err := scheme(cfg.RPC.GRPCListenAddress)
	if err != nil {
		return err
	}

	lis, err := net.Listen(network, address)
	if err != nil {
		return fmt.Errorf("failed to listen: %s", err)
	}
	orderApi := NewOrderApi(g.caller)
	pb.RegisterChainAPIServer(g.rpc, &chainApi{})
	pb.RegisterOrderAPIServer(g.rpc, orderApi)
	orchApi := NewOrchestraAPI(g.caller.GetOrchestra())
	pb.RegisterOrchestraAPIServer(g.rpc, orchApi)
	reflection.Register(g.rpc)
	go func() {
		if err := g.rpc.Serve(lis); err != nil {
			g.logger.Error(err)
		}
	}()
	go func() {
		if err := g.newGateway(address, cfg.RPC.ListenAddress); err != nil {
			g.logger.Errorf("start gateway: %s", err)
		}
	}()
	return nil
}

func (g *GRPCServer) newGateway(grpcAddress, gwAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	// no need proxy for internal gateway to internal grpc server
	optDial := grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
		network := "tcp"
		g.logger.Debugf("WithContextDialer addr %s", addr)
		return (&net.Dialer{}).DialContext(ctx, network, addr)
	})
	opts := []grpc.DialOption{grpc.WithInsecure(), optDial}
	err := pb.RegisterChainAPIHandlerFromEndpoint(ctx, gwmux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("gateway register: %s", err)
	}
	err = pb.RegisterOrderAPIHandlerFromEndpoint(ctx, gwmux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("gateway register: %s", err)
	}
	err = pb.RegisterOrchestraAPIHandlerFromEndpoint(ctx, gwmux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("gateway register: %s", err)
	}
	_, address, err := scheme(gwAddress)
	if err != nil {
		return err
	}
	if err := http.ListenAndServe(address, gwmux); err != nil {
		g.logger.Error(err)
	}
	return nil
}

func (g *GRPCServer) Stop() {
	g.rpc.Stop()
}

func scheme(endpoint string) (string, string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", "", err
	}
	return u.Scheme, u.Host, nil
}
