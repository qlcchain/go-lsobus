package grpcServer

import (
	"context"
	"testing"
	"time"

	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
	"github.com/qlcchain/go-lsobus/services/version"
)

func TestChainApi_Version(t *testing.T) {
	ca := chainApi{}
	versionRequest := &proto.VersionRequest{}
	version.BuildTime = time.Now().Format("2006-01-02 15:04:05-0700")
	version.GitRev = "c095c91ad2df"
	version.Version = "999"
	versionRsp, err := ca.Version(context.Background(), versionRequest)
	if err != nil {
		t.Fatal(err)
	}
	if versionRsp.Version != version.Version || versionRsp.Hash != version.GitRev || versionRsp.BuildTime != version.BuildTime {
		t.Fatal("version info error")
	}
}
