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

// NewOrgConcealMemberParams creates a new OrgConcealMemberParams object
// with the default values initialized.
func NewOrgConcealMemberParams() *OrgConcealMemberParams {
	var ()
	return &OrgConcealMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgConcealMemberParamsWithTimeout creates a new OrgConcealMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgConcealMemberParamsWithTimeout(timeout time.Duration) *OrgConcealMemberParams {
	var ()
	return &OrgConcealMemberParams{

		timeout: timeout,
	}
}

// NewOrgConcealMemberParamsWithContext creates a new OrgConcealMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgConcealMemberParamsWithContext(ctx context.Context) *OrgConcealMemberParams {
	var ()
	return &OrgConcealMemberParams{

		Context: ctx,
	}
}

// NewOrgConcealMemberParamsWithHTTPClient creates a new OrgConcealMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgConcealMemberParamsWithHTTPClient(client *http.Client) *OrgConcealMemberParams {
	var ()
	return &OrgConcealMemberParams{
		HTTPClient: client,
	}
}

/*OrgConcealMemberParams contains all the parameters to send to the API endpoint
for the org conceal member operation typically these are written to a http.Request
*/
type OrgConcealMemberParams struct {

	/*Org
	  name of the organization

	*/
	Org string
	/*Username
	  username of the user

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org conceal member params
func (o *OrgConcealMemberParams) WithTimeout(timeout time.Duration) *OrgConcealMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org conceal member params
func (o *OrgConcealMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org conceal member params
func (o *OrgConcealMemberParams) WithContext(ctx context.Context) *OrgConcealMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org conceal member params
func (o *OrgConcealMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org conceal member params
func (o *OrgConcealMemberParams) WithHTTPClient(client *http.Client) *OrgConcealMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org conceal member params
func (o *OrgConcealMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the org conceal member params
func (o *OrgConcealMemberParams) WithOrg(org string) *OrgConcealMemberParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org conceal member params
func (o *OrgConcealMemberParams) SetOrg(org string) {
	o.Org = org
}

// WithUsername adds the username to the org conceal member params
func (o *OrgConcealMemberParams) WithUsername(username string) *OrgConcealMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the org conceal member params
func (o *OrgConcealMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *OrgConcealMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
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