package orchestra

import (
	"github.com/qlcchain/go-lsobus/mock"
	sitcli "github.com/qlcchain/go-lsobus/sonata/site/client"
	sitapi "github.com/qlcchain/go-lsobus/sonata/site/client/geographic_site"
)

type sonataSiteImpl struct {
	sonataBaseImpl
}

func newSonataSiteImpl(p *PartnerImpl) *sonataSiteImpl {
	s := &sonataSiteImpl{}
	s.Partner = p
	s.Version = MEFAPIVersionSite
	return s
}

func (s *sonataSiteImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataSiteImpl) NewHTTPClient() *sitcli.APIGeographicSiteManagement {
	httpTran := s.NewHttpTransport(sitcli.DefaultBasePath)
	httpCli := sitcli.New(httpTran, nil)
	return httpCli
}

func (s *sonataSiteImpl) SendFindRequest(params *FindParams) error {
	reqParams := sitapi.NewGeographicSiteFindParams()
	//reqParams.GeographicAddressCountry = ""
	//reqParams.GeographicAddressCity = ""

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.GeographicSite.GeographicSiteFind(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateSiteFindResponse(reqParams)
	} else if err != nil {
		s.logger.Errorf("send request, error %s", err)
		s.handleResponseError(err)
		return err
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
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateSiteGetResponse(reqParams)
	} else if err != nil {
		s.logger.Errorf("send request, error %s", err)
		s.handleResponseError(err)
		return err
	}

	s.logger.Debugf("receive response, payload:", s.DumpValue(rspParams.GetPayload()))
	params.RspSite = rspParams.GetPayload()

	return nil
}

func (s *sonataSiteImpl) handleResponseError(rspErr error) {
	switch rspErr.(type) {
	case *sitapi.GeographicSiteFindUnauthorized:
		s.ClearApiToken()
	case *sitapi.GeographicSiteGetUnauthorized:
		s.ClearApiToken()
	}
}
