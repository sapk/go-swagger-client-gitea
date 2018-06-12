// Code generated by go-swagger; DO NOT EDIT.

package admin

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
)

// AdminEditUserReader is a Reader for the AdminEditUser structure.
type AdminEditUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminEditUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAdminEditUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewAdminEditUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewAdminEditUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdminEditUserOK creates a AdminEditUserOK with default headers values
func NewAdminEditUserOK() *AdminEditUserOK {
	return &AdminEditUserOK{}
}

/*AdminEditUserOK handles this case with default header values.

User
*/
type AdminEditUserOK struct {
	Payload AdminEditUserOKBody
}

func (o *AdminEditUserOK) Error() string {
	return fmt.Sprintf("[PATCH /admin/users/{username}][%d] adminEditUserOK  %+v", 200, o.Payload)
}

func (o *AdminEditUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminEditUserForbidden creates a AdminEditUserForbidden with default headers values
func NewAdminEditUserForbidden() *AdminEditUserForbidden {
	return &AdminEditUserForbidden{}
}

/*AdminEditUserForbidden handles this case with default header values.

APIForbiddenError is a forbidden error response
*/
type AdminEditUserForbidden struct {
	Message string

	URL string
}

func (o *AdminEditUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /admin/users/{username}][%d] adminEditUserForbidden ", 403)
}

func (o *AdminEditUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

// NewAdminEditUserUnprocessableEntity creates a AdminEditUserUnprocessableEntity with default headers values
func NewAdminEditUserUnprocessableEntity() *AdminEditUserUnprocessableEntity {
	return &AdminEditUserUnprocessableEntity{}
}

/*AdminEditUserUnprocessableEntity handles this case with default header values.

APIValidationError is error format response related to input validation
*/
type AdminEditUserUnprocessableEntity struct {
	Message string

	URL string
}

func (o *AdminEditUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /admin/users/{username}][%d] adminEditUserUnprocessableEntity ", 422)
}

func (o *AdminEditUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

/*AdminEditUserBody EditUserOption edit user options
swagger:model AdminEditUserBody
*/
type AdminEditUserBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// admin
	Admin bool `json:"admin,omitempty"`

	// allow git hook
	AllowGitHook bool `json:"allow_git_hook,omitempty"`

	// allow import local
	AllowImportLocal bool `json:"allow_import_local,omitempty"`

	// email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// login name
	LoginName string `json:"login_name,omitempty"`

	// max repo creation
	MaxRepoCreation int64 `json:"max_repo_creation,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// source ID
	SourceID int64 `json:"source_id,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// MarshalBinary interface implementation
func (o *AdminEditUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminEditUserBody) UnmarshalBinary(b []byte) error {
	var res AdminEditUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AdminEditUserOKBody User represents a user
swagger:model AdminEditUserOKBody
*/
type AdminEditUserOKBody struct {

	// URL to the user's avatar
	// Required: true
	AvatarURL *string `json:"avatar_url"`

	// email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// the user's full name
	// Required: true
	FullName *string `json:"full_name"`

	// the user's id
	// Required: true
	ID *int64 `json:"id"`

	// User locale
	// Required: true
	Language *string `json:"language"`

	// the user's username
	// Required: true
	UserName *string `json:"login"`
}

// Validate validates this admin edit user o k body
func (o *AdminEditUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvatarURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AdminEditUserOKBody) validateAvatarURL(formats strfmt.Registry) error {

	if err := validate.Required("adminEditUserOK"+"."+"avatar_url", "body", o.AvatarURL); err != nil {
		return err
	}

	return nil
}

func (o *AdminEditUserOKBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("adminEditUserOK"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("adminEditUserOK"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *AdminEditUserOKBody) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("adminEditUserOK"+"."+"full_name", "body", o.FullName); err != nil {
		return err
	}

	return nil
}

func (o *AdminEditUserOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("adminEditUserOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *AdminEditUserOKBody) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("adminEditUserOK"+"."+"language", "body", o.Language); err != nil {
		return err
	}

	return nil
}

func (o *AdminEditUserOKBody) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("adminEditUserOK"+"."+"login", "body", o.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AdminEditUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminEditUserOKBody) UnmarshalBinary(b []byte) error {
	var res AdminEditUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
