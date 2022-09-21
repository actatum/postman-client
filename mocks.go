package postman

import "net/http"

// MocksClient implements the functions that manage mock resources.
type MocksClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewMocksClient returns a new instance of MocksClient.
func NewMocksClient(apiKey string, httpClient *http.Client) *MocksClient {
	return &MocksClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/mocks",
	}
}
