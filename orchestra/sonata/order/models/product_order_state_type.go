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

// ProductOrderStateType An enumeration of valid order states
//
// swagger:model ProductOrderStateType
type ProductOrderStateType string

const (

	// ProductOrderStateTypeAcknowledged captures enum value "acknowledged"
	ProductOrderStateTypeAcknowledged ProductOrderStateType = "acknowledged"

	// ProductOrderStateTypeRejected captures enum value "rejected"
	ProductOrderStateTypeRejected ProductOrderStateType = "rejected"

	// ProductOrderStateTypeInProgress captures enum value "inProgress"
	ProductOrderStateTypeInProgress ProductOrderStateType = "inProgress"

	// ProductOrderStateTypePending captures enum value "pending"
	ProductOrderStateTypePending ProductOrderStateType = "pending"

	// ProductOrderStateTypeHeld captures enum value "held"
	ProductOrderStateTypeHeld ProductOrderStateType = "held"

	// ProductOrderStateTypeAssessingCancellation captures enum value "assessingCancellation"
	ProductOrderStateTypeAssessingCancellation ProductOrderStateType = "assessingCancellation"

	// ProductOrderStateTypePendingCancellation captures enum value "pendingCancellation"
	ProductOrderStateTypePendingCancellation ProductOrderStateType = "pendingCancellation"

	// ProductOrderStateTypeCancelled captures enum value "cancelled"
	ProductOrderStateTypeCancelled ProductOrderStateType = "cancelled"

	// ProductOrderStateTypeInProgressConfigured captures enum value "inProgress.configured"
	ProductOrderStateTypeInProgressConfigured ProductOrderStateType = "inProgress.configured"

	// ProductOrderStateTypeInProgressConfirmed captures enum value "inProgress.confirmed"
	ProductOrderStateTypeInProgressConfirmed ProductOrderStateType = "inProgress.confirmed"

	// ProductOrderStateTypeInProgressJeopardy captures enum value "inProgress.jeopardy"
	ProductOrderStateTypeInProgressJeopardy ProductOrderStateType = "inProgress.jeopardy"

	// ProductOrderStateTypeFailed captures enum value "failed"
	ProductOrderStateTypeFailed ProductOrderStateType = "failed"

	// ProductOrderStateTypePartial captures enum value "partial"
	ProductOrderStateTypePartial ProductOrderStateType = "partial"

	// ProductOrderStateTypeCompleted captures enum value "completed"
	ProductOrderStateTypeCompleted ProductOrderStateType = "completed"
)

// for schema
var productOrderStateTypeEnum []interface{}

func init() {
	var res []ProductOrderStateType
	if err := json.Unmarshal([]byte(`["acknowledged","rejected","inProgress","pending","held","assessingCancellation","pendingCancellation","cancelled","inProgress.configured","inProgress.confirmed","inProgress.jeopardy","failed","partial","completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productOrderStateTypeEnum = append(productOrderStateTypeEnum, v)
	}
}

func (m ProductOrderStateType) validateProductOrderStateTypeEnum(path, location string, value ProductOrderStateType) error {
	if err := validate.Enum(path, location, value, productOrderStateTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this product order state type
func (m ProductOrderStateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProductOrderStateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
