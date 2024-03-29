// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// RepoCreateStatusReader is a Reader for the RepoCreateStatus structure.
type RepoCreateStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreateStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRepoCreateStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoCreateStatusOK creates a RepoCreateStatusOK with default headers values
func NewRepoCreateStatusOK() *RepoCreateStatusOK {
	return &RepoCreateStatusOK{}
}

/*RepoCreateStatusOK handles this case with default header values.

StatusList
*/
type RepoCreateStatusOK struct {
	Payload []*models.RepoCreateStatusOKBodyItems0
}

func (o *RepoCreateStatusOK) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/statuses/{sha}][%d] repoCreateStatusOK  %+v", 200, o.Payload)
}

func (o *RepoCreateStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RepoCreateStatusBody CreateStatusOption holds the information needed to create a new Status for a Commit
swagger:model RepoCreateStatusBody
*/
type RepoCreateStatusBody struct {

	// context
	Context string `json:"context,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// target URL
	TargetURL string `json:"target_url,omitempty"`

	// StatusState holds the state of a Status
	// It can be "pending", "success", "error", "failure", and "warning"
	State string `json:"state,omitempty"`
}

// MarshalBinary interface implementation
func (o *RepoCreateStatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoCreateStatusBody) UnmarshalBinary(b []byte) error {
	var res RepoCreateStatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoCreateStatusOKBodyItems0 Status holds a single Status of a single Commit
swagger:model RepoCreateStatusOKBodyItems0
*/
type RepoCreateStatusOKBodyItems0 struct {

	// context
	Context string `json:"context,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// target URL
	TargetURL string `json:"target_url,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// creator
	Creator *RepoCreateStatusOKBodyItems0Creator `json:"creator,omitempty"`

	// StatusState holds the state of a Status
	// It can be "pending", "success", "error", "failure", and "warning"
	Status string `json:"status,omitempty"`
}

// Validate validates this repo create status o k body items0
func (o *RepoCreateStatusOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoCreateStatusOKBodyItems0) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateStatusOKBodyItems0) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(o.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", o.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateStatusOKBodyItems0) validateCreator(formats strfmt.Registry) error {

	if swag.IsZero(o.Creator) { // not required
		return nil
	}

	if o.Creator != nil {
		if err := o.Creator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoCreateStatusOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoCreateStatusOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res RepoCreateStatusOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoCreateStatusOKBodyItems0Creator User represents a user
swagger:model RepoCreateStatusOKBodyItems0Creator
*/
type RepoCreateStatusOKBodyItems0Creator struct {

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

// Validate validates this repo create status o k body items0 creator
func (o *RepoCreateStatusOKBodyItems0Creator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoCreateStatusOKBodyItems0Creator) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("creator"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoCreateStatusOKBodyItems0Creator) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoCreateStatusOKBodyItems0Creator) UnmarshalBinary(b []byte) error {
	var res RepoCreateStatusOKBodyItems0Creator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
