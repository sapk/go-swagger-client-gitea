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

	models "github.com/sapk/go-swagger-client-gitea/models"
)

// OrgListTeamsReader is a Reader for the OrgListTeams structure.
type OrgListTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrgListTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgListTeamsOK creates a OrgListTeamsOK with default headers values
func NewOrgListTeamsOK() *OrgListTeamsOK {
	return &OrgListTeamsOK{}
}

/*OrgListTeamsOK handles this case with default header values.

TeamList
*/
type OrgListTeamsOK struct {
	Payload []*models.OrgListTeamsOKBodyItems0
}

func (o *OrgListTeamsOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/teams][%d] orgListTeamsOK  %+v", 200, o.Payload)
}

func (o *OrgListTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OrgListTeamsOKBodyItems0 Team represents a team in an organization
swagger:model OrgListTeamsOKBodyItems0
*/
type OrgListTeamsOKBodyItems0 struct {

	// description
	Description string `json:"description,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// permission
	// Enum: [none read write admin owner]
	Permission string `json:"permission,omitempty"`
}

// Validate validates this org list teams o k body items0
func (o *OrgListTeamsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var orgListTeamsOKBodyItems0TypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","read","write","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orgListTeamsOKBodyItems0TypePermissionPropEnum = append(orgListTeamsOKBodyItems0TypePermissionPropEnum, v)
	}
}

const (

	// OrgListTeamsOKBodyItems0PermissionNone captures enum value "none"
	OrgListTeamsOKBodyItems0PermissionNone string = "none"

	// OrgListTeamsOKBodyItems0PermissionRead captures enum value "read"
	OrgListTeamsOKBodyItems0PermissionRead string = "read"

	// OrgListTeamsOKBodyItems0PermissionWrite captures enum value "write"
	OrgListTeamsOKBodyItems0PermissionWrite string = "write"

	// OrgListTeamsOKBodyItems0PermissionAdmin captures enum value "admin"
	OrgListTeamsOKBodyItems0PermissionAdmin string = "admin"

	// OrgListTeamsOKBodyItems0PermissionOwner captures enum value "owner"
	OrgListTeamsOKBodyItems0PermissionOwner string = "owner"
)

// prop value enum
func (o *OrgListTeamsOKBodyItems0) validatePermissionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orgListTeamsOKBodyItems0TypePermissionPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *OrgListTeamsOKBodyItems0) validatePermission(formats strfmt.Registry) error {

	if swag.IsZero(o.Permission) { // not required
		return nil
	}

	// value enum
	if err := o.validatePermissionEnum("permission", "body", o.Permission); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OrgListTeamsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrgListTeamsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res OrgListTeamsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
