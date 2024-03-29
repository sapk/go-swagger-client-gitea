// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// IssueClearLabelsReader is a Reader for the IssueClearLabels structure.
type IssueClearLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueClearLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewIssueClearLabelsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueClearLabelsNoContent creates a IssueClearLabelsNoContent with default headers values
func NewIssueClearLabelsNoContent() *IssueClearLabelsNoContent {
	return &IssueClearLabelsNoContent{}
}

/*IssueClearLabelsNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type IssueClearLabelsNoContent struct {
}

func (o *IssueClearLabelsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels][%d] issueClearLabelsNoContent ", 204)
}

func (o *IssueClearLabelsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
