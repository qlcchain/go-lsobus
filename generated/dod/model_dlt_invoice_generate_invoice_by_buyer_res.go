/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltInvoiceGenerateInvoiceByBuyerRes struct {
	InvoiceId            string                                      `json:"invoiceId,omitempty"`
	OrderCount           float64                                     `json:"orderCount,omitempty"`
	TotalConnectionCount float64                                     `json:"totalConnectionCount,omitempty"`
	TotalAmount          float64                                     `json:"totalAmount,omitempty"`
	Currency             string                                      `json:"currency,omitempty"`
	StartTime            float64                                     `json:"startTime,omitempty"`
	EndTime              float64                                     `json:"endTime,omitempty"`
	Buyer                *DltInvoiceGenerateInvoiceByOrderIdResBuyer `json:"buyer,omitempty"`
	Orders               []DltInvoiceGenerateInvoiceOrderModel       `json:"orders,omitempty"`
}
