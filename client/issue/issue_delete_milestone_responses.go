// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// IssueDeleteMilestoneReader is a Reader for the IssueDeleteMilestone structure.
type IssueDeleteMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueDeleteMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewIssueDeleteMilestoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueDeleteMilestoneNoContent creates a IssueDeleteMilestoneNoContent with default headers values
func NewIssueDeleteMilestoneNoContent() *IssueDeleteMilestoneNoContent {
	return &IssueDeleteMilestoneNoContent{}
}

/*IssueDeleteMilestoneNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type IssueDeleteMilestoneNoContent struct {
}

func (o *IssueDeleteMilestoneNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/milestones/{id}][%d] issueDeleteMilestoneNoContent ", 204)
}

func (o *IssueDeleteMilestoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
