// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Duration duration
//
// swagger:model duration
type Duration struct {

	// unit
	// Required: true
	// Enum: [DAY WEEK MONTH YEAR]
	Unit *string `json:"unit"`

	// value
	// Required: true
	Value *int32 `json:"value"`
}

// Validate validates this duration
func (m *Duration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var durationTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DAY","WEEK","MONTH","YEAR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		durationTypeUnitPropEnum = append(durationTypeUnitPropEnum, v)
	}
}

const (

	// DurationUnitDAY captures enum value "DAY"
	DurationUnitDAY string = "DAY"

	// DurationUnitWEEK captures enum value "WEEK"
	DurationUnitWEEK string = "WEEK"

	// DurationUnitMONTH captures enum value "MONTH"
	DurationUnitMONTH string = "MONTH"

	// DurationUnitYEAR captures enum value "YEAR"
	DurationUnitYEAR string = "YEAR"
)

// prop value enum
func (m *Duration) validateUnitEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, durationTypeUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Duration) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *Duration) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Duration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Duration) UnmarshalBinary(b []byte) error {
	var res Duration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
