// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// RepoCreateReleaseReader is a Reader for the RepoCreateRelease structure.
type RepoCreateReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreateReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewRepoCreateReleaseCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoCreateReleaseCreated creates a RepoCreateReleaseCreated with default headers values
func NewRepoCreateReleaseCreated() *RepoCreateReleaseCreated {
	return &RepoCreateReleaseCreated{}
}

/*RepoCreateReleaseCreated handles this case with default header values.

Release
*/
type RepoCreateReleaseCreated struct {
	Payload RepoCreateReleaseCreatedBody
}

func (o *RepoCreateReleaseCreated) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/releases][%d] repoCreateReleaseCreated  %+v", 201, o.Payload)
}

func (o *RepoCreateReleaseCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AssetsItems0 Attachment a generic attachment
swagger:model AssetsItems0
*/
type AssetsItems0 struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// download count
	DownloadCount int64 `json:"download_count,omitempty"`

	// download URL
	DownloadURL string `json:"browser_download_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// UUID
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this assets items0
func (o *AssetsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssetsItems0) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AssetsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssetsItems0) UnmarshalBinary(b []byte) error {
	var res AssetsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoCreateReleaseBody CreateReleaseOption options when creating a release
swagger:model RepoCreateReleaseBody
*/
type RepoCreateReleaseBody struct {

	// is draft
	IsDraft bool `json:"draft,omitempty"`

	// is prerelease
	IsPrerelease bool `json:"prerelease,omitempty"`

	// note
	Note string `json:"body,omitempty"`

	// tag name
	// Required: true
	TagName *string `json:"tag_name"`

	// target
	Target string `json:"target_commitish,omitempty"`

	// title
	Title string `json:"name,omitempty"`
}

// MarshalBinary interface implementation
func (o *RepoCreateReleaseBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoCreateReleaseBody) UnmarshalBinary(b []byte) error {
	var res RepoCreateReleaseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoCreateReleaseCreatedBody Release represents a repository release
swagger:model RepoCreateReleaseCreatedBody
*/
type RepoCreateReleaseCreatedBody struct {

	// attachments
	// Required: true
	Attachments []*models.AssetsItems0 `json:"assets"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// is draft
	// Required: true
	IsDraft *bool `json:"draft"`

	// is prerelease
	// Required: true
	IsPrerelease *bool `json:"prerelease"`

	// note
	// Required: true
	Note *string `json:"body"`

	// published at
	// Required: true
	// Format: date-time
	PublishedAt *strfmt.DateTime `json:"published_at"`

	// tag name
	// Required: true
	TagName *string `json:"tag_name"`

	// tar URL
	// Required: true
	TarURL *string `json:"tarball_url"`

	// target
	// Required: true
	Target *string `json:"target_commitish"`

	// title
	// Required: true
	Title *string `json:"name"`

	// URL
	// Required: true
	URL *string `json:"url"`

	// zip URL
	// Required: true
	ZipURL *string `json:"zipball_url"`

	// author
	// Required: true
	Author *RepoCreateReleaseCreatedBodyAuthor `json:"author"`
}

// Validate validates this repo create release created body
func (o *RepoCreateReleaseCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsDraft(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsPrerelease(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePublishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTagName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTarURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateZipURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateAttachments(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"assets", "body", o.Attachments); err != nil {
		return err
	}

	for i := 0; i < len(o.Attachments); i++ {
		if swag.IsZero(o.Attachments[i]) { // not required
			continue
		}

		if o.Attachments[i] != nil {
			if err := o.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repoCreateReleaseCreated" + "." + "assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("repoCreateReleaseCreated"+"."+"created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateIsDraft(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"draft", "body", o.IsDraft); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateIsPrerelease(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"prerelease", "body", o.IsPrerelease); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateNote(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"body", "body", o.Note); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validatePublishedAt(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"published_at", "body", o.PublishedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("repoCreateReleaseCreated"+"."+"published_at", "body", "date-time", o.PublishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateTagName(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"tag_name", "body", o.TagName); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateTarURL(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"tarball_url", "body", o.TarURL); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"target_commitish", "body", o.Target); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"name", "body", o.Title); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateZipURL(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"zipball_url", "body", o.ZipURL); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateReleaseCreatedBody) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateReleaseCreated"+"."+"author", "body", o.Author); err != nil {
		return err
	}

	if o.Author != nil {
		if err := o.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoCreateReleaseCreated" + "." + "author")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoCreateReleaseCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoCreateReleaseCreatedBody) UnmarshalBinary(b []byte) error {
	var res RepoCreateReleaseCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoCreateReleaseCreatedBodyAuthor User represents a user
swagger:model RepoCreateReleaseCreatedBodyAuthor
*/
type RepoCreateReleaseCreatedBodyAuthor struct {

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

// Validate validates this repo create release created body author
func (o *RepoCreateReleaseCreatedBodyAuthor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoCreateReleaseCreatedBodyAuthor) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("repoCreateReleaseCreated"+"."+"author"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoCreateReleaseCreatedBodyAuthor) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoCreateReleaseCreatedBodyAuthor) UnmarshalBinary(b []byte) error {
	var res RepoCreateReleaseCreatedBodyAuthor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}