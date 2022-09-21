package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/actatum/postman-client/gen/types"
)

// FullCollectionWrapper represents a wrapper around postman collections for requests and response that use the full collection object.
type FullCollectionWrapper struct {
	Collection types.CollectionsJson `json:"collection"`
}

// MiniCollectionWrapper represents a wrapper around postman collections for responses that use the minified collection object.
type MiniCollectionWrapper struct {
	Collection MiniCollectionsJson `json:"collection"`
}

// CollectionsWrapper represents a wrapper around a list of postman collections that use the minified collection object.
type CollectionsWrapper struct {
	Collections []MiniCollectionsJson `json:"collections"`
}

// MiniCollectionsJson represents a collection response.
type MiniCollectionsJson struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Owner     string        `json:"owner"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	UID       string        `json:"uid"`
	IsPublic  bool          `json:"isPublic"`
	Fork      *ForkResponse `json:"fork,omitempty"`
}

// ForkResponse represents a fork response.
type ForkResponse struct {
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"createdAt"`
	From      string    `json:"from"`
}

// CreateForkRequest represents the request body for creating a fork of a collection.
type CreateForkRequest struct {
	Label string `json:"label"`
}

// MergeRequest represents a request to merge a fork.
type MergeRequest struct {
	Strategy    string `json:"strategy"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
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

func (c *CollectionsClient) Create(ctx context.Context, in types.CollectionsJson, opts ...RequestOption) (MiniCollectionsJson, error) {
	wr := FullCollectionWrapper{Collection: in}
	data, err := json.Marshal(wr)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	return c.CreateWithJSON(ctx, data, opts...)
}

func (c *CollectionsClient) CreateWithJSON(ctx context.Context, jsonData []byte, opts ...RequestOption) (MiniCollectionsJson, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	fmt.Printf("Result: \n%s\n", string(respJSON))

	var wrapper MiniCollectionWrapper
	err = json.Unmarshal(respJSON, &wrapper)

	return wrapper.Collection, err
}

func (c *CollectionsClient) Fork(ctx context.Context, uid string, in CreateForkRequest, opts ...RequestOption) (MiniCollectionsJson, error) {
	data, err := json.Marshal(in)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	return c.ForkWithJSON(ctx, uid, data, opts...)
}
func (c *CollectionsClient) ForkWithJSON(ctx context.Context, uid string, jsonData []byte, opts ...RequestOption) (MiniCollectionsJson, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/fork/"+uid, bytes.NewBuffer(jsonData))
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	fmt.Printf("Result: \n%s\n", string(respJSON))

	var wrapper MiniCollectionWrapper
	err = json.Unmarshal(respJSON, &wrapper)

	return wrapper.Collection, err
}

func (c *CollectionsClient) Get(ctx context.Context, uid string, opts ...RequestOption) (types.CollectionsJson, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+uid, nil)
	if err != nil {
		return types.CollectionsJson{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return types.CollectionsJson{}, err
	}

	fmt.Printf("Result: \n%s\n", string(respJSON))

	var wrapper FullCollectionWrapper
	err = json.Unmarshal(respJSON, &wrapper)

	return wrapper.Collection, err
}

func (c *CollectionsClient) Delete(ctx context.Context, uid string, opts ...RequestOption) (MiniCollectionsJson, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+"/"+uid, nil)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return MiniCollectionsJson{}, nil
	}

	fmt.Printf("Result: \n%s\n", string(respJSON))

	var wrapper MiniCollectionWrapper
	err = json.Unmarshal(respJSON, &wrapper)

	return wrapper.Collection, err
}

func (c *CollectionsClient) Update(ctx context.Context, uid string, in types.CollectionsJson, opts ...RequestOption) (MiniCollectionsJson, error) {
	wr := FullCollectionWrapper{Collection: in}
	data, err := json.Marshal(wr)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	return c.UpdateWithJSON(ctx, uid, data, opts...)
}

func (c *CollectionsClient) UpdateWithJSON(ctx context.Context, uid string, jsonData []byte, opts ...RequestOption) (MiniCollectionsJson, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/"+uid, bytes.NewBuffer(jsonData))
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	fmt.Printf("Result: \n%s\n", string(respJSON))

	var wrapper MiniCollectionWrapper
	err = json.Unmarshal(respJSON, &wrapper)

	return wrapper.Collection, err
}

func (c *CollectionsClient) MergeFork(ctx context.Context, in MergeRequest, opts ...RequestOption) (MiniCollectionsJson, error) {
	data, err := json.Marshal(in)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	return c.MergeForkWithJSON(ctx, data, opts...)
}

func (c *CollectionsClient) MergeForkWithJSON(ctx context.Context, jsonData []byte, opts ...RequestOption) (MiniCollectionsJson, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/merge", bytes.NewBuffer(jsonData))
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return MiniCollectionsJson{}, err
	}

	fmt.Printf("Result: \n%s\n", string(respJSON))

	var wrapper MiniCollectionWrapper
	err = json.Unmarshal(respJSON, &wrapper)

	return wrapper.Collection, err
}

func (c *CollectionsClient) GetAll(ctx context.Context, opts ...RequestOption) ([]MiniCollectionsJson, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return nil, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return nil, err
	}

	var wrapper CollectionsWrapper
	err = json.Unmarshal(respJSON, &wrapper)

	return wrapper.Collections, err
}
