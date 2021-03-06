// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RelatedParty A related party defines party or party role linked to a quote.
//
// swagger:model RelatedParty
type RelatedParty struct {

	// Indicates the base (class) type of the party.
	AtReferredType string `json:"@referredType,omitempty"`

	// email of the related party
	EmailAddress string `json:"emailAddress,omitempty"`

	// Unique identifier of a related party
	ID string `json:"id,omitempty"`

	// Name of the related party
	Name string `json:"name,omitempty"`

	// Telephone number of the related party
	Number string `json:"number,omitempty"`

	// Telephone number extension of the related party
	NumberExtension string `json:"numberExtension,omitempty"`

	// Role of the related party for this quote or quoteItem
	// Required: true
	Role []string `json:"role"`
}

// Validate validates this related party
func (m *RelatedParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedParty) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelatedParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelatedParty) UnmarshalBinary(b []byte) error {
	var res RelatedParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
