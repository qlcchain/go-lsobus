package orchestra

import (
	"github.com/qlcchain/go-lsobus/sonata/offer"
)

type sonataOfferImpl struct {
	sonataBaseImpl
}

func newSonataOfferImpl(o *Orchestra) *sonataOfferImpl {
	s := &sonataOfferImpl{}
	s.Orch = o
	s.Version = MEFAPIVersionOffer
	return s
}

func (s *sonataOfferImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataOfferImpl) SendFindRequest(params *FindParams) error {
	offUrl := s.URL + "/api/mef/productOfferingManagement/v1/productOffering"
	offapi := offer.NewAPIProductOfferingManagement(offUrl)
	rspParams, err := offapi.ProductOfferingFind(nil)
	if err != nil {
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
	if err != nil {
		return nil
	}

	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.Data))
	params.RspOffer = rspParams.Data

	return nil
}
