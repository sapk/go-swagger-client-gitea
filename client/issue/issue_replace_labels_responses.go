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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sapk/go-swagger-client-gitea/models"
)

// IssueReplaceLabelsReader is a Reader for the IssueReplaceLabels structure.
type IssueReplaceLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueReplaceLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIssueReplaceLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueReplaceLabelsOK creates a IssueReplaceLabelsOK with default headers values
func NewIssueReplaceLabelsOK() *IssueReplaceLabelsOK {
	return &IssueReplaceLabelsOK{}
}

/*IssueReplaceLabelsOK handles this case with default header values.

LabelList
*/
type IssueReplaceLabelsOK struct {
	Payload []*models.IssueReplaceLabelsOKBodyItems0
}

func (o *IssueReplaceLabelsOK) Error() string {
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/issues/{index}/labels][%d] issueReplaceLabelsOK  %+v", 200, o.Payload)
}

func (o *IssueReplaceLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IssueReplaceLabelsBody IssueLabelsOption a collection of labels
swagger:model IssueReplaceLabelsBody
*/
type IssueReplaceLabelsBody struct {

	// list of label IDs
	Labels []int64 `json:"labels"`
}

// MarshalBinary interface implementation
func (o *IssueReplaceLabelsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueReplaceLabelsBody) UnmarshalBinary(b []byte) error {
	var res IssueReplaceLabelsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*IssueReplaceLabelsOKBodyItems0 Label a label to an issue or a pr
swagger:model IssueReplaceLabelsOKBodyItems0
*/
type IssueReplaceLabelsOKBodyItems0 struct {

	// color
	Color string `json:"color,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// URL
	URL string `json:"url,omitempty"`
}

// Validate validates this issue replace labels o k body items0
func (o *IssueReplaceLabelsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *IssueReplaceLabelsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IssueReplaceLabelsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res IssueReplaceLabelsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
