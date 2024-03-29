// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorespbAddProductRequest storespb add product request
//
// swagger:model storespbAddProductRequest
type StorespbAddProductRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// sku
	Sku string `json:"sku,omitempty"`

	// store Id
	StoreID string `json:"storeId,omitempty"`
}

// Validate validates this storespb add product request
func (m *StorespbAddProductRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storespb add product request based on context it is used
func (m *StorespbAddProductRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorespbAddProductRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorespbAddProductRequest) UnmarshalBinary(b []byte) error {
	var res StorespbAddProductRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
