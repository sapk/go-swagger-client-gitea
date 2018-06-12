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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrgGetHookParams creates a new OrgGetHookParams object
// with the default values initialized.
func NewOrgGetHookParams() *OrgGetHookParams {
	var ()
	return &OrgGetHookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgGetHookParamsWithTimeout creates a new OrgGetHookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgGetHookParamsWithTimeout(timeout time.Duration) *OrgGetHookParams {
	var ()
	return &OrgGetHookParams{

		timeout: timeout,
	}
}

// NewOrgGetHookParamsWithContext creates a new OrgGetHookParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgGetHookParamsWithContext(ctx context.Context) *OrgGetHookParams {
	var ()
	return &OrgGetHookParams{

		Context: ctx,
	}
}

// NewOrgGetHookParamsWithHTTPClient creates a new OrgGetHookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgGetHookParamsWithHTTPClient(client *http.Client) *OrgGetHookParams {
	var ()
	return &OrgGetHookParams{
		HTTPClient: client,
	}
}

/*OrgGetHookParams contains all the parameters to send to the API endpoint
for the org get hook operation typically these are written to a http.Request
*/
type OrgGetHookParams struct {

	/*ID
	  id of the hook to get

	*/
	ID int64
	/*Org
	  name of the organization

	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org get hook params
func (o *OrgGetHookParams) WithTimeout(timeout time.Duration) *OrgGetHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org get hook params
func (o *OrgGetHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org get hook params
func (o *OrgGetHookParams) WithContext(ctx context.Context) *OrgGetHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org get hook params
func (o *OrgGetHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org get hook params
func (o *OrgGetHookParams) WithHTTPClient(client *http.Client) *OrgGetHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org get hook params
func (o *OrgGetHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the org get hook params
func (o *OrgGetHookParams) WithID(id int64) *OrgGetHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the org get hook params
func (o *OrgGetHookParams) SetID(id int64) {
	o.ID = id
}

// WithOrg adds the org to the org get hook params
func (o *OrgGetHookParams) WithOrg(org string) *OrgGetHookParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org get hook params
func (o *OrgGetHookParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *OrgGetHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
