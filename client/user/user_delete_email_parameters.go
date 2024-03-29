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

// NewUserDeleteEmailParams creates a new UserDeleteEmailParams object
// with the default values initialized.
func NewUserDeleteEmailParams() *UserDeleteEmailParams {
	var ()
	return &UserDeleteEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserDeleteEmailParamsWithTimeout creates a new UserDeleteEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserDeleteEmailParamsWithTimeout(timeout time.Duration) *UserDeleteEmailParams {
	var ()
	return &UserDeleteEmailParams{

		timeout: timeout,
	}
}

// NewUserDeleteEmailParamsWithContext creates a new UserDeleteEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserDeleteEmailParamsWithContext(ctx context.Context) *UserDeleteEmailParams {
	var ()
	return &UserDeleteEmailParams{

		Context: ctx,
	}
}

// NewUserDeleteEmailParamsWithHTTPClient creates a new UserDeleteEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserDeleteEmailParamsWithHTTPClient(client *http.Client) *UserDeleteEmailParams {
	var ()
	return &UserDeleteEmailParams{
		HTTPClient: client,
	}
}

/*UserDeleteEmailParams contains all the parameters to send to the API endpoint
for the user delete email operation typically these are written to a http.Request
*/
type UserDeleteEmailParams struct {

	/*Body*/
	Body UserDeleteEmailBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user delete email params
func (o *UserDeleteEmailParams) WithTimeout(timeout time.Duration) *UserDeleteEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user delete email params
func (o *UserDeleteEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user delete email params
func (o *UserDeleteEmailParams) WithContext(ctx context.Context) *UserDeleteEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user delete email params
func (o *UserDeleteEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user delete email params
func (o *UserDeleteEmailParams) WithHTTPClient(client *http.Client) *UserDeleteEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user delete email params
func (o *UserDeleteEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user delete email params
func (o *UserDeleteEmailParams) WithBody(body UserDeleteEmailBody) *UserDeleteEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user delete email params
func (o *UserDeleteEmailParams) SetBody(body UserDeleteEmailBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UserDeleteEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
