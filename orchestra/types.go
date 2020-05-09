package orchestra

type OrderParams struct {
	SrcSiteID    string
	SrcPortSpeed uint
	DstSiteID    string
	DstPortSpeed uint

	SrcPortID string
	DstPortID string
	Bandwidth uint
	SVlanID   uint
	CosName   string
}
