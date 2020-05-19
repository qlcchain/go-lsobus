// Code generated by go-swagger; DO NOT EDIT.

package cancel_product_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-virtual-lsobus/sonata/order/models"
)

// CancelProductOrderFindReader is a Reader for the CancelProductOrderFind structure.
type CancelProductOrderFindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelProductOrderFindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelProductOrderFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelProductOrderFindBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelProductOrderFindUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelProductOrderFindForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelProductOrderFindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCancelProductOrderFindMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewCancelProductOrderFindRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCancelProductOrderFindUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelProductOrderFindInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCancelProductOrderFindServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelProductOrderFindOK creates a CancelProductOrderFindOK with default headers values
func NewCancelProductOrderFindOK() *CancelProductOrderFindOK {
	return &CancelProductOrderFindOK{}
}

/*CancelProductOrderFindOK handles this case with default header values.

Ok
*/
type CancelProductOrderFindOK struct {
	Payload []*models.CancelProductOrder
}

func (o *CancelProductOrderFindOK) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindOK  %+v", 200, o.Payload)
}

func (o *CancelProductOrderFindOK) GetPayload() []*models.CancelProductOrder {
	return o.Payload
}

func (o *CancelProductOrderFindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindBadRequest creates a CancelProductOrderFindBadRequest with default headers values
func NewCancelProductOrderFindBadRequest() *CancelProductOrderFindBadRequest {
	return &CancelProductOrderFindBadRequest{}
}

/*CancelProductOrderFindBadRequest handles this case with default header values.

Bad Request

List of supported error codes:
- 20: Invalid URL parameter value
- 21: Missing body
- 22: Invalid body
- 23: Missing body field
- 24: Invalid body field
- 25: Missing header
- 26: Invalid header value
- 27: Missing query-string parameter
- 28: Invalid query-string parameter value
*/
type CancelProductOrderFindBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindBadRequest) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindBadRequest  %+v", 400, o.Payload)
}

func (o *CancelProductOrderFindBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindUnauthorized creates a CancelProductOrderFindUnauthorized with default headers values
func NewCancelProductOrderFindUnauthorized() *CancelProductOrderFindUnauthorized {
	return &CancelProductOrderFindUnauthorized{}
}

/*CancelProductOrderFindUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type CancelProductOrderFindUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelProductOrderFindUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindForbidden creates a CancelProductOrderFindForbidden with default headers values
func NewCancelProductOrderFindForbidden() *CancelProductOrderFindForbidden {
	return &CancelProductOrderFindForbidden{}
}

/*CancelProductOrderFindForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type CancelProductOrderFindForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindForbidden) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindForbidden  %+v", 403, o.Payload)
}

func (o *CancelProductOrderFindForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindNotFound creates a CancelProductOrderFindNotFound with default headers values
func NewCancelProductOrderFindNotFound() *CancelProductOrderFindNotFound {
	return &CancelProductOrderFindNotFound{}
}

/*CancelProductOrderFindNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type CancelProductOrderFindNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindNotFound) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindNotFound  %+v", 404, o.Payload)
}

func (o *CancelProductOrderFindNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindMethodNotAllowed creates a CancelProductOrderFindMethodNotAllowed with default headers values
func NewCancelProductOrderFindMethodNotAllowed() *CancelProductOrderFindMethodNotAllowed {
	return &CancelProductOrderFindMethodNotAllowed{}
}

/*CancelProductOrderFindMethodNotAllowed handles this case with default header values.

Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type CancelProductOrderFindMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CancelProductOrderFindMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindRequestTimeout creates a CancelProductOrderFindRequestTimeout with default headers values
func NewCancelProductOrderFindRequestTimeout() *CancelProductOrderFindRequestTimeout {
	return &CancelProductOrderFindRequestTimeout{}
}

/*CancelProductOrderFindRequestTimeout handles this case with default header values.

Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type CancelProductOrderFindRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindRequestTimeout  %+v", 408, o.Payload)
}

func (o *CancelProductOrderFindRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindUnprocessableEntity creates a CancelProductOrderFindUnprocessableEntity with default headers values
func NewCancelProductOrderFindUnprocessableEntity() *CancelProductOrderFindUnprocessableEntity {
	return &CancelProductOrderFindUnprocessableEntity{}
}

/*CancelProductOrderFindUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error
*/
type CancelProductOrderFindUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CancelProductOrderFindUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindInternalServerError creates a CancelProductOrderFindInternalServerError with default headers values
func NewCancelProductOrderFindInternalServerError() *CancelProductOrderFindInternalServerError {
	return &CancelProductOrderFindInternalServerError{}
}

/*CancelProductOrderFindInternalServerError handles this case with default header values.

Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type CancelProductOrderFindInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelProductOrderFindInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelProductOrderFindServiceUnavailable creates a CancelProductOrderFindServiceUnavailable with default headers values
func NewCancelProductOrderFindServiceUnavailable() *CancelProductOrderFindServiceUnavailable {
	return &CancelProductOrderFindServiceUnavailable{}
}

/*CancelProductOrderFindServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type CancelProductOrderFindServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *CancelProductOrderFindServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /cancelProductOrder][%d] cancelProductOrderFindServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelProductOrderFindServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *CancelProductOrderFindServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
