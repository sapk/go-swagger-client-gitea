// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgIsMemberReader is a Reader for the OrgIsMember structure.
type OrgIsMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgIsMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewOrgIsMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewOrgIsMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgIsMemberNoContent creates a OrgIsMemberNoContent with default headers values
func NewOrgIsMemberNoContent() *OrgIsMemberNoContent {
	return &OrgIsMemberNoContent{}
}

/*OrgIsMemberNoContent handles this case with default header values.

user is a member
*/
type OrgIsMemberNoContent struct {
	Payload OrgIsMemberNoContentBody
}

func (o *OrgIsMemberNoContent) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/members/{username}][%d] orgIsMemberNoContent  %+v", 204, o.Payload)
}

func (o *OrgIsMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgIsMemberNotFound creates a OrgIsMemberNotFound with default headers values
func NewOrgIsMemberNotFound() *OrgIsMemberNotFound {
	return &OrgIsMemberNotFound{}
}

/*OrgIsMemberNotFound handles this case with default header values.

user is not a member
*/
type OrgIsMemberNotFound struct {
	Payload OrgIsMemberNotFoundBody
}

func (o *OrgIsMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/members/{username}][%d] orgIsMemberNotFound  %+v", 404, o.Payload)
}

func (o *OrgIsMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OrgIsMemberNoContentBody APIEmpty is an empty response
swagger:model OrgIsMemberNoContentBody
*/
type OrgIsMemberNoContentBody interface{}

/*OrgIsMemberNotFoundBody APIEmpty is an empty response
swagger:model OrgIsMemberNotFoundBody
*/
type OrgIsMemberNotFoundBody interface{}
