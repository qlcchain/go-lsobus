package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"

	"github.com/qlcchain/go-lsobus/config"
)

func TestNewContractService(t *testing.T) {
	dir := filepath.Join(config.TestDataDir(), uuid.New().String())
	cm := config.NewCfgManager(dir)
	_, err := cm.Load()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()
	cs, err := NewContractService(cm.ConfigFile)
	if err != nil {
		t.Fatal(err)
	}
	err = cs.Init()
	if err != nil {
		t.Fatal(err)
	}
	if cs.State() != 2 {
		t.Fatal("contract init failed")
	}
	err = cs.Start()
	if err != nil {
		t.Fatal(err)
	}
	if cs.State() != 4 {
		t.Fatal("contract start failed")
	}
	err = cs.Stop()
	if err != nil {
		t.Fatal(err)
	}

	if cs.Status() != 6 {
		t.Fatal("contract stop failed.")
	}
}
