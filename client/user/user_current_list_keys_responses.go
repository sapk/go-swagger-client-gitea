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

// UserCurrentListKeysReader is a Reader for the UserCurrentListKeys structure.
type UserCurrentListKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentListKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCurrentListKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentListKeysOK creates a UserCurrentListKeysOK with default headers values
func NewUserCurrentListKeysOK() *UserCurrentListKeysOK {
	return &UserCurrentListKeysOK{}
}

/*UserCurrentListKeysOK handles this case with default header values.

PublicKeyList
*/
type UserCurrentListKeysOK struct {
	Payload []*models.UserCurrentListKeysOKBodyItems0
}

func (o *UserCurrentListKeysOK) Error() string {
	return fmt.Sprintf("[GET /user/keys][%d] userCurrentListKeysOK  %+v", 200, o.Payload)
}

func (o *UserCurrentListKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UserCurrentListKeysOKBodyItems0 PublicKey publickey is a user key to push code to repository
swagger:model UserCurrentListKeysOKBodyItems0
*/
type UserCurrentListKeysOKBodyItems0 struct {

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

// Validate validates this user current list keys o k body items0
func (o *UserCurrentListKeysOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserCurrentListKeysOKBodyItems0) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserCurrentListKeysOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserCurrentListKeysOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res UserCurrentListKeysOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
