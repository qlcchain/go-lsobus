/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltInvoiceGenerateInvoiceByOrderIdResOrder struct {
	OrderId         string                                     `json:"orderId,omitempty"`
	InternalId      string                                     `json:"internalId,omitempty"`
	ConnectionCount float64                                    `json:"connectionCount,omitempty"`
	OrderAmount     float64                                    `json:"orderAmount,omitempty"`
	Connections     []DltInvoiceGenerateInvoiceConnectionModel `json:"connections,omitempty"`
}
