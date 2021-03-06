// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoOrchestraCommonRequest proto orchestra common request
//
// swagger:model protoOrchestraCommonRequest
type ProtoOrchestraCommonRequest struct {

	// action
	Action string `json:"action,omitempty"`

	// data
	Data string `json:"data,omitempty"`
}

// Validate validates this proto orchestra common request
func (m *ProtoOrchestraCommonRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoOrchestraCommonRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoOrchestraCommonRequest) UnmarshalBinary(b []byte) error {
	var res ProtoOrchestraCommonRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
