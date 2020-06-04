package contract

import (
	"testing"

	"github.com/qlcchain/go-lsobus/cmd/util"
	"github.com/qlcchain/go-lsobus/mock"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"
)

func TestContractService_ConvertProtoToTerminateOrderParam(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)

	proto := mock.ProtoTerminateOrderParams()
	dodCreateParams, err := cs.convertProtoToTerminateOrderParam(proto)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(util.ToIndentString(dodCreateParams))
}

func TestContractService_GetTerminateOrderBlock(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	_, pri, err := pkg.GenerateAddress()
	cs.account = pkg.NewAccount(pri)
	proto := mock.ProtoTerminateOrderParams()

	id, err := cs.GetTerminateOrderBlock(proto)
	if err == nil {
		t.Fatal("buyer address should be not match")
	}
	proto.Buyer.Address = cs.account.Address().String()
	id, err = cs.GetTerminateOrderBlock(proto)
	if err != chainNotReady {
		t.Fatal(err)
	}
	cs.SetFakeMode(true)
	id, err = cs.GetTerminateOrderBlock(proto)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
