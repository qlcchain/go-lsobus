package contract

import (
	"os"
	"path/filepath"
	"testing"

	ct "github.com/qlcchain/go-lsobus/services/context"

	"github.com/google/uuid"

	"github.com/qlcchain/go-lsobus/config"
)

func setupTestCase(t *testing.T) (func(t *testing.T), *ContractCaller) {
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
	cfg.Partner = &config.PartnerCfg{
		BackEndURL:     "http://127.0.0.1:7777",
		Implementation: "pccwg",
		IsFake:         true,
	}
}

//func TestContractService_GetOrderInfoByInternalId(t *testing.T) {
//	teardownTestCase, cs := setupTestCase(t)
//	defer teardownTestCase(t)
//	orderInfo, err := cs.seller.GetOrderInfoByInternalId(uuid.New().String())
//	if err == nil || err != chainNotReady {
//		t.Fatal("chain is not ready")
//	}
//	orderInfo, err = cs.seller.GetOrderInfoByInternalId(uuid.New().String())
//	if err != nil {
//		t.Fatal("fake mode should return orderInfo and no error")
//	}
//	t.Log(util.ToIndentString(orderInfo))
//}
