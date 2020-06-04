package contract

import (
	"testing"
	"time"
)

func TestContractService_ConnectRpcServer(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	cs.cfg.ChainUrl = "http://47.103.40.20:19735"
	go cs.connectRpcServer()
	time.Sleep(6 * time.Second)
	if !cs.chainReady {
		t.Fatal("connect to rpc server fail")
	}
	go cs.connectRpcServer()
	time.Sleep(6 * time.Second)
}
