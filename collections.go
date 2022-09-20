package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/moneymatchgg/postman-client/gen/types"
)

// Collection represents a postman collection.
type Collection struct {
	Collection types.CollectionsJson `json:"collection"`
}

// CollectionsClient implements the functions that manage collection resources.
type CollectionsClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewCollectionsClient returns a new instance of CollectionsClient
func NewCollectionsClient(apiKey string, httpClient *http.Client) *CollectionsClient {
	return &CollectionsClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/collections",
	}
}

func (c *CollectionsClient) Create(ctx context.Context, in types.CollectionsJson, opts ...RequestOption) (Collection, error) {
	coll := Collection{Collection: in}
	data, err := json.Marshal(coll)
	if err != nil {
		return Collection{}, err
	}

	return c.CreateWithJSON(ctx, data, opts...)
}

func (c *CollectionsClient) CreateWithJSON(ctx context.Context, jsonData []byte, opts ...RequestOption) (Collection, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return Collection{}, err
	}

	options := requestOptions{}

	for _, o := range opts {
		o.apply(&options)
	}

	respJSON, status, err := c.doRequest(r, options)
	if err != nil {
		return Collection{}, err
	}

	if status != http.StatusOK {
		var x interface{}
		err = json.Unmarshal(respJSON, &x)
		if err != nil {
			return Collection{}, err
		}
		fmt.Println(x)
	}

	fmt.Println(respJSON)

	return Collection{}, nil
}

func (c *CollectionsClient) doRequest(r *http.Request, options requestOptions) (data []byte, status int, err error) {
	query := r.URL.Query()

	if options.workspace != "" {
		query.Add("workspace", options.workspace)
	}

	r.URL.RawQuery = query.Encode()
	r.Header.Set("x-api-key", c.apiKey)
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	data, err = io.ReadAll(resp.Body)

	return data, resp.StatusCode, err
}
