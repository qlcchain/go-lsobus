package orchestra

import (
	"time"

	"github.com/iixlabs/virtual-lsobus/mock"

	poqcli "github.com/iixlabs/virtual-lsobus/sonata/poq/client"
	poqapi "github.com/iixlabs/virtual-lsobus/sonata/poq/client/product_offering_qualification"
	poqmod "github.com/iixlabs/virtual-lsobus/sonata/poq/models"
)

type sonataPOQImpl struct {
	sonataBaseImpl
}

func newSonataPOQImpl() *sonataPOQImpl {
	s := &sonataPOQImpl{}
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

	s.logger.Infof("send request, payload %s", s.DumpValue(reqParams.ProductOfferingQualification))

	rspParams, err := httpCli.ProductOfferingQualification.ProductOfferingQualificationCreate(reqParams)
	if err != nil {
		s.logger.Errorf("send request, error %s", err)
		//return err
		rspParams = mock.SonataGeneratePoqCreateResponse(reqParams)
	}
	s.logger.Infof("receive response, payload %s", s.DumpValue(rspParams.GetPayload()))

	orderParams.rspPoq = rspParams.GetPayload()

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
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())
	return nil
}

func (s *sonataPOQImpl) SendGetRequest(id string) error {
	reqParams := poqapi.NewProductOfferingQualificationGetParams()
	reqParams.ProductOfferingQualificationID = id

	httpCli := s.NewHTTPClient()

	rspParams, err := httpCli.ProductOfferingQualification.ProductOfferingQualificationGet(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return err
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())
	return nil
}

func (s *sonataPOQImpl) BuildCreateParams(orderParams *OrderParams) *poqapi.ProductOfferingQualificationCreateParams {
	reqParams := poqapi.NewProductOfferingQualificationCreateParams()

	reqParams.ProductOfferingQualification = new(poqmod.ProductOfferingQualificationCreate)
	reqParams.ProductOfferingQualification.ProjectID = orderParams.ProjectID

	isqVal := true
	reqParams.ProductOfferingQualification.InstantSyncQualification = &isqVal
	reqParams.ProductOfferingQualification.RequestedResponseDate.Scan(time.Now())

	// Source UNI
	srcUniItem := s.BuildUNIItem(orderParams, true)
	if srcUniItem != nil {
		reqParams.ProductOfferingQualification.ProductOfferingQualificationItem = append(reqParams.ProductOfferingQualification.ProductOfferingQualificationItem, srcUniItem)
	}

	// Destination UNI
	dstUniItem := s.BuildUNIItem(orderParams, false)
	if dstUniItem != nil {
		reqParams.ProductOfferingQualification.ProductOfferingQualificationItem = append(reqParams.ProductOfferingQualification.ProductOfferingQualificationItem, dstUniItem)
	}

	// ELine
	lineItem := s.BuildELineItem(orderParams)
	if lineItem != nil {
		reqParams.ProductOfferingQualification.ProductOfferingQualificationItem = append(reqParams.ProductOfferingQualification.ProductOfferingQualificationItem, lineItem)

		// Related Items
		if srcUniItem != nil {
			relItem := &poqmod.ProductOfferingQualificationItemRelationship{
				Type: poqmod.RelationshipTypeReliesOn,
				ID:   srcUniItem.ID,
			}
			lineItem.ProductOfferingQualificationItemRelationship = append(lineItem.ProductOfferingQualificationItemRelationship, relItem)
		}

		if dstUniItem != nil {
			relItem := &poqmod.ProductOfferingQualificationItemRelationship{
				Type: poqmod.RelationshipTypeReliesOn,
				ID:   dstUniItem.ID,
			}
			lineItem.ProductOfferingQualificationItemRelationship = append(lineItem.ProductOfferingQualificationItemRelationship, relItem)
		}

		// Related Products
		if orderParams.SrcPortID != "" {
			relProd := &poqmod.ProductRelationship{Type: poqmod.RelationshipTypeReliesOn}
			relProd.Product = &poqmod.ProductRef{}
			relProdID := orderParams.SrcPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}

		if orderParams.DstPortID != "" {
			relProd := &poqmod.ProductRelationship{Type: poqmod.RelationshipTypeReliesOn}
			relProd.Product = &poqmod.ProductRef{}
			relProdID := orderParams.DstPortID
			relProd.Product.ID = &relProdID
			lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, relProd)
		}
	}

	return reqParams
}

func (s *sonataPOQImpl) BuildUNIItem(orderParams *OrderParams, isDirSrc bool) *poqmod.ProductOfferingQualificationItemCreate {
	var siteID string
	if isDirSrc {
		siteID = orderParams.SrcSiteID
	} else {
		siteID = orderParams.DstSiteID
	}
	if siteID == "" {
		return nil
	}

	uniItem := &poqmod.ProductOfferingQualificationItemCreate{}

	uniItemID := s.NewItemID()
	uniItem.ID = &uniItemID
	uniItem.Action = poqmod.ProductActionType(orderParams.ItemAction)

	uniItem.ProductOffering = &poqmod.ProductOfferingRef{ID: MEFProductOfferingUNI}

	uniItem.Product = &poqmod.Product{}
	if uniItem.Action != poqmod.ProductActionTypeAdd {
		uniItem.Product.ID = orderParams.ProductID
	}

	// UNI Place
	uniPlace := &poqmod.ReferencedAddress{}
	uniPlace.ReferenceID = &siteID
	uniItem.Product.SetPlace([]poqmod.RelatedPlaceReforValue{uniPlace})

	// UNI Product Specification
	uniItem.Product.ProductSpecification = &poqmod.ProductSpecificationRef{}
	uniItem.Product.ProductSpecification.ID = "UNISpec"
	uniDesc := s.BuildUNIProductSpec(orderParams)
	uniItem.Product.ProductSpecification.SetDescribing(uniDesc)

	return uniItem
}

func (s *sonataPOQImpl) BuildELineItem(orderParams *OrderParams) *poqmod.ProductOfferingQualificationItemCreate {
	lineItem := &poqmod.ProductOfferingQualificationItemCreate{}

	lineItem.Action = poqmod.ProductActionType(orderParams.ItemAction)
	lineItemID := s.NewItemID()
	lineItem.ID = &lineItemID

	lineItem.ProductOffering = &poqmod.ProductOfferingRef{ID: MEFProductOfferingELine}

	// Product
	lineItem.Product = &poqmod.Product{}
	if lineItem.Action != poqmod.ProductActionTypeAdd {
		lineItem.Product.ID = orderParams.ProductID
	}

	//Product Specification
	lineItem.Product.ProductSpecification = &poqmod.ProductSpecificationRef{}
	lineItem.Product.ProductSpecification.ID = "ELineSpec"
	lineDesc := s.BuildELineProductSpec(orderParams)
	lineItem.Product.ProductSpecification.SetDescribing(lineDesc)

	return lineItem
}
