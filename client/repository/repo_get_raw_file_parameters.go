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

// NewRepoGetRawFileParams creates a new RepoGetRawFileParams object
// with the default values initialized.
func NewRepoGetRawFileParams() *RepoGetRawFileParams {
	var ()
	return &RepoGetRawFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetRawFileParamsWithTimeout creates a new RepoGetRawFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoGetRawFileParamsWithTimeout(timeout time.Duration) *RepoGetRawFileParams {
	var ()
	return &RepoGetRawFileParams{

		timeout: timeout,
	}
}

// NewRepoGetRawFileParamsWithContext creates a new RepoGetRawFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoGetRawFileParamsWithContext(ctx context.Context) *RepoGetRawFileParams {
	var ()
	return &RepoGetRawFileParams{

		Context: ctx,
	}
}

// NewRepoGetRawFileParamsWithHTTPClient creates a new RepoGetRawFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoGetRawFileParamsWithHTTPClient(client *http.Client) *RepoGetRawFileParams {
	var ()
	return &RepoGetRawFileParams{
		HTTPClient: client,
	}
}

/*RepoGetRawFileParams contains all the parameters to send to the API endpoint
for the repo get raw file operation typically these are written to a http.Request
*/
type RepoGetRawFileParams struct {

	/*Filepath
	  filepath of the file to get

	*/
	Filepath string
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

// WithTimeout adds the timeout to the repo get raw file params
func (o *RepoGetRawFileParams) WithTimeout(timeout time.Duration) *RepoGetRawFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get raw file params
func (o *RepoGetRawFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get raw file params
func (o *RepoGetRawFileParams) WithContext(ctx context.Context) *RepoGetRawFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get raw file params
func (o *RepoGetRawFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get raw file params
func (o *RepoGetRawFileParams) WithHTTPClient(client *http.Client) *RepoGetRawFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get raw file params
func (o *RepoGetRawFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilepath adds the filepath to the repo get raw file params
func (o *RepoGetRawFileParams) WithFilepath(filepath string) *RepoGetRawFileParams {
	o.SetFilepath(filepath)
	return o
}

// SetFilepath adds the filepath to the repo get raw file params
func (o *RepoGetRawFileParams) SetFilepath(filepath string) {
	o.Filepath = filepath
}

// WithOwner adds the owner to the repo get raw file params
func (o *RepoGetRawFileParams) WithOwner(owner string) *RepoGetRawFileParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get raw file params
func (o *RepoGetRawFileParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get raw file params
func (o *RepoGetRawFileParams) WithRepo(repo string) *RepoGetRawFileParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get raw file params
func (o *RepoGetRawFileParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetRawFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param filepath
	if err := r.SetPathParam("filepath", o.Filepath); err != nil {
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
