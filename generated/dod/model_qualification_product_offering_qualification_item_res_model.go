/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type QualificationProductOfferingQualificationItemResModel struct {
	Id                       string                                                                     `json:"id,omitempty"`
	State                    string                                                                     `json:"state,omitempty"`
	ServiceabilityConfidence string                                                                     `json:"serviceabilityConfidence,omitempty"`
	ServiceConfidenceReason  string                                                                     `json:"serviceConfidenceReason,omitempty"`
	InstallationInterval     *QualificationProductOfferingQualificationItemResModelInstallationInterval `json:"installationInterval,omitempty"`
	GuaranteedUntilDate      string                                                                     `json:"guaranteedUntilDate,omitempty"`
	Product                  *QualificationProductOfferingQualificationItemResModelProduct              `json:"product,omitempty"`
}
