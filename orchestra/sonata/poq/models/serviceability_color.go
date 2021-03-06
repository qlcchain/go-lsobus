// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ServiceabilityColor A color that indicates confidence to service the request.
//
// swagger:model ServiceabilityColor
type ServiceabilityColor string

const (

	// ServiceabilityColorGreen captures enum value "green"
	ServiceabilityColorGreen ServiceabilityColor = "green"

	// ServiceabilityColorRed captures enum value "red"
	ServiceabilityColorRed ServiceabilityColor = "red"

	// ServiceabilityColorYellow captures enum value "yellow"
	ServiceabilityColorYellow ServiceabilityColor = "yellow"
)

// for schema
var serviceabilityColorEnum []interface{}

func init() {
	var res []ServiceabilityColor
	if err := json.Unmarshal([]byte(`["green","red","yellow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceabilityColorEnum = append(serviceabilityColorEnum, v)
	}
}

func (m ServiceabilityColor) validateServiceabilityColorEnum(path, location string, value ServiceabilityColor) error {
	if err := validate.Enum(path, location, value, serviceabilityColorEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this serviceability color
func (m ServiceabilityColor) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServiceabilityColorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
