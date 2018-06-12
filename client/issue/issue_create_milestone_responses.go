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

// IssueCreateMilestoneReader is a Reader for the IssueCreateMilestone structure.
type IssueCreateMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueCreateMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewIssueCreateMilestoneCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueCreateMilestoneCreated creates a IssueCreateMilestoneCreated with default headers values
func NewIssueCreateMilestoneCreated() *IssueCreateMilestoneCreated {
	return &IssueCreateMilestoneCreated{}
}

/*IssueCreateMilestoneCreated handles this case with default header values.

Milestone
*/
type IssueCreateMilestoneCreated struct {
	Payload IssueCreateMilestoneCreatedBody
}

func (o *IssueCreateMilestoneCreated) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/milestones][%d] issueCreateMilestoneCreated  %+v", 201, o.Payload)
}

func (o *IssueCreateMilestoneCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IssueCreateMilestoneBody CreateMilestoneOption options for creating a milestone
swagger:model IssueCreateMilestoneBody
*/
type IssueCreateMilestoneBody struct {

	// deadline
	// Format: date-time
	Deadline strfmt.DateTime `json:"due_on,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// MarshalBinary interface implementation
func (o *IssueCreateMilestoneBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueCreateMilestoneBody) UnmarshalBinary(b []byte) error {
	var res IssueCreateMilestoneBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*IssueCreateMilestoneCreatedBody Milestone milestone is a collection of issues on one repository
swagger:model IssueCreateMilestoneCreatedBody
*/
type IssueCreateMilestoneCreatedBody struct {

	// closed
	// Required: true
	// Format: date-time
	Closed *strfmt.DateTime `json:"closed_at"`

	// closed issues
	// Required: true
	ClosedIssues *int64 `json:"closed_issues"`

	// deadline
	// Required: true
	// Format: date-time
	Deadline *strfmt.DateTime `json:"due_on"`

	// description
	// Required: true
	Description *string `json:"description"`

	// ID
	// Required: true
	ID *int64 `json:"id"`

	// open issues
	// Required: true
	OpenIssues *int64 `json:"open_issues"`

	// title
	// Required: true
	Title *string `json:"title"`

	// StateType issue state type
	// Required: true
	State *string `json:"state"`
}

// Validate validates this issue create milestone created body
func (o *IssueCreateMilestoneCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClosed(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateClosedIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOpenIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IssueCreateMilestoneCreatedBody) validateClosed(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateMilestoneCreated"+"."+"closed_at", "body", o.Closed); err != nil {
		return err
	}

	if err := validate.FormatOf("issueCreateMilestoneCreated"+"."+"closed_at", "body", "date-time", o.Closed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateMilestoneCreatedBody) validateClosedIssues(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateMilestoneCreated"+"."+"closed_issues", "body", o.ClosedIssues); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateMilestoneCreatedBody) validateDeadline(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateMilestoneCreated"+"."+"due_on", "body", o.Deadline); err != nil {
		return err
	}

	if err := validate.FormatOf("issueCreateMilestoneCreated"+"."+"due_on", "body", "date-time", o.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateMilestoneCreatedBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateMilestoneCreated"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateMilestoneCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateMilestoneCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateMilestoneCreatedBody) validateOpenIssues(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateMilestoneCreated"+"."+"open_issues", "body", o.OpenIssues); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateMilestoneCreatedBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateMilestoneCreated"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

func (o *IssueCreateMilestoneCreatedBody) validateState(formats strfmt.Registry) error {

	if err := validate.Required("issueCreateMilestoneCreated"+"."+"state", "body", o.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *IssueCreateMilestoneCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueCreateMilestoneCreatedBody) UnmarshalBinary(b []byte) error {
	var res IssueCreateMilestoneCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
