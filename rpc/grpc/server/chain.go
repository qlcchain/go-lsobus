package grpcServer

import (
	"context"

	"github.com/iixlabs/virtual-lsobus/services/version"

	"github.com/iixlabs/virtual-lsobus/rpc/grpc/proto"
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
