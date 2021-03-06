package pccwg

import (
	"time"

	"github.com/qlcchain/go-lsobus/api"
	"github.com/qlcchain/go-lsobus/mock"

	poqcli "github.com/qlcchain/go-lsobus/orchestra/sonata/poq/client"
	poqapi "github.com/qlcchain/go-lsobus/orchestra/sonata/poq/client/product_offering_qualification"
	poqmod "github.com/qlcchain/go-lsobus/orchestra/sonata/poq/models"
)

type sonataPOQImpl struct {
	sonataBaseImpl
}

func newSonataPOQImpl(p api.DoDSeller) *sonataPOQImpl {
	s := &sonataPOQImpl{}
	s.Partner = p
	s.Version = MEFAPIVersionPOQ
	return s
}

func (s *sonataPOQImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataPOQImpl) NewHTTPClient() *poqcli.APIProductOfferingQualificationManagement {
	httpTran := s.NewHttpTransport(poqcli.DefaultBasePath)
	httpCli := poqcli.New(httpTran, nil)
	return httpCli
}

func (s *sonataPOQImpl) SendCreateRequest(orderParams *api.OrderParams) error {
	reqParams := s.BuildCreateParams(orderParams)

	httpCli := s.NewHTTPClient()

	s.logger.Debugf("send request, payload %s", s.DumpValue(reqParams.ProductOfferingQualification))

	rspParams, err := httpCli.ProductOfferingQualification.ProductOfferingQualificationCreate(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGeneratePoqCreateResponse(reqParams)
	} else if err != nil {
		s.logger.Errorf("send request, error %s", err)
		s.handleResponseError(err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	orderParams.RspPoq = rspParams.GetPayload()

	return nil
}

func (s *sonataPOQImpl) SendFindRequest(params *api.FindParams) error {
	reqParams := poqapi.NewProductOfferingQualificationFindParams()
	if params.ProjectID != "" {
		reqParams.ProjectID = &params.ProjectID
	}
	if params.State != "" {
		reqParams.State = &params.State
	}
	if params.Offset != "" {
		reqParams.Offset = &params.Offset
	}
	if params.Limit != "" {
		reqParams.Limit = &params.Limit
	}

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.ProductOfferingQualification.ProductOfferingQualificationFind(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGeneratePoqFindResponse(reqParams)
	} else if err != nil {
		s.logger.Error("send request,", "error:", err)
		s.handleResponseError(err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspPoqList = rspParams.GetPayload()
	params.XResultCount = rspParams.XResultCount
	params.XTotalCount = rspParams.XTotalCount

	return nil
}

func (s *sonataPOQImpl) SendGetRequest(params *api.GetParams) error {
	reqParams := poqapi.NewProductOfferingQualificationGetParams()
	reqParams.ProductOfferingQualificationID = params.ID

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.ProductOfferingQualification.ProductOfferingQualificationGet(reqParams)
	if s.GetFakeMode() {
		rspParams = mock.SonataGeneratePoqGetResponse(reqParams)
	} else if err != nil {
		s.logger.Error("send request,", "error:", err)
		s.handleResponseError(err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspPoq = rspParams.GetPayload()
	return nil
}

func (s *sonataPOQImpl) BuildCreateParams(orderParams *api.OrderParams) *poqapi.ProductOfferingQualificationCreateParams {
	reqParams := poqapi.NewProductOfferingQualificationCreateParams()

	reqParams.ProductOfferingQualification = new(poqmod.ProductOfferingQualificationCreate)
	reqParams.ProductOfferingQualification.ProjectID = orderParams.ProjectID

	isqVal := true
	reqParams.ProductOfferingQualification.InstantSyncQualification = &isqVal
	reqParams.ProductOfferingQualification.RequestedResponseDate.Scan(time.Now())

	// UNI
	var allUniItems []*poqmod.ProductOfferingQualificationItemCreate
	for _, uniParams := range orderParams.UNIItems {
		uniItem := s.BuildUNIItem(uniParams)
		if uniItem == nil {
			continue
		}
		reqParams.ProductOfferingQualification.ProductOfferingQualificationItem = append(reqParams.ProductOfferingQualification.ProductOfferingQualificationItem, uniItem)
		allUniItems = append(allUniItems, uniItem)
	}

	// ELine
	var allLineItems []*poqmod.ProductOfferingQualificationItemCreate
	for _, lineParams := range orderParams.ELineItems {
		lineItem := s.BuildELineItem(lineParams)
		if lineItem == nil {
			continue
		}
		reqParams.ProductOfferingQualification.ProductOfferingQualificationItem = append(reqParams.ProductOfferingQualification.ProductOfferingQualificationItem, lineItem)
		allLineItems = append(allLineItems, lineItem)

		// Related Products
		if lineParams.SrcPortID != "" {
			relProd := &poqmod.ProductRelationship{Type: poqmod.RelationshipTypeReliesOn}
			relProd.Product = &poqmod.ProductRef{}
			relProdID := lineParams.SrcPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}

		if lineParams.DstPortID != "" {
			relProd := &poqmod.ProductRelationship{Type: poqmod.RelationshipTypeReliesOn}
			relProd.Product = &poqmod.ProductRef{}
			relProdID := lineParams.DstPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}
	}

	// Related Items
	if len(allLineItems) == 1 {
		lineItem := allLineItems[0]
		for _, uniItem := range allUniItems {
			relItem := &poqmod.ProductOfferingQualificationItemRelationship{
				Type: poqmod.RelationshipTypeReliesOn,
				ID:   uniItem.ID,
			}
			lineItem.ProductOfferingQualificationItemRelationship = append(lineItem.ProductOfferingQualificationItemRelationship, relItem)
		}
	}

	return reqParams
}

func (s *sonataPOQImpl) BuildUNIItem(params *api.UNIItemParams) *poqmod.ProductOfferingQualificationItemCreate {
	if params.ProdSpecID != "" && params.ProdSpecID != "UNISpec" {
		return nil
	}

	uniItem := &poqmod.ProductOfferingQualificationItemCreate{}

	if params.ItemID != "" {
		uniItem.ID = &params.ItemID
	} else {
		uniItemID := s.NewItemID()
		uniItem.ID = &uniItemID
	}

	uniItem.Action = poqmod.ProductActionType(params.Action)

	uniItem.ProductOffering = &poqmod.ProductOfferingRef{ID: params.ProdOfferID}

	uniItem.Product = &poqmod.Product{}
	if uniItem.Action != poqmod.ProductActionTypeAdd {
		uniItem.Product.ID = params.ProductID
	}

	// UNI Place
	if params.SiteID != "" {
		uniPlace := &poqmod.ReferencedAddress{}
		uniPlace.ReferenceID = &params.SiteID
		uniItem.Product.SetPlace([]poqmod.RelatedPlaceReforValue{uniPlace})
	}

	// UNI Product Specification
	if uniItem.Action != poqmod.ProductActionTypeRemove {
		uniItem.Product.ProductSpecification = &poqmod.ProductSpecificationRef{}
		uniItem.Product.ProductSpecification.ID = "UNISpec"
		uniDesc := s.BuildUNIProductSpec(params)
		uniItem.Product.ProductSpecification.SetDescribing(uniDesc)
	}

	return uniItem
}

func (s *sonataPOQImpl) BuildELineItem(params *api.ELineItemParams) *poqmod.ProductOfferingQualificationItemCreate {
	if params.ProdSpecID != "" && params.ProdSpecID != "ELineSpec" {
		return nil
	}

	if params.Action != string(poqmod.ProductActionTypeRemove) {
		if params.Bandwidth == 0 {
			return nil
		}
	}

	lineItem := &poqmod.ProductOfferingQualificationItemCreate{}

	lineItem.Action = poqmod.ProductActionType(params.Action)
	if params.ItemID != "" {
		lineItem.ID = &params.ItemID
	} else {
		lineItemID := s.NewItemID()
		lineItem.ID = &lineItemID
	}

	lineItem.ProductOffering = &poqmod.ProductOfferingRef{ID: params.ProdOfferID}

	// Product
	lineItem.Product = &poqmod.Product{}
	if lineItem.Action != poqmod.ProductActionTypeAdd {
		lineItem.Product.ID = params.ProductID
	}

	//Product Specification
	if params.Action != string(poqmod.ProductActionTypeRemove) {
		lineItem.Product.ProductSpecification = &poqmod.ProductSpecificationRef{}
		lineItem.Product.ProductSpecification.ID = "ELineSpec"
		lineDesc := s.BuildPCCWConnProductSpec(params)
		lineItem.Product.ProductSpecification.SetDescribing(lineDesc)
	}

	return lineItem
}

func (s *sonataPOQImpl) handleResponseError(rspErr error) {
	switch rspErr.(type) {
	case *poqapi.ProductOfferingQualificationCreateUnauthorized, *poqapi.ProductOfferingQualificationCreateForbidden:
		s.ClearApiToken()
	case *poqapi.ProductOfferingQualificationFindUnauthorized, *poqapi.ProductOfferingQualificationFindForbidden:
		s.ClearApiToken()
	case *poqapi.ProductOfferingQualificationGetUnauthorized, *poqapi.ProductOfferingQualificationGetForbidden:
		s.ClearApiToken()
	}
}
