// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// IssueDeleteCommentReader is a Reader for the IssueDeleteComment structure.
type IssueDeleteCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueDeleteCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewIssueDeleteCommentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueDeleteCommentNoContent creates a IssueDeleteCommentNoContent with default headers values
func NewIssueDeleteCommentNoContent() *IssueDeleteCommentNoContent {
	return &IssueDeleteCommentNoContent{}
}

/*IssueDeleteCommentNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type IssueDeleteCommentNoContent struct {
}

func (o *IssueDeleteCommentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/comments/{id}][%d] issueDeleteCommentNoContent ", 204)
}

func (o *IssueDeleteCommentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
