package orchestra

import (
	sitcli "github.com/iixlabs/virtual-lsobus/sonata/site/client"
	sitapi "github.com/iixlabs/virtual-lsobus/sonata/site/client/geographic_site"
)

type sonataSiteImpl struct {
	sonataBaseImpl
}

func newSonataSiteImpl() *sonataSiteImpl {
	s := &sonataSiteImpl{}
	s.Version = MEFAPIVersionSite
	return s
}

func (s *sonataSiteImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataSiteImpl) NewHTTPClient() *sitcli.APIGeographicSiteManagement {
	tranCfg := sitcli.DefaultTransportConfig().WithHost(s.Host).WithSchemes([]string{s.Scheme})
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
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())

	//rspOrder := rspParams.GetPayload()

	return nil
}

func (s *sonataSiteImpl) SendGetRequest(id string) error {
	reqParams := sitapi.NewGeographicSiteGetParams()
	reqParams.SiteID = id

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.GeographicSite.GeographicSiteGet(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())

	//rspOrder := rspParams.GetPayload()

	return nil
}
