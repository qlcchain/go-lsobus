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

// Contact Contact allow to capture contact information. It is used to capture billing account contact information
//
// swagger:model Contact
type Contact struct {

	// Technical attribut to extend API
	AtReferredType string `json:"@referredType,omitempty"`

	// Identifies the name of the person or office to be contacted on billing matters.
	// Required: true
	ContactName *string `json:"contactName"`

	// Identifies the electronic mail address of the Billing Contact when a Buyer profile does not already exist.
	// Required: true
	EmailAdress *string `json:"emailAdress"`

	// Identifies the telephone number (excluding extension) of the billing contact
	// Required: true
	PhoneNumber *string `json:"phoneNumber"`

	// Identifies the telephone number extension of the billing contact
	PhoneNumberExtension string `json:"phoneNumberExtension,omitempty"`

	// Identifies the address of the person or office to be contacted on billing matters.
	// Required: true
	StreetAdress *string `json:"streetAdress"`
}

// Validate validates this contact
func (m *Contact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailAdress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreetAdress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Contact) validateContactName(formats strfmt.Registry) error {

	if err := validate.Required("contactName", "body", m.ContactName); err != nil {
		return err
	}

	return nil
}

func (m *Contact) validateEmailAdress(formats strfmt.Registry) error {

	if err := validate.Required("emailAdress", "body", m.EmailAdress); err != nil {
		return err
	}

	return nil
}

func (m *Contact) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phoneNumber", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (m *Contact) validateStreetAdress(formats strfmt.Registry) error {

	if err := validate.Required("streetAdress", "body", m.StreetAdress); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Contact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Contact) UnmarshalBinary(b []byte) error {
	var res Contact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
