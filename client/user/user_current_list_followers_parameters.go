// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserCurrentListFollowersParams creates a new UserCurrentListFollowersParams object
// with the default values initialized.
func NewUserCurrentListFollowersParams() *UserCurrentListFollowersParams {

	return &UserCurrentListFollowersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentListFollowersParamsWithTimeout creates a new UserCurrentListFollowersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCurrentListFollowersParamsWithTimeout(timeout time.Duration) *UserCurrentListFollowersParams {

	return &UserCurrentListFollowersParams{

		timeout: timeout,
	}
}

// NewUserCurrentListFollowersParamsWithContext creates a new UserCurrentListFollowersParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCurrentListFollowersParamsWithContext(ctx context.Context) *UserCurrentListFollowersParams {

	return &UserCurrentListFollowersParams{

		Context: ctx,
	}
}

// NewUserCurrentListFollowersParamsWithHTTPClient creates a new UserCurrentListFollowersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCurrentListFollowersParamsWithHTTPClient(client *http.Client) *UserCurrentListFollowersParams {

	return &UserCurrentListFollowersParams{
		HTTPClient: client,
	}
}

/*UserCurrentListFollowersParams contains all the parameters to send to the API endpoint
for the user current list followers operation typically these are written to a http.Request
*/
type UserCurrentListFollowersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user current list followers params
func (o *UserCurrentListFollowersParams) WithTimeout(timeout time.Duration) *UserCurrentListFollowersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current list followers params
func (o *UserCurrentListFollowersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current list followers params
func (o *UserCurrentListFollowersParams) WithContext(ctx context.Context) *UserCurrentListFollowersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current list followers params
func (o *UserCurrentListFollowersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current list followers params
func (o *UserCurrentListFollowersParams) WithHTTPClient(client *http.Client) *UserCurrentListFollowersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current list followers params
func (o *UserCurrentListFollowersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentListFollowersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
