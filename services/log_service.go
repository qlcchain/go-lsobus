package services

import (
	"errors"

	"github.com/qlcchain/go-lsobus/config"
	"github.com/qlcchain/go-lsobus/log"

	"github.com/qlcchain/go-lsobus/common"
	"github.com/qlcchain/go-lsobus/services/context"
)

type LogService struct {
	common.ServiceLifecycle
	cfg *config.Config
}

func NewLogService(cfgFile string) *LogService {
	cc := context.NewServiceContext(cfgFile)
	cfg, _ := cc.Config()
	return &LogService{cfg: cfg}
}

func (ls *LogService) Init() error {
	if !ls.PreInit() {
		return errors.New("LogService pre init fail")
	}
	defer ls.PostInit()

	return log.Setup(ls.cfg)
}

func (ls *LogService) Start() error {
	if !ls.PreStart() {
		return errors.New("LogService pre start fail")
	}
	defer ls.PostStart()

	return nil
}

func (ls *LogService) Stop() error {
	if !ls.PreStop() {
		return errors.New("LogService pre stop fail")
	}
	defer ls.PostStop()

	return log.Teardown()
}

func (ls *LogService) Status() int32 {
	return ls.State()
}
