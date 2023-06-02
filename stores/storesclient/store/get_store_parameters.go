// Code generated by go-swagger; DO NOT EDIT.

package store

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

	"github.com/rezaAmiri123/microservice/stores/storesclient/models"
)

// NewGetStoreParams creates a new GetStoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStoreParams() *GetStoreParams {
	return &GetStoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoreParamsWithTimeout creates a new GetStoreParams object
// with the ability to set a timeout on a request.
func NewGetStoreParamsWithTimeout(timeout time.Duration) *GetStoreParams {
	return &GetStoreParams{
		timeout: timeout,
	}
}

// NewGetStoreParamsWithContext creates a new GetStoreParams object
// with the ability to set a context for a request.
func NewGetStoreParamsWithContext(ctx context.Context) *GetStoreParams {
	return &GetStoreParams{
		Context: ctx,
	}
}

// NewGetStoreParamsWithHTTPClient creates a new GetStoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStoreParamsWithHTTPClient(client *http.Client) *GetStoreParams {
	return &GetStoreParams{
		HTTPClient: client,
	}
}

/*
GetStoreParams contains all the parameters to send to the API endpoint

	for the get store operation.

	Typically these are written to a http.Request.
*/
type GetStoreParams struct {

	// Body.
	Body *models.StorespbGetStoreRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStoreParams) WithDefaults() *GetStoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get store params
func (o *GetStoreParams) WithTimeout(timeout time.Duration) *GetStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get store params
func (o *GetStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get store params
func (o *GetStoreParams) WithContext(ctx context.Context) *GetStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get store params
func (o *GetStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get store params
func (o *GetStoreParams) WithHTTPClient(client *http.Client) *GetStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get store params
func (o *GetStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get store params
func (o *GetStoreParams) WithBody(body *models.StorespbGetStoreRequest) *GetStoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get store params
func (o *GetStoreParams) SetBody(body *models.StorespbGetStoreRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
