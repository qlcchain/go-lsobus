package grpcServer

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/orchestra"

	"github.com/qlcchain/go-lsobus/log"
	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

type OrchestraApi struct {
	logger *zap.SugaredLogger
	orch   *orchestra.Orchestra
}

func NewOrchestraApi(orch *orchestra.Orchestra) *OrchestraApi {
	oa := &OrchestraApi{
		orch:   orch,
		logger: log.NewLogger("OrchestraApi"),
	}

	return oa
}

func (oa *OrchestraApi) ExecCreate(ctx context.Context, param *proto.OrchestraCommonRequest) (*proto.OrchestraCommonResponse, error) {
	orchParams := oa.Request2OrchCreateParams(param)

	var execErr error
	var execRspData interface{}

	switch param.GetAction() {
	case "ExecQuoteCreate":
		execErr = oa.orch.ExecQuoteCreate(orchParams)
		execRspData = orchParams.RspQuote
	case "ExecOrderCreate":
		execErr = oa.orch.ExecOrderCreate(orchParams)
		execRspData = orchParams.RspOrder
	default:
		return nil, errors.New("invalid ExecAction")
	}
	if execErr != nil {
		return nil, execErr
	}

	rsp := &proto.OrchestraCommonResponse{Action: param.GetAction()}
	rsp.TotalCount = 1
	rsp.ResultCount = 1

	dataBytes, err := json.Marshal(execRspData)
	if err != nil {
		return nil, err
	}
	rsp.Data = string(dataBytes)

	return rsp, nil
}

func (oa *OrchestraApi) ExecFind(ctx context.Context, param *proto.OrchestraCommonRequest) (*proto.OrchestraCommonResponse, error) {
	orchParams := oa.Request2OrchFindParams(param)

	var execErr error
	var execRspData interface{}

	switch param.GetAction() {
	case "ExecQuoteFind":
		execErr = oa.orch.ExecQuoteFind(orchParams)
		execRspData = orchParams.RspQuoteList
	case "ExecOrderFind":
		execErr = oa.orch.ExecOrderFind(orchParams)
		execRspData = orchParams.RspOrderList
	case "ExecInventoryFind":
		execErr = oa.orch.ExecInventoryFind(orchParams)
		execRspData = orchParams.RspInvList
	default:
		return nil, errors.New("invalid ExecAction")
	}
	if execErr != nil {
		return nil, execErr
	}

	rsp := &proto.OrchestraCommonResponse{Action: param.GetAction()}
	rsp.TotalCount = orchParams.XTotalCount
	rsp.ResultCount = orchParams.XResultCount

	dataBytes, err := json.Marshal(execRspData)
	if err != nil {
		return nil, err
	}
	rsp.Data = string(dataBytes)

	return rsp, nil
}

func (oa *OrchestraApi) ExecGet(ctx context.Context, param *proto.OrchestraCommonRequest) (*proto.OrchestraCommonResponse, error) {
	orchParams := oa.Request2OrchGetParams(param)

	var execErr error
	var execRspData interface{}

	switch param.GetAction() {
	case "ExecQuoteGet":
		execErr = oa.orch.ExecQuoteGet(orchParams)
		execRspData = orchParams.RspQuote
	case "ExecOrderGet":
		execErr = oa.orch.ExecOrderGet(orchParams)
		execRspData = orchParams.RspOrder
	case "ExecInventoryGet":
		execErr = oa.orch.ExecInventoryGet(orchParams)
		execRspData = orchParams.RspInv
	default:
		return nil, errors.New("invalid ExecAction")
	}
	if execErr != nil {
		return nil, execErr
	}

	rsp := &proto.OrchestraCommonResponse{Action: param.GetAction()}
	rsp.TotalCount = 1
	rsp.ResultCount = 1

	dataBytes, err := json.Marshal(execRspData)
	if err != nil {
		return nil, err
	}
	rsp.Data = string(dataBytes)

	return rsp, nil
}

func (oa *OrchestraApi) Request2OrchCreateParams(param *proto.OrchestraCommonRequest) *orchestra.OrderParams {
	orchParams := &orchestra.OrderParams{}

	err := json.Unmarshal([]byte(param.Data), orchParams)
	if err != nil {
		oa.logger.Errorf("json Unmarshal, err", err)
	}

	return orchParams
}

func (oa *OrchestraApi) Request2OrchFindParams(param *proto.OrchestraCommonRequest) *orchestra.FindParams {
	orchParams := &orchestra.FindParams{}

	err := json.Unmarshal([]byte(param.Data), orchParams)
	if err != nil {
		oa.logger.Errorf("json Unmarshal, err", err)
	}

	return orchParams
}

func (oa *OrchestraApi) Request2OrchGetParams(param *proto.OrchestraCommonRequest) *orchestra.GetParams {
	orchParams := &orchestra.GetParams{}

	err := json.Unmarshal([]byte(param.Data), orchParams)
	if err != nil {
		oa.logger.Errorf("json Unmarshal, err", err)
	}

	return orchParams
}
