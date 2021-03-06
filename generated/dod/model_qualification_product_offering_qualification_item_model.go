/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type QualificationProductOfferingQualificationItemModel struct {
	Id                                           string                                                     `json:"id,omitempty"`
	Product                                      *QualificationProductOfferingQualificationItemModelProduct `json:"product,omitempty"`
	ProductOffering                              *OrderOrderItemResModelProductOffering                     `json:"productOffering,omitempty"`
	ProductOfferingQualificationItemRelationship []QualificationProductOfferingQualificationItemModel       `json:"productOfferingQualificationItemRelationship,omitempty"`
}
