// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// OrgListTeamMembersReader is a Reader for the OrgListTeamMembers structure.
type OrgListTeamMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListTeamMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgListTeamMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgListTeamMembersOK creates a OrgListTeamMembersOK with default headers values
func NewOrgListTeamMembersOK() *OrgListTeamMembersOK {
	return &OrgListTeamMembersOK{}
}

/*OrgListTeamMembersOK handles this case with default header values.

UserList
*/
type OrgListTeamMembersOK struct {
	Payload []*models.OrgListTeamMembersOKBodyItems0
}

func (o *OrgListTeamMembersOK) Error() string {
	return fmt.Sprintf("[GET /teams/{id}/members][%d] orgListTeamMembersOK  %+v", 200, o.Payload)
}

func (o *OrgListTeamMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OrgListTeamMembersOKBodyItems0 User represents a user
swagger:model OrgListTeamMembersOKBodyItems0
*/
type OrgListTeamMembersOKBodyItems0 struct {

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

// Validate validates this org list team members o k body items0
func (o *OrgListTeamMembersOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrgListTeamMembersOKBodyItems0) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OrgListTeamMembersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrgListTeamMembersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res OrgListTeamMembersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
