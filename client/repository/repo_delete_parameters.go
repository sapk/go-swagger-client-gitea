// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewRepoDeleteParams creates a new RepoDeleteParams object
// with the default values initialized.
func NewRepoDeleteParams() *RepoDeleteParams {
	var ()
	return &RepoDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeleteParamsWithTimeout creates a new RepoDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoDeleteParamsWithTimeout(timeout time.Duration) *RepoDeleteParams {
	var ()
	return &RepoDeleteParams{

		timeout: timeout,
	}
}

// NewRepoDeleteParamsWithContext creates a new RepoDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoDeleteParamsWithContext(ctx context.Context) *RepoDeleteParams {
	var ()
	return &RepoDeleteParams{

		Context: ctx,
	}
}

// NewRepoDeleteParamsWithHTTPClient creates a new RepoDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoDeleteParamsWithHTTPClient(client *http.Client) *RepoDeleteParams {
	var ()
	return &RepoDeleteParams{
		HTTPClient: client,
	}
}

/*RepoDeleteParams contains all the parameters to send to the API endpoint
for the repo delete operation typically these are written to a http.Request
*/
type RepoDeleteParams struct {

	/*Owner
	  owner of the repo to delete

	*/
	Owner string
	/*Repo
	  name of the repo to delete

	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repo delete params
func (o *RepoDeleteParams) WithTimeout(timeout time.Duration) *RepoDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete params
func (o *RepoDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete params
func (o *RepoDeleteParams) WithContext(ctx context.Context) *RepoDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete params
func (o *RepoDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete params
func (o *RepoDeleteParams) WithHTTPClient(client *http.Client) *RepoDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete params
func (o *RepoDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo delete params
func (o *RepoDeleteParams) WithOwner(owner string) *RepoDeleteParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete params
func (o *RepoDeleteParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete params
func (o *RepoDeleteParams) WithRepo(repo string) *RepoDeleteParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete params
func (o *RepoDeleteParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}