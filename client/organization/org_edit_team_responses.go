// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgEditTeamReader is a Reader for the OrgEditTeam structure.
type OrgEditTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgEditTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgEditTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgEditTeamOK creates a OrgEditTeamOK with default headers values
func NewOrgEditTeamOK() *OrgEditTeamOK {
	return &OrgEditTeamOK{}
}

/*OrgEditTeamOK handles this case with default header values.

Team
*/
type OrgEditTeamOK struct {
	Payload OrgEditTeamOKBody
}

func (o *OrgEditTeamOK) Error() string {
	return fmt.Sprintf("[PATCH /teams/{id}][%d] orgEditTeamOK  %+v", 200, o.Payload)
}

func (o *OrgEditTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OrgEditTeamBody EditTeamOption options for editing a team
swagger:model OrgEditTeamBody
*/
type OrgEditTeamBody struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// permission
	// Enum: [read write admin]
	Permission string `json:"permission,omitempty"`
}

// MarshalBinary interface implementation
func (o *OrgEditTeamBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrgEditTeamBody) UnmarshalBinary(b []byte) error {
	var res OrgEditTeamBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*OrgEditTeamOKBody Team represents a team in an organization
swagger:model OrgEditTeamOKBody
*/
type OrgEditTeamOKBody struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// permission
	// Required: true
	// Enum: [none read write admin owner]
	Permission *string `json:"permission"`
}

// Validate validates this org edit team o k body
func (o *OrgEditTeamOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrgEditTeamOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("orgEditTeamOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *OrgEditTeamOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("orgEditTeamOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *OrgEditTeamOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("orgEditTeamOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

var orgEditTeamOKBodyTypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","read","write","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orgEditTeamOKBodyTypePermissionPropEnum = append(orgEditTeamOKBodyTypePermissionPropEnum, v)
	}
}

const (

	// OrgEditTeamOKBodyPermissionNone captures enum value "none"
	OrgEditTeamOKBodyPermissionNone string = "none"

	// OrgEditTeamOKBodyPermissionRead captures enum value "read"
	OrgEditTeamOKBodyPermissionRead string = "read"

	// OrgEditTeamOKBodyPermissionWrite captures enum value "write"
	OrgEditTeamOKBodyPermissionWrite string = "write"

	// OrgEditTeamOKBodyPermissionAdmin captures enum value "admin"
	OrgEditTeamOKBodyPermissionAdmin string = "admin"

	// OrgEditTeamOKBodyPermissionOwner captures enum value "owner"
	OrgEditTeamOKBodyPermissionOwner string = "owner"
)

// prop value enum
func (o *OrgEditTeamOKBody) validatePermissionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orgEditTeamOKBodyTypePermissionPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *OrgEditTeamOKBody) validatePermission(formats strfmt.Registry) error {

	if err := validate.Required("orgEditTeamOK"+"."+"permission", "body", o.Permission); err != nil {
		return err
	}

	// value enum
	if err := o.validatePermissionEnum("orgEditTeamOK"+"."+"permission", "body", *o.Permission); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OrgEditTeamOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrgEditTeamOKBody) UnmarshalBinary(b []byte) error {
	var res OrgEditTeamOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
