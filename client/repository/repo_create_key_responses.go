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

// RepoCreateKeyReader is a Reader for the RepoCreateKey structure.
type RepoCreateKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreateKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewRepoCreateKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoCreateKeyCreated creates a RepoCreateKeyCreated with default headers values
func NewRepoCreateKeyCreated() *RepoCreateKeyCreated {
	return &RepoCreateKeyCreated{}
}

/*RepoCreateKeyCreated handles this case with default header values.

DeployKey
*/
type RepoCreateKeyCreated struct {
	Payload RepoCreateKeyCreatedBody
}

func (o *RepoCreateKeyCreated) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/keys][%d] repoCreateKeyCreated  %+v", 201, o.Payload)
}

func (o *RepoCreateKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RepoCreateKeyBody CreateKeyOption options when creating a key
swagger:model RepoCreateKeyBody
*/
type RepoCreateKeyBody struct {

	// An armored SSH key to add
	// Required: true
	// Unique: true
	Key *string `json:"key"`

	// Describe if the key has only read access or read/write
	ReadOnly bool `json:"read_only,omitempty"`

	// Title of the key to add
	// Required: true
	// Unique: true
	Title *string `json:"title"`
}

// MarshalBinary interface implementation
func (o *RepoCreateKeyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoCreateKeyBody) UnmarshalBinary(b []byte) error {
	var res RepoCreateKeyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RepoCreateKeyCreatedBody DeployKey a deploy key
swagger:model RepoCreateKeyCreatedBody
*/
type RepoCreateKeyCreatedBody struct {

	// created
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created_at"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// key
	// Required: true
	Key *string `json:"key"`

	// read only
	// Required: true
	ReadOnly *bool `json:"read_only"`

	// title
	// Required: true
	Title *string `json:"title"`

	// URL
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this repo create key created body
func (o *RepoCreateKeyCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReadOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RepoCreateKeyCreatedBody) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateKeyCreated"+"."+"created_at", "body", o.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("repoCreateKeyCreated"+"."+"created_at", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateKeyCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateKeyCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateKeyCreatedBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateKeyCreated"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateKeyCreatedBody) validateReadOnly(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateKeyCreated"+"."+"read_only", "body", o.ReadOnly); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateKeyCreatedBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateKeyCreated"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

func (o *RepoCreateKeyCreatedBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("repoCreateKeyCreated"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RepoCreateKeyCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoCreateKeyCreatedBody) UnmarshalBinary(b []byte) error {
	var res RepoCreateKeyCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
