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

// OrgCreateTeamReader is a Reader for the OrgCreateTeam structure.
type OrgCreateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgCreateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewOrgCreateTeamCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgCreateTeamCreated creates a OrgCreateTeamCreated with default headers values
func NewOrgCreateTeamCreated() *OrgCreateTeamCreated {
	return &OrgCreateTeamCreated{}
}

/*OrgCreateTeamCreated handles this case with default header values.

Team
*/
type OrgCreateTeamCreated struct {
	Payload OrgCreateTeamCreatedBody
}

func (o *OrgCreateTeamCreated) Error() string {
	return fmt.Sprintf("[POST /orgs/{org}/teams][%d] orgCreateTeamCreated  %+v", 201, o.Payload)
}

func (o *OrgCreateTeamCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OrgCreateTeamBody CreateTeamOption options for creating a team
swagger:model OrgCreateTeamBody
*/
type OrgCreateTeamBody struct {

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
func (o *OrgCreateTeamBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrgCreateTeamBody) UnmarshalBinary(b []byte) error {
	var res OrgCreateTeamBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*OrgCreateTeamCreatedBody Team represents a team in an organization
swagger:model OrgCreateTeamCreatedBody
*/
type OrgCreateTeamCreatedBody struct {

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

// Validate validates this org create team created body
func (o *OrgCreateTeamCreatedBody) Validate(formats strfmt.Registry) error {
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

func (o *OrgCreateTeamCreatedBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("orgCreateTeamCreated"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *OrgCreateTeamCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("orgCreateTeamCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *OrgCreateTeamCreatedBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("orgCreateTeamCreated"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

var orgCreateTeamCreatedBodyTypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","read","write","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orgCreateTeamCreatedBodyTypePermissionPropEnum = append(orgCreateTeamCreatedBodyTypePermissionPropEnum, v)
	}
}

const (

	// OrgCreateTeamCreatedBodyPermissionNone captures enum value "none"
	OrgCreateTeamCreatedBodyPermissionNone string = "none"

	// OrgCreateTeamCreatedBodyPermissionRead captures enum value "read"
	OrgCreateTeamCreatedBodyPermissionRead string = "read"

	// OrgCreateTeamCreatedBodyPermissionWrite captures enum value "write"
	OrgCreateTeamCreatedBodyPermissionWrite string = "write"

	// OrgCreateTeamCreatedBodyPermissionAdmin captures enum value "admin"
	OrgCreateTeamCreatedBodyPermissionAdmin string = "admin"

	// OrgCreateTeamCreatedBodyPermissionOwner captures enum value "owner"
	OrgCreateTeamCreatedBodyPermissionOwner string = "owner"
)

// prop value enum
func (o *OrgCreateTeamCreatedBody) validatePermissionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orgCreateTeamCreatedBodyTypePermissionPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *OrgCreateTeamCreatedBody) validatePermission(formats strfmt.Registry) error {

	if err := validate.Required("orgCreateTeamCreated"+"."+"permission", "body", o.Permission); err != nil {
		return err
	}

	// value enum
	if err := o.validatePermissionEnum("orgCreateTeamCreated"+"."+"permission", "body", *o.Permission); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OrgCreateTeamCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrgCreateTeamCreatedBody) UnmarshalBinary(b []byte) error {
	var res OrgCreateTeamCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}