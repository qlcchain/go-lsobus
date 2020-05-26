package orchestra

import (
	invmod "github.com/qlcchain/go-lsobus/sonata/inventory/models"
	ordmod "github.com/qlcchain/go-lsobus/sonata/order/models"
	poqmod "github.com/qlcchain/go-lsobus/sonata/poq/models"
	quomod "github.com/qlcchain/go-lsobus/sonata/quote/models"
	sitmod "github.com/qlcchain/go-lsobus/sonata/site/models"
)

const (
	BillingTypePAYG  = "PAYG"
	BillingTypeDOD   = "DOD"
	BillingTypeUsage = "USAGE"
)

type Partner struct {
	ID   string
	Name string
}

type BillingParams struct {
	BillingType string
	BillingUnit string // used for PAYG, etc day/month/year
	MeasureUnit string // used for USAGE, etc minute/hour/Mbps/MByte
	StartTime   int64  // used for DOD Duration, unix seconds
	EndTime     int64  // used for DOD Duration, unix seconds
	Currency    string // etc USA/HKD/CNY
	Price       float32
}

type BaseItemParams struct {
	ItemID string
	Action string

	ProdSpecID  string
	ProdOfferID string
	ProdQuoteID string

	ProductID      string
	BuyerProductID string
	Description    string

	DurationUnit   string
	DurationAmount uint

	BillingParams *BillingParams
}

type UNIItemParams struct {
	BaseItemParams

	SiteID    string
	PortSpeed uint
}

type ELineItemParams struct {
	BaseItemParams

	SrcPortID string
	DstPortID string
	Bandwidth uint
	BwUnit    string
	SVlanID   uint
	CosName   string
}

type OrderParams struct {
	OrderActivity string

	ContractID string
	Buyer      *Partner
	Seller     *Partner

	Description string
	ProjectID   string
	ExternalID  string

	UNIItems   []*UNIItemParams
	ELineItems []*ELineItemParams

	RspPoq   *poqmod.ProductOfferingQualification
	RspQuote *quomod.Quote
	RspOrder *ordmod.ProductOrder
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
	ProductOfferingID      string
	ProductOrderID         string

	RspSiteList  []*sitmod.GeographicSiteFindResp
	RspPoqList   []*poqmod.ProductOfferingQualificationFind
	RspQuoteList []*quomod.QuoteFind
	RspOrderList []*ordmod.ProductOrderSummary
	RspInvList   []*invmod.ProductSummary
}

type GetParams struct {
	ID string

	RspSite  *sitmod.GeographicSite
	RspPoq   *poqmod.ProductOfferingQualification
	RspQuote *quomod.Quote
	RspOrder *ordmod.ProductOrder
	RspInv   *invmod.Product
}
