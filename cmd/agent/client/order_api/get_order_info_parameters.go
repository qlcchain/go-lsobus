// Code generated by go-swagger; DO NOT EDIT.

package order_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetOrderInfoParams creates a new GetOrderInfoParams object
// with the default values initialized.
func NewGetOrderInfoParams() *GetOrderInfoParams {
	var ()
	return &GetOrderInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrderInfoParamsWithTimeout creates a new GetOrderInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrderInfoParamsWithTimeout(timeout time.Duration) *GetOrderInfoParams {
	var ()
	return &GetOrderInfoParams{

		timeout: timeout,
	}
}

// NewGetOrderInfoParamsWithContext creates a new GetOrderInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrderInfoParamsWithContext(ctx context.Context) *GetOrderInfoParams {
	var ()
	return &GetOrderInfoParams{

		Context: ctx,
	}
}

// NewGetOrderInfoParamsWithHTTPClient creates a new GetOrderInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrderInfoParamsWithHTTPClient(client *http.Client) *GetOrderInfoParams {
	var ()
	return &GetOrderInfoParams{
		HTTPClient: client,
	}
}

/*GetOrderInfoParams contains all the parameters to send to the API endpoint
for the get order info operation typically these are written to a http.Request
*/
type GetOrderInfoParams struct {

	/*InternalID*/
	InternalID *string
	/*OrderID*/
	OrderID *string
	/*SellerAddress*/
	SellerAddress *string
	/*SellerName*/
	SellerName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get order info params
func (o *GetOrderInfoParams) WithTimeout(timeout time.Duration) *GetOrderInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get order info params
func (o *GetOrderInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get order info params
func (o *GetOrderInfoParams) WithContext(ctx context.Context) *GetOrderInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get order info params
func (o *GetOrderInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get order info params
func (o *GetOrderInfoParams) WithHTTPClient(client *http.Client) *GetOrderInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get order info params
func (o *GetOrderInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInternalID adds the internalID to the get order info params
func (o *GetOrderInfoParams) WithInternalID(internalID *string) *GetOrderInfoParams {
	o.SetInternalID(internalID)
	return o
}

// SetInternalID adds the internalId to the get order info params
func (o *GetOrderInfoParams) SetInternalID(internalID *string) {
	o.InternalID = internalID
}

// WithOrderID adds the orderID to the get order info params
func (o *GetOrderInfoParams) WithOrderID(orderID *string) *GetOrderInfoParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the get order info params
func (o *GetOrderInfoParams) SetOrderID(orderID *string) {
	o.OrderID = orderID
}

// WithSellerAddress adds the sellerAddress to the get order info params
func (o *GetOrderInfoParams) WithSellerAddress(sellerAddress *string) *GetOrderInfoParams {
	o.SetSellerAddress(sellerAddress)
	return o
}

// SetSellerAddress adds the sellerAddress to the get order info params
func (o *GetOrderInfoParams) SetSellerAddress(sellerAddress *string) {
	o.SellerAddress = sellerAddress
}

// WithSellerName adds the sellerName to the get order info params
func (o *GetOrderInfoParams) WithSellerName(sellerName *string) *GetOrderInfoParams {
	o.SetSellerName(sellerName)
	return o
}

// SetSellerName adds the sellerName to the get order info params
func (o *GetOrderInfoParams) SetSellerName(sellerName *string) {
	o.SellerName = sellerName
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrderInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InternalID != nil {

		// query param internalId
		var qrInternalID string
		if o.InternalID != nil {
			qrInternalID = *o.InternalID
		}
		qInternalID := qrInternalID
		if qInternalID != "" {
			if err := r.SetQueryParam("internalId", qInternalID); err != nil {
				return err
			}
		}

	}

	if o.OrderID != nil {

		// query param orderId
		var qrOrderID string
		if o.OrderID != nil {
			qrOrderID = *o.OrderID
		}
		qOrderID := qrOrderID
		if qOrderID != "" {
			if err := r.SetQueryParam("orderId", qOrderID); err != nil {
				return err
			}
		}

	}

	if o.SellerAddress != nil {

		// query param seller.address
		var qrSellerAddress string
		if o.SellerAddress != nil {
			qrSellerAddress = *o.SellerAddress
		}
		qSellerAddress := qrSellerAddress
		if qSellerAddress != "" {
			if err := r.SetQueryParam("seller.address", qSellerAddress); err != nil {
				return err
			}
		}

	}

	if o.SellerName != nil {

		// query param seller.name
		var qrSellerName string
		if o.SellerName != nil {
			qrSellerName = *o.SellerName
		}
		qSellerName := qrSellerName
		if qSellerName != "" {
			if err := r.SetQueryParam("seller.name", qSellerName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
