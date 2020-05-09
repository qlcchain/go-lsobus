package orchestra

import (
	"strconv"

	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/iixlabs/virtual-lsobus/log"
)

const (
	MEFSchemaLocationRoot      = "https://github.com/MEF-GIT/MEF-LSO-Sonata-SDK/blob/working-draft"
	MEFSchemaLocationSpecRoot  = MEFSchemaLocationRoot + "/payload_descriptions/ProductSpecDescription"
	MEFSchemaLocationSpecUNI   = MEFSchemaLocationSpecRoot + "/MEF_UNISpec_v3.json"
	MEFSchemaLocationSpecELine = MEFSchemaLocationSpecRoot + "/MEF_ELineSpec_v3.json"
)

type sonataBaseImpl struct {
	logger *zap.SugaredLogger
	itemID atomic.Int32
}

func (s *sonataBaseImpl) Init() error {
	s.logger = log.NewLogger("sonataImpl")
	return nil
}

func (s *sonataBaseImpl) NewItemID() string {
	return strconv.Itoa(int(s.itemID.Inc()))
}
