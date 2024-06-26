// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCreated, error)

	GetID(params *GetIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Create creates a new customer to track
*/
func (a *Client) Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create",
		Method:             "POST",
		PathPattern:        "/customers",
		ProducesMediaTypes: []string{"application/keyauth.api.v1+json"},
		ConsumesMediaTypes: []string{"application/keyauth.api.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*CreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetID gets a customer Id given an s s n
*/
func (a *Client) GetID(params *GetIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getId",
		Method:             "GET",
		PathPattern:        "/customers",
		ProducesMediaTypes: []string{"application/keyauth.api.v1+json"},
		ConsumesMediaTypes: []string{"application/keyauth.api.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIDReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
