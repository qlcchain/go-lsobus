// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductOrderItemCreate A ProductOrderItem_Create object is provided as input to the product order item create operation.
//
// swagger:model ProductOrderItem_Create
type ProductOrderItemCreate struct {

	// Technical attribute to extend API
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Technical attribute to extend API
	AtType string `json:"@type,omitempty"`

	// action
	// Required: true
	Action ProductActionType `json:"action"`

	// billing account
	BillingAccount *BillingAccountRef `json:"billingAccount,omitempty"`

	// Identifier of the line item (generally it is a sequence number 01, 02, 03, ...)
	// Required: true
	ID *string `json:"id"`

	// order item price
	// Min Items: 1
	OrderItemPrice []*OrderItemPrice `json:"orderItemPrice,omitempty"`

	// order item relationship
	OrderItemRelationship []*OrderItemRelationShip `json:"orderItemRelationship,omitempty"`

	// pricing method
	PricingMethod PricingMethod `json:"pricingMethod,omitempty"`

	// The identifier references the previously agreed upon pricing terms, as applicable, based on the pricingMethod (e.g. a contract id or tariff id.
	PricingReference string `json:"pricingReference,omitempty"`

	// The length of the contract in months
	// Minimum: 0
	PricingTerm *int32 `json:"pricingTerm,omitempty"`

	// product
	Product *Product `json:"product,omitempty"`

	// product offering
	// Required: true
	ProductOffering *ProductOfferingRef `json:"productOffering,omitempty"`

	// qualification
	Qualification *QualificationRef `json:"qualification,omitempty"`

	// quote
	Quote *QuoteRef `json:"quote,omitempty"`

	// Identifier of then reference order
	RefOrderId *QuoteRef `json:"refOrderId,omitempty"`

	// Identifier of then reference order item
	RefOrderItemId *QuoteRef `json:"refOrderItemId,omitempty"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty,omitempty"`
}

// Validate validates this product order item create
func (m *ProductOrderItemCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItemPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItemRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductOffering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQualification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductOrderItemCreate) validateAction(formats strfmt.Registry) error {

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *ProductOrderItemCreate) validateBillingAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingAccount) { // not required
		return nil
	}

	if m.BillingAccount != nil {
		if err := m.BillingAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingAccount")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOrderItemCreate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderItemCreate) validateOrderItemPrice(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderItemPrice) { // not required
		return nil
	}

	iOrderItemPriceSize := int64(len(m.OrderItemPrice))

	if err := validate.MinItems("orderItemPrice", "body", iOrderItemPriceSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.OrderItemPrice); i++ {
		if swag.IsZero(m.OrderItemPrice[i]) { // not required
			continue
		}

		if m.OrderItemPrice[i] != nil {
			if err := m.OrderItemPrice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderItemPrice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOrderItemCreate) validateOrderItemRelationship(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderItemRelationship) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderItemRelationship); i++ {
		if swag.IsZero(m.OrderItemRelationship[i]) { // not required
			continue
		}

		if m.OrderItemRelationship[i] != nil {
			if err := m.OrderItemRelationship[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderItemRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOrderItemCreate) validatePricingMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.PricingMethod) { // not required
		return nil
	}

	if err := m.PricingMethod.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pricingMethod")
		}
		return err
	}

	return nil
}

func (m *ProductOrderItemCreate) validatePricingTerm(formats strfmt.Registry) error {

	if swag.IsZero(m.PricingTerm) { // not required
		return nil
	}

	if err := validate.MinimumInt("pricingTerm", "body", int64(*m.PricingTerm), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ProductOrderItemCreate) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOrderItemCreate) validateProductOffering(formats strfmt.Registry) error {

	if err := validate.Required("productOffering", "body", m.ProductOffering); err != nil {
		return err
	}

	if m.ProductOffering != nil {
		if err := m.ProductOffering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productOffering")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOrderItemCreate) validateQualification(formats strfmt.Registry) error {

	if swag.IsZero(m.Qualification) { // not required
		return nil
	}

	if m.Qualification != nil {
		if err := m.Qualification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qualification")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOrderItemCreate) validateQuote(formats strfmt.Registry) error {

	if swag.IsZero(m.Quote) { // not required
		return nil
	}

	if m.Quote != nil {
		if err := m.Quote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quote")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOrderItemCreate) validateRelatedParty(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedParty) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedParty); i++ {
		if swag.IsZero(m.RelatedParty[i]) { // not required
			continue
		}

		if m.RelatedParty[i] != nil {
			if err := m.RelatedParty[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductOrderItemCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductOrderItemCreate) UnmarshalBinary(b []byte) error {
	var res ProductOrderItemCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
