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

// QuoteItem A quote items describe an action to be performed on a productOffering or a product in order to get pricing elements and condition
//
// swagger:model QuoteItem
type QuoteItem struct {

	// Link to the schema describing this REST resource
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Indicates the base (class) type of the quote Item.
	AtType string `json:"@type,omitempty"`

	// action
	// Required: true
	Action ProductActionType `json:"action"`

	// ID given by the consumer and only understandable by him (to facilitate his searches afterwards)
	ExternalID string `json:"externalId,omitempty"`

	// Identifier of the quote item (generally it is a sequence number 01, 02, 03, ...).
	// Required: true
	ID *string `json:"id"`

	// note
	Note []*Note `json:"note"`

	// product
	Product *Product `json:"product,omitempty"`

	// product offering
	ProductOffering *ProductOfferingRef `json:"productOffering,omitempty"`

	// qualification
	Qualification []*ProductOfferingQualificationRef `json:"qualification"`

	// pre-calculated price
	PreCalculatedPrice *QuotePrice `json:"preCalculatedPrice,omitempty"`

	// quote item price
	QuoteItemPrice []*QuotePrice `json:"quoteItemPrice"`

	// quote item relationship
	QuoteItemRelationship []*QuoteItemRelationship `json:"quoteItemRelationship"`

	// quote item term
	QuoteItemTerm *ItemTerm `json:"quoteItemTerm,omitempty"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// requested quote item term
	RequestedQuoteItemTerm *ItemTerm `json:"requestedQuoteItemTerm,omitempty"`

	// state
	// Required: true
	State QuoteItemStateType `json:"state"`
}

// Validate validates this quote item
func (m *QuoteItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductOffering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQualification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteItemPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteItemRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteItemTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedQuoteItemTerm(formats); err != nil {
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

func (m *QuoteItem) validateAction(formats strfmt.Registry) error {

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *QuoteItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *QuoteItem) validateNote(formats strfmt.Registry) error {

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

func (m *QuoteItem) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItem) validateProductOffering(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductOffering) { // not required
		return nil
	}

	if m.ProductOffering != nil {
		if err := m.ProductOffering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productOffering")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItem) validateQualification(formats strfmt.Registry) error {

	if swag.IsZero(m.Qualification) { // not required
		return nil
	}

	for i := 0; i < len(m.Qualification); i++ {
		if swag.IsZero(m.Qualification[i]) { // not required
			continue
		}

		if m.Qualification[i] != nil {
			if err := m.Qualification[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("qualification" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItem) validateQuoteItemPrice(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteItemPrice) { // not required
		return nil
	}

	for i := 0; i < len(m.QuoteItemPrice); i++ {
		if swag.IsZero(m.QuoteItemPrice[i]) { // not required
			continue
		}

		if m.QuoteItemPrice[i] != nil {
			if err := m.QuoteItemPrice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quoteItemPrice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItem) validateQuoteItemRelationship(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteItemRelationship) { // not required
		return nil
	}

	for i := 0; i < len(m.QuoteItemRelationship); i++ {
		if swag.IsZero(m.QuoteItemRelationship[i]) { // not required
			continue
		}

		if m.QuoteItemRelationship[i] != nil {
			if err := m.QuoteItemRelationship[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quoteItemRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItem) validateQuoteItemTerm(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteItemTerm) { // not required
		return nil
	}

	if m.QuoteItemTerm != nil {
		if err := m.QuoteItemTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteItemTerm")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItem) validateRelatedParty(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedParty) { // not required
		return nil
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

func (m *QuoteItem) validateRequestedQuoteItemTerm(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedQuoteItemTerm) { // not required
		return nil
	}

	if m.RequestedQuoteItemTerm != nil {
		if err := m.RequestedQuoteItemTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedQuoteItemTerm")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItem) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuoteItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteItem) UnmarshalBinary(b []byte) error {
	var res QuoteItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
