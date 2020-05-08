package contract

import (
	"context"

	"github.com/iixlabs/virtual-lsobus/config"
	qlcchain "github.com/qlcchain/qlc-go-sdk"
	qlctypes "github.com/qlcchain/qlc-go-sdk/pkg/types"
	"go.uber.org/zap"
)

type ContractService struct {
	cfg     *config.Config
	account *qlctypes.Account
	logger  *zap.SugaredLogger
	client  *qlcchain.QLCClient
	ctx     context.Context
	cancel  context.CancelFunc
}
