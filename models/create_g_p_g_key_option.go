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

// CreateGPGKeyOption CreateGPGKeyOption options create user GPG key
// swagger:model CreateGPGKeyOption
type CreateGPGKeyOption struct {

	// An armored GPG key to add
	// Required: true
	// Unique: true
	ArmoredKey *string `json:"armored_public_key"`
}

// Validate validates this create g p g key option
func (m *CreateGPGKeyOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArmoredKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGPGKeyOption) validateArmoredKey(formats strfmt.Registry) error {

	if err := validate.Required("armored_public_key", "body", m.ArmoredKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateGPGKeyOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGPGKeyOption) UnmarshalBinary(b []byte) error {
	var res CreateGPGKeyOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
