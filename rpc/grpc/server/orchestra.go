package grpcServer

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/api"

	"github.com/qlcchain/go-lsobus/log"
	"github.com/qlcchain/go-lsobus/rpc/grpc/proto"
)

type OrchestraApi struct {
	logger *zap.SugaredLogger
	seller api.DoDSeller
}

func NewOrchestraAPI(seller api.DoDSeller) *OrchestraApi {
	oa := &OrchestraApi{
		seller: seller,
		logger: log.NewLogger("OrchestraApi"),
	}

	return oa
}

func (oa *OrchestraApi) ExecCreate(
	ctx context.Context, param *proto.OrchestraCommonRequest,
) (*proto.OrchestraCommonResponse, error) {
	oa.logger.Debugf("ExecCreate request %+v", param)

	orchParams := oa.Request2OrchCreateParams(param)

	var execErr error
	var execRspData interface{}
	switch param.GetAction() {
	case "ExecQuoteCreate":
		execErr = oa.seller.ExecQuoteCreate(orchParams)
		execRspData = orchParams.RspQuote
	case "ExecOrderCreate":
		execErr = oa.seller.ExecOrderCreate(orchParams)
		execRspData = orchParams.RspOrder
	default:
		return nil, errors.New("invalid ExecAction")
	}
	if execErr != nil {
		oa.logger.Errorf("%s err %s", param.GetAction(), execErr)
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

	oa.logger.Debugf("ExecCreate response %+v", rsp)

	return rsp, nil
}

func (oa *OrchestraApi) ExecFind(
	ctx context.Context, param *proto.OrchestraCommonRequest,
) (*proto.OrchestraCommonResponse, error) {
	oa.logger.Debugf("ExecFind request %+v", param)

	orchParams := oa.Request2OrchFindParams(param)

	var execErr error
	var execRspData interface{}
	switch param.GetAction() {
	case "ExecQuoteFind":
		execErr = oa.seller.ExecQuoteFind(orchParams)
		execRspData = orchParams.RspQuoteList
	case "ExecOrderFind":
		execErr = oa.seller.ExecOrderFind(orchParams)
		execRspData = orchParams.RspOrderList
	case "ExecInventoryFind":
		execErr = oa.seller.ExecInventoryFind(orchParams)
		execRspData = orchParams.RspInvList
	default:
		return nil, errors.New("invalid ExecAction")
	}
	if execErr != nil {
		oa.logger.Errorf("%s err %s", param.GetAction(), execErr)
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

	oa.logger.Debugf("ExecFind response %+v", rsp)

	return rsp, nil
}

func (oa *OrchestraApi) ExecGet(
	ctx context.Context, param *proto.OrchestraCommonRequest,
) (*proto.OrchestraCommonResponse, error) {
	oa.logger.Debugf("ExecGet request %+v", param)

	orchParams := oa.Request2OrchGetParams(param)

	var execErr error
	var execRspData interface{}
	switch param.GetAction() {
	case "ExecQuoteGet":
		execErr = oa.seller.ExecQuoteGet(orchParams)
		execRspData = orchParams.RspQuote
	case "ExecOrderGet":
		execErr = oa.seller.ExecOrderGet(orchParams)
		execRspData = orchParams.RspOrder
	case "ExecInventoryGet":
		execErr = oa.seller.ExecInventoryGet(orchParams)
		execRspData = orchParams.RspInv
	default:
		return nil, errors.New("invalid ExecAction")
	}
	if execErr != nil {
		oa.logger.Errorf("%s err %s", param.GetAction(), execErr)
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

	oa.logger.Debugf("ExecGet response %+v", rsp)

	return rsp, nil
}

func (oa *OrchestraApi) Request2OrchCreateParams(param *proto.OrchestraCommonRequest) *api.OrderParams {
	orchParams := &api.OrderParams{}

	err := json.Unmarshal([]byte(param.Data), orchParams)
	if err != nil {
		oa.logger.Errorf("json Unmarshal, err", err)
	}

	return orchParams
}

func (oa *OrchestraApi) Request2OrchFindParams(param *proto.OrchestraCommonRequest) *api.FindParams {
	orchParams := &api.FindParams{}

	err := json.Unmarshal([]byte(param.Data), orchParams)
	if err != nil {
		oa.logger.Errorf("json Unmarshal, err", err)
	}

	return orchParams
}

func (oa *OrchestraApi) Request2OrchGetParams(param *proto.OrchestraCommonRequest) *api.GetParams {
	orchParams := &api.GetParams{}

	err := json.Unmarshal([]byte(param.Data), orchParams)
	if err != nil {
		oa.logger.Errorf("json Unmarshal, err", err)
	}

	return orchParams
}
