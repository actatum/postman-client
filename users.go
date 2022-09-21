package postman

import (
	"context"
	"encoding/json"
	"net/http"
)

// UsersClient implements the functions that manage users resources.
type UsersClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewUsersClient returns a new instance of UsersClient.
func NewUsersClient(apiKey string, httpClient *http.Client) *UsersClient {
	return &UsersClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/me",
	}
}

func (c *UsersClient) GetAuthenticatedUser(ctx context.Context, opts ...RequestOption) (GetAuthenticatedUserResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return GetAuthenticatedUserResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetAuthenticatedUserResponse{}, err
	}

	var response GetAuthenticatedUserResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type GetAuthenticatedUserResponse struct {
	User struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		FullName string `json:"fullName"`
		Avatar   string `json:"avatar"`
		IsPublic bool   `json:"isPublic"`
	} `json:"user"`
	Operations []struct {
		Name    string `json:"name"`
		Limit   int    `json:"limit"`
		Usage   int    `json:"usage"`
		Overage int    `json:"overage"`
	} `json:"operations"`
}
