package grpcServer

import (
	"context"
	"encoding/json"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"

	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"

	"github.com/qlcchain/go-lsobus/config"
	"github.com/qlcchain/go-lsobus/orchestra"
	ct "github.com/qlcchain/go-lsobus/services/context"
)

func setupTestCaseOrchestraAPI(t *testing.T) (func(t *testing.T), *OrchestraApi) {
	cfgFile := filepath.Join(config.TestDataDir(), uuid.New().String(), config.CfgFileName)
	cc := ct.NewServiceContext(cfgFile)
	cfg, err := cc.Config()
	setupOrchestraAPIConfig(cfg)
	orch, err := orchestra.NewSeller(context.Background(), cfgFile)
	if err != nil {
		t.Fatal(err)
	}
	orchApi := NewOrchestraAPI(orch)
	return func(t *testing.T) {
		err = os.RemoveAll(filepath.Join(config.TestDataDir(), uuid.New().String()))
		if err != nil {
			t.Fatal(err)
		}
	}, orchApi
}

func setupOrchestraAPIConfig(cfg *config.Config) {
	cfg.Partner = &config.PartnerCfg{
		Name:      "PCCWG",
		SonataUrl: "http://127.0.0.1:7777",
		Username:  "test",
		Password:  "test",
	}
}

func TestOrchestraApi_Create(t *testing.T) {
	teardownTestCase, oa := setupTestCaseOrchestraAPI(t)
	defer teardownTestCase(t)

	orchParams := api.OrderParams{}
	orchParams.Seller = &api.PartnerParams{Name: "PCCWG", ID: "PCCWG"}
	orchJson, _ := json.Marshal(orchParams)

	req := &proto.OrchestraCommonRequest{}
	req.Action = "ExecQuoteCreate"
	req.Data = string(orchJson)
	rsp, err := oa.ExecCreate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	req = &proto.OrchestraCommonRequest{}
	req.Action = "ExecOrderCreate"
	req.Data = string(orchJson)
	rsp, err = oa.ExecCreate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp)
}

func TestOrchestraApi_Find(t *testing.T) {
	teardownTestCase, oa := setupTestCaseOrchestraAPI(t)
	defer teardownTestCase(t)

	orchParams := api.FindParams{}
	orchParams.Seller = &api.PartnerParams{Name: "PCCWG", ID: "PCCWG"}
	orchJson, _ := json.Marshal(orchParams)

	req := &proto.OrchestraCommonRequest{}
	req.Action = "ExecQuoteFind"
	req.Data = string(orchJson)
	rsp, err := oa.ExecFind(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	req = &proto.OrchestraCommonRequest{}
	req.Action = "ExecOrderFind"
	req.Data = string(orchJson)
	rsp, err = oa.ExecFind(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	req = &proto.OrchestraCommonRequest{}
	req.Action = "ExecInventoryFind"
	req.Data = string(orchJson)
	rsp, err = oa.ExecFind(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp)
}

func TestOrchestraApi_Get(t *testing.T) {
	teardownTestCase, oa := setupTestCaseOrchestraAPI(t)
	defer teardownTestCase(t)

	orchParams := api.GetParams{}
	orchParams.Seller = &api.PartnerParams{Name: "PCCWG", ID: "PCCWG"}
	orchJson, _ := json.Marshal(orchParams)

	req := &proto.OrchestraCommonRequest{}
	req.Action = "ExecQuoteGet"
	req.Data = string(orchJson)
	rsp, err := oa.ExecGet(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	req = &proto.OrchestraCommonRequest{}
	req.Action = "ExecOrderGet"
	req.Data = string(orchJson)
	rsp, err = oa.ExecGet(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	req = &proto.OrchestraCommonRequest{}
	req.Action = "ExecInventoryGet"
	req.Data = string(orchJson)
	rsp, err = oa.ExecGet(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp)
}
