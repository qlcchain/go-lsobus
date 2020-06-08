// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoOrchestraCommonResponse proto orchestra common response
//
// swagger:model protoOrchestraCommonResponse
type ProtoOrchestraCommonResponse struct {

	// action
	Action string `json:"action,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// result count
	ResultCount int32 `json:"resultCount,omitempty"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this proto orchestra common response
func (m *ProtoOrchestraCommonResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoOrchestraCommonResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoOrchestraCommonResponse) UnmarshalBinary(b []byte) error {
	var res ProtoOrchestraCommonResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
