package orchestra

import (
	ordmod "github.com/qlcchain/go-lsobus/sonata/order/models"
	poqmod "github.com/qlcchain/go-lsobus/sonata/poq/models"
	quomod "github.com/qlcchain/go-lsobus/sonata/quote/models"
)

type Partner struct {
	ID   string
	Name string
}

type OrderParams struct {
	OrderActivity string
	ItemAction    string
	ProductID     string

	ContractID string
	Buyer      *Partner
	Seller     *Partner

	ExternalID  string
	Description string
	ProjectID   string

	SrcSiteID    string
	SrcPortSpeed uint
	DstSiteID    string
	DstPortSpeed uint

	SrcPortID string
	SrcVlanID []uint
	DstPortID string
	DstVlanID []uint

	Bandwidth uint
	SVlanID   uint
	CosName   string

	rspPoq   *poqmod.ProductOfferingQualification
	rspQuote *quomod.Quote
	rspOrder *ordmod.ProductOrder
}

type FindParams struct {
	Offset     string
	Limit      string
	ProjectID  string
	ExternalID string
	BuyerID    string
	SiteID     string
	State      string

	ProductSpecificationID string
}

type GetParams struct {
	ID string
}
