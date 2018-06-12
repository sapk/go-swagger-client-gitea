// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UserAddEmail adds email addresses
*/
func (a *Client) UserAddEmail(params *UserAddEmailParams, authInfo runtime.ClientAuthInfoWriter) (*UserAddEmailCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAddEmailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userAddEmail",
		Method:             "POST",
		PathPattern:        "/user/emails",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserAddEmailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserAddEmailCreated), nil

}

/*
UserCheckFollowing checks if one user is following another user
*/
func (a *Client) UserCheckFollowing(params *UserCheckFollowingParams, authInfo runtime.ClientAuthInfoWriter) (*UserCheckFollowingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCheckFollowingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCheckFollowing",
		Method:             "GET",
		PathPattern:        "/users/{follower}/following/{followee}",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCheckFollowingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCheckFollowingNoContent), nil

}

/*
UserCreateToken creates an access token
*/
func (a *Client) UserCreateToken(params *UserCreateTokenParams, authInfo runtime.ClientAuthInfoWriter) (*UserCreateTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCreateTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCreateToken",
		Method:             "POST",
		PathPattern:        "/users/{username}/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCreateTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCreateTokenOK), nil

}

/*
UserCurrentCheckFollowing checks whether a user is followed by the authenticated user
*/
func (a *Client) UserCurrentCheckFollowing(params *UserCurrentCheckFollowingParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentCheckFollowingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentCheckFollowingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentCheckFollowing",
		Method:             "GET",
		PathPattern:        "/user/following/{username}",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentCheckFollowingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentCheckFollowingNoContent), nil

}

/*
UserCurrentCheckStarring whethers the authenticated is starring the repo
*/
func (a *Client) UserCurrentCheckStarring(params *UserCurrentCheckStarringParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentCheckStarringNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentCheckStarringParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentCheckStarring",
		Method:             "GET",
		PathPattern:        "/user/starred/{owner}/{repo}",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentCheckStarringReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentCheckStarringNoContent), nil

}

/*
UserCurrentDeleteFollow unfollows a user
*/
func (a *Client) UserCurrentDeleteFollow(params *UserCurrentDeleteFollowParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentDeleteFollowNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentDeleteFollowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentDeleteFollow",
		Method:             "DELETE",
		PathPattern:        "/user/following/{username}",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentDeleteFollowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentDeleteFollowNoContent), nil

}

/*
UserCurrentDeleteGPGKey removes a g p g key
*/
func (a *Client) UserCurrentDeleteGPGKey(params *UserCurrentDeleteGPGKeyParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentDeleteGPGKeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentDeleteGPGKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentDeleteGPGKey",
		Method:             "DELETE",
		PathPattern:        "/user/gpg_keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentDeleteGPGKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentDeleteGPGKeyNoContent), nil

}

/*
UserCurrentDeleteKey deletes a public key
*/
func (a *Client) UserCurrentDeleteKey(params *UserCurrentDeleteKeyParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentDeleteKeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentDeleteKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentDeleteKey",
		Method:             "DELETE",
		PathPattern:        "/user/keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentDeleteKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentDeleteKeyNoContent), nil

}

/*
UserCurrentDeleteStar unstars the given repo
*/
func (a *Client) UserCurrentDeleteStar(params *UserCurrentDeleteStarParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentDeleteStarNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentDeleteStarParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentDeleteStar",
		Method:             "DELETE",
		PathPattern:        "/user/starred/{owner}/{repo}",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentDeleteStarReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentDeleteStarNoContent), nil

}

/*
UserCurrentGetGPGKey gets a g p g key
*/
func (a *Client) UserCurrentGetGPGKey(params *UserCurrentGetGPGKeyParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentGetGPGKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentGetGPGKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentGetGPGKey",
		Method:             "GET",
		PathPattern:        "/user/gpg_keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentGetGPGKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentGetGPGKeyOK), nil

}

/*
UserCurrentGetKey gets a public key
*/
func (a *Client) UserCurrentGetKey(params *UserCurrentGetKeyParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentGetKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentGetKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentGetKey",
		Method:             "GET",
		PathPattern:        "/user/keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentGetKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentGetKeyOK), nil

}

/*
UserCurrentListFollowers lists the authenticated user s followers
*/
func (a *Client) UserCurrentListFollowers(params *UserCurrentListFollowersParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentListFollowersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentListFollowersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentListFollowers",
		Method:             "GET",
		PathPattern:        "/user/followers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentListFollowersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentListFollowersOK), nil

}

