package grpcServer

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/qlcchain/go-lsobus/cmd/util"
	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"github.com/google/uuid"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/config"
	"github.com/qlcchain/go-lsobus/contract"
	"github.com/qlcchain/go-lsobus/mock"
	ct "github.com/qlcchain/go-lsobus/services/context"
)

func setupTestCase(t *testing.T) (func(t *testing.T), *OrderApi) {
	cfgFile := filepath.Join(config.TestDataDir(), uuid.New().String(), config.CfgFileName)
	cc := ct.NewServiceContext(cfgFile)
	cfg, err := cc.Config()
	setupOrchestraConfig(cfg)
	cs, err := contract.NewContractService(cfgFile)
	if err != nil {
		t.Fatal(err)
	}
	cs.SetFakeMode(true)

	if err = cs.Init(); err != nil {
		t.Fatal(err)
	}
	orderApi := NewOrderApi(cs)
	return func(t *testing.T) {
		err = cs.Stop()
		if err != nil {
			t.Fatal(err)
		}
		err = os.RemoveAll(filepath.Join(config.TestDataDir(), uuid.New().String()))
		if err != nil {
			t.Fatal(err)
		}
	}, orderApi
}

func setupOrchestraConfig(cfg *config.Config) {
	cfg.Partners = nil
	p1 := &config.PartnerCfg{
		Name:      "PCCWG",
		ID:        "PCCWG",
		SonataUrl: "http://127.0.0.1:7777",
		Username:  "test",
		Password:  "test",
	}
	cfg.Partners = append(cfg.Partners, p1)
}

func TestOrderApi_CreateOrder(t *testing.T) {
	teardownTestCase, oa := setupTestCase(t)
	defer teardownTestCase(t)
	createParam := mock.ProtoCreateOrderParams()
	_, pri, err := pkg.GenerateAddress()
	oa.cs.SetAccount(pkg.NewAccount(pri))
	createParam.Buyer.Address = oa.cs.GetAccount().Address().String()
	orderId, err := oa.CreateOrder(context.Background(), createParam)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(orderId)
}

func TestOrderApi_ChangeOrder(t *testing.T) {
	teardownTestCase, oa := setupTestCase(t)
	defer teardownTestCase(t)
	changeParam := mock.ProtoChangeOrderParams()
	_, pri, err := pkg.GenerateAddress()
	oa.cs.SetAccount(pkg.NewAccount(pri))
	changeParam.Buyer.Address = oa.cs.GetAccount().Address().String()
	orderId, err := oa.ChangeOrder(context.Background(), changeParam)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(orderId)
}

func TestOrderApi_TerminateOrder(t *testing.T) {
	teardownTestCase, oa := setupTestCase(t)
	defer teardownTestCase(t)
	terminateParam := mock.ProtoTerminateOrderParams()
	_, pri, err := pkg.GenerateAddress()
	oa.cs.SetAccount(pkg.NewAccount(pri))
	terminateParam.Buyer.Address = oa.cs.GetAccount().Address().String()
	orderId, err := oa.TerminateOrder(context.Background(), terminateParam)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(orderId)
}

func TestOrderApi_GetOrderInfo(t *testing.T) {
	teardownTestCase, oa := setupTestCase(t)
	defer teardownTestCase(t)
	id := &proto.GetOrderInfoByInternalId{
		InternalId: uuid.New().String(),
	}
	info, err := oa.GetOrderInfo(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(util.ToIndentString(info))
}
