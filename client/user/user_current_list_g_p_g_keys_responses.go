// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sapk/go-swagger-client-gitea/models"
)

// UserCurrentListGPGKeysReader is a Reader for the UserCurrentListGPGKeys structure.
type UserCurrentListGPGKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentListGPGKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCurrentListGPGKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentListGPGKeysOK creates a UserCurrentListGPGKeysOK with default headers values
func NewUserCurrentListGPGKeysOK() *UserCurrentListGPGKeysOK {
	return &UserCurrentListGPGKeysOK{}
}

/*UserCurrentListGPGKeysOK handles this case with default header values.

GPGKeyList
*/
type UserCurrentListGPGKeysOK struct {
	Payload []*models.UserCurrentListGPGKeysOKBodyItems0
}

func (o *UserCurrentListGPGKeysOK) Error() string {
	return fmt.Sprintf("[GET /user/gpg_keys][%d] userCurrentListGPGKeysOK  %+v", 200, o.Payload)
}

func (o *UserCurrentListGPGKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UserCurrentListGPGKeysOKBodyItems0 GPGKey a user GPG key to sign commit and tag in repository
swagger:model UserCurrentListGPGKeysOKBodyItems0
*/
type UserCurrentListGPGKeysOKBodyItems0 struct {

	// can certify
	CanCertify bool `json:"can_certify,omitempty"`

	// can encrypt comms
	CanEncryptComms bool `json:"can_encrypt_comms,omitempty"`

	// can encrypt storage
	CanEncryptStorage bool `json:"can_encrypt_storage,omitempty"`

	// can sign
	CanSign bool `json:"can_sign,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// emails
	Emails []*models.UserCurrentListGPGKeysOKBodyItems0EmailsItems0 `json:"emails"`

	// expires
	// Format: date-time
	Expires strfmt.DateTime `json:"expires_at,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// key ID
	KeyID string `json:"key_id,omitempty"`

	// primary key ID
	PrimaryKeyID string `json:"primary_key_id,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// subs key
	SubsKey []*models.UserCurrentListGPGKeysOKBodyItems0SubkeysItems0 `json:"subkeys"`
}

// Validate validates this user current list g p g keys o k body items0
func (o *UserCurrentListGPGKeysOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExpires(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubsKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(o.Emails) { // not required
		return nil
	}

	for i := 0; i < len(o.Emails); i++ {
		if swag.IsZero(o.Emails[i]) { // not required
			continue
		}

		if o.Emails[i] != nil {
			if err := o.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0) validateExpires(formats strfmt.Registry) error {

	if swag.IsZero(o.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", o.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0) validateSubsKey(formats strfmt.Registry) error {

	if swag.IsZero(o.SubsKey) { // not required
		return nil
	}

	for i := 0; i < len(o.SubsKey); i++ {
		if swag.IsZero(o.SubsKey[i]) { // not required
			continue
		}

		if o.SubsKey[i] != nil {
			if err := o.SubsKey[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subkeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res UserCurrentListGPGKeysOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UserCurrentListGPGKeysOKBodyItems0EmailsItems0 GPGKeyEmail an email attached to a GPGKey
swagger:model UserCurrentListGPGKeysOKBodyItems0EmailsItems0
*/
type UserCurrentListGPGKeysOKBodyItems0EmailsItems0 struct {

	// email
	Email string `json:"email,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this user current list g p g keys o k body items0 emails items0
func (o *UserCurrentListGPGKeysOKBodyItems0EmailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0EmailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0EmailsItems0) UnmarshalBinary(b []byte) error {
	var res UserCurrentListGPGKeysOKBodyItems0EmailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UserCurrentListGPGKeysOKBodyItems0SubkeysItems0 GPGKey a user GPG key to sign commit and tag in repository
swagger:model UserCurrentListGPGKeysOKBodyItems0SubkeysItems0
*/
type UserCurrentListGPGKeysOKBodyItems0SubkeysItems0 struct {

	// can certify
	CanCertify bool `json:"can_certify,omitempty"`

	// can encrypt comms
	CanEncryptComms bool `json:"can_encrypt_comms,omitempty"`

	// can encrypt storage
	CanEncryptStorage bool `json:"can_encrypt_storage,omitempty"`

	// can sign
	CanSign bool `json:"can_sign,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// emails
	Emails []*models.UserCurrentListGPGKeysOKBodyItems0SubkeysItems0EmailsItems0 `json:"emails"`

	// expires
	// Format: date-time
	Expires strfmt.DateTime `json:"expires_at,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// key ID
	KeyID string `json:"key_id,omitempty"`

	// primary key ID
	PrimaryKeyID string `json:"primary_key_id,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// subs key
	SubsKey []*models.UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0 `json:"subkeys"`
}

// Validate validates this user current list g p g keys o k body items0 subkeys items0
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExpires(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubsKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(o.Emails) { // not required
		return nil
	}

	for i := 0; i < len(o.Emails); i++ {
		if swag.IsZero(o.Emails[i]) { // not required
			continue
		}

		if o.Emails[i] != nil {
			if err := o.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0) validateExpires(formats strfmt.Registry) error {

	if swag.IsZero(o.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", o.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0) validateSubsKey(formats strfmt.Registry) error {

	if swag.IsZero(o.SubsKey) { // not required
		return nil
	}

	for i := 0; i < len(o.SubsKey); i++ {
		if swag.IsZero(o.SubsKey[i]) { // not required
			continue
		}

		if o.SubsKey[i] != nil {
			if err := o.SubsKey[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subkeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0) UnmarshalBinary(b []byte) error {
	var res UserCurrentListGPGKeysOKBodyItems0SubkeysItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UserCurrentListGPGKeysOKBodyItems0SubkeysItems0EmailsItems0 GPGKeyEmail an email attached to a GPGKey
swagger:model UserCurrentListGPGKeysOKBodyItems0SubkeysItems0EmailsItems0
*/
type UserCurrentListGPGKeysOKBodyItems0SubkeysItems0EmailsItems0 struct {

	// email
	Email string `json:"email,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this user current list g p g keys o k body items0 subkeys items0 emails items0
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0EmailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0EmailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0EmailsItems0) UnmarshalBinary(b []byte) error {
	var res UserCurrentListGPGKeysOKBodyItems0SubkeysItems0EmailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0 GPGKey a user GPG key to sign commit and tag in repository
swagger:model UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0
*/
type UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0 struct {

	// can certify
	CanCertify bool `json:"can_certify,omitempty"`

	// can encrypt comms
	CanEncryptComms bool `json:"can_encrypt_comms,omitempty"`

	// can encrypt storage
	CanEncryptStorage bool `json:"can_encrypt_storage,omitempty"`

	// can sign
	CanSign bool `json:"can_sign,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// emails
	Emails []*models.UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0EmailsItems0 `json:"emails"`

	// expires
	// Format: date-time
	Expires strfmt.DateTime `json:"expires_at,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// key ID
	KeyID string `json:"key_id,omitempty"`

	// primary key ID
	PrimaryKeyID string `json:"primary_key_id,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// subs key
	SubsKey []*models.GPGKey `json:"subkeys"`
}

// Validate validates this user current list g p g keys o k body items0 subkeys items0 subkeys items0
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExpires(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubsKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(o.Emails) { // not required
		return nil
	}

	for i := 0; i < len(o.Emails); i++ {
		if swag.IsZero(o.Emails[i]) { // not required
			continue
		}

		if o.Emails[i] != nil {
			if err := o.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0) validateExpires(formats strfmt.Registry) error {

	if swag.IsZero(o.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", o.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0) validateSubsKey(formats strfmt.Registry) error {

	if swag.IsZero(o.SubsKey) { // not required
		return nil
	}

	for i := 0; i < len(o.SubsKey); i++ {
		if swag.IsZero(o.SubsKey[i]) { // not required
			continue
		}

		if o.SubsKey[i] != nil {
			if err := o.SubsKey[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subkeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0) UnmarshalBinary(b []byte) error {
	var res UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0EmailsItems0 GPGKeyEmail an email attached to a GPGKey
swagger:model UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0EmailsItems0
*/
type UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0EmailsItems0 struct {

	// email
	Email string `json:"email,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this user current list g p g keys o k body items0 subkeys items0 subkeys items0 emails items0
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0EmailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0EmailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0EmailsItems0) UnmarshalBinary(b []byte) error {
	var res UserCurrentListGPGKeysOKBodyItems0SubkeysItems0SubkeysItems0EmailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
