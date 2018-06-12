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

// NewUserCurrentPutFollowParams creates a new UserCurrentPutFollowParams object
// with the default values initialized.
func NewUserCurrentPutFollowParams() *UserCurrentPutFollowParams {
	var ()
	return &UserCurrentPutFollowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentPutFollowParamsWithTimeout creates a new UserCurrentPutFollowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCurrentPutFollowParamsWithTimeout(timeout time.Duration) *UserCurrentPutFollowParams {
	var ()
	return &UserCurrentPutFollowParams{

		timeout: timeout,
	}
}

// NewUserCurrentPutFollowParamsWithContext creates a new UserCurrentPutFollowParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCurrentPutFollowParamsWithContext(ctx context.Context) *UserCurrentPutFollowParams {
	var ()
	return &UserCurrentPutFollowParams{

		Context: ctx,
	}
}

// NewUserCurrentPutFollowParamsWithHTTPClient creates a new UserCurrentPutFollowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCurrentPutFollowParamsWithHTTPClient(client *http.Client) *UserCurrentPutFollowParams {
	var ()
	return &UserCurrentPutFollowParams{
		HTTPClient: client,
	}
}

/*UserCurrentPutFollowParams contains all the parameters to send to the API endpoint
for the user current put follow operation typically these are written to a http.Request
*/
type UserCurrentPutFollowParams struct {

	/*Username
	  username of user to follow

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user current put follow params
func (o *UserCurrentPutFollowParams) WithTimeout(timeout time.Duration) *UserCurrentPutFollowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current put follow params
func (o *UserCurrentPutFollowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current put follow params
func (o *UserCurrentPutFollowParams) WithContext(ctx context.Context) *UserCurrentPutFollowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current put follow params
func (o *UserCurrentPutFollowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current put follow params
func (o *UserCurrentPutFollowParams) WithHTTPClient(client *http.Client) *UserCurrentPutFollowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current put follow params
func (o *UserCurrentPutFollowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the user current put follow params
func (o *UserCurrentPutFollowParams) WithUsername(username string) *UserCurrentPutFollowParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user current put follow params
func (o *UserCurrentPutFollowParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentPutFollowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}