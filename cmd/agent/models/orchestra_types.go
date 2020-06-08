package models

const (
	BillingTypePAYG  = "PAYG"
	BillingTypeDOD   = "DOD"
	BillingTypeUsage = "USAGE"
)

type PartnerParams struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BillingParams struct {
	PaymentType string  `json:"paymentType,omitempty"`
	BillingType string  `json:"billingType,omitempty"`
	BillingUnit string  `json:"billingUnit,omitempty"` // used for PAYG, etc day/month/year
	MeasureUnit string  `json:"measureUnit,omitempty"` // used for USAGE, etc minute/hour/Mbps/MByte
	StartTime   int64   `json:"startTime,omitempty"`   // used for DOD Duration, unix seconds
	EndTime     int64   `json:"endTime,omitempty"`     // used for DOD Duration, unix seconds
	Currency    string  `json:"currency,omitempty"`    // etc USA/HKD/CNY
	Price       float32 `json:"price,omitempty"`
}

type BaseItemParams struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	ItemID string `json:"itemID,omitempty"`
	Action string `json:"action,omitempty"`

	ProdSpecID  string `json:"prodSpecID,omitempty"`
	ProdOfferID string `json:"prodOfferID,omitempty"`

	ProductID      string `json:"productID,omitempty"`
	BuyerProductID string `json:"buyerProductID,omitempty"`
	QuoteID        string `json:"quoteID,omitempty"`
	QuoteItemID    string `json:"quoteItemID,omitempty"`

	BillingParams *BillingParams `json:"billingParams,omitempty"`
}

type UNIItemParams struct {
	BaseItemParams

	SiteID    string `json:"siteID,omitempty"`
	PortSpeed uint   `json:"portSpeed,omitempty"`
}

type ELineItemParams struct {
	BaseItemParams

	SrcPortID    string `json:"srcPortID,omitempty"`
	DstPortID    string `json:"dstPortID,omitempty"`
	DstCompanyID string `json:"dstCompanyID,omitempty"`
	DstMetroID   string `json:"dstMetroID,omitempty"`
	Bandwidth    uint   `json:"bandwidth,omitempty"`
	BwUnit       string `json:"bwUnit,omitempty"`
	SVlanID      uint   `json:"sVlanID,omitempty"`
	CosName      string `json:"cosName,omitempty"`

	SrcLocationID string `json:"srcLocationID,omitempty"`
	DstLocationID string `json:"dstLocationID,omitempty"`
}

type OrderParams struct {
	Buyer  *PartnerParams `json:"buyer,omitempty"`
	Seller *PartnerParams `json:"seller,omitempty"`

	OrderActivity string `json:"orderActivity,omitempty"`

	Description string `json:"description,omitempty"`
	ProjectID   string `json:"projectID,omitempty"`
	ExternalID  string `json:"externalID,omitempty"`

	UNIItems   []*UNIItemParams   `json:"uniItems,omitempty"`
	ELineItems []*ELineItemParams `json:"elineItems,omitempty"`

	BillingType string `json:"billingType,omitempty"`
	PaymentType string `json:"paymentType,omitempty"`
}

type FindParams struct {
	Buyer  *PartnerParams `json:"buyer,omitempty"`
	Seller *PartnerParams `json:"seller,omitempty"`

	Offset     string `json:"offset,omitempty"`
	Limit      string `json:"limit,omitempty"`
	ProjectID  string `json:"projectID,omitempty"`
	ExternalID string `json:"externalID,omitempty"`
	BuyerID    string `json:"buyerID,omitempty"`
	SiteID     string `json:"siteID,omitempty"`
	State      string `json:"state,omitempty"`

	ProductSpecificationID string `json:"productSpecificationID,omitempty"`
	ProductOfferingID      string `json:"productOfferingID,omitempty"`
	ProductOrderID         string `json:"productOrderID,omitempty"`
}

type GetParams struct {
	Buyer  *PartnerParams `json:"buyer,omitempty"`
	Seller *PartnerParams `json:"seller,omitempty"`

	ID string `json:"id,omitempty"`
}
