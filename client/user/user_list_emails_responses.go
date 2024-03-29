// Code generated by go-swagger; DO NOT EDIT.

package user

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

// UserListEmailsReader is a Reader for the UserListEmails structure.
type UserListEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserListEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserListEmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserListEmailsOK creates a UserListEmailsOK with default headers values
func NewUserListEmailsOK() *UserListEmailsOK {
	return &UserListEmailsOK{}
}

/*UserListEmailsOK handles this case with default header values.

EmailList
*/
type UserListEmailsOK struct {
	Payload []*models.UserListEmailsOKBodyItems0
}

func (o *UserListEmailsOK) Error() string {
	return fmt.Sprintf("[GET /user/emails][%d] userListEmailsOK  %+v", 200, o.Payload)
}

func (o *UserListEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UserListEmailsOKBodyItems0 Email an email address belonging to a user
swagger:model UserListEmailsOKBodyItems0
*/
type UserListEmailsOKBodyItems0 struct {

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// primary
	Primary bool `json:"primary,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this user list emails o k body items0
func (o *UserListEmailsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserListEmailsOKBodyItems0) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserListEmailsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserListEmailsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res UserListEmailsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
