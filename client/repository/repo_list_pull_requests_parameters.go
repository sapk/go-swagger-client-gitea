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

// NewRepoListPullRequestsParams creates a new RepoListPullRequestsParams object
// with the default values initialized.
func NewRepoListPullRequestsParams() *RepoListPullRequestsParams {
	var ()
	return &RepoListPullRequestsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListPullRequestsParamsWithTimeout creates a new RepoListPullRequestsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoListPullRequestsParamsWithTimeout(timeout time.Duration) *RepoListPullRequestsParams {
	var ()
	return &RepoListPullRequestsParams{

		timeout: timeout,
	}
}

// NewRepoListPullRequestsParamsWithContext creates a new RepoListPullRequestsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoListPullRequestsParamsWithContext(ctx context.Context) *RepoListPullRequestsParams {
	var ()
	return &RepoListPullRequestsParams{

		Context: ctx,
	}
}

// NewRepoListPullRequestsParamsWithHTTPClient creates a new RepoListPullRequestsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoListPullRequestsParamsWithHTTPClient(client *http.Client) *RepoListPullRequestsParams {
	var ()
	return &RepoListPullRequestsParams{
		HTTPClient: client,
	}
}

/*RepoListPullRequestsParams contains all the parameters to send to the API endpoint
for the repo list pull requests operation typically these are written to a http.Request
*/
type RepoListPullRequestsParams struct {

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

// WithTimeout adds the timeout to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithTimeout(timeout time.Duration) *RepoListPullRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithContext(ctx context.Context) *RepoListPullRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithHTTPClient(client *http.Client) *RepoListPullRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithOwner(owner string) *RepoListPullRequestsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithRepo(repo string) *RepoListPullRequestsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListPullRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
