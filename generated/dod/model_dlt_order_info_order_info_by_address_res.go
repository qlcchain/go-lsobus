/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderInfoOrderInfoByAddressRes struct {
	Code   float64                                `json:"code,omitempty"`
	Error_ *interface{}                           `json:"error,omitempty"`
	Data   *DltOrderInfoOrderInfoByAddressResData `json:"data,omitempty"`
	Meta   *interface{}                           `json:"meta,omitempty"`
}
