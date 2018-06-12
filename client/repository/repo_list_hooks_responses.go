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

// RepoListHooksReader is a Reader for the RepoListHooks structure.
type RepoListHooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListHooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRepoListHooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoListHooksOK creates a RepoListHooksOK with default headers values
func NewRepoListHooksOK() *RepoListHooksOK {
	return &RepoListHooksOK{}
}

/*RepoListHooksOK handles this case with default header values.

HookList
*/
type RepoListHooksOK struct {
	Payload []*models.RepoListHooksOKBodyItems0
}

func (o *RepoListHooksOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks][%d] repoListHooksOK  %+v", 200, o.Payload)
}

func (o *RepoListHooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RepoListHooksOKBodyItems0 Branch represents a repository branch
swagger:model RepoListHooksOKBodyItems0
*/
type RepoListHooksOKBodyItems0 struct {

	// name
	Name string `json:"name,omitempty"`

	// commit
	Commit *RepoListHooksOKBodyItems0Commit `json:"commit,omitempty"`
}

// Validate validates this repo list hooks o k body items0
func (o *RepoListHooksOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoListHooksOKBodyItems0) validateCommit(formats strfmt.Registry) error {

	if swag.IsZero(o.Commit) { // not required
		return nil
	}

	if o.Commit != nil {
		if err := o.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res RepoListHooksOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoListHooksOKBodyItems0Commit PayloadCommit represents a commit
swagger:model RepoListHooksOKBodyItems0Commit
*/
type RepoListHooksOKBodyItems0Commit struct {

	// sha1 hash of the commit
	ID string `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// author
	Author *RepoListHooksOKBodyItems0CommitAuthor `json:"author,omitempty"`

	// committer
	Committer *RepoListHooksOKBodyItems0CommitCommitter `json:"committer,omitempty"`

	// verification
	Verification *RepoListHooksOKBodyItems0CommitVerification `json:"verification,omitempty"`
}

// Validate validates this repo list hooks o k body items0 commit
func (o *RepoListHooksOKBodyItems0Commit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCommitter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVerification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoListHooksOKBodyItems0Commit) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("commit"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoListHooksOKBodyItems0Commit) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(o.Author) { // not required
		return nil
	}

	if o.Author != nil {
		if err := o.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit" + "." + "author")
			}
			return err
		}
	}

	return nil
}

func (o *RepoListHooksOKBodyItems0Commit) validateCommitter(formats strfmt.Registry) error {

	if swag.IsZero(o.Committer) { // not required
		return nil
	}

	if o.Committer != nil {
		if err := o.Committer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit" + "." + "committer")
			}
			return err
		}
	}

	return nil
}

func (o *RepoListHooksOKBodyItems0Commit) validateVerification(formats strfmt.Registry) error {

	if swag.IsZero(o.Verification) { // not required
		return nil
	}

	if o.Verification != nil {
		if err := o.Verification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit" + "." + "verification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0Commit) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0Commit) UnmarshalBinary(b []byte) error {
	var res RepoListHooksOKBodyItems0Commit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoListHooksOKBodyItems0CommitAuthor PayloadUser represents the author or committer of a commit
swagger:model RepoListHooksOKBodyItems0CommitAuthor
*/
type RepoListHooksOKBodyItems0CommitAuthor struct {

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// Full name of the commit author
	Name string `json:"name,omitempty"`

	// user name
	UserName string `json:"username,omitempty"`
}

// Validate validates this repo list hooks o k body items0 commit author
func (o *RepoListHooksOKBodyItems0CommitAuthor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoListHooksOKBodyItems0CommitAuthor) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("commit"+"."+"author"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0CommitAuthor) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0CommitAuthor) UnmarshalBinary(b []byte) error {
	var res RepoListHooksOKBodyItems0CommitAuthor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoListHooksOKBodyItems0CommitCommitter PayloadUser represents the author or committer of a commit
swagger:model RepoListHooksOKBodyItems0CommitCommitter
*/
type RepoListHooksOKBodyItems0CommitCommitter struct {

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// Full name of the commit author
	Name string `json:"name,omitempty"`

	// user name
	UserName string `json:"username,omitempty"`
}

// Validate validates this repo list hooks o k body items0 commit committer
func (o *RepoListHooksOKBodyItems0CommitCommitter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoListHooksOKBodyItems0CommitCommitter) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("commit"+"."+"committer"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0CommitCommitter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0CommitCommitter) UnmarshalBinary(b []byte) error {
	var res RepoListHooksOKBodyItems0CommitCommitter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoListHooksOKBodyItems0CommitVerification PayloadCommitVerification represents the GPG verification of a commit
swagger:model RepoListHooksOKBodyItems0CommitVerification
*/
type RepoListHooksOKBodyItems0CommitVerification struct {

	// payload
	Payload string `json:"payload,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// signature
	Signature string `json:"signature,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this repo list hooks o k body items0 commit verification
func (o *RepoListHooksOKBodyItems0CommitVerification) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0CommitVerification) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoListHooksOKBodyItems0CommitVerification) UnmarshalBinary(b []byte) error {
	var res RepoListHooksOKBodyItems0CommitVerification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}