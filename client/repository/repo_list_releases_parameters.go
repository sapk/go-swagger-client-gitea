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

// NewRepoListReleasesParams creates a new RepoListReleasesParams object
// with the default values initialized.
func NewRepoListReleasesParams() *RepoListReleasesParams {
	var ()
	return &RepoListReleasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListReleasesParamsWithTimeout creates a new RepoListReleasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoListReleasesParamsWithTimeout(timeout time.Duration) *RepoListReleasesParams {
	var ()
	return &RepoListReleasesParams{

		timeout: timeout,
	}
}

// NewRepoListReleasesParamsWithContext creates a new RepoListReleasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoListReleasesParamsWithContext(ctx context.Context) *RepoListReleasesParams {
	var ()
	return &RepoListReleasesParams{

		Context: ctx,
	}
}

// NewRepoListReleasesParamsWithHTTPClient creates a new RepoListReleasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoListReleasesParamsWithHTTPClient(client *http.Client) *RepoListReleasesParams {
	var ()
	return &RepoListReleasesParams{
		HTTPClient: client,
	}
}

/*RepoListReleasesParams contains all the parameters to send to the API endpoint
for the repo list releases operation typically these are written to a http.Request
*/
type RepoListReleasesParams struct {

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

// WithTimeout adds the timeout to the repo list releases params
func (o *RepoListReleasesParams) WithTimeout(timeout time.Duration) *RepoListReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list releases params
func (o *RepoListReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list releases params
func (o *RepoListReleasesParams) WithContext(ctx context.Context) *RepoListReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list releases params
func (o *RepoListReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list releases params
func (o *RepoListReleasesParams) WithHTTPClient(client *http.Client) *RepoListReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list releases params
func (o *RepoListReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo list releases params
func (o *RepoListReleasesParams) WithOwner(owner string) *RepoListReleasesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list releases params
func (o *RepoListReleasesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list releases params
func (o *RepoListReleasesParams) WithRepo(repo string) *RepoListReleasesParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list releases params
func (o *RepoListReleasesParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
