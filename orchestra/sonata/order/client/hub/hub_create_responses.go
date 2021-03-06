// Code generated by go-swagger; DO NOT EDIT.

package hub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-lsobus/orchestra/sonata/order/models"
)

// HubCreateReader is a Reader for the HubCreate structure.
type HubCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HubCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewHubCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHubCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHubCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHubCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHubCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewHubCreateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewHubCreateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewHubCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHubCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewHubCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHubCreateCreated creates a HubCreateCreated with default headers values
func NewHubCreateCreated() *HubCreateCreated {
	return &HubCreateCreated{}
}

/*HubCreateCreated handles this case with default header values.

Created
*/
type HubCreateCreated struct {
	Payload *models.Hub
}

func (o *HubCreateCreated) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateCreated  %+v", 201, o.Payload)
}

func (o *HubCreateCreated) GetPayload() *models.Hub {
	return o.Payload
}

func (o *HubCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hub)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateBadRequest creates a HubCreateBadRequest with default headers values
func NewHubCreateBadRequest() *HubCreateBadRequest {
	return &HubCreateBadRequest{}
}

/*HubCreateBadRequest handles this case with default header values.

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
type HubCreateBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateBadRequest  %+v", 400, o.Payload)
}

func (o *HubCreateBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateUnauthorized creates a HubCreateUnauthorized with default headers values
func NewHubCreateUnauthorized() *HubCreateUnauthorized {
	return &HubCreateUnauthorized{}
}

/*HubCreateUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type HubCreateUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *HubCreateUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateForbidden creates a HubCreateForbidden with default headers values
func NewHubCreateForbidden() *HubCreateForbidden {
	return &HubCreateForbidden{}
}

/*HubCreateForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type HubCreateForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateForbidden  %+v", 403, o.Payload)
}

func (o *HubCreateForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateNotFound creates a HubCreateNotFound with default headers values
func NewHubCreateNotFound() *HubCreateNotFound {
	return &HubCreateNotFound{}
}

/*HubCreateNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type HubCreateNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateNotFound  %+v", 404, o.Payload)
}

func (o *HubCreateNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateMethodNotAllowed creates a HubCreateMethodNotAllowed with default headers values
func NewHubCreateMethodNotAllowed() *HubCreateMethodNotAllowed {
	return &HubCreateMethodNotAllowed{}
}

/*HubCreateMethodNotAllowed handles this case with default header values.

Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type HubCreateMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *HubCreateMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateRequestTimeout creates a HubCreateRequestTimeout with default headers values
func NewHubCreateRequestTimeout() *HubCreateRequestTimeout {
	return &HubCreateRequestTimeout{}
}

/*HubCreateRequestTimeout handles this case with default header values.

Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type HubCreateRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateRequestTimeout  %+v", 408, o.Payload)
}

func (o *HubCreateRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateUnprocessableEntity creates a HubCreateUnprocessableEntity with default headers values
func NewHubCreateUnprocessableEntity() *HubCreateUnprocessableEntity {
	return &HubCreateUnprocessableEntity{}
}

/*HubCreateUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error
*/
type HubCreateUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *HubCreateUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateInternalServerError creates a HubCreateInternalServerError with default headers values
func NewHubCreateInternalServerError() *HubCreateInternalServerError {
	return &HubCreateInternalServerError{}
}

/*HubCreateInternalServerError handles this case with default header values.

Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type HubCreateInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *HubCreateInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubCreateServiceUnavailable creates a HubCreateServiceUnavailable with default headers values
func NewHubCreateServiceUnavailable() *HubCreateServiceUnavailable {
	return &HubCreateServiceUnavailable{}
}

/*HubCreateServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type HubCreateServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *HubCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /hub][%d] hubCreateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *HubCreateServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
