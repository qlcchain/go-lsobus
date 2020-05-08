package contract

import (
	"context"

	qlcchain "github.com/qlcchain/qlc-go-sdk"
	qlctypes "github.com/qlcchain/qlc-go-sdk/pkg/types"
	"go.uber.org/zap"

	"github.com/iixlabs/virtual-lsobus/config"
)

type ContractService struct {
	cfg     *config.Config
	account *qlctypes.Account
	logger  *zap.SugaredLogger
	client  *qlcchain.QLCClient
	ctx     context.Context
	cancel  context.CancelFunc
}
