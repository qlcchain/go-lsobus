package contract

import (
	"testing"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"

	"github.com/google/uuid"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/mock"
)

func TestContractService_GetProductId(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	_, pri, _ := pkg.GenerateAddress()
	cs.account = pkg.NewAccount(pri)
	orderInfo, err := mock.OrderInfo()
	if err != nil {
		t.Fatal(err)
	}
	orderInfo.Seller.Address = cs.account.Address()
	id1 := uuid.New().String()
	cs.orderIdFromSonata.Store(id1, orderInfo)
	cs.SetFakeMode(true)
	cs.getProductId()
	if _, ok := cs.orderIdFromSonata.Load(id1); ok {
		t.Fatal("id should not exit")
	}
	orderInfo.OrderType, _ = qlcSdk.ParseDoDSettleOrderType("change")
	id2 := uuid.New().String()
	cs.orderIdFromSonata.Store(id2, orderInfo)
	cs.getProductId()
	if _, ok := cs.orderIdFromSonata.Load(id2); ok {
		t.Fatal("id should not exit")
	}
}
