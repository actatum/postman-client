package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// APISecurityClient implements the functions that manage api security resources.
type APISecurityClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewAPISecurityClient returns a new instance of APISecurityClient.
func NewAPISecurityClient(apiKey string, httpClient *http.Client) *APISecurityClient {
	return &APISecurityClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/security",
	}
}

func (c *APISecurityClient) ValidateSchema(
	ctx context.Context,
	req APISecuritySchemaValidationRequest,
	opts ...RequestOption,
) (APISecuritySchemaValidationResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return APISecuritySchemaValidationResponse{}, err
	}

	return c.ValidateSchemaWithJSON(ctx, data, opts...)
}

func (c *APISecurityClient) ValidateSchemaWithJSON(
	ctx context.Context,
	jsonData []byte,
	opts ...RequestOption,
) (APISecuritySchemaValidationResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api-validation", bytes.NewBuffer(jsonData))
	if err != nil {
		return APISecuritySchemaValidationResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return APISecuritySchemaValidationResponse{}, err
	}

	var response APISecuritySchemaValidationResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type APISecuritySchemaValidationRequest struct {
	Schema struct {
		Type     string `json:"type"`
		Language string `json:"language"`
		Schema   string `json:"schema"`
	} `json:"schema"`
}

type APISecuritySchemaValidationResponse struct {
	Warnings []struct {
		Severity string `json:"severity"`
		Message  string `json:"message"`
		Location struct {
			Start struct {
				Line   int `json:"line"`
				Column int `json:"column"`
			} `json:"start"`
			End struct {
				Line   int `json:"line"`
				Column int `json:"column"`
			} `json:"end"`
		} `json:"location"`
		DataPath       []string `json:"dataPath"`
		PossibleFixUrl string   `json:"possibleFixUrl"`
		Category       struct {
			Name string `json:"name"`
		} `json:"category"`
	} `json:"warnings"`
}
