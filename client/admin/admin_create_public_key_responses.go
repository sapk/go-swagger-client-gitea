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

// AdminCreatePublicKeyReader is a Reader for the AdminCreatePublicKey structure.
type AdminCreatePublicKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminCreatePublicKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAdminCreatePublicKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewAdminCreatePublicKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewAdminCreatePublicKeyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdminCreatePublicKeyCreated creates a AdminCreatePublicKeyCreated with default headers values
func NewAdminCreatePublicKeyCreated() *AdminCreatePublicKeyCreated {
	return &AdminCreatePublicKeyCreated{}
}

/*AdminCreatePublicKeyCreated handles this case with default header values.

PublicKey
*/
type AdminCreatePublicKeyCreated struct {
	Payload AdminCreatePublicKeyCreatedBody
}

func (o *AdminCreatePublicKeyCreated) Error() string {
	return fmt.Sprintf("[POST /admin/users/{username}/keys][%d] adminCreatePublicKeyCreated  %+v", 201, o.Payload)
}

func (o *AdminCreatePublicKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreatePublicKeyForbidden creates a AdminCreatePublicKeyForbidden with default headers values
func NewAdminCreatePublicKeyForbidden() *AdminCreatePublicKeyForbidden {
	return &AdminCreatePublicKeyForbidden{}
}

/*AdminCreatePublicKeyForbidden handles this case with default header values.

APIForbiddenError is a forbidden error response
*/
type AdminCreatePublicKeyForbidden struct {
	Message string

	URL string
}

func (o *AdminCreatePublicKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users/{username}/keys][%d] adminCreatePublicKeyForbidden ", 403)
}

func (o *AdminCreatePublicKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

// NewAdminCreatePublicKeyUnprocessableEntity creates a AdminCreatePublicKeyUnprocessableEntity with default headers values
func NewAdminCreatePublicKeyUnprocessableEntity() *AdminCreatePublicKeyUnprocessableEntity {
	return &AdminCreatePublicKeyUnprocessableEntity{}
}

/*AdminCreatePublicKeyUnprocessableEntity handles this case with default header values.

APIValidationError is error format response related to input validation
*/
type AdminCreatePublicKeyUnprocessableEntity struct {
	Message string

	URL string
}

func (o *AdminCreatePublicKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /admin/users/{username}/keys][%d] adminCreatePublicKeyUnprocessableEntity ", 422)
}

func (o *AdminCreatePublicKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

/*AdminCreatePublicKeyCreatedBody PublicKey publickey is a user key to push code to repository
swagger:model AdminCreatePublicKeyCreatedBody
*/
type AdminCreatePublicKeyCreatedBody struct {

	// created
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created_at"`

	// fingerprint
	// Required: true
	Fingerprint *string `json:"fingerprint"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// key
	// Required: true
	Key *string `json:"key"`

	// title
	// Required: true
	Title *string `json:"title"`

	// URL
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this admin create public key created body
func (o *AdminCreatePublicKeyCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFingerprint(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AdminCreatePublicKeyCreatedBody) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("adminCreatePublicKeyCreated"+"."+"created_at", "body", o.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("adminCreatePublicKeyCreated"+"."+"created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreatePublicKeyCreatedBody) validateFingerprint(formats strfmt.Registry) error {

	if err := validate.Required("adminCreatePublicKeyCreated"+"."+"fingerprint", "body", o.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreatePublicKeyCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("adminCreatePublicKeyCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreatePublicKeyCreatedBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("adminCreatePublicKeyCreated"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreatePublicKeyCreatedBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("adminCreatePublicKeyCreated"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

func (o *AdminCreatePublicKeyCreatedBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("adminCreatePublicKeyCreated"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AdminCreatePublicKeyCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminCreatePublicKeyCreatedBody) UnmarshalBinary(b []byte) error {
	var res AdminCreatePublicKeyCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
