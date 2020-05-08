package orchestra

import (
	"go.uber.org/zap"

	"github.com/iixlabs/virtual-lsobus/log"
)

type sonataBaseImpl struct {
	logger *zap.SugaredLogger
}

func (s *sonataBaseImpl) Init() error {
	s.logger = log.NewLogger("sonataPOQImpl")
	return nil
}
