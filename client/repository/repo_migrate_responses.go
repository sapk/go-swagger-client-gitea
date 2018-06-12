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

// RepoMigrateReader is a Reader for the RepoMigrate structure.
type RepoMigrateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoMigrateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewRepoMigrateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoMigrateCreated creates a RepoMigrateCreated with default headers values
func NewRepoMigrateCreated() *RepoMigrateCreated {
	return &RepoMigrateCreated{}
}

/*RepoMigrateCreated handles this case with default header values.

Repository
*/
type RepoMigrateCreated struct {
	Payload RepoMigrateCreatedBody
}

func (o *RepoMigrateCreated) Error() string {
	return fmt.Sprintf("[POST /repos/migrate][%d] repoMigrateCreated  %+v", 201, o.Payload)
}

func (o *RepoMigrateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RepoMigrateBody MigrateRepoForm form for migrating repository
swagger:model RepoMigrateBody
*/
type RepoMigrateBody struct {

	// auth password
	AuthPassword string `json:"auth_password,omitempty"`

	// auth username
	AuthUsername string `json:"auth_username,omitempty"`

	// clone addr
	// Required: true
	CloneAddr *string `json:"clone_addr"`

	// description
	Description string `json:"description,omitempty"`

	// mirror
	Mirror bool `json:"mirror,omitempty"`

	// private
	Private bool `json:"private,omitempty"`

	// repo name
	// Required: true
	RepoName *string `json:"repo_name"`

	// UID
	// Required: true
	UID *int64 `json:"uid"`
}

// MarshalBinary interface implementation
func (o *RepoMigrateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateBody) UnmarshalBinary(b []byte) error {
	var res RepoMigrateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBody Repository represents a repository
swagger:model RepoMigrateCreatedBody
*/
type RepoMigrateCreatedBody struct {

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
	Owner *RepoMigrateCreatedBodyOwner `json:"owner"`

	// parent
	// Required: true
	Parent *RepoMigrateCreatedBodyParent `json:"parent"`

	// permissions
	// Required: true
	Permissions *RepoMigrateCreatedBodyPermissions `json:"permissions"`
}

// Validate validates this repo migrate created body
func (o *RepoMigrateCreatedBody) Validate(formats strfmt.Registry) error {
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

func (o *RepoMigrateCreatedBody) validateCloneURL(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"clone_url", "body", o.CloneURL); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"created_at", "body", o.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateDefaultBranch(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"default_branch", "body", o.DefaultBranch); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateEmpty(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"empty", "body", o.Empty); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateFork(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"fork", "body", o.Fork); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateForks(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"forks_count", "body", o.Forks); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"full_name", "body", o.FullName); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateHTMLURL(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"html_url", "body", o.HTMLURL); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateMirror(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"mirror", "body", o.Mirror); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateOpenIssues(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"open_issues_count", "body", o.OpenIssues); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validatePrivate(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"private", "body", o.Private); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateSSHURL(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"ssh_url", "body", o.SSHURL); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateStars(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"stars_count", "body", o.Stars); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"updated_at", "body", o.Updated); err != nil {
		return err
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"updated_at", "body", "date-time", o.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateWatchers(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"watchers_count", "body", o.Watchers); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateWebsite(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"website", "body", o.Website); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"owner", "body", o.Owner); err != nil {
		return err
	}

	if o.Owner != nil {
		if err := o.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validateParent(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"parent", "body", o.Parent); err != nil {
		return err
	}

	if o.Parent != nil {
		if err := o.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "parent")
			}
			return err
		}
	}

	return nil
}

func (o *RepoMigrateCreatedBody) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("repoMigrateCreated"+"."+"permissions", "body", o.Permissions); err != nil {
		return err
	}

	if o.Permissions != nil {
		if err := o.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBody) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBodyOwner User represents a user
swagger:model RepoMigrateCreatedBodyOwner
*/
type RepoMigrateCreatedBodyOwner struct {

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

// Validate validates this repo migrate created body owner
func (o *RepoMigrateCreatedBodyOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoMigrateCreatedBodyOwner) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"owner"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyOwner) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBodyOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBodyParent Repository represents a repository
swagger:model RepoMigrateCreatedBodyParent
*/
type RepoMigrateCreatedBodyParent struct {

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
	Owner *RepoMigrateCreatedBodyParentOwner `json:"owner,omitempty"`

	// parent
	Parent *RepoMigrateCreatedBodyParentParent `json:"parent,omitempty"`

	// permissions
	Permissions *RepoMigrateCreatedBodyParentPermissions `json:"permissions,omitempty"`
}

