package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// EnvironmentsClient implements the functions that manage environment resources.
type EnvironmentsClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewEnvironmentsClient returns a new instance of EnvironmentsClient.
func NewEnvironmentsClient(apiKey string, httpClient *http.Client) *EnvironmentsClient {
	return &EnvironmentsClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/environments",
	}
}

func (c *EnvironmentsClient) Create(ctx context.Context, req CreateEnvironmentRequest, opts ...RequestOption) (CreateEnvironmentResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return CreateEnvironmentResponse{}, err
	}

	return c.CreateWithJSON(ctx, data, opts...)
}

func (c *EnvironmentsClient) CreateWithJSON(ctx context.Context, jsonData []byte, opts ...RequestOption) (CreateEnvironmentResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return CreateEnvironmentResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return CreateEnvironmentResponse{}, err
	}

	var response CreateEnvironmentResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *EnvironmentsClient) Get(ctx context.Context, uid string, opts ...RequestOption) (GetEnvironmentResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+uid, nil)
	if err != nil {
		return GetEnvironmentResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetEnvironmentResponse{}, err
	}

	var response GetEnvironmentResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *EnvironmentsClient) GetAll(ctx context.Context, opts ...RequestOption) (GetAllEnvironmentsResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return GetAllEnvironmentsResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetAllEnvironmentsResponse{}, err
	}

	var response GetAllEnvironmentsResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *EnvironmentsClient) Update(ctx context.Context, req UpdateEnvironmentRequest, opts ...RequestOption) (UpdateEnvironmentResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return UpdateEnvironmentResponse{}, err
	}

	return c.UpdateWithJSON(ctx, data, opts...)
}

func (c *EnvironmentsClient) UpdateWithJSON(ctx context.Context, jsonData []byte, opts ...RequestOption) (UpdateEnvironmentResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPut, c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return UpdateEnvironmentResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return UpdateEnvironmentResponse{}, err
	}

	var response UpdateEnvironmentResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *EnvironmentsClient) Delete(ctx context.Context, uid string, opts ...RequestOption) (DeleteEnvironmentResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+"/"+uid, nil)
	if err != nil {
		return DeleteEnvironmentResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return DeleteEnvironmentResponse{}, err
	}

	var response DeleteEnvironmentResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type CreateEnvironmentRequest struct {
	Environment struct {
		Name   string `json:"name"`
		Values []struct {
			Key     string `json:"key"`
			Value   string `json:"value"`
			Enabled bool   `json:"enabled"`
			Type    string `json:"type"`
		} `json:"values"`
	} `json:"environment"`
}

type CreateEnvironmentResponse struct {
	Environment struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"environment"`
}

type GetEnvironmentResponse struct {
	Environment struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		Owner     string    `json:"owner"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Values    []struct {
			Key     string `json:"key"`
			Value   string `json:"value"`
			Enabled bool   `json:"enabled"`
			Type    string `json:"type"`
		} `json:"values"`
		IsPublic bool `json:"isPublic"`
	} `json:"environment"`
}

type GetAllEnvironmentsResponse struct {
	Environments []struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Owner     string    `json:"owner"`
		Uid       string    `json:"uid"`
		IsPublic  bool      `json:"isPublic"`
	} `json:"environments"`
}

type UpdateEnvironmentRequest struct {
	Environment struct {
		Name   string `json:"name"`
		Values []struct {
			Key     string `json:"key"`
			Value   string `json:"value"`
			Enabled bool   `json:"enabled"`
			Type    string `json:"type"`
		} `json:"values"`
	} `json:"environment"`
}

type UpdateEnvironmentResponse struct {
	Environment struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"environment"`
}

type DeleteEnvironmentResponse struct {
	Environment struct {
		Id  string `json:"id"`
		Uid string `json:"uid"`
	} `json:"environment"`
}
