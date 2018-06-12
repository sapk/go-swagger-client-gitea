// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgAddTeamMemberReader is a Reader for the OrgAddTeamMember structure.
type OrgAddTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgAddTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewOrgAddTeamMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgAddTeamMemberNoContent creates a OrgAddTeamMemberNoContent with default headers values
func NewOrgAddTeamMemberNoContent() *OrgAddTeamMemberNoContent {
	return &OrgAddTeamMemberNoContent{}
}

/*OrgAddTeamMemberNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type OrgAddTeamMemberNoContent struct {
}

func (o *OrgAddTeamMemberNoContent) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}/members/{username}][%d] orgAddTeamMemberNoContent ", 204)
}

func (o *OrgAddTeamMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
