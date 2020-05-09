package orchestra

import (
	"time"

	cmnmod "github.com/iixlabs/virtual-lsobus/sonata/common/models"
	poqcli "github.com/iixlabs/virtual-lsobus/sonata/poq/client"
	poqapi "github.com/iixlabs/virtual-lsobus/sonata/poq/client/product_offering_qualification"
	poqmod "github.com/iixlabs/virtual-lsobus/sonata/poq/models"
)

type sonataPOQImpl struct {
	sonataBaseImpl
}

func newSonataPOQImpl() *sonataPOQImpl {
	s := &sonataPOQImpl{}
	return s
}

func (s *sonataPOQImpl) Init() error {
	return s.sonataBaseImpl.Init()
}

func (s *sonataPOQImpl) SendCreateRequest(orderParams *OrderParams) error {
	reqParams := s.BuildCreateParams(orderParams)

	tranCfg := poqcli.DefaultTransportConfig().WithHost("localhost").WithSchemes([]string{"http"})
	poqCli := poqcli.NewHTTPClientWithConfig(nil, tranCfg)

	rspParams, err := poqCli.ProductOfferingQualification.ProductOfferingQualificationCreate(reqParams)
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
	uniItem.Action = poqmod.ProductActionTypeAdd

	uniItem.ProductOffering = &poqmod.ProductOfferingRef{ID: "LSO_Sonata_ProviderOnDemand_EthernetPort_UNI"}

	uniItem.Product = &poqmod.Product{}

	// UNI Place
	uniPlace := &poqmod.ReferencedAddress{}
	uniPlace.ReferenceID = &siteID
	uniItem.Product.SetPlace([]poqmod.RelatedPlaceReforValue{uniPlace})

	// UNI Product Specification
	uniItem.Product.ProductSpecification = &poqmod.ProductSpecificationRef{}
	uniItem.Product.ProductSpecification.ID = "UNISpec"
	uniDesc := &cmnmod.UNIProductSpecification{}
	s.FillUNIProductSpec(uniDesc, orderParams)
	uniItem.Product.ProductSpecification.SetDescribing(uniDesc)

	return uniItem
}

func (s *sonataPOQImpl) BuildELineItem(orderParams *OrderParams) *poqmod.ProductOfferingQualificationItemCreate {
	lineItem := &poqmod.ProductOfferingQualificationItemCreate{}

	lineItem.Action = poqmod.ProductActionTypeAdd
	lineItemID := s.NewItemID()
	lineItem.ID = &lineItemID
	lineItem.ProductOffering = &poqmod.ProductOfferingRef{ID: "LSO_Sonata_ProviderOnDemand_EthernetConnection"}
	lineItem.Product = &poqmod.Product{}

	//Product Specification
	lineItem.Product.ProductSpecification = &poqmod.ProductSpecificationRef{}
	lineItem.Product.ProductSpecification.ID = "ELineSpec"
	lineDesc := &cmnmod.ELineProductSpecification{}
	s.FillELineProductSpec(lineDesc, orderParams)
	lineItem.Product.ProductSpecification.SetDescribing(lineDesc)

	return lineItem
}
