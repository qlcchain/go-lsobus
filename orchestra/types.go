package orchestra

type Partner struct {
	ID   string
	Name string
}

type OrderParams struct {
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
