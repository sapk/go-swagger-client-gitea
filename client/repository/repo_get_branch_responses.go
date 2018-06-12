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
)

// RepoGetBranchReader is a Reader for the RepoGetBranch structure.
type RepoGetBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRepoGetBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoGetBranchOK creates a RepoGetBranchOK with default headers values
func NewRepoGetBranchOK() *RepoGetBranchOK {
	return &RepoGetBranchOK{}
}

/*RepoGetBranchOK handles this case with default header values.

Branch
*/
type RepoGetBranchOK struct {
	Payload RepoGetBranchOKBody
}

func (o *RepoGetBranchOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/branches/{branch}][%d] repoGetBranchOK  %+v", 200, o.Payload)
}

func (o *RepoGetBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RepoGetBranchOKBody Branch represents a repository branch
swagger:model RepoGetBranchOKBody
*/
type RepoGetBranchOKBody struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// commit
	// Required: true
	Commit *RepoGetBranchOKBodyCommit `json:"commit"`
}

// Validate validates this repo get branch o k body
func (o *RepoGetBranchOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCommit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoGetBranchOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("repoGetBranchOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *RepoGetBranchOKBody) validateCommit(formats strfmt.Registry) error {

	if err := validate.Required("repoGetBranchOK"+"."+"commit", "body", o.Commit); err != nil {
		return err
	}

	if o.Commit != nil {
		if err := o.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoGetBranchOK" + "." + "commit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoGetBranchOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoGetBranchOKBody) UnmarshalBinary(b []byte) error {
	var res RepoGetBranchOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoGetBranchOKBodyCommit PayloadCommit represents a commit
swagger:model RepoGetBranchOKBodyCommit
*/
type RepoGetBranchOKBodyCommit struct {

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
	Author *RepoGetBranchOKBodyCommitAuthor `json:"author,omitempty"`

	// committer
	Committer *RepoGetBranchOKBodyCommitCommitter `json:"committer,omitempty"`

	// verification
	Verification *RepoGetBranchOKBodyCommitVerification `json:"verification,omitempty"`
}

// Validate validates this repo get branch o k body commit
func (o *RepoGetBranchOKBodyCommit) Validate(formats strfmt.Registry) error {
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

func (o *RepoGetBranchOKBodyCommit) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("repoGetBranchOK"+"."+"commit"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoGetBranchOKBodyCommit) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(o.Author) { // not required
		return nil
	}

	if o.Author != nil {
		if err := o.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoGetBranchOK" + "." + "commit" + "." + "author")
			}
			return err
		}
	}

	return nil
}

func (o *RepoGetBranchOKBodyCommit) validateCommitter(formats strfmt.Registry) error {

	if swag.IsZero(o.Committer) { // not required
		return nil
	}

	if o.Committer != nil {
		if err := o.Committer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoGetBranchOK" + "." + "commit" + "." + "committer")
			}
			return err
		}
	}

	return nil
}

func (o *RepoGetBranchOKBodyCommit) validateVerification(formats strfmt.Registry) error {

	if swag.IsZero(o.Verification) { // not required
		return nil
	}

	if o.Verification != nil {
		if err := o.Verification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repoGetBranchOK" + "." + "commit" + "." + "verification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoGetBranchOKBodyCommit) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoGetBranchOKBodyCommit) UnmarshalBinary(b []byte) error {
	var res RepoGetBranchOKBodyCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoGetBranchOKBodyCommitAuthor PayloadUser represents the author or committer of a commit
swagger:model RepoGetBranchOKBodyCommitAuthor
*/
type RepoGetBranchOKBodyCommitAuthor struct {

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// Full name of the commit author
	Name string `json:"name,omitempty"`

	// user name
	UserName string `json:"username,omitempty"`
}

// Validate validates this repo get branch o k body commit author
func (o *RepoGetBranchOKBodyCommitAuthor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoGetBranchOKBodyCommitAuthor) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("repoGetBranchOK"+"."+"commit"+"."+"author"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoGetBranchOKBodyCommitAuthor) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoGetBranchOKBodyCommitAuthor) UnmarshalBinary(b []byte) error {
	var res RepoGetBranchOKBodyCommitAuthor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoGetBranchOKBodyCommitCommitter PayloadUser represents the author or committer of a commit
swagger:model RepoGetBranchOKBodyCommitCommitter
*/
type RepoGetBranchOKBodyCommitCommitter struct {

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// Full name of the commit author
	Name string `json:"name,omitempty"`

	// user name
	UserName string `json:"username,omitempty"`
}

// Validate validates this repo get branch o k body commit committer
func (o *RepoGetBranchOKBodyCommitCommitter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoGetBranchOKBodyCommitCommitter) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("repoGetBranchOK"+"."+"commit"+"."+"committer"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoGetBranchOKBodyCommitCommitter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoGetBranchOKBodyCommitCommitter) UnmarshalBinary(b []byte) error {
	var res RepoGetBranchOKBodyCommitCommitter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoGetBranchOKBodyCommitVerification PayloadCommitVerification represents the GPG verification of a commit
swagger:model RepoGetBranchOKBodyCommitVerification
*/
type RepoGetBranchOKBodyCommitVerification struct {

	// payload
	Payload string `json:"payload,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// signature
	Signature string `json:"signature,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this repo get branch o k body commit verification
func (o *RepoGetBranchOKBodyCommitVerification) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RepoGetBranchOKBodyCommitVerification) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoGetBranchOKBodyCommitVerification) UnmarshalBinary(b []byte) error {
	var res RepoGetBranchOKBodyCommitVerification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
