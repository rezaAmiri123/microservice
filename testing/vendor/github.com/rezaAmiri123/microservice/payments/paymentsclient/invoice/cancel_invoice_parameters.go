// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/rezaAmiri123/microservice/payments/paymentsclient/models"
)

// NewCancelInvoiceParams creates a new CancelInvoiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelInvoiceParams() *CancelInvoiceParams {
	return &CancelInvoiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelInvoiceParamsWithTimeout creates a new CancelInvoiceParams object
// with the ability to set a timeout on a request.
func NewCancelInvoiceParamsWithTimeout(timeout time.Duration) *CancelInvoiceParams {
	return &CancelInvoiceParams{
		timeout: timeout,
	}
}

// NewCancelInvoiceParamsWithContext creates a new CancelInvoiceParams object
// with the ability to set a context for a request.
func NewCancelInvoiceParamsWithContext(ctx context.Context) *CancelInvoiceParams {
	return &CancelInvoiceParams{
		Context: ctx,
	}
}

// NewCancelInvoiceParamsWithHTTPClient creates a new CancelInvoiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelInvoiceParamsWithHTTPClient(client *http.Client) *CancelInvoiceParams {
	return &CancelInvoiceParams{
		HTTPClient: client,
	}
}

/*
CancelInvoiceParams contains all the parameters to send to the API endpoint

	for the cancel invoice operation.

	Typically these are written to a http.Request.
*/
type CancelInvoiceParams struct {

	// Body.
	Body *models.PaymentspbCancelInvoiceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelInvoiceParams) WithDefaults() *CancelInvoiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelInvoiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel invoice params
func (o *CancelInvoiceParams) WithTimeout(timeout time.Duration) *CancelInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel invoice params
func (o *CancelInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel invoice params
func (o *CancelInvoiceParams) WithContext(ctx context.Context) *CancelInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel invoice params
func (o *CancelInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel invoice params
func (o *CancelInvoiceParams) WithHTTPClient(client *http.Client) *CancelInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel invoice params
func (o *CancelInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cancel invoice params
func (o *CancelInvoiceParams) WithBody(body *models.PaymentspbCancelInvoiceRequest) *CancelInvoiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cancel invoice params
func (o *CancelInvoiceParams) SetBody(body *models.PaymentspbCancelInvoiceRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CancelInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
