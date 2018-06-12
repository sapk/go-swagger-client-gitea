// Code generated by go-swagger; DO NOT EDIT.

package issue

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

// IssueCreateLabelReader is a Reader for the IssueCreateLabel structure.
type IssueCreateLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueCreateLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewIssueCreateLabelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueCreateLabelCreated creates a IssueCreateLabelCreated with default headers values
func NewIssueCreateLabelCreated() *IssueCreateLabelCreated {
	return &IssueCreateLabelCreated{}
}

/*IssueCreateLabelCreated handles this case with default header values.

Label
*/
type IssueCreateLabelCreated struct {
	Payload IssueCreateLabelCreatedBody
}

func (o *IssueCreateLabelCreated) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/labels][%d] issueCreateLabelCreated  %+v", 201, o.Payload)
}

func (o *IssueCreateLabelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IssueCreateLabelBody CreateLabelOption options for creating a label
swagger:model IssueCreateLabelBody
*/
type IssueCreateLabelBody struct {

	// color
	// Required: true
	Color *string `json:"color"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// MarshalBinary interface implementation
func (o *IssueCreateLabelBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueCreateLabelBody) UnmarshalBinary(b []byte) error {
	var res IssueCreateLabelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*IssueCreateLabelCreatedBody Label a label to an issue or a pr
swagger:model IssueCreateLabelCreatedBody
*/
type IssueCreateLabelCreatedBody struct {

	// color
	// Required: true
	Color *string `json:"color"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// URL
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this issue create label created body
func (o *IssueCreateLabelCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
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

func (o *IssueCreateLabelCreatedBody) validateColor(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateLabelCreated"+"."+"color", "body", o.Color); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateLabelCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateLabelCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateLabelCreatedBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateLabelCreated"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateLabelCreatedBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateLabelCreated"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *IssueCreateLabelCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueCreateLabelCreatedBody) UnmarshalBinary(b []byte) error {
	var res IssueCreateLabelCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
