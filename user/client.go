// Package user provides types/client for making requests to /me.
package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/actatum/postman-client/rest"
)

const path = "/me"

// Client handles user operations.
type Client struct {
	restClient *rest.Client
}

// NewClient returns a new instance of Client.
func NewClient(restClient *rest.Client) *Client {
	return &Client{
		restClient: restClient,
	}
}

// GetAuthenticatedUser sends a GET request to /me.
func (c *Client) GetAuthenticatedUser(ctx context.Context) (User, []Operation, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		nil,
	)
	if err != nil {
		return User{}, nil, err
	}

	var response authenticatedUserWrapper
	err = c.restClient.DoRequest(r, &response)

	return response.User, response.Operations, err
}
