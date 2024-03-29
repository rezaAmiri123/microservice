// Code generated by go-swagger; DO NOT EDIT.

package basket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new basket API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for basket API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelBasket(params *CancelBasketParams, opts ...ClientOption) (*CancelBasketOK, error)

	CheckoutBasket(params *CheckoutBasketParams, opts ...ClientOption) (*CheckoutBasketOK, error)

	GetBasket(params *GetBasketParams, opts ...ClientOption) (*GetBasketOK, error)

	StartBasket(params *StartBasketParams, opts ...ClientOption) (*StartBasketOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CancelBasket cancels a basket

Use this API to cancel a basket
*/
func (a *Client) CancelBasket(params *CancelBasketParams, opts ...ClientOption) (*CancelBasketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelBasketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancelBasket",
		Method:             "POST",
		PathPattern:        "/v1/api/baskets/cancel_basket",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CancelBasketReader{formats: a.formats},
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
	success, ok := result.(*CancelBasketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CancelBasketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CheckoutBasket checkouts a basket

Use this API to checkout a basket
*/
func (a *Client) CheckoutBasket(params *CheckoutBasketParams, opts ...ClientOption) (*CheckoutBasketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckoutBasketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "checkoutBasket",
		Method:             "POST",
		PathPattern:        "/v1/api/baskets/checkout_basket",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckoutBasketReader{formats: a.formats},
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
	success, ok := result.(*CheckoutBasketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CheckoutBasketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBasket gets a basket

Use this API to get a basket
*/
func (a *Client) GetBasket(params *GetBasketParams, opts ...ClientOption) (*GetBasketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBasketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBasket",
		Method:             "POST",
		PathPattern:        "/v1/api/baskets/get_basket",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBasketReader{formats: a.formats},
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
	success, ok := result.(*GetBasketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBasketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StartBasket creates new basket

Use this API to start a new basket
*/
func (a *Client) StartBasket(params *StartBasketParams, opts ...ClientOption) (*StartBasketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartBasketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startBasket",
		Method:             "POST",
		PathPattern:        "/v1/api/baskets/start_basket",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartBasketReader{formats: a.formats},
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
	success, ok := result.(*StartBasketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartBasketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
