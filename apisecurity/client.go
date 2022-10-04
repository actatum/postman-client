// Package apisecurity provides types/client for making requests to /security/api-validation.
package apisecurity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/actatum/postman-client/rest"
)

const path = "/security/api-validation"

// Client handles api security operations.
type Client struct {
	restClient *rest.Client
}

// NewClient returns a new instance of Client.
func NewClient(restClient *rest.Client) *Client {
	return &Client{
		restClient: restClient,
	}
}

// ValidateAPISchema sends a POST request to /security/api-validation.
func (c *Client) ValidateAPISchema(
	ctx context.Context,
	req ValidateAPISchemaRequest,
	opts ...rest.RequestOption,
) ([]Warning, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		req,
	)
	if err != nil {
		return nil, err
	}

	var response ValidateAPISchemaResponse
	err = c.restClient.DoRequest(r, &response, opts...)

	return response.Warnings, err
}
