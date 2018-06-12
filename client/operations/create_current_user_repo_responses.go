// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// CreateCurrentUserRepoReader is a Reader for the CreateCurrentUserRepo structure.
type CreateCurrentUserRepoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCurrentUserRepoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateCurrentUserRepoCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCurrentUserRepoCreated creates a CreateCurrentUserRepoCreated with default headers values
func NewCreateCurrentUserRepoCreated() *CreateCurrentUserRepoCreated {
	return &CreateCurrentUserRepoCreated{}
}

/*CreateCurrentUserRepoCreated handles this case with default header values.

Repository
*/
type CreateCurrentUserRepoCreated struct {
	Payload CreateCurrentUserRepoCreatedBody
}

func (o *CreateCurrentUserRepoCreated) Error() string {
	return fmt.Sprintf("[POST /user/repos][%d] createCurrentUserRepoCreated  %+v", 201, o.Payload)
}

func (o *CreateCurrentUserRepoCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateCurrentUserRepoBody CreateRepoOption options when creating repository
swagger:model CreateCurrentUserRepoBody
*/
type CreateCurrentUserRepoBody struct {

	// Whether the repository should be auto-intialized?
	AutoInit bool `json:"auto_init,omitempty"`

	// Description of the repository to create
	Description string `json:"description,omitempty"`

	// Gitignores to use
	Gitignores string `json:"gitignores,omitempty"`

	// License to use
	License string `json:"license,omitempty"`

	// Name of the repository to create
	// Required: true
	// Unique: true
	Name *string `json:"name"`

	// Whether the repository is private
	Private bool `json:"private,omitempty"`

	// Readme of the repository to create
	Readme string `json:"readme,omitempty"`
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoBody) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBody Repository represents a repository
swagger:model CreateCurrentUserRepoCreatedBody
*/
type CreateCurrentUserRepoCreatedBody struct {

	// clone URL
	// Required: true
	CloneURL *string `json:"clone_url"`

	// created
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created_at"`

	// default branch
	// Required: true
	DefaultBranch *string `json:"default_branch"`

	// description
	// Required: true
	Description *string `json:"description"`

	// empty
	// Required: true
	Empty *bool `json:"empty"`

	// fork
	// Required: true
	Fork *bool `json:"fork"`

	// forks
	// Required: true
	Forks *int64 `json:"forks_count"`

	// full name
	// Required: true
	FullName *string `json:"full_name"`

	// HTML URL
	// Required: true
	HTMLURL *string `json:"html_url"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// mirror
	// Required: true
	Mirror *bool `json:"mirror"`

	// name
	// Required: true
	Name *string `json:"name"`

	// open issues
	// Required: true
	OpenIssues *int64 `json:"open_issues_count"`

	// private
	// Required: true
	Private *bool `json:"private"`

	// SSH URL
	// Required: true
	SSHURL *string `json:"ssh_url"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// stars
	// Required: true
	Stars *int64 `json:"stars_count"`

	// updated
	// Required: true
	// Format: date-time
	Updated *strfmt.DateTime `json:"updated_at"`

	// watchers
	// Required: true
	Watchers *int64 `json:"watchers_count"`

	// website
	// Required: true
	Website *string `json:"website"`

	// owner
	// Required: true
	Owner *CreateCurrentUserRepoCreatedBodyOwner `json:"owner"`

	// parent
	// Required: true
	Parent *CreateCurrentUserRepoCreatedBodyParent `json:"parent"`

	// permissions
	// Required: true
	Permissions *CreateCurrentUserRepoCreatedBodyPermissions `json:"permissions"`
}

// Validate validates this create current user repo created body
func (o *CreateCurrentUserRepoCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCloneURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefaultBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmpty(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFork(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateForks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHTMLURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMirror(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOpenIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrivate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSSHURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStars(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWatchers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWebsite(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateCloneURL(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"clone_url", "body", o.CloneURL); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"created_at", "body", o.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateDefaultBranch(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"default_branch", "body", o.DefaultBranch); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateEmpty(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"empty", "body", o.Empty); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateFork(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"fork", "body", o.Fork); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateForks(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"forks_count", "body", o.Forks); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"full_name", "body", o.FullName); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateHTMLURL(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"html_url", "body", o.HTMLURL); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateMirror(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"mirror", "body", o.Mirror); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateOpenIssues(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"open_issues_count", "body", o.OpenIssues); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validatePrivate(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"private", "body", o.Private); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateSSHURL(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"ssh_url", "body", o.SSHURL); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateStars(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"stars_count", "body", o.Stars); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"updated_at", "body", o.Updated); err != nil {
		return err
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"updated_at", "body", "date-time", o.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateWatchers(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"watchers_count", "body", o.Watchers); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateWebsite(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"website", "body", o.Website); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"owner", "body", o.Owner); err != nil {
		return err
	}

	if o.Owner != nil {
		if err := o.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validateParent(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"parent", "body", o.Parent); err != nil {
		return err
	}

	if o.Parent != nil {
		if err := o.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "parent")
			}
			return err
		}
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBody) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("createCurrentUserRepoCreated"+"."+"permissions", "body", o.Permissions); err != nil {
		return err
	}

	if o.Permissions != nil {
		if err := o.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBodyOwner User represents a user
swagger:model CreateCurrentUserRepoCreatedBodyOwner
*/
type CreateCurrentUserRepoCreatedBodyOwner struct {

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

// Validate validates this create current user repo created body owner
func (o *CreateCurrentUserRepoCreatedBodyOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyOwner) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"owner"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyOwner) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBodyOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBodyParent Repository represents a repository
swagger:model CreateCurrentUserRepoCreatedBodyParent
*/
type CreateCurrentUserRepoCreatedBodyParent struct {

	// clone URL
	CloneURL string `json:"clone_url,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// default branch
	DefaultBranch string `json:"default_branch,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// fork
	Fork bool `json:"fork,omitempty"`

	// forks
	Forks int64 `json:"forks_count,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// mirror
	Mirror bool `json:"mirror,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// open issues
	OpenIssues int64 `json:"open_issues_count,omitempty"`

	// private
	Private bool `json:"private,omitempty"`

	// SSH URL
	SSHURL string `json:"ssh_url,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// stars
	Stars int64 `json:"stars_count,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// watchers
	Watchers int64 `json:"watchers_count,omitempty"`

	// website
	Website string `json:"website,omitempty"`

	// owner
	Owner *CreateCurrentUserRepoCreatedBodyParentOwner `json:"owner,omitempty"`

	// parent
	Parent *CreateCurrentUserRepoCreatedBodyParentParent `json:"parent,omitempty"`

	// permissions
	Permissions *CreateCurrentUserRepoCreatedBodyParentPermissions `json:"permissions,omitempty"`
}

// Validate validates this create current user repo created body parent
func (o *CreateCurrentUserRepoCreatedBodyParent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParent) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"parent"+"."+"created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParent) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(o.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"parent"+"."+"updated_at", "body", "date-time", o.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParent) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(o.Owner) { // not required
		return nil
	}

	if o.Owner != nil {
		if err := o.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "parent" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParent) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(o.Parent) { // not required
		return nil
	}

	if o.Parent != nil {
		if err := o.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "parent" + "." + "parent")
			}
			return err
		}
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParent) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(o.Permissions) { // not required
		return nil
	}

	if o.Permissions != nil {
		if err := o.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "parent" + "." + "permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParent) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBodyParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBodyParentOwner User represents a user
swagger:model CreateCurrentUserRepoCreatedBodyParentOwner
*/
type CreateCurrentUserRepoCreatedBodyParentOwner struct {

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

// Validate validates this create current user repo created body parent owner
func (o *CreateCurrentUserRepoCreatedBodyParentOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParentOwner) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"parent"+"."+"owner"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentOwner) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBodyParentOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBodyParentParent Repository represents a repository
swagger:model CreateCurrentUserRepoCreatedBodyParentParent
*/
type CreateCurrentUserRepoCreatedBodyParentParent struct {

	// clone URL
	CloneURL string `json:"clone_url,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// default branch
	DefaultBranch string `json:"default_branch,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// fork
	Fork bool `json:"fork,omitempty"`

	// forks
	Forks int64 `json:"forks_count,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// mirror
	Mirror bool `json:"mirror,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// open issues
	OpenIssues int64 `json:"open_issues_count,omitempty"`

	// private
	Private bool `json:"private,omitempty"`

	// SSH URL
	SSHURL string `json:"ssh_url,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// stars
	Stars int64 `json:"stars_count,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// watchers
	Watchers int64 `json:"watchers_count,omitempty"`

	// website
	Website string `json:"website,omitempty"`

	// owner
	Owner *CreateCurrentUserRepoCreatedBodyParentParentOwner `json:"owner,omitempty"`

	// parent
	Parent *models.Repository `json:"parent,omitempty"`

	// permissions
	Permissions *CreateCurrentUserRepoCreatedBodyParentParentPermissions `json:"permissions,omitempty"`
}

// Validate validates this create current user repo created body parent parent
func (o *CreateCurrentUserRepoCreatedBodyParentParent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParentParent) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"parent"+"."+"parent"+"."+"created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParentParent) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(o.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"parent"+"."+"parent"+"."+"updated_at", "body", "date-time", o.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParentParent) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(o.Owner) { // not required
		return nil
	}

	if o.Owner != nil {
		if err := o.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "parent" + "." + "parent" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParentParent) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(o.Parent) { // not required
		return nil
	}

	if o.Parent != nil {
		if err := o.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "parent" + "." + "parent" + "." + "parent")
			}
			return err
		}
	}

	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParentParent) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(o.Permissions) { // not required
		return nil
	}

	if o.Permissions != nil {
		if err := o.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCurrentUserRepoCreated" + "." + "parent" + "." + "parent" + "." + "permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentParent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentParent) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBodyParentParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBodyParentParentOwner User represents a user
