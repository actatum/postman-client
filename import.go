package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// ImportClient implements the functions that manage import resources.
type ImportClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewImportClient returns a new instance of ImportClient.
func NewImportClient(apiKey string, httpClient *http.Client) *ImportClient {
	return &ImportClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/import",
	}
}

func (c *ImportClient) OpenAPIJSON(
	ctx context.Context,
	req ImportOpenAPIJSONRequest,
	opts ...RequestOption,
) (ImportOpenAPIResponse, error) {
	req.Type = "json"
	data, err := json.Marshal(req)
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/openapi", bytes.NewBuffer(data))
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}

	var response ImportOpenAPIResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *ImportClient) OpenAPIStringified(
	ctx context.Context,
	req ImportOpenAPIStringifiedRequest,
	opts ...RequestOption,
) (ImportOpenAPIResponse, error) {
	req.Type = "string"
	data, err := json.Marshal(req)
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/openapi", bytes.NewBuffer(data))
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}

	var response ImportOpenAPIResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *ImportClient) OpenAPIFile(
	ctx context.Context,
	fileName string,
	file io.Reader,
	opts ...RequestOption,
) (ImportOpenAPIResponse, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	fw, err := w.CreateFormFile("input", fileName)
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return ImportOpenAPIResponse{}, err
	}
	fw, err = w.CreateFormField("type")
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}
	if _, err = io.Copy(fw, strings.NewReader("file")); err != nil {
		return ImportOpenAPIResponse{}, err
	}
	if err = w.Close(); err != nil {
		return ImportOpenAPIResponse{}, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/openapi", buf)
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return ImportOpenAPIResponse{}, err
	}

	var response ImportOpenAPIResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *ImportClient) Exported(
	ctx context.Context,
	fileName string,
	file io.Reader,
	opts ...RequestOption,
) (ImportExportedResponse, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	fw, err := w.CreateFormFile("input", fileName)
	if err != nil {
		return ImportExportedResponse{}, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return ImportExportedResponse{}, err
	}
	fw, err = w.CreateFormField("type")
	if err != nil {
		return ImportExportedResponse{}, err
	}
	if _, err = io.Copy(fw, strings.NewReader("file")); err != nil {
		return ImportExportedResponse{}, err
	}
	if err = w.Close(); err != nil {
		return ImportExportedResponse{}, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/openapi", buf)
	if err != nil {
		return ImportExportedResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return ImportExportedResponse{}, err
	}

	var response ImportExportedResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type ImportOpenAPIJSONRequest struct {
	Type  string                 `json:"type"`
	Input map[string]interface{} `json:"input"`
}

type ImportOpenAPIStringifiedRequest struct {
	Type  string `json:"type"`
	Input string `json:"input"`
}

type ImportOpenAPIResponse struct {
	Collections []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"collections"`
	Environments []interface{} `json:"environments"`
}

type ImportExportedResponse struct {
	Collections []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"collections"`
}
