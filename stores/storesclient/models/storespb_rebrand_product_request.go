// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorespbRebrandProductRequest storespb rebrand product request
//
// swagger:model storespbRebrandProductRequest
type StorespbRebrandProductRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this storespb rebrand product request
func (m *StorespbRebrandProductRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storespb rebrand product request based on context it is used
func (m *StorespbRebrandProductRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorespbRebrandProductRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorespbRebrandProductRequest) UnmarshalBinary(b []byte) error {
	var res StorespbRebrandProductRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
