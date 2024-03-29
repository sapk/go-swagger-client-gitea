// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserCreateTokenReader is a Reader for the UserCreateToken structure.
type UserCreateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCreateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCreateTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCreateTokenOK creates a UserCreateTokenOK with default headers values
func NewUserCreateTokenOK() *UserCreateTokenOK {
	return &UserCreateTokenOK{}
}

/*UserCreateTokenOK handles this case with default header values.

AccessToken represents a API access token.
*/
type UserCreateTokenOK struct {
	Name string

	Sha1 string
}

func (o *UserCreateTokenOK) Error() string {
	return fmt.Sprintf("[POST /users/{username}/tokens][%d] userCreateTokenOK ", 200)
}

func (o *UserCreateTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header name
	o.Name = response.GetHeader("name")

	// response header sha1
	o.Sha1 = response.GetHeader("sha1")

	return nil
}
