// Package workspaces provides types/client for making requests to /workspaces.
package workspaces

import (
	"context"
	"fmt"
	"net/http"

	"github.com/actatum/postman-client/rest"
)

const path = "/workspaces"

// Client handles workspaces operations.
type Client struct {
	restClient *rest.Client
}

// NewClient returns a new instance of Client.
func NewClient(restClient *rest.Client) *Client {
	return &Client{
		restClient: restClient,
	}
}

// Create sends a POST request to /workspaces.
func (c *Client) Create(
	ctx context.Context,
	workspace Workspace,
	opts ...rest.RequestOption,
) (Workspace, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		workspaceWrapper{Workspace: workspace},
	)
	if err != nil {
		return Workspace{}, err
	}

	var response workspaceWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Workspace, err
}

// Get sends a GET request to /workspaces/:id.
func (c *Client) Get(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (Workspace, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return Workspace{}, err
	}

	var response workspaceWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Workspace, err
}

// GetAll sends a GET request to /workspaces.
func (c *Client) GetAll(
	ctx context.Context,
	req GetAllWorkspacesRequest,
	opts ...rest.RequestOption,
) ([]Workspace, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		nil,
	)
	if err != nil {
		return nil, err
	}
	q := r.URL.Query()
	if req.Type != nil {
		q.Set("type", *req.Type)
	}
	r.URL.RawQuery = q.Encode()

	var response workspaceWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Workspaces, err
}

// Update sends a PUT request to /workspaces/:id.
func (c *Client) Update(
	ctx context.Context,
	id string,
	workspace Workspace,
	opts ...rest.RequestOption,
) (Workspace, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		workspaceWrapper{Workspace: workspace},
	)
	if err != nil {
		return Workspace{}, err
	}

	var response workspaceWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Workspace, err
}

// Delete sends a DELETE request to /workspaces/:id.
func (c *Client) Delete(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (Workspace, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return Workspace{}, err
	}

	var response workspaceWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Workspace, err
}
