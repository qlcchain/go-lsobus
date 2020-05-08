package orchestra

import (
	"fmt"
	"strconv"
	"time"

	cmnmod "github.com/iixlabs/virtual-lsobus/sonata/common/models"
	poqapi "github.com/iixlabs/virtual-lsobus/sonata/poq/client"
	poqcli "github.com/iixlabs/virtual-lsobus/sonata/poq/client/product_offering_qualification"
	poqmod "github.com/iixlabs/virtual-lsobus/sonata/poq/models"
)

func SendSonataPOQCreateRequest() {
	reqParams := BuildSonataPOQCreateParams()

	tranCfg := poqapi.DefaultTransportConfig().WithHost("localhost").WithSchemes([]string{"http"})
	poqCli := poqapi.NewHTTPClientWithConfig(nil, tranCfg)

	rspParams, err := poqCli.ProductOfferingQualification.ProductOfferingQualificationCreate(reqParams)
	if err != nil {
		fmt.Println("send request,", "error:", err)
		return
	}
	fmt.Println("receive response,", "error:", rspParams.Error(), "Payload:", rspParams.GetPayload())
}

func BuildSonataPOQCreateParams() *poqcli.ProductOfferingQualificationCreateParams {
	reqParams := poqcli.NewProductOfferingQualificationCreateParams()
	poqCreate := new(poqmod.ProductOfferingQualificationCreate)
	isqVal := true
	poqCreate.ProjectID = "DoD-CBC-PCCW"
	poqCreate.InstantSyncQualification = &isqVal
	poqCreate.RequestedResponseDate.Scan(time.Now())
	reqParams.SetProductOfferingQualification(poqCreate)

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
	poqCreate.ProductOfferingQualificationItem = append(poqCreate.ProductOfferingQualificationItem, uniItem)

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

	poqCreate.ProductOfferingQualificationItem = append(poqCreate.ProductOfferingQualificationItem, lineItem)

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
