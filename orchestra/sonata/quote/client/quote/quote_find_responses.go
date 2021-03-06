// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-lsobus/orchestra/sonata/quote/models"
)

// QuoteFindReader is a Reader for the QuoteFind structure.
type QuoteFindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteFindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteFindBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQuoteFindUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuoteFindForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQuoteFindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewQuoteFindMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewQuoteFindUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuoteFindInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewQuoteFindServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQuoteFindOK creates a QuoteFindOK with default headers values
func NewQuoteFindOK() *QuoteFindOK {
	return &QuoteFindOK{}
}

/*QuoteFindOK handles this case with default header values.

Ok
*/
type QuoteFindOK struct {
	/*The number of resources retrieved in the response
	 */
	XResultCount int32
	/*Total number of items matching criteria
	 */
	XTotalCount int32

	Payload []*models.QuoteFind
}

func (o *QuoteFindOK) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindOK  %+v", 200, o.Payload)
}

func (o *QuoteFindOK) GetPayload() []*models.QuoteFind {
	return o.Payload
}

func (o *QuoteFindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Result-Count
	xResultCount, err := swag.ConvertInt32(response.GetHeader("X-Result-Count"))
	if err != nil {
		return errors.InvalidType("X-Result-Count", "header", "int32", response.GetHeader("X-Result-Count"))
	}
	o.XResultCount = xResultCount

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt32(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int32", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteFindBadRequest creates a QuoteFindBadRequest with default headers values
func NewQuoteFindBadRequest() *QuoteFindBadRequest {
	return &QuoteFindBadRequest{}
}

/*QuoteFindBadRequest handles this case with default header values.

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
type QuoteFindBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteFindBadRequest) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindBadRequest  %+v", 400, o.Payload)
}

func (o *QuoteFindBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteFindBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteFindUnauthorized creates a QuoteFindUnauthorized with default headers values
func NewQuoteFindUnauthorized() *QuoteFindUnauthorized {
	return &QuoteFindUnauthorized{}
}

/*QuoteFindUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type QuoteFindUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteFindUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteFindUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteFindUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteFindForbidden creates a QuoteFindForbidden with default headers values
func NewQuoteFindForbidden() *QuoteFindForbidden {
	return &QuoteFindForbidden{}
}

/*QuoteFindForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type QuoteFindForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteFindForbidden) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindForbidden  %+v", 403, o.Payload)
}

func (o *QuoteFindForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteFindForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteFindNotFound creates a QuoteFindNotFound with default headers values
func NewQuoteFindNotFound() *QuoteFindNotFound {
	return &QuoteFindNotFound{}
}

/*QuoteFindNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type QuoteFindNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteFindNotFound) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindNotFound  %+v", 404, o.Payload)
}

func (o *QuoteFindNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteFindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteFindMethodNotAllowed creates a QuoteFindMethodNotAllowed with default headers values
func NewQuoteFindMethodNotAllowed() *QuoteFindMethodNotAllowed {
	return &QuoteFindMethodNotAllowed{}
}

/*QuoteFindMethodNotAllowed handles this case with default header values.

Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type QuoteFindMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteFindMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *QuoteFindMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteFindMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteFindUnprocessableEntity creates a QuoteFindUnprocessableEntity with default headers values
func NewQuoteFindUnprocessableEntity() *QuoteFindUnprocessableEntity {
	return &QuoteFindUnprocessableEntity{}
}

/*QuoteFindUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error





 - code: 100
message: Too many records retrieved - please restrict requested parameter value(s)
description:
*/
type QuoteFindUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteFindUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *QuoteFindUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteFindUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteFindInternalServerError creates a QuoteFindInternalServerError with default headers values
func NewQuoteFindInternalServerError() *QuoteFindInternalServerError {
	return &QuoteFindInternalServerError{}
}

/*QuoteFindInternalServerError handles this case with default header values.

Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type QuoteFindInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteFindInternalServerError) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindInternalServerError  %+v", 500, o.Payload)
}

func (o *QuoteFindInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteFindInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteFindServiceUnavailable creates a QuoteFindServiceUnavailable with default headers values
func NewQuoteFindServiceUnavailable() *QuoteFindServiceUnavailable {
	return &QuoteFindServiceUnavailable{}
}

/*QuoteFindServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type QuoteFindServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteFindServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteFindServiceUnavailable  %+v", 503, o.Payload)
}

func (o *QuoteFindServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteFindServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
