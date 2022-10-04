// Package collections provides types/client for making requests to /collections.
package collections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/actatum/postman-client/rest"
)

const path = "/collections"

// Client handles collections operations.
type Client struct {
	restClient *rest.Client
}

// NewClient returns a new instance of Client.
func NewClient(restClient *rest.Client) *Client {
	return &Client{
		restClient: restClient,
	}
}

// Create sends a POST request to /collections.
func (c *Client) Create(
	ctx context.Context,
	details CollectionDetails,
	opts ...rest.RequestOption,
) (Collection, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		collectionDetailsWrapper{Details: details},
	)
	if err != nil {
		return Collection{}, err
	}

	var response collectionWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Collection, err
}

// Get sends a GET request to /collections/:id.
func (c *Client) Get(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (CollectionDetails, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return CollectionDetails{}, err
	}

	var response collectionDetailsWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Details, err
}

// GetAll sends a GET request to /collections.
func (c *Client) GetAll(
	ctx context.Context,
	opts ...rest.RequestOption,
) ([]Collection, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var response collectionWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Collections, err
}

// Update sends a PUT request to /collections/:id.
func (c *Client) Update(
	ctx context.Context,
	id string,
	details CollectionDetails,
	opts ...rest.RequestOption,
) (Collection, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		collectionDetailsWrapper{Details: details},
	)
	if err != nil {
		return Collection{}, err
	}

	var response collectionWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Collection, err
}

// Delete sends a DELETE request to /collections/:id.
func (c *Client) Delete(
	ctx context.Context,
	id string,
	opts ...rest.RequestOption,
) (Collection, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s%s/%s", c.restClient.BaseURL(), path, id),
		nil,
	)
	if err != nil {
		return Collection{}, err
	}

	var response collectionWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Collection, err
}

// CreateFork sends a POST request to /collections/fork/:id
func (c *Client) CreateFork(
	ctx context.Context,
	id string,
	label string,
	opts ...rest.RequestOption,
) (Collection, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s%s/%s", c.restClient.BaseURL(), path, "/fork", id),
		map[string]string{"label": label},
	)
	if err != nil {
		return Collection{}, err
	}

	var response collectionWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Collection, err
}

// MergeFork sends a POST request to /collections/merge
func (c *Client) MergeFork(
	ctx context.Context,
	req MergeForkRequest,
	opts ...rest.RequestOption,
) (Collection, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s%s", c.restClient.BaseURL(), path, "/merge"),
		req,
	)
	if err != nil {
		return Collection{}, err
	}

	var response collectionWrapper
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Collection, err
}