// Validate validates this repo migrate created body parent
func (o *RepoMigrateCreatedBodyParent) Validate(formats strfmt.Registry) error {
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

func (o *RepoMigrateCreatedBodyParent) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"parent"+"."+"created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBodyParent) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(o.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"parent"+"."+"updated_at", "body", "date-time", o.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBodyParent) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(o.Owner) { // not required
		return nil
	}

	if o.Owner != nil {
		if err := o.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "parent" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

func (o *RepoMigrateCreatedBodyParent) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(o.Parent) { // not required
		return nil
	}

	if o.Parent != nil {
		if err := o.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "parent" + "." + "parent")
			}
			return err
		}
	}

	return nil
}

func (o *RepoMigrateCreatedBodyParent) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(o.Permissions) { // not required
		return nil
	}

	if o.Permissions != nil {
		if err := o.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "parent" + "." + "permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParent) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBodyParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBodyParentOwner User represents a user
swagger:model RepoMigrateCreatedBodyParentOwner
*/
type RepoMigrateCreatedBodyParentOwner struct {

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

// Validate validates this repo migrate created body parent owner
func (o *RepoMigrateCreatedBodyParentOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoMigrateCreatedBodyParentOwner) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"parent"+"."+"owner"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentOwner) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBodyParentOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBodyParentParent Repository represents a repository
swagger:model RepoMigrateCreatedBodyParentParent
*/
type RepoMigrateCreatedBodyParentParent struct {

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
	Owner *RepoMigrateCreatedBodyParentParentOwner `json:"owner,omitempty"`

	// parent
	Parent *models.Repository `json:"parent,omitempty"`

	// permissions
	Permissions *RepoMigrateCreatedBodyParentParentPermissions `json:"permissions,omitempty"`
}

// Validate validates this repo migrate created body parent parent
func (o *RepoMigrateCreatedBodyParentParent) Validate(formats strfmt.Registry) error {
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

func (o *RepoMigrateCreatedBodyParentParent) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"parent"+"."+"parent"+"."+"created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBodyParentParent) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(o.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"parent"+"."+"parent"+"."+"updated_at", "body", "date-time", o.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoMigrateCreatedBodyParentParent) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(o.Owner) { // not required
		return nil
	}

	if o.Owner != nil {
		if err := o.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "parent" + "." + "parent" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

func (o *RepoMigrateCreatedBodyParentParent) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(o.Parent) { // not required
		return nil
	}

	if o.Parent != nil {
		if err := o.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "parent" + "." + "parent" + "." + "parent")
			}
			return err
		}
	}

	return nil
}

func (o *RepoMigrateCreatedBodyParentParent) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(o.Permissions) { // not required
		return nil
	}

	if o.Permissions != nil {
		if err := o.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoMigrateCreated" + "." + "parent" + "." + "parent" + "." + "permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentParent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentParent) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBodyParentParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBodyParentParentOwner User represents a user
swagger:model RepoMigrateCreatedBodyParentParentOwner
*/
type RepoMigrateCreatedBodyParentParentOwner struct {

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

// Validate validates this repo migrate created body parent parent owner
func (o *RepoMigrateCreatedBodyParentParentOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoMigrateCreatedBodyParentParentOwner) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("repoMigrateCreated"+"."+"parent"+"."+"parent"+"."+"owner"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentParentOwner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentParentOwner) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBodyParentParentOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBodyParentParentPermissions Permission represents a set of permissions
swagger:model RepoMigrateCreatedBodyParentParentPermissions
*/
type RepoMigrateCreatedBodyParentParentPermissions struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// pull
	Pull bool `json:"pull,omitempty"`

	// push
	Push bool `json:"push,omitempty"`
}

// Validate validates this repo migrate created body parent parent permissions
func (o *RepoMigrateCreatedBodyParentParentPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentParentPermissions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentParentPermissions) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBodyParentParentPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBodyParentPermissions Permission represents a set of permissions
swagger:model RepoMigrateCreatedBodyParentPermissions
*/
type RepoMigrateCreatedBodyParentPermissions struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// pull
	Pull bool `json:"pull,omitempty"`

	// push
	Push bool `json:"push,omitempty"`
}

// Validate validates this repo migrate created body parent permissions
func (o *RepoMigrateCreatedBodyParentPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentPermissions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyParentPermissions) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBodyParentPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoMigrateCreatedBodyPermissions Permission represents a set of permissions
swagger:model RepoMigrateCreatedBodyPermissions
*/
type RepoMigrateCreatedBodyPermissions struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// pull
	Pull bool `json:"pull,omitempty"`

	// push
	Push bool `json:"push,omitempty"`
}

// Validate validates this repo migrate created body permissions
func (o *RepoMigrateCreatedBodyPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyPermissions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoMigrateCreatedBodyPermissions) UnmarshalBinary(b []byte) error {
	var res RepoMigrateCreatedBodyPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
