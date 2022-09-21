package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// WebhooksClient implements the functions that manage webhooks resources.
type WebhooksClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewWebhooksClient returns a new instance of WebhooksClient.
func NewWebhooksClient(apiKey string, httpClient *http.Client) *WebhooksClient {
	return &WebhooksClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/webhooks",
	}
}

func (c *WebhooksClient) Create(ctx context.Context, req CreateWebhookRequest, opts ...RequestOption) (CreateWebhookResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return CreateWebhookResponse{}, err
	}

	return c.CreateWithJSON(ctx, data, opts...)
}

func (c *WebhooksClient) CreateWithJSON(ctx context.Context, jsonData []byte, opts ...RequestOption) (CreateWebhookResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return CreateWebhookResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return CreateWebhookResponse{}, err
	}

	var response CreateWebhookResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type CreateWebhookRequest struct {
	Webhook struct {
		Name       string `json:"name"`
		Collection string `json:"collection"`
	} `json:"webhook"`
}

type CreateWebhookResponse struct {
	Webhook struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Collection string `json:"collection"`
		WebhookUrl string `json:"webhookUrl"`
		Uid        string `json:"uid"`
	} `json:"webhook"`
}
