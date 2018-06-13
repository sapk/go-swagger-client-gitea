// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewAdminEditUserParams creates a new AdminEditUserParams object
// with the default values initialized.
func NewAdminEditUserParams() *AdminEditUserParams {
	var ()
	return &AdminEditUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminEditUserParamsWithTimeout creates a new AdminEditUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminEditUserParamsWithTimeout(timeout time.Duration) *AdminEditUserParams {
	var ()
	return &AdminEditUserParams{

		timeout: timeout,
	}
}

// NewAdminEditUserParamsWithContext creates a new AdminEditUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminEditUserParamsWithContext(ctx context.Context) *AdminEditUserParams {
	var ()
	return &AdminEditUserParams{

		Context: ctx,
	}
}

// NewAdminEditUserParamsWithHTTPClient creates a new AdminEditUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminEditUserParamsWithHTTPClient(client *http.Client) *AdminEditUserParams {
	var ()
	return &AdminEditUserParams{
		HTTPClient: client,
	}
}

/*AdminEditUserParams contains all the parameters to send to the API endpoint
for the admin edit user operation typically these are written to a http.Request
*/
type AdminEditUserParams struct {

	/*Body*/
	Body AdminEditUserBody
	/*Username
	  username of user to edit

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin edit user params
func (o *AdminEditUserParams) WithTimeout(timeout time.Duration) *AdminEditUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin edit user params
func (o *AdminEditUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin edit user params
func (o *AdminEditUserParams) WithContext(ctx context.Context) *AdminEditUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin edit user params
func (o *AdminEditUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin edit user params
func (o *AdminEditUserParams) WithHTTPClient(client *http.Client) *AdminEditUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin edit user params
func (o *AdminEditUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin edit user params
func (o *AdminEditUserParams) WithBody(body AdminEditUserBody) *AdminEditUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin edit user params
func (o *AdminEditUserParams) SetBody(body AdminEditUserBody) {
	o.Body = body
}

// WithUsername adds the username to the admin edit user params
func (o *AdminEditUserParams) WithUsername(username string) *AdminEditUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the admin edit user params
func (o *AdminEditUserParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *AdminEditUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
