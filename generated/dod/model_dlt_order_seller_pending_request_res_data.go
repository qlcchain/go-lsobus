/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderSellerPendingRequestResData struct {
	Hash  string                                    `json:"hash,omitempty"`
	Order *DltOrderSellerPendingRequestResDataOrder `json:"order,omitempty"`
}
