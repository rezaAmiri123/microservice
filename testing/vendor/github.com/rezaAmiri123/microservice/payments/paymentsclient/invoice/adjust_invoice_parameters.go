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

// NewAdjustInvoiceParams creates a new AdjustInvoiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdjustInvoiceParams() *AdjustInvoiceParams {
	return &AdjustInvoiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdjustInvoiceParamsWithTimeout creates a new AdjustInvoiceParams object
// with the ability to set a timeout on a request.
func NewAdjustInvoiceParamsWithTimeout(timeout time.Duration) *AdjustInvoiceParams {
	return &AdjustInvoiceParams{
		timeout: timeout,
	}
}

// NewAdjustInvoiceParamsWithContext creates a new AdjustInvoiceParams object
// with the ability to set a context for a request.
func NewAdjustInvoiceParamsWithContext(ctx context.Context) *AdjustInvoiceParams {
	return &AdjustInvoiceParams{
		Context: ctx,
	}
}

// NewAdjustInvoiceParamsWithHTTPClient creates a new AdjustInvoiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdjustInvoiceParamsWithHTTPClient(client *http.Client) *AdjustInvoiceParams {
	return &AdjustInvoiceParams{
		HTTPClient: client,
	}
}

/*
AdjustInvoiceParams contains all the parameters to send to the API endpoint

	for the adjust invoice operation.

	Typically these are written to a http.Request.
*/
type AdjustInvoiceParams struct {

	// Body.
	Body *models.PaymentspbAdjustInvoiceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the adjust invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdjustInvoiceParams) WithDefaults() *AdjustInvoiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the adjust invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdjustInvoiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the adjust invoice params
func (o *AdjustInvoiceParams) WithTimeout(timeout time.Duration) *AdjustInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the adjust invoice params
func (o *AdjustInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the adjust invoice params
func (o *AdjustInvoiceParams) WithContext(ctx context.Context) *AdjustInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the adjust invoice params
func (o *AdjustInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the adjust invoice params
func (o *AdjustInvoiceParams) WithHTTPClient(client *http.Client) *AdjustInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the adjust invoice params
func (o *AdjustInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the adjust invoice params
func (o *AdjustInvoiceParams) WithBody(body *models.PaymentspbAdjustInvoiceRequest) *AdjustInvoiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the adjust invoice params
func (o *AdjustInvoiceParams) SetBody(body *models.PaymentspbAdjustInvoiceRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AdjustInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
