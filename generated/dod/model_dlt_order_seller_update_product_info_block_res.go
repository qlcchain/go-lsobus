/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderSellerUpdateProductInfoBlockRes struct {
	Jsonrpc string                                `json:"jsonrpc"`
	ID      string                                `json:"id"`
	Result  *DltOrderBuyerCreateOrderBlockResData `json:"result,omitempty"`
}