/*
UserCurrentListFollowing lists the users that the authenticated user is following
*/
func (a *Client) UserCurrentListFollowing(params *UserCurrentListFollowingParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentListFollowingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentListFollowingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentListFollowing",
		Method:             "GET",
		PathPattern:        "/user/following",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentListFollowingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentListFollowingOK), nil

}

/*
UserCurrentListGPGKeys lists the authenticated user s g p g keys
*/
func (a *Client) UserCurrentListGPGKeys(params *UserCurrentListGPGKeysParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentListGPGKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentListGPGKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentListGPGKeys",
		Method:             "GET",
		PathPattern:        "/user/gpg_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentListGPGKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentListGPGKeysOK), nil

}

/*
UserCurrentListKeys lists the authenticated user s public keys
*/
func (a *Client) UserCurrentListKeys(params *UserCurrentListKeysParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentListKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentListKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentListKeys",
		Method:             "GET",
		PathPattern:        "/user/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentListKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentListKeysOK), nil

}

/*
UserCurrentListRepos lists the repos that the authenticated user owns or has access to
*/
func (a *Client) UserCurrentListRepos(params *UserCurrentListReposParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentListReposOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentListReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentListRepos",
		Method:             "GET",
		PathPattern:        "/user/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentListReposReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentListReposOK), nil

}

/*
UserCurrentListStarred thes repos that the authenticated user has starred
*/
func (a *Client) UserCurrentListStarred(params *UserCurrentListStarredParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentListStarredOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentListStarredParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentListStarred",
		Method:             "GET",
		PathPattern:        "/user/starred",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentListStarredReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentListStarredOK), nil

}

/*
UserCurrentListSubscriptions lists repositories watched by the authenticated user
*/
func (a *Client) UserCurrentListSubscriptions(params *UserCurrentListSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentListSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentListSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentListSubscriptions",
		Method:             "GET",
		PathPattern:        "/user/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentListSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentListSubscriptionsOK), nil

}

/*
UserCurrentPostGPGKey creates a g p g key
*/
func (a *Client) UserCurrentPostGPGKey(params *UserCurrentPostGPGKeyParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentPostGPGKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentPostGPGKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentPostGPGKey",
		Method:             "POST",
		PathPattern:        "/user/gpg_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentPostGPGKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentPostGPGKeyCreated), nil

}

/*
UserCurrentPostKey creates a public key
*/
func (a *Client) UserCurrentPostKey(params *UserCurrentPostKeyParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentPostKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentPostKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentPostKey",
		Method:             "POST",
		PathPattern:        "/user/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentPostKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentPostKeyCreated), nil

}

/*
UserCurrentPutFollow follows a user
*/
func (a *Client) UserCurrentPutFollow(params *UserCurrentPutFollowParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentPutFollowNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentPutFollowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentPutFollow",
		Method:             "PUT",
		PathPattern:        "/user/following/{username}",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentPutFollowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentPutFollowNoContent), nil

}

/*
UserCurrentPutStar stars the given repo
*/
func (a *Client) UserCurrentPutStar(params *UserCurrentPutStarParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentPutStarNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentPutStarParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentPutStar",
		Method:             "PUT",
		PathPattern:        "/user/starred/{owner}/{repo}",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentPutStarReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentPutStarNoContent), nil

}

/*
UserCurrentTrackedTimes lists the current user s tracked times
*/
func (a *Client) UserCurrentTrackedTimes(params *UserCurrentTrackedTimesParams, authInfo runtime.ClientAuthInfoWriter) (*UserCurrentTrackedTimesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCurrentTrackedTimesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userCurrentTrackedTimes",
		Method:             "GET",
		PathPattern:        "/user/times",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserCurrentTrackedTimesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCurrentTrackedTimesOK), nil

}

/*
UserDeleteEmail deletes email addresses
*/
func (a *Client) UserDeleteEmail(params *UserDeleteEmailParams, authInfo runtime.ClientAuthInfoWriter) (*UserDeleteEmailNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserDeleteEmailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userDeleteEmail",
		Method:             "DELETE",
		PathPattern:        "/user/emails",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserDeleteEmailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserDeleteEmailNoContent), nil

}

