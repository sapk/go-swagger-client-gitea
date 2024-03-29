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

// NewRepoTrackedTimesParams creates a new RepoTrackedTimesParams object
// with the default values initialized.
func NewRepoTrackedTimesParams() *RepoTrackedTimesParams {
	var ()
	return &RepoTrackedTimesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoTrackedTimesParamsWithTimeout creates a new RepoTrackedTimesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoTrackedTimesParamsWithTimeout(timeout time.Duration) *RepoTrackedTimesParams {
	var ()
	return &RepoTrackedTimesParams{

		timeout: timeout,
	}
}

// NewRepoTrackedTimesParamsWithContext creates a new RepoTrackedTimesParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoTrackedTimesParamsWithContext(ctx context.Context) *RepoTrackedTimesParams {
	var ()
	return &RepoTrackedTimesParams{

		Context: ctx,
	}
}

// NewRepoTrackedTimesParamsWithHTTPClient creates a new RepoTrackedTimesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoTrackedTimesParamsWithHTTPClient(client *http.Client) *RepoTrackedTimesParams {
	var ()
	return &RepoTrackedTimesParams{
		HTTPClient: client,
	}
}

/*RepoTrackedTimesParams contains all the parameters to send to the API endpoint
for the repo tracked times operation typically these are written to a http.Request
*/
type RepoTrackedTimesParams struct {

	/*Owner
	  owner of the repo

	*/
	Owner string
	/*Repo
	  name of the repo

	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repo tracked times params
func (o *RepoTrackedTimesParams) WithTimeout(timeout time.Duration) *RepoTrackedTimesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo tracked times params
func (o *RepoTrackedTimesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo tracked times params
func (o *RepoTrackedTimesParams) WithContext(ctx context.Context) *RepoTrackedTimesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo tracked times params
func (o *RepoTrackedTimesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo tracked times params
func (o *RepoTrackedTimesParams) WithHTTPClient(client *http.Client) *RepoTrackedTimesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo tracked times params
func (o *RepoTrackedTimesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo tracked times params
func (o *RepoTrackedTimesParams) WithOwner(owner string) *RepoTrackedTimesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo tracked times params
func (o *RepoTrackedTimesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo tracked times params
func (o *RepoTrackedTimesParams) WithRepo(repo string) *RepoTrackedTimesParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo tracked times params
func (o *RepoTrackedTimesParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoTrackedTimesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
