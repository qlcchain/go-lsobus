// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoCreateOrderParam proto create order param
//
// swagger:model protoCreateOrderParam
type ProtoCreateOrderParam struct {

	// buyer
	Buyer *ProtoUser `json:"buyer,omitempty"`

	// connection param
	ConnectionParam []*ProtoConnectionParam `json:"connectionParam"`

	// privacy
	Privacy *ProtoContractPrivacyParam `json:"privacy,omitempty"`

	// seller
	Seller *ProtoUser `json:"seller,omitempty"`
}

// Validate validates this proto create order param
func (m *ProtoCreateOrderParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuyer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionParam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeller(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoCreateOrderParam) validateBuyer(formats strfmt.Registry) error {

	if swag.IsZero(m.Buyer) { // not required
		return nil
	}

	if m.Buyer != nil {
		if err := m.Buyer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buyer")
			}
			return err
		}
	}

	return nil
}

func (m *ProtoCreateOrderParam) validateConnectionParam(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionParam) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectionParam); i++ {
		if swag.IsZero(m.ConnectionParam[i]) { // not required
			continue
		}

		if m.ConnectionParam[i] != nil {
			if err := m.ConnectionParam[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectionParam" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtoCreateOrderParam) validatePrivacy(formats strfmt.Registry) error {

	if swag.IsZero(m.Privacy) { // not required
		return nil
	}

	if m.Privacy != nil {
		if err := m.Privacy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtoCreateOrderParam) validateSeller(formats strfmt.Registry) error {

	if swag.IsZero(m.Seller) { // not required
		return nil
	}

	if m.Seller != nil {
		if err := m.Seller.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seller")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtoCreateOrderParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoCreateOrderParam) UnmarshalBinary(b []byte) error {
	var res ProtoCreateOrderParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
