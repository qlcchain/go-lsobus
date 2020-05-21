package orchestra

import (
	"time"

	"github.com/qlcchain/go-lsobus/mock"

	poqcli "github.com/qlcchain/go-lsobus/sonata/poq/client"
	poqapi "github.com/qlcchain/go-lsobus/sonata/poq/client/product_offering_qualification"
	poqmod "github.com/qlcchain/go-lsobus/sonata/poq/models"
)

type sonataPOQImpl struct {
	sonataBaseImpl
}

func newSonataPOQImpl(o *Orchestra) *sonataPOQImpl {
	s := &sonataPOQImpl{}
	s.Orch = o
	s.Version = MEFAPIVersionPOQ
	return s
}

func (s *sonataPOQImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataPOQImpl) NewHTTPClient() *poqcli.APIProductOfferingQualificationManagement {
	tranCfg := poqcli.DefaultTransportConfig().WithHost(s.GetHost()).WithSchemes([]string{s.GetScheme()})
	httpCli := poqcli.NewHTTPClientWithConfig(nil, tranCfg)
	return httpCli
}

func (s *sonataPOQImpl) SendCreateRequest(orderParams *OrderParams) error {
	reqParams := s.BuildCreateParams(orderParams)

	httpCli := s.NewHTTPClient()

	s.logger.Debugf("send request, payload %s", s.DumpValue(reqParams.ProductOfferingQualification))

	rspParams, err := httpCli.ProductOfferingQualification.ProductOfferingQualificationCreate(reqParams)
	if err != nil {
		s.logger.Errorf("send request, error %s", err)
		//return err
		rspParams = mock.SonataGeneratePoqCreateResponse(reqParams)
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	orderParams.RspPoq = rspParams.GetPayload()

	return nil
}

func (s *sonataPOQImpl) SendFindRequest(params *FindParams) error {
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
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspPoqList = rspParams.GetPayload()
	return nil
}

func (s *sonataPOQImpl) SendGetRequest(params *GetParams) error {
	reqParams := poqapi.NewProductOfferingQualificationGetParams()
	reqParams.ProductOfferingQualificationID = params.ID

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.ProductOfferingQualification.ProductOfferingQualificationGet(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Debugf("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))
	params.RspPoq = rspParams.GetPayload()
	return nil
}

func (s *sonataPOQImpl) BuildCreateParams(orderParams *OrderParams) *poqapi.ProductOfferingQualificationCreateParams {
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

func (s *sonataPOQImpl) BuildUNIItem(params *UNIItemParams) *poqmod.ProductOfferingQualificationItemCreate {
	if params.ProdSpecID != "" && params.ProdSpecID != "UNISpec" {
		return nil
	}

	uniItem := &poqmod.ProductOfferingQualificationItemCreate{}

	uniItemID := s.NewItemID()
	uniItem.ID = &uniItemID
	uniItem.Action = poqmod.ProductActionType(params.Action)

	uniItem.ProductOffering = &poqmod.ProductOfferingRef{ID: MEFProductOfferingUNI}

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

func (s *sonataPOQImpl) BuildELineItem(params *ELineItemParams) *poqmod.ProductOfferingQualificationItemCreate {
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
	lineItemID := s.NewItemID()
	lineItem.ID = &lineItemID

	lineItem.ProductOffering = &poqmod.ProductOfferingRef{ID: MEFProductOfferingELine}

	// Product
	lineItem.Product = &poqmod.Product{}
	if lineItem.Action != poqmod.ProductActionTypeAdd {
		lineItem.Product.ID = params.ProductID
	}

	//Product Specification
	if params.Action != string(poqmod.ProductActionTypeRemove) {
		lineItem.Product.ProductSpecification = &poqmod.ProductSpecificationRef{}
		lineItem.Product.ProductSpecification.ID = "ELineSpec"
		lineDesc := s.BuildELineProductSpec(params)
		lineItem.Product.ProductSpecification.SetDescribing(lineDesc)
	}

	return lineItem
}
