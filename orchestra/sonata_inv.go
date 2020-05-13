package orchestra

import (
	invcli "github.com/iixlabs/virtual-lsobus/sonata/inventory/client"
	invapi "github.com/iixlabs/virtual-lsobus/sonata/inventory/client/product"
)

type sonataInvImpl struct {
	sonataBaseImpl
}

func newSonataInvImpl() *sonataInvImpl {
	s := &sonataInvImpl{}
	s.Version = MEFAPIVersionInv
	return s
}

func (s *sonataInvImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataInvImpl) NewHTTPClient() *invcli.APIProductInventoryManagement {
	tranCfg := invcli.DefaultTransportConfig().WithHost(s.GetHost()).WithSchemes([]string{s.GetScheme()})
	httpCli := invcli.NewHTTPClientWithConfig(nil, tranCfg)
	return httpCli
}

func (s *sonataInvImpl) SendFindRequest(params *FindParams) error {
	reqParams := invapi.NewProductFindParams()
	if params.BuyerID != "" {
		reqParams.BuyerID = &params.BuyerID
	}
	if params.State != "" {
		reqParams.Status = &params.State
	}
	if params.SiteID != "" {
		reqParams.GeographicalSiteID = &params.SiteID
	}
	if params.ProductSpecificationID != "" {
		reqParams.ProductSpecificationID = &params.ProductSpecificationID
	}
	if params.Offset != "" {
		reqParams.Offset = &params.Offset
	}
	if params.Limit != "" {
		reqParams.Limit = &params.Limit
	}

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.Product.ProductFind(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())

	return nil
}

func (s *sonataInvImpl) SendGetRequest(id string) error {
	reqParams := invapi.NewProductGetParams()
	reqParams.ProductID = id

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.Product.ProductGet(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())

	//rspOrder := rspParams.GetPayload()

	return nil
}
