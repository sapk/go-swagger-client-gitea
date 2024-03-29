// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreateForkOption CreateForkOption options for creating a fork
// swagger:model CreateForkOption
type CreateForkOption struct {

	// organization name, if forking into an organization
	Organization string `json:"organization,omitempty"`
}

// Validate validates this create fork option
func (m *CreateForkOption) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateForkOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateForkOption) UnmarshalBinary(b []byte) error {
	var res CreateForkOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
