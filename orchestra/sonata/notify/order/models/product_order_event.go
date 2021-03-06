// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductOrderEvent ProductOrder structure used for notification
//
// swagger:model ProductOrderEvent
type ProductOrderEvent struct {

	// at base type
	AtBaseType string `json:"@baseType,omitempty"`

	// at schema location
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// at type
	AtType string `json:"@type,omitempty"`

	// The date the order is completed. Format is YYYY-MM-DDThh:mmTZD (e.g. 1997-07-16T19:20+01:00)
	// Format: date-time
	CompletionDate strfmt.DateTime `json:"completionDate,omitempty"`

	// Expected delivery date amended by the provider
	// Format: date-time
	ExpectedCompletionDate strfmt.DateTime `json:"expectedCompletionDate,omitempty"`

	// A number that uniquely identifies an order within the buyer's enterprise.
	// Required: true
	ExternalID *string `json:"externalId"`

	// Unique (within the ordering domain) identifier for the order that is generated by the seller when the order is initially accepted.
	// Required: true
	ID *string `json:"id"`

	// note
	Note []*Note `json:"note"`

	// order item
	// Required: true
	OrderItem []*OrderItemEvent `json:"orderItem"`

	// order message
	OrderMessage []*OrderMessage `json:"orderMessage"`

	// The version number that the Buyer uses to refer to this particular version of the order
	// Required: true
	OrderVersion *string `json:"orderVersion"`

	// An identifier that is used to group Orders that represent a unit of functionality that is important to a Buyer.  A Project can be used to relate multiple Orders together.
	ProjectID string `json:"projectId,omitempty"`

	// related party
	// Required: true
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// Identifies the buyer's desired due date (requested delivery date). Cannot be requested on cancelled orders.  Format is YYYY-MM-DDThh:mmTZD (e.g. 1997-07-16T19:20+01:00).
	// Required: true
	// Format: date-time
	RequestedCompletionDate *strfmt.DateTime `json:"requestedCompletionDate"`

	// The buyer's requested date that order processing should start. Format is YYYY-MM-DDThh:mmTZD (e.g. 1997-07-16T19:20+01:00).
	// Format: date-time
	RequestedStartDate strfmt.DateTime `json:"requestedStartDate,omitempty"`

	// state
	// Required: true
	State ProductOrderStateType `json:"state"`
}

// Validate validates this product order event
func (m *ProductOrderEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductOrderEvent) validateCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completionDate", "body", "date-time", m.CompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderEvent) validateExpectedCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpectedCompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedCompletionDate", "body", "date-time", m.ExpectedCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderEvent) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("externalId", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderEvent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderEvent) validateNote(formats strfmt.Registry) error {

	if swag.IsZero(m.Note) { // not required
		return nil
	}

	for i := 0; i < len(m.Note); i++ {
		if swag.IsZero(m.Note[i]) { // not required
			continue
		}

		if m.Note[i] != nil {
			if err := m.Note[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("note" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOrderEvent) validateOrderItem(formats strfmt.Registry) error {

	if err := validate.Required("orderItem", "body", m.OrderItem); err != nil {
		return err
	}

	for i := 0; i < len(m.OrderItem); i++ {
		if swag.IsZero(m.OrderItem[i]) { // not required
			continue
		}

		if m.OrderItem[i] != nil {
			if err := m.OrderItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOrderEvent) validateOrderMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderMessage) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderMessage); i++ {
		if swag.IsZero(m.OrderMessage[i]) { // not required
			continue
		}

		if m.OrderMessage[i] != nil {
			if err := m.OrderMessage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderMessage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOrderEvent) validateOrderVersion(formats strfmt.Registry) error {

	if err := validate.Required("orderVersion", "body", m.OrderVersion); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderEvent) validateRelatedParty(formats strfmt.Registry) error {

	if err := validate.Required("relatedParty", "body", m.RelatedParty); err != nil {
		return err
	}

	for i := 0; i < len(m.RelatedParty); i++ {
		if swag.IsZero(m.RelatedParty[i]) { // not required
			continue
		}

		if m.RelatedParty[i] != nil {
			if err := m.RelatedParty[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOrderEvent) validateRequestedCompletionDate(formats strfmt.Registry) error {

	if err := validate.Required("requestedCompletionDate", "body", m.RequestedCompletionDate); err != nil {
		return err
	}

	if err := validate.FormatOf("requestedCompletionDate", "body", "date-time", m.RequestedCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderEvent) validateRequestedStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedStartDate", "body", "date-time", m.RequestedStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderEvent) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductOrderEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductOrderEvent) UnmarshalBinary(b []byte) error {
	var res ProductOrderEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
