// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorespbDisableParticipationRequest storespb disable participation request
//
// swagger:model storespbDisableParticipationRequest
type StorespbDisableParticipationRequest struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this storespb disable participation request
func (m *StorespbDisableParticipationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storespb disable participation request based on context it is used
func (m *StorespbDisableParticipationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorespbDisableParticipationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorespbDisableParticipationRequest) UnmarshalBinary(b []byte) error {
	var res StorespbDisableParticipationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
