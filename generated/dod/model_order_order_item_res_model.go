/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OrderOrderItemResModel struct {
	Id              string                                 `json:"id,omitempty"`
	Action          string                                 `json:"action,omitempty"`
	ProductOffering *OrderOrderItemResModelProductOffering `json:"productOffering,omitempty"`
	RelatedParty    []OrderRelatedPartyModel               `json:"relatedParty,omitempty"`
	Product         []OrderProductModel                    `json:"product,omitempty"`
	Qualification   []OrderQualificationModel              `json:"qualification,omitempty"`
	Quote           *OrderOrderItemResModelQuote           `json:"quote,omitempty"`
}
