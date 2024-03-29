// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserCurrentCheckStarringReader is a Reader for the UserCurrentCheckStarring structure.
type UserCurrentCheckStarringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentCheckStarringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUserCurrentCheckStarringNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUserCurrentCheckStarringNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentCheckStarringNoContent creates a UserCurrentCheckStarringNoContent with default headers values
func NewUserCurrentCheckStarringNoContent() *UserCurrentCheckStarringNoContent {
	return &UserCurrentCheckStarringNoContent{}
}

/*UserCurrentCheckStarringNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type UserCurrentCheckStarringNoContent struct {
}

func (o *UserCurrentCheckStarringNoContent) Error() string {
	return fmt.Sprintf("[GET /user/starred/{owner}/{repo}][%d] userCurrentCheckStarringNoContent ", 204)
}

func (o *UserCurrentCheckStarringNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserCurrentCheckStarringNotFound creates a UserCurrentCheckStarringNotFound with default headers values
func NewUserCurrentCheckStarringNotFound() *UserCurrentCheckStarringNotFound {
	return &UserCurrentCheckStarringNotFound{}
}

/*UserCurrentCheckStarringNotFound handles this case with default header values.

APINotFound is a not found empty response
*/
type UserCurrentCheckStarringNotFound struct {
}

func (o *UserCurrentCheckStarringNotFound) Error() string {
	return fmt.Sprintf("[GET /user/starred/{owner}/{repo}][%d] userCurrentCheckStarringNotFound ", 404)
}

func (o *UserCurrentCheckStarringNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
