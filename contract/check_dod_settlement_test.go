package contract

import (
	"testing"

	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"
)

func TestContractService_ProcessDoDContract(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	_, pri, err := pkg.GenerateAddress()
	if err != nil {
		t.Fatal(err)
	}
	cs.account = pkg.NewAccount(pri)
	cs.SetFakeMode(true)
	cs.processDoDContract()
}
