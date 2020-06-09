package contract

import (
	"testing"

	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/google/uuid"
)

func TestContractService_GetContractStatus(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	internalId := uuid.New().String()
	cs.orderIdOnChainBuyer.Store(internalId, "")
	_, pri, err := pkg.GenerateAddress()
	if err != nil {
		t.Fatal(err)
	}
	cs.account = pkg.NewAccount(pri)
	cs.SetFakeMode(true)
	cs.getContractStatus()
	if _, ok := cs.orderIdOnChainBuyer.Load(internalId); ok {
		t.Fatal("id should not exit")
	}
}
