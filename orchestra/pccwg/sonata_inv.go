package pccwg

import (
	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/mock"
	invcli "github.com/qlcchain/go-lsobus/orchestra/sonata/inventory/client"
	invapi "github.com/qlcchain/go-lsobus/orchestra/sonata/inventory/client/product"
)

type sonataInvImpl struct {
	sonataBaseImpl
}

func newSonataInvImpl(p api.DoDSeller) *sonataInvImpl {
	s := &sonataInvImpl{}
	s.Partner = p
	s.Version = MEFAPIVersionInv
	return s
}

func (s *sonataInvImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataInvImpl) NewHTTPClient() *invcli.APIProductInventoryManagement {
	httpTran := s.NewHttpTransport(invcli.DefaultBasePath)
	httpCli := invcli.New(httpTran, nil)
	return httpCli
}

func (s *sonataInvImpl) SendFindRequest(params *api.FindParams) error {
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
	if params.ProductOfferingID != "" {
		reqParams.ProductOfferingID = &params.ProductOfferingID
	}
	if params.ProductOrderID != "" {
		reqParams.ProductOrderID = &params.ProductOrderID
	}
	if params.Offset != "" {
		reqParams.Offset = &params.Offset
	}
	if params.Limit != "" {
		reqParams.Limit = &params.Limit
	}

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.Product.ProductFind(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateInvFindResponse(reqParams)
	} else if err != nil {
		s.logger.Errorf("send request, error %s", err)
		s.handleResponseError(err)
		return err
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspInvList = rspParams.GetPayload()

	return nil
}

func (s *sonataInvImpl) SendGetRequest(params *api.GetParams) error {
	reqParams := invapi.NewProductGetParams()
	reqParams.ProductID = params.ID

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.Product.ProductGet(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateInvGetResponse(reqParams)
	} else if err != nil {
		s.logger.Errorf("send request, error %s", err)
		s.handleResponseError(err)
		return err
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspInv = rspParams.GetPayload()

	return nil
}

func (s *sonataInvImpl) handleResponseError(rspErr error) {
	switch rspErr.(type) {
	case *invapi.ProductFindUnauthorized, *invapi.ProductFindForbidden:
		s.ClearApiToken()
	case *invapi.ProductGetUnauthorized, *invapi.ProductGetForbidden:
		s.ClearApiToken()
	}
}
