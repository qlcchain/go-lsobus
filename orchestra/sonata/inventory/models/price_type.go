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

// PriceType price type
//
// swagger:model PriceType
type PriceType string

const (

	// PriceTypeRecurring captures enum value "recurring"
	PriceTypeRecurring PriceType = "recurring"

	// PriceTypeNonRecurring captures enum value "nonRecurring"
	PriceTypeNonRecurring PriceType = "nonRecurring"
)

// for schema
var priceTypeEnum []interface{}

func init() {
	var res []PriceType
	if err := json.Unmarshal([]byte(`["recurring","nonRecurring"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		priceTypeEnum = append(priceTypeEnum, v)
	}
}

func (m PriceType) validatePriceTypeEnum(path, location string, value PriceType) error {
	if err := validate.Enum(path, location, value, priceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this price type
func (m PriceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePriceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
