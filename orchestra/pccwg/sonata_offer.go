package pccwg

import (
	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-lsobus/api"

	"github.com/qlcchain/go-lsobus/mock"
	"github.com/qlcchain/go-lsobus/orchestra/sonata/offer"
)

type sonataOfferImpl struct {
	sonataBaseImpl
}

func newSonataOfferImpl(p api.DoDSeller) *sonataOfferImpl {
	s := &sonataOfferImpl{}
	s.Partner = p
	s.Version = MEFAPIVersionOffer
	return s
}

func (s *sonataOfferImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataOfferImpl) SendFindRequest(params *api.FindParams) error {
	offUrl := s.URL + "/api/mef/productOfferingManagement/v1/productOffering"
	offapi := offer.NewAPIProductOfferingManagement(offUrl)

	reqParams := &offer.ProductOfferingFindParams{}
	reqParams.ApiToken = s.GetApiToken()

	rspParams, err := offapi.ProductOfferingFind(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateOfferFindResponse(reqParams)
	} else if err != nil {
		s.handleResponseError(err)
		return err
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.Payload))
	params.RspOfferList = rspParams.Payload
	params.XResultCount, err = swag.ConvertInt32(rspParams.XResultCount)
	params.XTotalCount, err = swag.ConvertInt32(rspParams.XTotalCount)

	return nil
}

func (s *sonataOfferImpl) SendGetRequest(params *api.GetParams) error {
	offUrl := s.URL + "/api/mef/productOfferingManagement/v1/productOffering"
	offapi := offer.NewAPIProductOfferingManagement(offUrl)

	reqParams := &offer.ProductOfferingGetParams{ProductOfferingID: params.ID}
	reqParams.ApiToken = s.GetApiToken()

	rspParams, err := offapi.ProductOfferingGet(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateOfferGetResponse(reqParams)
	} else if err != nil {
		s.handleResponseError(err)
		return nil
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.Payload))
	params.RspOffer = rspParams.Payload

	return nil
}

func (s *sonataOfferImpl) handleResponseError(rspErr error) {
}
