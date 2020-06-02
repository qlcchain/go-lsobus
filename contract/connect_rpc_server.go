package contract

import (
	"time"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
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
					client, err := qlcSdk.NewQLCClient(cs.cfg.ChainUrl)
					if err != nil || client == nil {
						continue
					} else {
						cs.client = client
						s, err := cs.client.Pov.GetPovStatus()
						if err != nil {
							continue
						} else if s.SyncState == 2 {
							cs.chainReady = true
							cs.quit <- true
						}
					}
				} else {
					s, err := cs.client.Pov.GetPovStatus()
					if err != nil {
						continue
					} else if s.SyncState == 2 {
						cs.chainReady = true
						cs.quit <- true
					}
				}
			}
		}
	}
}
