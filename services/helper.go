package services

import (
	"github.com/iixlabs/virtual-lsobus/services/context"
)

//RegisterServices register services to chain context
func RegisterServices(cs *context.ServiceContext) error {
	cfgFile := cs.ConfigFile()

	logService := NewLogService(cfgFile)
	_ = cs.Register(context.LogService, logService)
	_ = logService.Init()

	if rpcService, err := NewRPCService(cfgFile); err != nil {
		return err
	} else {
		_ = cs.Register(context.RPCService, rpcService)
	}

	return nil
}
