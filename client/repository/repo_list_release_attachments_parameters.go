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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRepoListReleaseAttachmentsParams creates a new RepoListReleaseAttachmentsParams object
// with the default values initialized.
func NewRepoListReleaseAttachmentsParams() *RepoListReleaseAttachmentsParams {
	var ()
	return &RepoListReleaseAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListReleaseAttachmentsParamsWithTimeout creates a new RepoListReleaseAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoListReleaseAttachmentsParamsWithTimeout(timeout time.Duration) *RepoListReleaseAttachmentsParams {
	var ()
	return &RepoListReleaseAttachmentsParams{

		timeout: timeout,
	}
}

// NewRepoListReleaseAttachmentsParamsWithContext creates a new RepoListReleaseAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoListReleaseAttachmentsParamsWithContext(ctx context.Context) *RepoListReleaseAttachmentsParams {
	var ()
	return &RepoListReleaseAttachmentsParams{

		Context: ctx,
	}
}

// NewRepoListReleaseAttachmentsParamsWithHTTPClient creates a new RepoListReleaseAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoListReleaseAttachmentsParamsWithHTTPClient(client *http.Client) *RepoListReleaseAttachmentsParams {
	var ()
	return &RepoListReleaseAttachmentsParams{
		HTTPClient: client,
	}
}

/*RepoListReleaseAttachmentsParams contains all the parameters to send to the API endpoint
for the repo list release attachments operation typically these are written to a http.Request
*/
type RepoListReleaseAttachmentsParams struct {

	/*ID
	  id of the release

	*/
	ID int64
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

// WithTimeout adds the timeout to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) WithTimeout(timeout time.Duration) *RepoListReleaseAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) WithContext(ctx context.Context) *RepoListReleaseAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) WithHTTPClient(client *http.Client) *RepoListReleaseAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) WithID(id int64) *RepoListReleaseAttachmentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) WithOwner(owner string) *RepoListReleaseAttachmentsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) WithRepo(repo string) *RepoListReleaseAttachmentsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list release attachments params
func (o *RepoListReleaseAttachmentsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListReleaseAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

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
