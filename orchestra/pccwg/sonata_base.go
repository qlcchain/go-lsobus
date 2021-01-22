package pccwg

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/qlcchain/qlc-go-sdk/pkg/util"

	"github.com/qlcchain/go-lsobus/orchestra/sonata/common/models"

	"github.com/qlcchain/go-lsobus/api"

	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/log"
)

const (
	MEFAPIVersionSite  = "3"
	MEFAPIVersionPOQ   = "3"
	MEFAPIVersionQuote = "2"
	MEFAPIVersionOrder = "3"
	MEFAPIVersionInv   = "3"
	MEFAPIVersionOffer = "1"

	MEFSchemaLocationRoot      = "https://github.com/MEF-GIT/MEF-LSO-Sonata-SDK/blob/working-draft"
	MEFSchemaLocationSpecRoot  = MEFSchemaLocationRoot + "/payload_descriptions/ProductSpecDescription"
	MEFSchemaLocationSpecUNI   = MEFSchemaLocationSpecRoot + "/MEF_UNISpec_v3.json"
	MEFSchemaLocationSpecELine = MEFSchemaLocationSpecRoot + "/MEF_ELineSpec_v3.json"

	MEFProductOfferingUNI   = "LSO_Sonata_DataOnDemand_EthernetPort_UNI"
	MEFProductOfferingELine = "LSO_Sonata_DataOnDemand_EthernetConnection"
)

type sonataBaseImpl struct {
	Partner api.DoDSeller

	URL     string
	Scheme  string
	Host    string
	Version string

	logger *zap.SugaredLogger
	itemID atomic.Int32
}

func (s *sonataBaseImpl) Init() error {
	s.URL = s.Partner.GetSellerConfig().BackEndURL
	if s.URL != "" {
		retUrl, err := url.Parse(s.URL)
		if err != nil {
			return fmt.Errorf("sonata url parse err %s", err)
		}
		s.Scheme = retUrl.Scheme
		s.Host = retUrl.Host
	}

	if s.Scheme == "" {
		s.Scheme = "http"
	}
	if s.Host == "" {
		s.Host = "127.0.0.1:8080"
	}
	if s.Version == "" {
		s.Version = "1"
	}
	s.logger = log.NewLogger("sonataImpl")
	return nil
}

func (s *sonataBaseImpl) GetFakeMode() bool {
	return s.Partner.GetSellerConfig().IsFake
}

func (s *sonataBaseImpl) GetHost() string {
	return s.Host
}

func (s *sonataBaseImpl) GetScheme() string {
	return s.Scheme
}

func (s *sonataBaseImpl) GetApiToken() string {
	return s.Partner.GetAPIToken()
}

func (s *sonataBaseImpl) RenewApiToken() string {
	return s.Partner.RenewAPIToken()
}

func (s *sonataBaseImpl) ClearApiToken() {
	s.Partner.ClearAPIToken()
}

func (s *sonataBaseImpl) NewHttpTransport(basePath string) *httptransport.Runtime {
	httpTran := httptransport.New(s.GetHost(), basePath, []string{s.GetScheme()})
	httpTran.DefaultAuthentication = httptransport.BearerToken(s.GetApiToken())
	return httpTran
}

func (s *sonataBaseImpl) NewItemID() string {
	return strconv.Itoa(int(s.itemID.Inc()))
}

func (s *sonataBaseImpl) BuildUNIProductSpec(params *api.UNIItemParams) *models.UNISpec {
	uniSpec := &models.UNISpec{}
	uniSpec.SetAtSchemaLocation(MEFSchemaLocationSpecUNI)
	uniSpec.SetAtType("UNISpec")

	if params.PortSpeed == 1000 {
		uniSpec.PhysicalLayer = []models.PhysicalLayer{models.PhysicalLayerNr1000BASET}
	} else if params.PortSpeed == 10000 {
		uniSpec.PhysicalLayer = []models.PhysicalLayer{models.PhysicalLayerNr10GBASESR}
	} else {
		uniSpec.PhysicalLayer = []models.PhysicalLayer{models.PhysicalLayerNr100BASETX}
	}
	uniSpec.MaxServiceFrameSize = 1522
	uniSpec.NumberOfLinks = 1

	return uniSpec
}

func (s *sonataBaseImpl) BuildELineProductSpec(params *api.ELineItemParams) *models.ELineSpec {
	lineSpec := &models.ELineSpec{}
	//lineSpec.SetAtSchemaLocation(MEFSchemaLocationSpecELine)
	lineSpec.SetAtType("ELineSpec")

	lineSpec.ClassOfServiceName = params.CosName
	lineSpec.MaximumFrameSize = 1526
	lineSpec.SVlanID = int32(params.SVlanID)
	bwMbps := int32(params.Bandwidth)
	bwProfile := &models.BandwidthProfile{
		Cir: &models.InformationRate{Unit: models.InformationRateUnit(params.BwUnit), Amount: &bwMbps},
	}
	lineSpec.ENNIIngressBWProfile = []*models.BandwidthProfile{bwProfile}
	lineSpec.UNIIngressBWProfile = []*models.BandwidthProfile{bwProfile}

	return lineSpec
}

func (s *sonataBaseImpl) BuildPCCWConnProductSpec(params *api.ELineItemParams) *models.PCCWConnSpec {
	lineSpec := &models.PCCWConnSpec{}
	//lineSpec.SetAtSchemaLocation(MEFSchemaLocationSpecELine)
	lineSpec.SetAtType("PCCWConnSpec")

	//lineSpec.Type = "ELINE"
	lineSpec.Name = params.Name
	lineSpec.ClassOfService = params.CosName
	lineSpec.Bandwidth = int32(params.Bandwidth)
	lineSpec.SrcPortID = params.SrcPortID
	lineSpec.DestPortID = params.DstPortID
	lineSpec.DestCompanyID = params.DstCompanyID
	lineSpec.DestMetroID = params.DstMetroID
	lineSpec.SrcLocationID = params.SrcLocationID
	lineSpec.DestLocationID = params.DstLocationID

	if params.BillingParams != nil {
		lineSpec.StartedAt.Scan(time.Unix(params.BillingParams.StartTime, 0))
		lineSpec.TerminatedAt.Scan(time.Unix(params.BillingParams.EndTime, 0))
	}

	return lineSpec
}

func (s *sonataBaseImpl) DumpValue(v interface{}) string {
	return util.ToIndentString(v)
}
