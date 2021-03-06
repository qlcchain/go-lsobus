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

// Resiliency resiliency
//
// swagger:model resiliency
type Resiliency string

const (

	// ResiliencyNONE captures enum value "NONE"
	ResiliencyNONE Resiliency = "NONE"

	// ResiliencyNr2LINKACTIVESTANDBY captures enum value "2_LINK_ACTIVE_STANDBY"
	ResiliencyNr2LINKACTIVESTANDBY Resiliency = "2_LINK_ACTIVE_STANDBY"

	// ResiliencyALLACTIVE captures enum value "ALL_ACTIVE"
	ResiliencyALLACTIVE Resiliency = "ALL_ACTIVE"

	// ResiliencyOTHER captures enum value "OTHER"
	ResiliencyOTHER Resiliency = "OTHER"
)

// for schema
var resiliencyEnum []interface{}

func init() {
	var res []Resiliency
	if err := json.Unmarshal([]byte(`["NONE","2_LINK_ACTIVE_STANDBY","ALL_ACTIVE","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resiliencyEnum = append(resiliencyEnum, v)
	}
}

func (m Resiliency) validateResiliencyEnum(path, location string, value Resiliency) error {
	if err := validate.Enum(path, location, value, resiliencyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this resiliency
func (m Resiliency) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResiliencyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
