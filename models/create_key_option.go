// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateKeyOption CreateKeyOption options when creating a key
// swagger:model CreateKeyOption
type CreateKeyOption struct {

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

// Validate validates this create key option
func (m *CreateKeyOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateKeyOption) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *CreateKeyOption) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateKeyOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateKeyOption) UnmarshalBinary(b []byte) error {
	var res CreateKeyOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
