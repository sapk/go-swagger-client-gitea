// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewOrgListUserOrgsParams creates a new OrgListUserOrgsParams object
// with the default values initialized.
func NewOrgListUserOrgsParams() *OrgListUserOrgsParams {
	var ()
	return &OrgListUserOrgsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListUserOrgsParamsWithTimeout creates a new OrgListUserOrgsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgListUserOrgsParamsWithTimeout(timeout time.Duration) *OrgListUserOrgsParams {
	var ()
	return &OrgListUserOrgsParams{

		timeout: timeout,
	}
}

// NewOrgListUserOrgsParamsWithContext creates a new OrgListUserOrgsParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgListUserOrgsParamsWithContext(ctx context.Context) *OrgListUserOrgsParams {
	var ()
	return &OrgListUserOrgsParams{

		Context: ctx,
	}
}

// NewOrgListUserOrgsParamsWithHTTPClient creates a new OrgListUserOrgsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgListUserOrgsParamsWithHTTPClient(client *http.Client) *OrgListUserOrgsParams {
	var ()
	return &OrgListUserOrgsParams{
		HTTPClient: client,
	}
}

/*OrgListUserOrgsParams contains all the parameters to send to the API endpoint
for the org list user orgs operation typically these are written to a http.Request
*/
type OrgListUserOrgsParams struct {

	/*Username
	  username of user

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org list user orgs params
func (o *OrgListUserOrgsParams) WithTimeout(timeout time.Duration) *OrgListUserOrgsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list user orgs params
func (o *OrgListUserOrgsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list user orgs params
func (o *OrgListUserOrgsParams) WithContext(ctx context.Context) *OrgListUserOrgsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list user orgs params
func (o *OrgListUserOrgsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list user orgs params
func (o *OrgListUserOrgsParams) WithHTTPClient(client *http.Client) *OrgListUserOrgsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list user orgs params
func (o *OrgListUserOrgsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the org list user orgs params
func (o *OrgListUserOrgsParams) WithUsername(username string) *OrgListUserOrgsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the org list user orgs params
func (o *OrgListUserOrgsParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListUserOrgsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
