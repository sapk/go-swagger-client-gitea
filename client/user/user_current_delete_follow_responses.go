// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserCurrentDeleteFollowReader is a Reader for the UserCurrentDeleteFollow structure.
type UserCurrentDeleteFollowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentDeleteFollowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUserCurrentDeleteFollowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentDeleteFollowNoContent creates a UserCurrentDeleteFollowNoContent with default headers values
func NewUserCurrentDeleteFollowNoContent() *UserCurrentDeleteFollowNoContent {
	return &UserCurrentDeleteFollowNoContent{}
}

/*UserCurrentDeleteFollowNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type UserCurrentDeleteFollowNoContent struct {
}

func (o *UserCurrentDeleteFollowNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user/following/{username}][%d] userCurrentDeleteFollowNoContent ", 204)
}

func (o *UserCurrentDeleteFollowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
