/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type QuoteQuoteItemModel struct {
	Id                     string                              `json:"id"`
	Action                 string                              `json:"action"`
	ProductOffering        *OrderOrderItemModelProductOffering `json:"productOffering"`
	RequestedQuoteItemTerm []QuoteRequestedQuoteItemTermModel  `json:"requestedQuoteItemTerm,omitempty"`
	Product                []QuoteProductModel                 `json:"product"`
	Qualification          []QuoteQualificationModel           `json:"qualification"`
	QuoteItemRelationship  []QuoteItemRelationshipModel        `json:"quoteItemRelationship"`
	Note                   []QuoteNoteModel                    `json:"note,omitempty"`
	RelatedParty           []QuoteRelatedPartyModel            `json:"relatedParty"`
}
