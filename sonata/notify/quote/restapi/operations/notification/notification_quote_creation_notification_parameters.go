// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-virtual-lsobus/sonata/notify/quote/models"
)

// NewNotificationQuoteCreationNotificationParams creates a new NotificationQuoteCreationNotificationParams object
// no default values defined in spec.
func NewNotificationQuoteCreationNotificationParams() NotificationQuoteCreationNotificationParams {

	return NotificationQuoteCreationNotificationParams{}
}

// NotificationQuoteCreationNotificationParams contains all the bound params for the notification quote creation notification operation
// typically these are obtained from a http.Request
//
// swagger:parameters notificationQuoteCreationNotification
type NotificationQuoteCreationNotificationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	QuoteCreationNotification models.Event
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNotificationQuoteCreationNotificationParams() beforehand.
func (o *NotificationQuoteCreationNotificationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		body, err := models.UnmarshalEvent(r.Body, route.Consumer)
		if err != nil {
			if err == io.EOF {
				err = errors.Required("quoteCreationNotification", "body")
			}
			res = append(res, err)
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.QuoteCreationNotification = body
			}
		}
	} else {
		res = append(res, errors.Required("quoteCreationNotification", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
