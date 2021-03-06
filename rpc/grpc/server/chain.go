package grpcServer

import (
	"context"

	"github.com/qlcchain/go-lsobus/services/version"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

type chainApi struct {
}

func (chainApi) Version(context.Context, *proto.VersionRequest) (*proto.VersionResponse, error) {
	return &proto.VersionResponse{
		BuildTime: version.BuildTime,
		Version:   version.Version,
		Hash:      version.GitRev,
	}, nil
}
