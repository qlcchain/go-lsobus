/*
 * QLC DOD Service
 *
 * REST Api for QLC DOD Service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DltOrderSellerPendingRequestResDataOrder struct {
	Buyer         *DltInvoiceGenerateInvoiceByOrderIdResBuyer   `json:"buyer,omitempty"`
	Seller        *DltInvoiceGenerateInvoiceByOrderIdResBuyer   `json:"seller,omitempty"`
	OrderId       string                                        `json:"orderId,omitempty"`
	OrderType     string                                        `json:"orderType,omitempty"`
	OrderState    string                                        `json:"orderState,omitempty"`
	ContractState string                                        `json:"contractState,omitempty"`
	Connections   []DltOrderSellerPendingRequestConnectionModel `json:"connections,omitempty"`
	Track         []DltOrderSellerPendingRequestTrackModel      `json:"track,omitempty"`
}
