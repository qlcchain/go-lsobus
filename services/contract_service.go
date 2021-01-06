package services

import (
	"errors"

	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/contract"
	"github.com/qlcchain/go-lsobus/log"

	"github.com/qlcchain/go-lsobus/common"
)

type ContractService struct {
	common.ServiceLifecycle
	cs     *contract.ContractCaller
	logger *zap.SugaredLogger
}

func NewContractService(cfgFile string) (*ContractService, error) {
	cs, err := contract.NewContractService(cfgFile)
	if err != nil {
		return nil, err
	}
	return &ContractService{cs: cs, logger: log.NewLogger("contract_service")}, nil
}

func (cs *ContractService) Init() error {
	if !cs.PreInit() {
		return errors.New("pre init fail")
	}
	defer cs.PostInit()
	err := cs.cs.Init()
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractService) Start() error {
	if !cs.PreStart() {
		return errors.New("pre start fail")
	}
	err := cs.cs.Start()
	if err != nil {
		cs.logger.Error(err)
		return err
	}
	cs.PostStart()
	return nil
}

func (cs *ContractService) Stop() error {
	if !cs.PreStop() {
		return errors.New("pre stop fail")
	}
	defer cs.PostStop()

	err := cs.cs.Stop()
	if err != nil {
		cs.logger.Error(err)
		return err
	}
	cs.logger.Info("cs stopped")
	return nil
}

func (cs *ContractService) Status() int32 {
	return cs.State()
}
