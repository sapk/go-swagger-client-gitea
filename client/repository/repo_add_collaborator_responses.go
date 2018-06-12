// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// RepoAddCollaboratorReader is a Reader for the RepoAddCollaborator structure.
type RepoAddCollaboratorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoAddCollaboratorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRepoAddCollaboratorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoAddCollaboratorNoContent creates a RepoAddCollaboratorNoContent with default headers values
func NewRepoAddCollaboratorNoContent() *RepoAddCollaboratorNoContent {
	return &RepoAddCollaboratorNoContent{}
}

/*RepoAddCollaboratorNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type RepoAddCollaboratorNoContent struct {
}

func (o *RepoAddCollaboratorNoContent) Error() string {
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/collaborators/{collaborator}][%d] repoAddCollaboratorNoContent ", 204)
}

func (o *RepoAddCollaboratorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*RepoAddCollaboratorBody AddCollaboratorOption options when adding a user as a collaborator of a repository
swagger:model RepoAddCollaboratorBody
*/
type RepoAddCollaboratorBody struct {

	// permission
	Permission string `json:"permission,omitempty"`
}

// MarshalBinary interface implementation
func (o *RepoAddCollaboratorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RepoAddCollaboratorBody) UnmarshalBinary(b []byte) error {
	var res RepoAddCollaboratorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}