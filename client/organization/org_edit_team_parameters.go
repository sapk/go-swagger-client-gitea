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

// NewOrgEditTeamParams creates a new OrgEditTeamParams object
// with the default values initialized.
func NewOrgEditTeamParams() *OrgEditTeamParams {
	var ()
	return &OrgEditTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgEditTeamParamsWithTimeout creates a new OrgEditTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgEditTeamParamsWithTimeout(timeout time.Duration) *OrgEditTeamParams {
	var ()
	return &OrgEditTeamParams{

		timeout: timeout,
	}
}

// NewOrgEditTeamParamsWithContext creates a new OrgEditTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgEditTeamParamsWithContext(ctx context.Context) *OrgEditTeamParams {
	var ()
	return &OrgEditTeamParams{

		Context: ctx,
	}
}

// NewOrgEditTeamParamsWithHTTPClient creates a new OrgEditTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgEditTeamParamsWithHTTPClient(client *http.Client) *OrgEditTeamParams {
	var ()
	return &OrgEditTeamParams{
		HTTPClient: client,
	}
}

/*OrgEditTeamParams contains all the parameters to send to the API endpoint
for the org edit team operation typically these are written to a http.Request
*/
type OrgEditTeamParams struct {

	/*Body*/
	Body OrgEditTeamBody
	/*ID
	  id of the team to edit

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org edit team params
func (o *OrgEditTeamParams) WithTimeout(timeout time.Duration) *OrgEditTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org edit team params
func (o *OrgEditTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org edit team params
func (o *OrgEditTeamParams) WithContext(ctx context.Context) *OrgEditTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org edit team params
func (o *OrgEditTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org edit team params
func (o *OrgEditTeamParams) WithHTTPClient(client *http.Client) *OrgEditTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org edit team params
func (o *OrgEditTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the org edit team params
func (o *OrgEditTeamParams) WithBody(body OrgEditTeamBody) *OrgEditTeamParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the org edit team params
func (o *OrgEditTeamParams) SetBody(body OrgEditTeamBody) {
	o.Body = body
}

// WithID adds the id to the org edit team params
func (o *OrgEditTeamParams) WithID(id int64) *OrgEditTeamParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the org edit team params
func (o *OrgEditTeamParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrgEditTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