swagger:model CreateCurrentUserRepoCreatedBodyParentParentOwner
*/
type CreateCurrentUserRepoCreatedBodyParentParentOwner struct {

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

// Validate validates this create current user repo created body parent parent owner
func (o *CreateCurrentUserRepoCreatedBodyParentParentOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateCurrentUserRepoCreatedBodyParentParentOwner) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("createCurrentUserRepoCreated"+"."+"parent"+"."+"parent"+"."+"owner"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentParentOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentParentOwner) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBodyParentParentOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBodyParentParentPermissions Permission represents a set of permissions
swagger:model CreateCurrentUserRepoCreatedBodyParentParentPermissions
*/
type CreateCurrentUserRepoCreatedBodyParentParentPermissions struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// pull
	Pull bool `json:"pull,omitempty"`

	// push
	Push bool `json:"push,omitempty"`
}

// Validate validates this create current user repo created body parent parent permissions
func (o *CreateCurrentUserRepoCreatedBodyParentParentPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentParentPermissions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentParentPermissions) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBodyParentParentPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBodyParentPermissions Permission represents a set of permissions
swagger:model CreateCurrentUserRepoCreatedBodyParentPermissions
*/
type CreateCurrentUserRepoCreatedBodyParentPermissions struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// pull
	Pull bool `json:"pull,omitempty"`

	// push
	Push bool `json:"push,omitempty"`
}

// Validate validates this create current user repo created body parent permissions
func (o *CreateCurrentUserRepoCreatedBodyParentPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentPermissions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyParentPermissions) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBodyParentPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateCurrentUserRepoCreatedBodyPermissions Permission represents a set of permissions
swagger:model CreateCurrentUserRepoCreatedBodyPermissions
*/
type CreateCurrentUserRepoCreatedBodyPermissions struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// pull
	Pull bool `json:"pull,omitempty"`

	// push
	Push bool `json:"push,omitempty"`
}

// Validate validates this create current user repo created body permissions
func (o *CreateCurrentUserRepoCreatedBodyPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyPermissions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCurrentUserRepoCreatedBodyPermissions) UnmarshalBinary(b []byte) error {
	var res CreateCurrentUserRepoCreatedBodyPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}