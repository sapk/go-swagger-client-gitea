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

// IssueAddTimeReader is a Reader for the IssueAddTime structure.
type IssueAddTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueAddTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIssueAddTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewIssueAddTimeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewIssueAddTimeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueAddTimeOK creates a IssueAddTimeOK with default headers values
func NewIssueAddTimeOK() *IssueAddTimeOK {
	return &IssueAddTimeOK{}
}

/*IssueAddTimeOK handles this case with default header values.

TrackedTime
*/
type IssueAddTimeOK struct {
	Payload IssueAddTimeOKBody
}

func (o *IssueAddTimeOK) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{id}/times][%d] issueAddTimeOK  %+v", 200, o.Payload)
}

func (o *IssueAddTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueAddTimeBadRequest creates a IssueAddTimeBadRequest with default headers values
func NewIssueAddTimeBadRequest() *IssueAddTimeBadRequest {
	return &IssueAddTimeBadRequest{}
}

/*IssueAddTimeBadRequest handles this case with default header values.

APIError is error format response
*/
type IssueAddTimeBadRequest struct {
	Message string

	URL string
}

func (o *IssueAddTimeBadRequest) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{id}/times][%d] issueAddTimeBadRequest ", 400)
}

func (o *IssueAddTimeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

// NewIssueAddTimeForbidden creates a IssueAddTimeForbidden with default headers values
func NewIssueAddTimeForbidden() *IssueAddTimeForbidden {
	return &IssueAddTimeForbidden{}
}

/*IssueAddTimeForbidden handles this case with default header values.

APIError is error format response
*/
type IssueAddTimeForbidden struct {
	Message string

	URL string
}

func (o *IssueAddTimeForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{id}/times][%d] issueAddTimeForbidden ", 403)
}

func (o *IssueAddTimeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

/*IssueAddTimeBody AddTimeOption options for adding time to an issue
swagger:model IssueAddTimeBody
*/
type IssueAddTimeBody struct {

	// time in seconds
	// Required: true
	Time *int64 `json:"time"`
}

// MarshalBinary interface implementation
func (o *IssueAddTimeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueAddTimeBody) UnmarshalBinary(b []byte) error {
	var res IssueAddTimeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*IssueAddTimeOKBody TrackedTime worked time for an issue / pr
swagger:model IssueAddTimeOKBody
*/
type IssueAddTimeOKBody struct {

	// created
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// issue ID
	// Required: true
	IssueID *int64 `json:"issue_id"`

	// Time in seconds
	// Required: true
	Time *int64 `json:"time"`

	// user ID
	// Required: true
	UserID *int64 `json:"user_id"`
}

// Validate validates this issue add time o k body
func (o *IssueAddTimeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIssueID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IssueAddTimeOKBody) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("issueAddTimeOK"+"."+"created", "body", o.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("issueAddTimeOK"+"."+"created", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IssueAddTimeOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("issueAddTimeOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *IssueAddTimeOKBody) validateIssueID(formats strfmt.Registry) error {

	if err := validate.Required("issueAddTimeOK"+"."+"issue_id", "body", o.IssueID); err != nil {
		return err
	}

	return nil
}

func (o *IssueAddTimeOKBody) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("issueAddTimeOK"+"."+"time", "body", o.Time); err != nil {
		return err
	}

	return nil
}

func (o *IssueAddTimeOKBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("issueAddTimeOK"+"."+"user_id", "body", o.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *IssueAddTimeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueAddTimeOKBody) UnmarshalBinary(b []byte) error {
	var res IssueAddTimeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
