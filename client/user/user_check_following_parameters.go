// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUserCheckFollowingParams creates a new UserCheckFollowingParams object
// with the default values initialized.
func NewUserCheckFollowingParams() *UserCheckFollowingParams {
	var ()
	return &UserCheckFollowingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCheckFollowingParamsWithTimeout creates a new UserCheckFollowingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCheckFollowingParamsWithTimeout(timeout time.Duration) *UserCheckFollowingParams {
	var ()
	return &UserCheckFollowingParams{

		timeout: timeout,
	}
}

// NewUserCheckFollowingParamsWithContext creates a new UserCheckFollowingParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCheckFollowingParamsWithContext(ctx context.Context) *UserCheckFollowingParams {
	var ()
	return &UserCheckFollowingParams{

		Context: ctx,
	}
}

// NewUserCheckFollowingParamsWithHTTPClient creates a new UserCheckFollowingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCheckFollowingParamsWithHTTPClient(client *http.Client) *UserCheckFollowingParams {
	var ()
	return &UserCheckFollowingParams{
		HTTPClient: client,
	}
}

/*UserCheckFollowingParams contains all the parameters to send to the API endpoint
for the user check following operation typically these are written to a http.Request
*/
type UserCheckFollowingParams struct {

	/*Followee
	  username of followed user

	*/
	Followee string
	/*Follower
	  username of following user

	*/
	Follower string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user check following params
func (o *UserCheckFollowingParams) WithTimeout(timeout time.Duration) *UserCheckFollowingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user check following params
func (o *UserCheckFollowingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user check following params
func (o *UserCheckFollowingParams) WithContext(ctx context.Context) *UserCheckFollowingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user check following params
func (o *UserCheckFollowingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user check following params
func (o *UserCheckFollowingParams) WithHTTPClient(client *http.Client) *UserCheckFollowingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user check following params
func (o *UserCheckFollowingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFollowee adds the followee to the user check following params
func (o *UserCheckFollowingParams) WithFollowee(followee string) *UserCheckFollowingParams {
	o.SetFollowee(followee)
	return o
}

// SetFollowee adds the followee to the user check following params
func (o *UserCheckFollowingParams) SetFollowee(followee string) {
	o.Followee = followee
}

// WithFollower adds the follower to the user check following params
func (o *UserCheckFollowingParams) WithFollower(follower string) *UserCheckFollowingParams {
	o.SetFollower(follower)
	return o
}

// SetFollower adds the follower to the user check following params
func (o *UserCheckFollowingParams) SetFollower(follower string) {
	o.Follower = follower
}

// WriteToRequest writes these params to a swagger request
func (o *UserCheckFollowingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param followee
	if err := r.SetPathParam("followee", o.Followee); err != nil {
		return err
	}

	// path param follower
	if err := r.SetPathParam("follower", o.Follower); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}