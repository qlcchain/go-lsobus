/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import qlcSdk "github.com/qlcchain/qlc-go-sdk"

type DltOrderSellerPendingRequestRes struct {
	Jsonrpc string                         `json:"jsonrpc"`
	ID      string                         `json:"id"`
	Result  []*qlcSdk.DoDPendingRequestRsp `json:"result"`
}
