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

// User User represents a user
// swagger:model User
type User struct {

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

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
