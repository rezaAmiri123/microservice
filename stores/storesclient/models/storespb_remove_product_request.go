// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorespbRemoveProductRequest storespb remove product request
//
// swagger:model storespbRemoveProductRequest
type StorespbRemoveProductRequest struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this storespb remove product request
func (m *StorespbRemoveProductRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storespb remove product request based on context it is used
func (m *StorespbRemoveProductRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorespbRemoveProductRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorespbRemoveProductRequest) UnmarshalBinary(b []byte) error {
	var res StorespbRemoveProductRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
