// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sapk/go-swagger-client-gitea/models"
)

// IssueGetRepoCommentsReader is a Reader for the IssueGetRepoComments structure.
type IssueGetRepoCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueGetRepoCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIssueGetRepoCommentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueGetRepoCommentsOK creates a IssueGetRepoCommentsOK with default headers values
func NewIssueGetRepoCommentsOK() *IssueGetRepoCommentsOK {
	return &IssueGetRepoCommentsOK{}
}

/*IssueGetRepoCommentsOK handles this case with default header values.

CommentList
*/
type IssueGetRepoCommentsOK struct {
	Payload []*models.IssueGetRepoCommentsOKBodyItems0
}

func (o *IssueGetRepoCommentsOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/comments][%d] issueGetRepoCommentsOK  %+v", 200, o.Payload)
}

func (o *IssueGetRepoCommentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IssueGetRepoCommentsOKBodyItems0 Comment represents a comment on a commit or issue
swagger:model IssueGetRepoCommentsOKBodyItems0
*/
type IssueGetRepoCommentsOKBodyItems0 struct {

	// body
	Body string `json:"body,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// issue URL
	IssueURL string `json:"issue_url,omitempty"`

	// p r URL
	PRURL string `json:"pull_request_url,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// user
	User *IssueGetRepoCommentsOKBodyItems0User `json:"user,omitempty"`
}

// Validate validates this issue get repo comments o k body items0
func (o *IssueGetRepoCommentsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IssueGetRepoCommentsOKBodyItems0) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IssueGetRepoCommentsOKBodyItems0) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(o.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", o.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IssueGetRepoCommentsOKBodyItems0) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(o.User) { // not required
		return nil
	}

	if o.User != nil {
		if err := o.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *IssueGetRepoCommentsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueGetRepoCommentsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res IssueGetRepoCommentsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*IssueGetRepoCommentsOKBodyItems0User User represents a user
swagger:model IssueGetRepoCommentsOKBodyItems0User
*/
type IssueGetRepoCommentsOKBodyItems0User struct {

	// URL to the user's avatar
	AvatarURL string `json:"avatar_url,omitempty"`

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// the user's full name
	FullName string `json:"full_name,omitempty"`

	// the user's id
	ID int64 `json:"id,omitempty"`

	// User locale
	Language string `json:"language,omitempty"`

	// the user's username
	UserName string `json:"login,omitempty"`
}

// Validate validates this issue get repo comments o k body items0 user
func (o *IssueGetRepoCommentsOKBodyItems0User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IssueGetRepoCommentsOKBodyItems0User) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("user"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *IssueGetRepoCommentsOKBodyItems0User) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueGetRepoCommentsOKBodyItems0User) UnmarshalBinary(b []byte) error {
	var res IssueGetRepoCommentsOKBodyItems0User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
