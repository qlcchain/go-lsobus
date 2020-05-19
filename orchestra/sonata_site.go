package orchestra

import (
	"github.com/qlcchain/go-lsobus/mock"
	sitcli "github.com/qlcchain/go-lsobus/sonata/site/client"
	sitapi "github.com/qlcchain/go-lsobus/sonata/site/client/geographic_site"
)

type sonataSiteImpl struct {
	sonataBaseImpl
}

func newSonataSiteImpl(o *Orchestra) *sonataSiteImpl {
	s := &sonataSiteImpl{}
	s.Orch = o
	s.Version = MEFAPIVersionSite
	return s
}

func (s *sonataSiteImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataSiteImpl) NewHTTPClient() *sitcli.APIGeographicSiteManagement {
	tranCfg := sitcli.DefaultTransportConfig().WithHost(s.GetHost()).WithSchemes([]string{s.GetScheme()})
	httpCli := sitcli.NewHTTPClientWithConfig(nil, tranCfg)
	return httpCli
}

func (s *sonataSiteImpl) SendFindRequest(params *FindParams) error {
	reqParams := sitapi.NewGeographicSiteFindParams()
	//reqParams.GeographicAddressCountry = ""
	//reqParams.GeographicAddressCity = ""

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.GeographicSite.GeographicSiteFind(reqParams)
	if err != nil {
		s.logger.Errorf("send request, error %s", err)
		//return err
		rspParams = mock.SonataGenerateSiteFindResponse(reqParams)
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspSiteList = rspParams.GetPayload()

	return nil
}

func (s *sonataSiteImpl) SendGetRequest(params *GetParams) error {
	reqParams := sitapi.NewGeographicSiteGetParams()
	reqParams.SiteID = params.ID

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.GeographicSite.GeographicSiteGet(reqParams)
	if err != nil {
		s.logger.Errorf("send request, error %s", err)
		return err
	}

	s.logger.Debugf("receive response, payload:", s.DumpValue(rspParams.GetPayload()))
	params.RspSite = rspParams.GetPayload()

	return nil
}
