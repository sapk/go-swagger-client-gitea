// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserCurrentPutFollowReader is a Reader for the UserCurrentPutFollow structure.
type UserCurrentPutFollowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentPutFollowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUserCurrentPutFollowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentPutFollowNoContent creates a UserCurrentPutFollowNoContent with default headers values
func NewUserCurrentPutFollowNoContent() *UserCurrentPutFollowNoContent {
	return &UserCurrentPutFollowNoContent{}
}

/*UserCurrentPutFollowNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type UserCurrentPutFollowNoContent struct {
}

func (o *UserCurrentPutFollowNoContent) Error() string {
	return fmt.Sprintf("[PUT /user/following/{username}][%d] userCurrentPutFollowNoContent ", 204)
}

func (o *UserCurrentPutFollowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
