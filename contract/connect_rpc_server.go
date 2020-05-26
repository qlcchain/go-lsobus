package contract

import (
	"time"

	qlcchain "github.com/qlcchain/go-qlc/rpc/api"
	rpc "github.com/qlcchain/jsonrpc2"
)

func (cs *ContractService) connectRpcServer() {
	ticker := time.NewTicker(connectRpcServerInterval)
	for {
		select {
		case <-cs.quit:
			return
		case <-ticker.C:
			if cs.cfg.ChainUrl != "" {
				if cs.client == nil {
					client, err := rpc.Dial(cs.cfg.ChainUrl)
					if err != nil || client == nil {
						continue
					} else {
						cs.client = client
						var pov qlcchain.PovStatus
						err := cs.client.Call(&pov, "pov_getPovStatus")
						if err != nil {
							continue
						} else if pov.SyncState == 2 {
							cs.chainReady = true
							cs.quit <- true
						}
					}
				} else {
					var pov qlcchain.PovStatus
					err := cs.client.Call(&pov, "pov_getPovStatus")
					if err != nil {
						continue
					} else if pov.SyncState == 2 {
						cs.chainReady = true
						cs.quit <- true
					}
				}
			}
		}
	}
}
