// Code generated by go-swagger; DO NOT EDIT.

package issue

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

// NewIssueDeleteCommentDeprecatedParams creates a new IssueDeleteCommentDeprecatedParams object
// with the default values initialized.
func NewIssueDeleteCommentDeprecatedParams() *IssueDeleteCommentDeprecatedParams {
	var ()
	return &IssueDeleteCommentDeprecatedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIssueDeleteCommentDeprecatedParamsWithTimeout creates a new IssueDeleteCommentDeprecatedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIssueDeleteCommentDeprecatedParamsWithTimeout(timeout time.Duration) *IssueDeleteCommentDeprecatedParams {
	var ()
	return &IssueDeleteCommentDeprecatedParams{

		timeout: timeout,
	}
}

// NewIssueDeleteCommentDeprecatedParamsWithContext creates a new IssueDeleteCommentDeprecatedParams object
// with the default values initialized, and the ability to set a context for a request
func NewIssueDeleteCommentDeprecatedParamsWithContext(ctx context.Context) *IssueDeleteCommentDeprecatedParams {
	var ()
	return &IssueDeleteCommentDeprecatedParams{

		Context: ctx,
	}
}

// NewIssueDeleteCommentDeprecatedParamsWithHTTPClient creates a new IssueDeleteCommentDeprecatedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIssueDeleteCommentDeprecatedParamsWithHTTPClient(client *http.Client) *IssueDeleteCommentDeprecatedParams {
	var ()
	return &IssueDeleteCommentDeprecatedParams{
		HTTPClient: client,
	}
}

/*IssueDeleteCommentDeprecatedParams contains all the parameters to send to the API endpoint
for the issue delete comment deprecated operation typically these are written to a http.Request
*/
type IssueDeleteCommentDeprecatedParams struct {

	/*ID
	  id of comment to delete

	*/
	ID int64
	/*Index
	  this parameter is ignored

	*/
	Index int64
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

// WithTimeout adds the timeout to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) WithTimeout(timeout time.Duration) *IssueDeleteCommentDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) WithContext(ctx context.Context) *IssueDeleteCommentDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) WithHTTPClient(client *http.Client) *IssueDeleteCommentDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) WithID(id int64) *IssueDeleteCommentDeprecatedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) SetID(id int64) {
	o.ID = id
}

// WithIndex adds the index to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) WithIndex(index int64) *IssueDeleteCommentDeprecatedParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) WithOwner(owner string) *IssueDeleteCommentDeprecatedParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) WithRepo(repo string) *IssueDeleteCommentDeprecatedParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue delete comment deprecated params
func (o *IssueDeleteCommentDeprecatedParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueDeleteCommentDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
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