// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Hub This resource is used to manage notification subscription.
//
// swagger:model Hub
type Hub struct {

	// callback urn, for instance an url http://yourdomain/listener/api/v1/event
	// Required: true
	Callback *string `json:"callback"`

	// id of the Hub
	// Required: true
	ID *string `json:"id"`

	// attribute selection & search criteria
	// Required: true
	Query *string `json:"query"`
}

// Validate validates this hub
func (m *Hub) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallback(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hub) validateCallback(formats strfmt.Registry) error {

	if err := validate.Required("callback", "body", m.Callback); err != nil {
		return err
	}

	return nil
}

func (m *Hub) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Hub) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Hub) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hub) UnmarshalBinary(b []byte) error {
	var res Hub
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
