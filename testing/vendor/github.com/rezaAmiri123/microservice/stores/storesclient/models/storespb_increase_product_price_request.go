// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorespbIncreaseProductPriceRequest storespb increase product price request
//
// swagger:model storespbIncreaseProductPriceRequest
type StorespbIncreaseProductPriceRequest struct {

	// id
	ID string `json:"id,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`
}

// Validate validates this storespb increase product price request
func (m *StorespbIncreaseProductPriceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storespb increase product price request based on context it is used
func (m *StorespbIncreaseProductPriceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorespbIncreaseProductPriceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorespbIncreaseProductPriceRequest) UnmarshalBinary(b []byte) error {
	var res StorespbIncreaseProductPriceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
