// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateCurrentUserRepo creates a repository
*/
func (a *Client) CreateCurrentUserRepo(params *CreateCurrentUserRepoParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCurrentUserRepoCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCurrentUserRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCurrentUserRepo",
		Method:             "POST",
		PathPattern:        "/user/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCurrentUserRepoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateCurrentUserRepoCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
