// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new store API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for store API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateStore(params *CreateStoreParams, opts ...ClientOption) (*CreateStoreOK, error)

	GetStore(params *GetStoreParams, opts ...ClientOption) (*GetStoreOK, error)

	GetStores(params *GetStoresParams, opts ...ClientOption) (*GetStoresOK, error)

	RebrandStore(params *RebrandStoreParams, opts ...ClientOption) (*RebrandStoreOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateStore creates new store

Use this API to create a new store
*/
func (a *Client) CreateStore(params *CreateStoreParams, opts ...ClientOption) (*CreateStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createStore",
		Method:             "POST",
		PathPattern:        "/v1/api/stores/create_store",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateStoreDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStore gets a store

Use this API to get a store
*/
func (a *Client) GetStore(params *GetStoreParams, opts ...ClientOption) (*GetStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStore",
		Method:             "POST",
		PathPattern:        "/v1/api/stores/get_store",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoreDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStores gets stores

Use this API to get stores
*/
func (a *Client) GetStores(params *GetStoresParams, opts ...ClientOption) (*GetStoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStores",
		Method:             "POST",
		PathPattern:        "/v1/api/stores/get_stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStoresReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStoresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RebrandStore rebrands a store

Use this API to rebrand a store
*/
func (a *Client) RebrandStore(params *RebrandStoreParams, opts ...ClientOption) (*RebrandStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebrandStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "rebrandStore",
		Method:             "POST",
		PathPattern:        "/v1/api/stores/rebrand_store",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RebrandStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RebrandStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RebrandStoreDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
