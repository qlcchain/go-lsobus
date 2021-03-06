// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Quantity An amount in a given unit
//
// swagger:model Quantity
type Quantity struct {

	// Numeric value in a given unit
	Amount float32 `json:"amount,omitempty"`

	// Unit
	Units string `json:"units,omitempty"`
}

// Validate validates this quantity
func (m *Quantity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Quantity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Quantity) UnmarshalBinary(b []byte) error {
	var res Quantity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
