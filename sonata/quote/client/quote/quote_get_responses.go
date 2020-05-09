// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/iixlabs/virtual-lsobus/sonata/quote/models"
)

// QuoteGetReader is a Reader for the QuoteGet structure.
type QuoteGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQuoteGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuoteGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQuoteGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewQuoteGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewQuoteGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuoteGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewQuoteGetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQuoteGetOK creates a QuoteGetOK with default headers values
func NewQuoteGetOK() *QuoteGetOK {
	return &QuoteGetOK{}
}

/*QuoteGetOK handles this case with default header values.

Ok
*/
type QuoteGetOK struct {
	Payload *models.Quote
}

func (o *QuoteGetOK) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetOK  %+v", 200, o.Payload)
}

func (o *QuoteGetOK) GetPayload() *models.Quote {
	return o.Payload
}

func (o *QuoteGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Quote)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetBadRequest creates a QuoteGetBadRequest with default headers values
func NewQuoteGetBadRequest() *QuoteGetBadRequest {
	return &QuoteGetBadRequest{}
}

/*QuoteGetBadRequest handles this case with default header values.

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
type QuoteGetBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteGetBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetUnauthorized creates a QuoteGetUnauthorized with default headers values
func NewQuoteGetUnauthorized() *QuoteGetUnauthorized {
	return &QuoteGetUnauthorized{}
}

/*QuoteGetUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type QuoteGetUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteGetUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetForbidden creates a QuoteGetForbidden with default headers values
func NewQuoteGetForbidden() *QuoteGetForbidden {
	return &QuoteGetForbidden{}
}

/*QuoteGetForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type QuoteGetForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteGetForbidden) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetForbidden  %+v", 403, o.Payload)
}

func (o *QuoteGetForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetNotFound creates a QuoteGetNotFound with default headers values
func NewQuoteGetNotFound() *QuoteGetNotFound {
	return &QuoteGetNotFound{}
}

/*QuoteGetNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type QuoteGetNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteGetNotFound) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetNotFound  %+v", 404, o.Payload)
}

func (o *QuoteGetNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetMethodNotAllowed creates a QuoteGetMethodNotAllowed with default headers values
func NewQuoteGetMethodNotAllowed() *QuoteGetMethodNotAllowed {
	return &QuoteGetMethodNotAllowed{}
}

/*QuoteGetMethodNotAllowed handles this case with default header values.

Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type QuoteGetMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *QuoteGetMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetUnprocessableEntity creates a QuoteGetUnprocessableEntity with default headers values
func NewQuoteGetUnprocessableEntity() *QuoteGetUnprocessableEntity {
	return &QuoteGetUnprocessableEntity{}
}

/*QuoteGetUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error
*/
type QuoteGetUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *QuoteGetUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetInternalServerError creates a QuoteGetInternalServerError with default headers values
func NewQuoteGetInternalServerError() *QuoteGetInternalServerError {
	return &QuoteGetInternalServerError{}
}

/*QuoteGetInternalServerError handles this case with default header values.

Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type QuoteGetInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetInternalServerError  %+v", 500, o.Payload)
}

func (o *QuoteGetInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetServiceUnavailable creates a QuoteGetServiceUnavailable with default headers values
func NewQuoteGetServiceUnavailable() *QuoteGetServiceUnavailable {
	return &QuoteGetServiceUnavailable{}
}

/*QuoteGetServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type QuoteGetServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteGetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /quote/{id}][%d] quoteGetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *QuoteGetServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteGetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
