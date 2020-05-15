package grpcServer

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/iixlabs/virtual-lsobus/contract"

	"github.com/iixlabs/virtual-lsobus/config"

	pb "github.com/iixlabs/virtual-lsobus/rpc/grpc/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/iixlabs/virtual-lsobus/log"
)

type GRPCServer struct {
	rpc    *grpc.Server
	cs     *contract.ContractService
	logger *zap.SugaredLogger
}

func NewGRPCServer(cs *contract.ContractService) *GRPCServer {
	gRpcServer := grpc.NewServer()
	r := &GRPCServer{
		rpc:    gRpcServer,
		cs:     cs,
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
	orderApi := NewOrderApi(g.cs)
	pb.RegisterChainAPIServer(g.rpc, &chainApi{})
	pb.RegisterOrderAPIServer(g.rpc, orderApi)
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
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterChainAPIHandlerFromEndpoint(ctx, gwmux, grpcAddress, opts)
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
