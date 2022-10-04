// Package environments provides types/client for making requests to /environments.
package environments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/actatum/postman-client/rest"
)

const path = "/environments"

// Client handles environment operations.
type Client struct {
	restClient *rest.Client
}

// NewClient returns a new instance of Client.
func NewClient(restClient *rest.Client) *Client {
	return &Client{
		restClient: restClient,
	}
}

// Create sends a POST request to /environments.
func (c *Client) Create(
	ctx context.Context,
	env Environment,
	opts ...rest.RequestOption,
) (Environment, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		environmentWrapper{Environment: env},
	)
	if err != nil {
		return Environment{}, err
	}

	var response environmentWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Environment, err
}

// Get sends a GET request to /environments/:id.
func (c *Client) Get(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (Environment, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return Environment{}, err
	}

	var response environmentWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Environment, err
}

// GetAll sends a GET request to /environments.
func (c *Client) GetAll(
	ctx context.Context,
	opts ...rest.RequestOption,
) ([]Environment, error) {
	r, err := c.restClient.NewRequest(
		ctx, http.MethodGet,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var response environmentWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Environments, err
}

// Update sends a PUT request to /environments/:id.
func (c *Client) Update(
	ctx context.Context,
	id string,
	env Environment,
	opts ...rest.RequestOption,
) (Environment, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		environmentWrapper{Environment: env},
	)
	if err != nil {
		return Environment{}, err
	}

	var response environmentWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Environment, err
}

// Delete sends a DELETE request to /environments/:id.
func (c *Client) Delete(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (Environment, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return Environment{}, err
	}

	var response environmentWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Environment, err
}
