package postman

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// AuditLogsClient implements the functions that manage aAuditLogs resources.
type AuditLogsClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewAuditLogsClient returns a new instance of AuditLogsClient.
func NewAuditLogsClient(apiKey string, httpClient *http.Client) *AuditLogsClient {
	return &AuditLogsClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/audit/logs",
	}
}

func (c *AuditLogsClient) Get(ctx context.Context, req GetAuditLogsRequest, opts ...RequestOption) (GetAuditLogsResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return GetAuditLogsResponse{}, err
	}
	q := r.URL.Query()
	if req.Since != nil {
		q.Add("since", *req.Since)
	}
	if req.Limit != nil {
		q.Add("limit", strconv.Itoa(*req.Limit))
	}
	r.URL.RawQuery = q.Encode()

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetAuditLogsResponse{}, err
	}

	var response GetAuditLogsResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type GetAuditLogsRequest struct {
	// Return logs created after the given time, in YYYY-MM-DD format.
	Since *string
	// The maximum number of audit events to return at once. The maximum value is 300.
	Limit *int
}

type GetAuditLogsResponse struct {
	Trails []struct {
		Id        int       `json:"id"`
		Ip        string    `json:"ip"`
		UserAgent string    `json:"userAgent"`
		Action    string    `json:"action"`
		Timestamp time.Time `json:"timestamp"`
		Message   string    `json:"message"`
		Data      struct {
			Actor struct {
				Name     string `json:"name"`
				Username string `json:"username"`
				Email    string `json:"email"`
				Id       int    `json:"id"`
				Active   bool   `json:"active"`
			} `json:"actor"`
			User struct {
				Name     string `json:"name"`
				Username string `json:"username"`
				Email    string `json:"email"`
				Id       int    `json:"id"`
			} `json:"user"`
			Team struct {
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"team"`
		} `json:"data"`
	} `json:"trails"`
}
