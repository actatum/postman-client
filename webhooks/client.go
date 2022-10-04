// Package webhooks provides types/client for making requests to /webhooks.
package webhooks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/actatum/postman-client/rest"
)

const path = "/webhooks"

// Client handles webhook operations.
type Client struct {
	restClient *rest.Client
}

// NewClient returns a new instance of Client.
func NewClient(restClient *rest.Client) *Client {
	return &Client{
		restClient: restClient,
	}
}

// Create sends a POST request to /webhooks.
func (c *Client) Create(ctx context.Context, webhook Webhook) (Webhook, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		webhookWrapper{Webhook: webhook},
	)
	if err != nil {
		return Webhook{}, err
	}

	var response webhookWrapper
	err = c.restClient.DoRequest(r, &response)

	return response.Webhook, err
}
