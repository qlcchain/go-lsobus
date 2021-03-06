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

// QuoteLevel Quote level
//
// swagger:model QuoteLevel
type QuoteLevel string

const (

	// QuoteLevelBUDGETARY captures enum value "BUDGETARY"
	QuoteLevelBUDGETARY QuoteLevel = "BUDGETARY"

	// QuoteLevelINDICATIVE captures enum value "INDICATIVE"
	QuoteLevelINDICATIVE QuoteLevel = "INDICATIVE"

	// QuoteLevelFIRM captures enum value "FIRM"
	QuoteLevelFIRM QuoteLevel = "FIRM"
)

// for schema
var quoteLevelEnum []interface{}

func init() {
	var res []QuoteLevel
	if err := json.Unmarshal([]byte(`["BUDGETARY","INDICATIVE","FIRM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		quoteLevelEnum = append(quoteLevelEnum, v)
	}
}

func (m QuoteLevel) validateQuoteLevelEnum(path, location string, value QuoteLevel) error {
	if err := validate.Enum(path, location, value, quoteLevelEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this quote level
func (m QuoteLevel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateQuoteLevelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
