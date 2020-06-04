package contract

import (
	"testing"

	"github.com/google/uuid"
)

func TestContractService_GetContractStatus(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	internalId := uuid.New().String()
	cs.orderIdOnChain.Store(internalId, "")
	cs.SetFakeMode(true)
	cs.getContractStatus()
	if _, ok := cs.orderIdOnChain.Load(internalId); ok {
		t.Fatal("id should not exit")
	}
	if _, ok := cs.orderIdFromSonata.Load(internalId); !ok {
		t.Fatal("order id from sonata should be store")
	}
}
