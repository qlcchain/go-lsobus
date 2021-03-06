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

// PoqEventType Indicates the type of product offering qualification event.
//
// swagger:model PoqEventType
type PoqEventType string

const (

	// PoqEventTypeProductOfferingQualificationCreateEventNotification captures enum value "ProductOfferingQualificationCreateEventNotification"
	PoqEventTypeProductOfferingQualificationCreateEventNotification PoqEventType = "ProductOfferingQualificationCreateEventNotification"

	// PoqEventTypeProductOfferingQualificationStateChangeEventNotification captures enum value "ProductOfferingQualificationStateChangeEventNotification"
	PoqEventTypeProductOfferingQualificationStateChangeEventNotification PoqEventType = "ProductOfferingQualificationStateChangeEventNotification"
)

// for schema
var poqEventTypeEnum []interface{}

func init() {
	var res []PoqEventType
	if err := json.Unmarshal([]byte(`["ProductOfferingQualificationCreateEventNotification","ProductOfferingQualificationStateChangeEventNotification"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		poqEventTypeEnum = append(poqEventTypeEnum, v)
	}
}

func (m PoqEventType) validatePoqEventTypeEnum(path, location string, value PoqEventType) error {
	if err := validate.Enum(path, location, value, poqEventTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this poq event type
func (m PoqEventType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePoqEventTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
