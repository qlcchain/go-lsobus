package orchestra

import (
	"github.com/qlcchain/go-lsobus/mock"
	"github.com/qlcchain/go-lsobus/sonata/offer"
)

type sonataOfferImpl struct {
	sonataBaseImpl
}

func newSonataOfferImpl(p *PartnerImpl) *sonataOfferImpl {
	s := &sonataOfferImpl{}
	s.Partner = p
	s.Version = MEFAPIVersionOffer
	return s
}

func (s *sonataOfferImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataOfferImpl) SendFindRequest(params *FindParams) error {
	offUrl := s.URL + "/api/mef/productOfferingManagement/v1/productOffering"
	offapi := offer.NewAPIProductOfferingManagement(offUrl)

	reqParams := &offer.ProductOfferingFindParams{}
	rspParams, err := offapi.ProductOfferingFind(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateOfferFindResponse(reqParams)
	} else if err != nil {
		return err
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.Data))
	params.RspOfferList = rspParams.Data

	return nil
}

func (s *sonataOfferImpl) SendGetRequest(params *GetParams) error {
	offUrl := s.URL + "/api/mef/productOfferingManagement/v1/productOffering"
	offapi := offer.NewAPIProductOfferingManagement(offUrl)

	reqParams := &offer.ProductOfferingGetParams{ProductOfferingID: params.ID}
	rspParams, err := offapi.ProductOfferingGet(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGenerateOfferGetResponse(reqParams)
	} else if err != nil {
		return nil
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.Data))
	params.RspOffer = rspParams.Data

	return nil
}
