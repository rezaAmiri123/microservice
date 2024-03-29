// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserspbLoginUserRequest userspb login user request
//
// swagger:model userspbLoginUserRequest
type UserspbLoginUserRequest struct {

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this userspb login user request
func (m *UserspbLoginUserRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this userspb login user request based on context it is used
func (m *UserspbLoginUserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserspbLoginUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserspbLoginUserRequest) UnmarshalBinary(b []byte) error {
	var res UserspbLoginUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
