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

// AdminCreateOrgReader is a Reader for the AdminCreateOrg structure.
type AdminCreateOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminCreateOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAdminCreateOrgCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewAdminCreateOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewAdminCreateOrgUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdminCreateOrgCreated creates a AdminCreateOrgCreated with default headers values
func NewAdminCreateOrgCreated() *AdminCreateOrgCreated {
	return &AdminCreateOrgCreated{}
}

/*AdminCreateOrgCreated handles this case with default header values.

Organization
*/
type AdminCreateOrgCreated struct {
	Payload AdminCreateOrgCreatedBody
}

func (o *AdminCreateOrgCreated) Error() string {
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgCreated  %+v", 201, o.Payload)
}

func (o *AdminCreateOrgCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateOrgForbidden creates a AdminCreateOrgForbidden with default headers values
func NewAdminCreateOrgForbidden() *AdminCreateOrgForbidden {
	return &AdminCreateOrgForbidden{}
}

/*AdminCreateOrgForbidden handles this case with default header values.

APIForbiddenError is a forbidden error response
*/
type AdminCreateOrgForbidden struct {
	Message string

	URL string
}

func (o *AdminCreateOrgForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgForbidden ", 403)
}

func (o *AdminCreateOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

// NewAdminCreateOrgUnprocessableEntity creates a AdminCreateOrgUnprocessableEntity with default headers values
func NewAdminCreateOrgUnprocessableEntity() *AdminCreateOrgUnprocessableEntity {
	return &AdminCreateOrgUnprocessableEntity{}
}

/*AdminCreateOrgUnprocessableEntity handles this case with default header values.

APIValidationError is error format response related to input validation
*/
type AdminCreateOrgUnprocessableEntity struct {
	Message string

	URL string
}

func (o *AdminCreateOrgUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgUnprocessableEntity ", 422)
}

func (o *AdminCreateOrgUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

/*AdminCreateOrgCreatedBody Organization represents an organization
swagger:model AdminCreateOrgCreatedBody
*/
type AdminCreateOrgCreatedBody struct {

	// avatar URL
	// Required: true
	AvatarURL *string `json:"avatar_url"`

	// description
	// Required: true
	Description *string `json:"description"`

	// full name
	// Required: true
	FullName *string `json:"full_name"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// location
	// Required: true
	Location *string `json:"location"`

	// user name
	// Required: true
	UserName *string `json:"username"`

	// website
	// Required: true
	Website *string `json:"website"`
}

// Validate validates this admin create org created body
func (o *AdminCreateOrgCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvatarURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWebsite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AdminCreateOrgCreatedBody) validateAvatarURL(formats strfmt.Registry) error {

	if err := validate.Required("adminCreateOrgCreated"+"."+"avatar_url", "body", o.AvatarURL); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreateOrgCreatedBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("adminCreateOrgCreated"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreateOrgCreatedBody) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("adminCreateOrgCreated"+"."+"full_name", "body", o.FullName); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreateOrgCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("adminCreateOrgCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreateOrgCreatedBody) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("adminCreateOrgCreated"+"."+"location", "body", o.Location); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreateOrgCreatedBody) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("adminCreateOrgCreated"+"."+"username", "body", o.UserName); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreateOrgCreatedBody) validateWebsite(formats strfmt.Registry) error {

	if err := validate.Required("adminCreateOrgCreated"+"."+"website", "body", o.Website); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AdminCreateOrgCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminCreateOrgCreatedBody) UnmarshalBinary(b []byte) error {
	var res AdminCreateOrgCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
