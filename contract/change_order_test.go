package contract

import (
	"testing"

	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/cmd/util"
	"github.com/qlcchain/go-lsobus/mock"
)

func TestContractService_ConvertProtoToChangeOrderParam(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)

	proto := mock.ProtoChangeOrderParams()
	dodChangeParams, err := cs.convertProtoToChangeOrderParam(proto)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(util.ToIndentString(dodChangeParams))
}

func TestContractService_GetChangeOrderBlock(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	_, pri, err := pkg.GenerateAddress()
	cs.account = pkg.NewAccount(pri)
	proto := mock.ProtoChangeOrderParams()

	id, err := cs.GetChangeOrderBlock(proto)
	if err == nil {
		t.Fatal("buyer address should be not match")
	}
	proto.Buyer.Address = cs.account.Address().String()
	id, err = cs.GetChangeOrderBlock(proto)
	if err != chainNotReady {
		t.Fatal(err)
	}
	cs.SetFakeMode(true)
	id, err = cs.GetChangeOrderBlock(proto)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
