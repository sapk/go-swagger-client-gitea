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

// NewUserCurrentListKeysParams creates a new UserCurrentListKeysParams object
// with the default values initialized.
func NewUserCurrentListKeysParams() *UserCurrentListKeysParams {

	return &UserCurrentListKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentListKeysParamsWithTimeout creates a new UserCurrentListKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCurrentListKeysParamsWithTimeout(timeout time.Duration) *UserCurrentListKeysParams {

	return &UserCurrentListKeysParams{

		timeout: timeout,
	}
}

// NewUserCurrentListKeysParamsWithContext creates a new UserCurrentListKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCurrentListKeysParamsWithContext(ctx context.Context) *UserCurrentListKeysParams {

	return &UserCurrentListKeysParams{

		Context: ctx,
	}
}

// NewUserCurrentListKeysParamsWithHTTPClient creates a new UserCurrentListKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCurrentListKeysParamsWithHTTPClient(client *http.Client) *UserCurrentListKeysParams {

	return &UserCurrentListKeysParams{
		HTTPClient: client,
	}
}

/*UserCurrentListKeysParams contains all the parameters to send to the API endpoint
for the user current list keys operation typically these are written to a http.Request
*/
type UserCurrentListKeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user current list keys params
func (o *UserCurrentListKeysParams) WithTimeout(timeout time.Duration) *UserCurrentListKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current list keys params
func (o *UserCurrentListKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current list keys params
func (o *UserCurrentListKeysParams) WithContext(ctx context.Context) *UserCurrentListKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current list keys params
func (o *UserCurrentListKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current list keys params
func (o *UserCurrentListKeysParams) WithHTTPClient(client *http.Client) *UserCurrentListKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current list keys params
func (o *UserCurrentListKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentListKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}