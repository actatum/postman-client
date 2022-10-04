// Package auditlogs provides types/client for making requests to /audit/logs.
package auditlogs

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/actatum/postman-client/rest"
)

const path = "/audit/logs"

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

// Get sends a GET request to /audit/logs.
func (c *Client) Get(ctx context.Context, req GetAuditLogsRequest) (AuditLogs, error) {
	r, err := c.restClient.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s", c.restClient.BaseURL(), path),
		nil,
	)
	if err != nil {
		return AuditLogs{}, err
	}
	q := r.URL.Query()
	if req.Since != nil {
		q.Add("since", *req.Since)
	}
	if req.Until != nil {
		q.Add("until", *req.Until)
	}
	if req.Limit != nil {
		q.Add("limit", strconv.Itoa(*req.Limit))
	}
	if req.Cursor != nil {
		q.Add("cursor", strconv.Itoa(*req.Cursor))
	}
	if req.OrderBy != nil {
		q.Add("order_by", *req.OrderBy)
	}
	r.URL.RawQuery = q.Encode()

	var response AuditLogs
	err = c.restClient.DoRequest(r, &response)

	return response, err
}
