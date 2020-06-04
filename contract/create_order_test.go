package contract

import (
	"testing"

	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/cmd/util"
	"github.com/qlcchain/go-lsobus/mock"
)

func TestContractService_ConvertProtoToCreateOrderParam(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)

	proto := mock.ProtoCreateOrderParams()
	dodCreateParams, err := cs.convertProtoToCreateOrderParam(proto)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(util.ToIndentString(dodCreateParams))
}

func TestContractService_GetCreateOrderBlock(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	_, pri, err := pkg.GenerateAddress()
	cs.account = pkg.NewAccount(pri)
	proto := mock.ProtoCreateOrderParams()

	id, err := cs.GetCreateOrderBlock(proto)
	if err == nil {
		t.Fatal("buyer address should be not match")
	}
	proto.Buyer.Address = cs.account.Address().String()
	id, err = cs.GetCreateOrderBlock(proto)
	if err != chainNotReady {
		t.Fatal(err)
	}
	cs.SetFakeMode(true)
	id, err = cs.GetCreateOrderBlock(proto)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
