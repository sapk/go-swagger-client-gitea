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

// NewOrgListMembersParams creates a new OrgListMembersParams object
// with the default values initialized.
func NewOrgListMembersParams() *OrgListMembersParams {
	var ()
	return &OrgListMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListMembersParamsWithTimeout creates a new OrgListMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgListMembersParamsWithTimeout(timeout time.Duration) *OrgListMembersParams {
	var ()
	return &OrgListMembersParams{

		timeout: timeout,
	}
}

// NewOrgListMembersParamsWithContext creates a new OrgListMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgListMembersParamsWithContext(ctx context.Context) *OrgListMembersParams {
	var ()
	return &OrgListMembersParams{

		Context: ctx,
	}
}

// NewOrgListMembersParamsWithHTTPClient creates a new OrgListMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgListMembersParamsWithHTTPClient(client *http.Client) *OrgListMembersParams {
	var ()
	return &OrgListMembersParams{
		HTTPClient: client,
	}
}

/*OrgListMembersParams contains all the parameters to send to the API endpoint
for the org list members operation typically these are written to a http.Request
*/
type OrgListMembersParams struct {

	/*Org
	  name of the organization

	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org list members params
func (o *OrgListMembersParams) WithTimeout(timeout time.Duration) *OrgListMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list members params
func (o *OrgListMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list members params
func (o *OrgListMembersParams) WithContext(ctx context.Context) *OrgListMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list members params
func (o *OrgListMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list members params
func (o *OrgListMembersParams) WithHTTPClient(client *http.Client) *OrgListMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list members params
func (o *OrgListMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the org list members params
func (o *OrgListMembersParams) WithOrg(org string) *OrgListMembersParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org list members params
func (o *OrgListMembersParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
