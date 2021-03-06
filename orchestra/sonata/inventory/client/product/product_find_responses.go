// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-lsobus/orchestra/sonata/inventory/models"
)

// ProductFindReader is a Reader for the ProductFind structure.
type ProductFindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductFindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProductFindBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProductFindUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProductFindForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProductFindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewProductFindUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewProductFindServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProductFindOK creates a ProductFindOK with default headers values
func NewProductFindOK() *ProductFindOK {
	return &ProductFindOK{}
}

/*ProductFindOK handles this case with default header values.

Ok
*/
type ProductFindOK struct {
	/*The number of resources retrieved in the response
	 */
	XResultCount string
	/*The total number of matching resources
	 */
	XTotalCount string

	Payload []*models.ProductSummary
}

func (o *ProductFindOK) Error() string {
	return fmt.Sprintf("[GET /product][%d] productFindOK  %+v", 200, o.Payload)
}

func (o *ProductFindOK) GetPayload() []*models.ProductSummary {
	return o.Payload
}

func (o *ProductFindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Result-Count
	o.XResultCount = response.GetHeader("X-Result-Count")

	// response header X-Total_Count
	o.XTotalCount = response.GetHeader("X-Total_Count")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductFindBadRequest creates a ProductFindBadRequest with default headers values
func NewProductFindBadRequest() *ProductFindBadRequest {
	return &ProductFindBadRequest{}
}

/*ProductFindBadRequest handles this case with default header values.

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
type ProductFindBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductFindBadRequest) Error() string {
	return fmt.Sprintf("[GET /product][%d] productFindBadRequest  %+v", 400, o.Payload)
}

func (o *ProductFindBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductFindBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductFindUnauthorized creates a ProductFindUnauthorized with default headers values
func NewProductFindUnauthorized() *ProductFindUnauthorized {
	return &ProductFindUnauthorized{}
}

/*ProductFindUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type ProductFindUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductFindUnauthorized) Error() string {
	return fmt.Sprintf("[GET /product][%d] productFindUnauthorized  %+v", 401, o.Payload)
}

func (o *ProductFindUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductFindUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductFindForbidden creates a ProductFindForbidden with default headers values
func NewProductFindForbidden() *ProductFindForbidden {
	return &ProductFindForbidden{}
}

/*ProductFindForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type ProductFindForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductFindForbidden) Error() string {
	return fmt.Sprintf("[GET /product][%d] productFindForbidden  %+v", 403, o.Payload)
}

func (o *ProductFindForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductFindForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductFindNotFound creates a ProductFindNotFound with default headers values
func NewProductFindNotFound() *ProductFindNotFound {
	return &ProductFindNotFound{}
}

/*ProductFindNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type ProductFindNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductFindNotFound) Error() string {
	return fmt.Sprintf("[GET /product][%d] productFindNotFound  %+v", 404, o.Payload)
}

func (o *ProductFindNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductFindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductFindUnprocessableEntity creates a ProductFindUnprocessableEntity with default headers values
func NewProductFindUnprocessableEntity() *ProductFindUnprocessableEntity {
	return &ProductFindUnprocessableEntity{}
}

/*ProductFindUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error





 - code: 100
message: Too many records retrieved - please restrict requested parameter value(s)
description:


 - code: 103
message: Incomplete request - If place.id is filled, place.type must be filled
description:


 - code: 104
message: Incomplete request - If place.type is filled, place.id must be filled
description:


 - code: 105
message: Incomplete request - If partyRole.role is filled, partyRole.relatedPartyId must be filled
description:


 - code: 106
message: Incomplete request - If partyRole.relatedPartyId is filled, partyRole.role must be filled
description:
*/
type ProductFindUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductFindUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /product][%d] productFindUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ProductFindUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductFindUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductFindServiceUnavailable creates a ProductFindServiceUnavailable with default headers values
func NewProductFindServiceUnavailable() *ProductFindServiceUnavailable {
	return &ProductFindServiceUnavailable{}
}

/*ProductFindServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type ProductFindServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductFindServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /product][%d] productFindServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ProductFindServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductFindServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