/*
UserGet gets a user
*/
func (a *Client) UserGet(params *UserGetParams, authInfo runtime.ClientAuthInfoWriter) (*UserGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userGet",
		Method:             "GET",
		PathPattern:        "/users/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserGetOK), nil

}

/*
UserGetCurrent gets the authenticated user
*/
func (a *Client) UserGetCurrent(params *UserGetCurrentParams, authInfo runtime.ClientAuthInfoWriter) (*UserGetCurrentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGetCurrentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userGetCurrent",
		Method:             "GET",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserGetCurrentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserGetCurrentOK), nil

}

/*
UserGetTokens lists the authenticated user s access tokens
*/
func (a *Client) UserGetTokens(params *UserGetTokensParams, authInfo runtime.ClientAuthInfoWriter) (*UserGetTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGetTokensParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userGetTokens",
		Method:             "GET",
		PathPattern:        "/users/{username}/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserGetTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserGetTokensOK), nil

}

/*
UserListEmails lists the authenticated user s email addresses
*/
func (a *Client) UserListEmails(params *UserListEmailsParams, authInfo runtime.ClientAuthInfoWriter) (*UserListEmailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListEmailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userListEmails",
		Method:             "GET",
		PathPattern:        "/user/emails",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListEmailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListEmailsOK), nil

}

/*
UserListFollowers lists the given user s followers
*/
func (a *Client) UserListFollowers(params *UserListFollowersParams, authInfo runtime.ClientAuthInfoWriter) (*UserListFollowersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListFollowersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userListFollowers",
		Method:             "GET",
		PathPattern:        "/users/{username}/followers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListFollowersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListFollowersOK), nil

}

/*
UserListFollowing lists the users that the given user is following
*/
func (a *Client) UserListFollowing(params *UserListFollowingParams, authInfo runtime.ClientAuthInfoWriter) (*UserListFollowingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListFollowingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userListFollowing",
		Method:             "GET",
		PathPattern:        "/users/{username}/following",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListFollowingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListFollowingOK), nil

}

/*
UserListGPGKeys lists the given user s g p g keys
*/
func (a *Client) UserListGPGKeys(params *UserListGPGKeysParams, authInfo runtime.ClientAuthInfoWriter) (*UserListGPGKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListGPGKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userListGPGKeys",
		Method:             "GET",
		PathPattern:        "/users/{username}/gpg_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListGPGKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListGPGKeysOK), nil

}

/*
UserListKeys lists the given user s public keys
*/
func (a *Client) UserListKeys(params *UserListKeysParams, authInfo runtime.ClientAuthInfoWriter) (*UserListKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userListKeys",
		Method:             "GET",
		PathPattern:        "/users/{username}/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListKeysOK), nil

}

/*
UserListRepos lists the repos owned by the given user
*/
func (a *Client) UserListRepos(params *UserListReposParams, authInfo runtime.ClientAuthInfoWriter) (*UserListReposOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userListRepos",
		Method:             "GET",
		PathPattern:        "/users/{username}/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListReposReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListReposOK), nil

}

/*
UserListStarred thes repos that the given user has starred
*/
func (a *Client) UserListStarred(params *UserListStarredParams, authInfo runtime.ClientAuthInfoWriter) (*UserListStarredOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListStarredParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userListStarred",
		Method:             "GET",
		PathPattern:        "/users/{username}/starred",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListStarredReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListStarredOK), nil

}

/*
UserListSubscriptions lists the repositories watched by a user
*/
func (a *Client) UserListSubscriptions(params *UserListSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*UserListSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userListSubscriptions",
		Method:             "GET",
		PathPattern:        "/users/{username}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserListSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListSubscriptionsOK), nil

}

/*
UserSearch searches for users
*/
func (a *Client) UserSearch(params *UserSearchParams, authInfo runtime.ClientAuthInfoWriter) (*UserSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userSearch",
		Method:             "GET",
		PathPattern:        "/users/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserSearchOK), nil

}

/*
UserTrackedTimes lists a user s tracked times in a repo
*/
func (a *Client) UserTrackedTimes(params *UserTrackedTimesParams, authInfo runtime.ClientAuthInfoWriter) (*UserTrackedTimesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserTrackedTimesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userTrackedTimes",
		Method:             "GET",
		PathPattern:        "/repos/{owner}/{repo}/times/{user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserTrackedTimesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserTrackedTimesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
