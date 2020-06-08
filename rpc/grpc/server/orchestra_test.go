package grpcServer

import (
	"context"
	"encoding/json"

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

	orch := orchestra.NewOrchestra(cfgFile)
	orch.SetFakeMode(true)

	if err = orch.Init(); err != nil {
		t.Fatal(err)
	}

	orchApi := NewOrchestraApi(orch)
	return func(t *testing.T) {
		err = os.RemoveAll(filepath.Join(config.TestDataDir(), uuid.New().String()))
		if err != nil {
			t.Fatal(err)
		}
	}, orchApi
}

func setupOrchestraAPIConfig(cfg *config.Config) {
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

func TestOrchestraApi_Create(t *testing.T) {
	teardownTestCase, oa := setupTestCaseOrchestraAPI(t)
	defer teardownTestCase(t)

	orchParams := orchestra.OrderParams{}
	orchParams.Seller = &orchestra.PartnerParams{Name: "PCCWG", ID: "PCCWG"}
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

	orchParams := orchestra.FindParams{}
	orchParams.Seller = &orchestra.PartnerParams{Name: "PCCWG", ID: "PCCWG"}
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

	orchParams := orchestra.GetParams{}
	orchParams.Seller = &orchestra.PartnerParams{Name: "PCCWG", ID: "PCCWG"}
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
