// Code generated by go-swagger; DO NOT EDIT.

package hub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-virtual-lsobus/sonata/order/models"
)

// HubDeleteReader is a Reader for the HubDelete structure.
type HubDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HubDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewHubDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHubDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHubDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHubDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHubDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewHubDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewHubDeleteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewHubDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHubDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewHubDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHubDeleteNoContent creates a HubDeleteNoContent with default headers values
func NewHubDeleteNoContent() *HubDeleteNoContent {
	return &HubDeleteNoContent{}
}

/*HubDeleteNoContent handles this case with default header values.

No Content
*/
type HubDeleteNoContent struct {
}

func (o *HubDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteNoContent ", 204)
}

func (o *HubDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHubDeleteBadRequest creates a HubDeleteBadRequest with default headers values
func NewHubDeleteBadRequest() *HubDeleteBadRequest {
	return &HubDeleteBadRequest{}
}

/*HubDeleteBadRequest handles this case with default header values.

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
type HubDeleteBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *HubDeleteBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubDeleteUnauthorized creates a HubDeleteUnauthorized with default headers values
func NewHubDeleteUnauthorized() *HubDeleteUnauthorized {
	return &HubDeleteUnauthorized{}
}

/*HubDeleteUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type HubDeleteUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *HubDeleteUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubDeleteForbidden creates a HubDeleteForbidden with default headers values
func NewHubDeleteForbidden() *HubDeleteForbidden {
	return &HubDeleteForbidden{}
}

/*HubDeleteForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type HubDeleteForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteForbidden  %+v", 403, o.Payload)
}

func (o *HubDeleteForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubDeleteNotFound creates a HubDeleteNotFound with default headers values
func NewHubDeleteNotFound() *HubDeleteNotFound {
	return &HubDeleteNotFound{}
}

/*HubDeleteNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type HubDeleteNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteNotFound  %+v", 404, o.Payload)
}

func (o *HubDeleteNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubDeleteMethodNotAllowed creates a HubDeleteMethodNotAllowed with default headers values
func NewHubDeleteMethodNotAllowed() *HubDeleteMethodNotAllowed {
	return &HubDeleteMethodNotAllowed{}
}

/*HubDeleteMethodNotAllowed handles this case with default header values.

Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type HubDeleteMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *HubDeleteMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubDeleteRequestTimeout creates a HubDeleteRequestTimeout with default headers values
func NewHubDeleteRequestTimeout() *HubDeleteRequestTimeout {
	return &HubDeleteRequestTimeout{}
}

/*HubDeleteRequestTimeout handles this case with default header values.

Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type HubDeleteRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteRequestTimeout  %+v", 408, o.Payload)
}

func (o *HubDeleteRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubDeleteUnprocessableEntity creates a HubDeleteUnprocessableEntity with default headers values
func NewHubDeleteUnprocessableEntity() *HubDeleteUnprocessableEntity {
	return &HubDeleteUnprocessableEntity{}
}

/*HubDeleteUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error
*/
type HubDeleteUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *HubDeleteUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubDeleteInternalServerError creates a HubDeleteInternalServerError with default headers values
func NewHubDeleteInternalServerError() *HubDeleteInternalServerError {
	return &HubDeleteInternalServerError{}
}

/*HubDeleteInternalServerError handles this case with default header values.

Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type HubDeleteInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *HubDeleteInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHubDeleteServiceUnavailable creates a HubDeleteServiceUnavailable with default headers values
func NewHubDeleteServiceUnavailable() *HubDeleteServiceUnavailable {
	return &HubDeleteServiceUnavailable{}
}

/*HubDeleteServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type HubDeleteServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *HubDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /hub/{HubId}][%d] hubDeleteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *HubDeleteServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *HubDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
