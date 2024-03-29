// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepoDeleteReleaseReader is a Reader for the RepoDeleteRelease structure.
type RepoDeleteReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoDeleteReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRepoDeleteReleaseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoDeleteReleaseNoContent creates a RepoDeleteReleaseNoContent with default headers values
func NewRepoDeleteReleaseNoContent() *RepoDeleteReleaseNoContent {
	return &RepoDeleteReleaseNoContent{}
}

/*RepoDeleteReleaseNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type RepoDeleteReleaseNoContent struct {
}

func (o *RepoDeleteReleaseNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/releases/{id}][%d] repoDeleteReleaseNoContent ", 204)
}

func (o *RepoDeleteReleaseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
