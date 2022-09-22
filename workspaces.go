package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// WorkspacesClient implements the functions that manage workspace resources.
type WorkspacesClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewWorkspacesClient returns a new instance of WorkspacesClient.
func NewWorkspacesClient(apiKey string, httpClient *http.Client) *WorkspacesClient {
	return &WorkspacesClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/workspaces",
	}
}

func (c *WorkspacesClient) Create(
	ctx context.Context,
	req CreateWorkspaceRequest,
	opts ...RequestOption,
) (CreateWorkspaceResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return CreateWorkspaceResponse{}, err
	}

	return c.CreateWithJSON(ctx, data, opts...)
}

func (c *WorkspacesClient) CreateWithJSON(
	ctx context.Context,
	jsonData []byte,
	opts ...RequestOption,
) (CreateWorkspaceResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return CreateWorkspaceResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return CreateWorkspaceResponse{}, err
	}

	var response CreateWorkspaceResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *WorkspacesClient) Get(
	ctx context.Context,
	id string,
	opts ...RequestOption,
) (GetWorkspaceResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+id, nil)
	if err != nil {
		return GetWorkspaceResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetWorkspaceResponse{}, err
	}

	var response GetWorkspaceResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *WorkspacesClient) GetAll(
	ctx context.Context,
	opts ...RequestOption,
) (GetAllWorkspacesResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return GetAllWorkspacesResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetAllWorkspacesResponse{}, err
	}

	var response GetAllWorkspacesResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *WorkspacesClient) Update(
	ctx context.Context,
	id string,
	req UpdateWorkspaceRequest,
	opts ...RequestOption,
) (UpdateWorkspaceResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return UpdateWorkspaceResponse{}, err
	}

	return c.UpdateWithJSON(ctx, id, data, opts...)
}

func (c *WorkspacesClient) UpdateWithJSON(
	ctx context.Context,
	id string,
	jsonData []byte,
	opts ...RequestOption,
) (UpdateWorkspaceResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPut, c.baseURL+"/"+id, bytes.NewBuffer(jsonData))
	if err != nil {
		return UpdateWorkspaceResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return UpdateWorkspaceResponse{}, err
	}

	var response UpdateWorkspaceResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *WorkspacesClient) Delete(
	ctx context.Context,
	id string,
	opts ...RequestOption,
) (DeleteWorkspaceResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+"/"+id, nil)
	if err != nil {
		return DeleteWorkspaceResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return DeleteWorkspaceResponse{}, err
	}

	var response DeleteWorkspaceResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type CreateWorkspaceRequest struct {
	Workspace struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"workspace"`
}

type CreateWorkspaceResponse struct {
	Workspace struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"workspace"`
}

type GetWorkspaceResponse struct {
	Workspace struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Type        string    `json:"type"`
		Description string    `json:"description"`
		Visibility  string    `json:"visibility"`
		CreatedBy   string    `json:"createdBy"`
		UpdatedBy   string    `json:"updatedBy"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		Collections []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Uid  string `json:"uid"`
		} `json:"collections"`
		Environments []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Uid  string `json:"uid"`
		} `json:"environments"`
		Mocks []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Uid  string `json:"uid"`
		} `json:"mocks"`
		Monitors []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Uid  string `json:"uid"`
		} `json:"monitors"`
		Apis []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Uid  string `json:"uid"`
		} `json:"apis"`
	} `json:"workspace"`
}

type GetAllWorkspacesResponse struct {
	Workspaces []struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Visibility string `json:"visibility"`
	} `json:"workspaces"`
}

type UpdateWorkspaceRequest struct {
	Workspace struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Type        string `json:"type"`
	} `json:"workspace"`
}

type UpdateWorkspaceResponse struct {
	Workspace struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"workspace"`
}

type DeleteWorkspaceResponse struct {
	Workspace struct {
		Id string `json:"id"`
	} `json:"workspace"`
}
