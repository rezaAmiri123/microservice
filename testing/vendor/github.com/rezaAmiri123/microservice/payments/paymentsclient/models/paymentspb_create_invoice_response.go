// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentspbCreateInvoiceResponse paymentspb create invoice response
//
// swagger:model paymentspbCreateInvoiceResponse
type PaymentspbCreateInvoiceResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this paymentspb create invoice response
func (m *PaymentspbCreateInvoiceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this paymentspb create invoice response based on context it is used
func (m *PaymentspbCreateInvoiceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentspbCreateInvoiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentspbCreateInvoiceResponse) UnmarshalBinary(b []byte) error {
	var res PaymentspbCreateInvoiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}