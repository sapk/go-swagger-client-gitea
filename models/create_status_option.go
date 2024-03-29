// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreateStatusOption CreateStatusOption holds the information needed to create a new Status for a Commit
// swagger:model CreateStatusOption
type CreateStatusOption struct {

	// context
	Context string `json:"context,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// target URL
	TargetURL string `json:"target_url,omitempty"`

	// StatusState holds the state of a Status
	// It can be "pending", "success", "error", "failure", and "warning"
	State string `json:"state,omitempty"`
}

// Validate validates this create status option
func (m *CreateStatusOption) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateStatusOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateStatusOption) UnmarshalBinary(b []byte) error {
	var res CreateStatusOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
