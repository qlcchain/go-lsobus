package services

import (
	"github.com/qlcchain/go-virtual-lsobus/services/context"
)

//RegisterServices register services to chain context
func RegisterServices(cs *context.ServiceContext) error {
	cfgFile := cs.ConfigFile()

	logService := NewLogService(cfgFile)
	_ = cs.Register(context.LogService, logService)
	_ = logService.Init()

	if contractService, err := NewContractService(cfgFile); err != nil {
		return err
	} else {
		_ = cs.Register(context.ContractService, contractService)
		if rpcService, err := NewRPCService(cfgFile, contractService); err != nil {
			return err
		} else {
			_ = cs.Register(context.RPCService, rpcService)
		}
	}

	return nil
}
