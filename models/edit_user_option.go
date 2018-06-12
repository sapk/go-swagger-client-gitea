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

// EditUserOption EditUserOption edit user options
// swagger:model EditUserOption
type EditUserOption struct {

	// active
	Active bool `json:"active,omitempty"`

	// admin
	Admin bool `json:"admin,omitempty"`

	// allow git hook
	AllowGitHook bool `json:"allow_git_hook,omitempty"`

	// allow import local
	AllowImportLocal bool `json:"allow_import_local,omitempty"`

	// email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// login name
	LoginName string `json:"login_name,omitempty"`

	// max repo creation
	MaxRepoCreation int64 `json:"max_repo_creation,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// source ID
	SourceID int64 `json:"source_id,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this edit user option
func (m *EditUserOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditUserOption) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EditUserOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditUserOption) UnmarshalBinary(b []byte) error {
	var res EditUserOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
