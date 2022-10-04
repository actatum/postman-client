// Package monitors provides types/client for making requests to /monitors.
package monitors

import (
	"context"
	"fmt"
	"net/http"

	"github.com/actatum/postman-client/rest"
)

const path = "/monitors"

// Client handles monitors operations.
type Client struct {
	restClient *rest.Client
}

// NewClient returns a new instance of Client.
func NewClient(restClient *rest.Client) *Client {
	return &Client{
		restClient: restClient,
	}
}

// Create sends a POST request to /monitors.
func (c *Client) Create(
	ctx context.Context,
	monitor Monitor,
	opts ...rest.RequestOption,
) (Monitor, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		monitorWrapper{Monitor: monitor},
	)
	if err != nil {
		return Monitor{}, err
	}

	var response monitorWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Monitor, err
}

// Get sends a GET request to /monitors/:id.
func (c *Client) Get(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (Monitor, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return Monitor{}, err
	}

	var response monitorWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Monitor, err
}

// GetAll sends a GET request to /monitors.
func (c *Client) GetAll(
	ctx context.Context,
	opts ...rest.RequestOption,
) ([]Monitor, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var response monitorWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Monitors, err
}

// Update sends a PUT request to /monitors/:id.
func (c *Client) Update(
	ctx context.Context,
	id string,
	monitor Monitor,
	opts ...rest.RequestOption,
) (Monitor, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		monitorWrapper{Monitor: monitor},
	)
	if err != nil {
		return Monitor{}, err
	}

	var response monitorWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Monitor, err
}

// Delete sends a DELETE request to /monitors/:id.
func (c *Client) Delete(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (Monitor, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return Monitor{}, err
	}

	var response monitorWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Monitor, err
}

// RunMonitor sends a POST request to /monitors/:id/run
func (c *Client) RunMonitor(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (Run, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s/%s/run", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return Run{}, err
	}

	var response runWrapper
	err = c.restClient.DoRequest(r, &response)

	return response.Run, err
}
