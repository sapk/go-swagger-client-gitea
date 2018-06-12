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

// UserListKeysReader is a Reader for the UserListKeys structure.
type UserListKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserListKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserListKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserListKeysOK creates a UserListKeysOK with default headers values
func NewUserListKeysOK() *UserListKeysOK {
	return &UserListKeysOK{}
}

/*UserListKeysOK handles this case with default header values.

PublicKeyList
*/
type UserListKeysOK struct {
	Payload []*models.UserListKeysOKBodyItems0
}

func (o *UserListKeysOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/keys][%d] userListKeysOK  %+v", 200, o.Payload)
}

func (o *UserListKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UserListKeysOKBodyItems0 PublicKey publickey is a user key to push code to repository
swagger:model UserListKeysOKBodyItems0
*/
type UserListKeysOKBodyItems0 struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// URL
	URL string `json:"url,omitempty"`
}

// Validate validates this user list keys o k body items0
func (o *UserListKeysOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserListKeysOKBodyItems0) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserListKeysOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserListKeysOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res UserListKeysOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}