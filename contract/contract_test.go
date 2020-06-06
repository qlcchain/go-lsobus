package contract

import (
	"os"
	"path/filepath"
	"testing"

	ct "github.com/qlcchain/go-lsobus/services/context"

	"github.com/google/uuid"

	"github.com/qlcchain/go-lsobus/cmd/util"
	"github.com/qlcchain/go-lsobus/config"
)

func setupTestCase(t *testing.T) (func(t *testing.T), *ContractService) {
	cfgFile := filepath.Join(config.TestDataDir(), uuid.New().String(), config.CfgFileName)
	cc := ct.NewServiceContext(cfgFile)
	cfg, err := cc.Config()
	setupOrchestraConfig(cfg)
	cs, err := NewContractService(cfgFile)
	if err != nil {
		t.Fatal(err)
	}
	if err = cs.Init(); err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		err = cs.Stop()
		if err != nil {
			t.Fatal(err)
		}
		err = os.RemoveAll(filepath.Join(config.TestDataDir(), uuid.New().String()))
		if err != nil {
			t.Fatal(err)
		}
	}, cs
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
	cfg.FakeMode = true
}

func TestContractService_GetOrderInfoByInternalId(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	orderInfo, err := cs.GetOrderInfoByInternalId(uuid.New().String())
	if err == nil || err != chainNotReady {
		t.Fatal("chain is not ready")
	}
	cs.SetFakeMode(true)
	orderInfo, err = cs.GetOrderInfoByInternalId(uuid.New().String())
	if err != nil {
		t.Fatal("fake mode should return orderInfo and no error")
	}
	t.Log(util.ToIndentString(orderInfo))
}
