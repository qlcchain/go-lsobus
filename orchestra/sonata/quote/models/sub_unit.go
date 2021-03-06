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

// SubUnit sub unit
//
// swagger:model SubUnit
type SubUnit struct {

	// The discriminator used for the subunit, often just a simple number but may also be a range.
	// Required: true
	SubUnitIdentifier *string `json:"subUnitIdentifier"`

	// The type of subunit e.g.BERTH, FLAT, PIER, SUITE, SHOP, TOWER, UNIT, WHARF.
	// Required: true
	SubUnitType *string `json:"subUnitType"`
}

// Validate validates this sub unit
func (m *SubUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubUnitIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubUnitType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubUnit) validateSubUnitIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("subUnitIdentifier", "body", m.SubUnitIdentifier); err != nil {
		return err
	}

	return nil
}

func (m *SubUnit) validateSubUnitType(formats strfmt.Registry) error {

	if err := validate.Required("subUnitType", "body", m.SubUnitType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubUnit) UnmarshalBinary(b []byte) error {
	var res SubUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
