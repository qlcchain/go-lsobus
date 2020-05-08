package orchestra

import (
	"strconv"
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

func (s *sonataPOQImpl) SendCreateRequest() {
	reqParams := s.BuildCreateParams()

	tranCfg := poqcli.DefaultTransportConfig().WithHost("localhost").WithSchemes([]string{"http"})
	poqCli := poqcli.NewHTTPClientWithConfig(nil, tranCfg)

	rspParams, err := poqCli.ProductOfferingQualification.ProductOfferingQualificationCreate(reqParams)
	if err != nil {
		s.logger.Error("send request,", "error:", err)
		return
	}
	s.logger.Info("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())
}

func (s *sonataPOQImpl) BuildCreateParams() *poqapi.ProductOfferingQualificationCreateParams {
	reqParams := poqapi.NewProductOfferingQualificationCreateParams()

	reqParams.ProductOfferingQualification = new(poqmod.ProductOfferingQualificationCreate)
	isqVal := true
	reqParams.ProductOfferingQualification.ProjectID = "DoD-CBC-PCCW"
	reqParams.ProductOfferingQualification.InstantSyncQualification = &isqVal
	reqParams.ProductOfferingQualification.RequestedResponseDate.Scan(time.Now())

	itemIdSeq := int(0)

	// UNI
	itemIdSeq++
	uniItem := new(poqmod.ProductOfferingQualificationItemCreate)
	uniItemID := strconv.Itoa(itemIdSeq)
	uniItem.ID = &uniItemID
	uniItem.Action = "INSTALL"
	uniItem.ProductOffering = &poqmod.ProductOfferingRef{ID: "LSO_Sonata_ProviderOnDemand_EthernetPort_UNI"}
	uniItem.Product = new(poqmod.Product)

	// UNI Place
	uniPlace := &poqmod.ReferencedAddress{}
	addrId := "1111-2222-3333"
	uniPlace.ReferenceID = &addrId
	uniItem.Product.SetPlace([]poqmod.RelatedPlaceReforValue{uniPlace})

	// UNI Product Specification
	uniItem.Product.ProductSpecification = new(poqmod.ProductSpecificationRef)
	uniItem.Product.ProductSpecification.ID = "UNISpec"
	uniItem.Product.ProductSpecification.Describing = new(poqmod.Describing)
	uniItem.Product.ProductSpecification.Describing.AtSchemaLocation = "https://github.com/MEF-GIT/MEF-LSO-Sonata-SDK/blob/working-draft/payload_descriptions/ProductSpecDescription/MEF_UNISpec_v3.json"
	uniItem.Product.ProductSpecification.Describing.AtType = "UNISpec"
	uniItem.Product.ProductSpecification.Describing.MEFUNISpecV3 = new(cmnmod.MEFUNISpecV3)
	uniItem.Product.ProductSpecification.Describing.MEFUNISpecV3.PhysicalLayer = []cmnmod.PhysicalLayer{cmnmod.PhysicalLayerNr1000BASET}
	uniItem.Product.ProductSpecification.Describing.MEFUNISpecV3.MaxServiceFrameSize = 1522
	uniItem.Product.ProductSpecification.Describing.MEFUNISpecV3.NumberOfLinks = 1
	reqParams.ProductOfferingQualification.ProductOfferingQualificationItem = append(reqParams.ProductOfferingQualification.ProductOfferingQualificationItem, uniItem)

	// ELine
	itemIdSeq++
	lineItem := new(poqmod.ProductOfferingQualificationItemCreate)
	lineItem.Action = "INSTALL"
	lineItemID := strconv.Itoa(itemIdSeq)
	lineItem.ID = &lineItemID
	lineItem.ProductOffering = &poqmod.ProductOfferingRef{ID: "LSO_Sonata_ProviderOnDemand_EthernetConnection"}
	lineItem.Product = new(poqmod.Product)

	//Product Specification
	lineItem.Product.ProductSpecification = new(poqmod.ProductSpecificationRef)
	lineItem.Product.ProductSpecification.ID = "ELineSpec"
	lineItem.Product.ProductSpecification.Describing = new(poqmod.Describing)
	lineItem.Product.ProductSpecification.Describing.AtSchemaLocation = "https://github.com/MEF-GIT/MEF-LSO-Sonata-SDK/blob/working-draft/payload_descriptions/ProductSpecDescription/MEF_ELineSpec_v3.json"
	lineItem.Product.ProductSpecification.Describing.AtType = "ELineSpec"
	lineItem.Product.ProductSpecification.Describing.MEFELineSpecV3 = new(cmnmod.MEFELineSpecV3)
	lineItem.Product.ProductSpecification.Describing.MEFELineSpecV3.ClassOfServiceName = "Gold"
	lineItem.Product.ProductSpecification.Describing.MEFELineSpecV3.MaximumFrameSize = 1526
	lineItem.Product.ProductSpecification.Describing.MEFELineSpecV3.SVlanID = 101
	bwMbps := int32(10)
	bwProfile := &cmnmod.BandwidthProfile{
		Cir: &cmnmod.InformationRate{Unit: "Mbps", Amount: &bwMbps},
	}
	lineItem.Product.ProductSpecification.Describing.MEFELineSpecV3.ENNIIngressBWProfile = []*cmnmod.BandwidthProfile{bwProfile}
	lineItem.Product.ProductSpecification.Describing.MEFELineSpecV3.UNIIngressBWProfile = []*cmnmod.BandwidthProfile{bwProfile}

	// Related Products
	lineItemRelOnEnni := &poqmod.ProductRelationship{Type: poqmod.RelationshipTypeReliesOn}
	lineItemRelOnEnni.Product = &poqmod.ProductRef{}
	enniId := "8888"
	lineItemRelOnEnni.Product.ID = &enniId
	lineItem.Product.ProductRelationship = append(lineItem.Product.ProductRelationship, lineItemRelOnEnni)

	lineItemRelOnUni := &poqmod.ProductOfferingQualificationItemRelationship{Type: poqmod.RelationshipTypeReliesOn}
	lineItemRelOnUni.ID = &uniItemID
	lineItem.ProductOfferingQualificationItemRelationship = append(lineItem.ProductOfferingQualificationItemRelationship, lineItemRelOnUni)

	reqParams.ProductOfferingQualification.ProductOfferingQualificationItem = append(reqParams.ProductOfferingQualification.ProductOfferingQualificationItem, lineItem)

	// Related Parties
	/*
		partyCBC := &poqmod.RelatedParty{}
		partyCBC.Role = []string{"Buyer"}
		partyCBC.ID = "Partner111"
		cbcName := "CBC"
		partyCBC.Name = &cbcName
		cbcNumber := "12345678"
		partyCBC.Number = &cbcNumber
		poqCreate.RelatedParty = append(poqCreate.RelatedParty, partyCBC)*/

	return reqParams
}
