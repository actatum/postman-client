package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// CollectionsClient implements the functions that manage collection resources.
type CollectionsClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewCollectionsClient returns a new instance of CollectionsClient.
func NewCollectionsClient(apiKey string, httpClient *http.Client) *CollectionsClient {
	return &CollectionsClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/collections",
	}
}

func (c *CollectionsClient) Create(
	ctx context.Context,
	req CreateCollectionRequest,
	opts ...RequestOption,
) (CreateCollectionResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return CreateCollectionResponse{}, err
	}

	return c.CreateWithJSON(ctx, data, opts...)
}

func (c *CollectionsClient) CreateWithJSON(
	ctx context.Context,
	jsonData []byte,
	opts ...RequestOption,
) (CreateCollectionResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return CreateCollectionResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return CreateCollectionResponse{}, err
	}

	var response CreateCollectionResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *CollectionsClient) Get(
	ctx context.Context,
	uid string,
	opts ...RequestOption,
) (GetCollectionResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+uid, nil)
	if err != nil {
		return GetCollectionResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetCollectionResponse{}, err
	}

	var response GetCollectionResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *CollectionsClient) GetAll(
	ctx context.Context,
	opts ...RequestOption,
) (GetAllCollectionsResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return GetAllCollectionsResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetAllCollectionsResponse{}, err
	}

	var response GetAllCollectionsResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *CollectionsClient) Update(
	ctx context.Context,
	uid string,
	req UpdateCollectionRequest,
	opts ...RequestOption,
) (UpdateCollectionResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return UpdateCollectionResponse{}, err
	}

	return c.UpdateWithJSON(ctx, uid, data, opts...)
}

func (c *CollectionsClient) UpdateWithJSON(
	ctx context.Context,
	uid string,
	jsonData []byte,
	opts ...RequestOption,
) (UpdateCollectionResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/"+uid, bytes.NewBuffer(jsonData))
	if err != nil {
		return UpdateCollectionResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return UpdateCollectionResponse{}, err
	}

	var response UpdateCollectionResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *CollectionsClient) Delete(
	ctx context.Context,
	uid string,
	opts ...RequestOption,
) (DeleteCollectionResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+"/"+uid, nil)
	if err != nil {
		return DeleteCollectionResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return DeleteCollectionResponse{}, nil
	}

	var response DeleteCollectionResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *CollectionsClient) Fork(
	ctx context.Context,
	uid string,
	req CreateForkRequest,
	opts ...RequestOption,
) (CreateForkResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return CreateForkResponse{}, err
	}

	return c.ForkWithJSON(ctx, uid, data, opts...)
}
func (c *CollectionsClient) ForkWithJSON(
	ctx context.Context,
	uid string,
	jsonData []byte,
	opts ...RequestOption,
) (CreateForkResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/fork/"+uid, bytes.NewBuffer(jsonData))
	if err != nil {
		return CreateForkResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return CreateForkResponse{}, err
	}

	var response CreateForkResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *CollectionsClient) MergeFork(
	ctx context.Context,
	req MergeForkRequest,
	opts ...RequestOption,
) (MergeForkResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return MergeForkResponse{}, err
	}

	return c.MergeForkWithJSON(ctx, data, opts...)
}

func (c *CollectionsClient) MergeForkWithJSON(
	ctx context.Context,
	jsonData []byte,
	opts ...RequestOption,
) (MergeForkResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/merge", bytes.NewBuffer(jsonData))
	if err != nil {
		return MergeForkResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return MergeForkResponse{}, err
	}

	var response MergeForkResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type CreateCollectionRequest struct {
	Collection struct {
		Info struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Schema      string `json:"schema"`
		} `json:"info"`
		Item []struct {
			Name  string `json:"name"`
			Event []struct {
				Listen string `json:"listen"`
				Script struct {
					ID   string   `json:"id"`
					Exec []string `json:"exec"`
					Type string   `json:"type"`
				} `json:"script"`
			} `json:"event"`
			Request struct {
				URL    string `json:"url"`
				Method string `json:"method"`
				Header []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"header"`
			} `json:"request"`
		} `json:"item"`
	} `json:"collection"`
}

type CreateCollectionResponse struct {
	Collection struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		UID  string `json:"uid"`
	} `json:"collection"`
}

type GetCollectionResponse struct {
	Collection struct {
		Info struct {
			Name        string    `json:"name"`
			Description string    `json:"description"`
			PostmanId   string    `json:"_postman_id"`
			Schema      string    `json:"schema"`
			UpdatedAt   time.Time `json:"updatedAt"`
			Fork        struct {
				Label     string    `json:"label"`
				CreatedAt time.Time `json:"createdAt"`
				From      string    `json:"from"`
			} `json:"fork"`
		} `json:"info"`
		Item []struct {
			Name  string `json:"name"`
			Id    string `json:"id"`
			Event []struct {
				Listen string `json:"listen"`
				Script struct {
					Id   string   `json:"id"`
					Exec []string `json:"exec"`
					Type string   `json:"type"`
				} `json:"script"`
			} `json:"event"`
			Request struct {
				Url    string `json:"url"`
				Method string `json:"method"`
				Header []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"header"`
			} `json:"request"`
			Response []interface{} `json:"response"`
		} `json:"item"`
	} `json:"collection"`
}

type GetAllCollectionsResponse struct {
	Collections []struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		Owner     string    `json:"owner"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Uid       string    `json:"uid"`
		IsPublic  bool      `json:"isPublic"`
		Fork      struct {
			Label     string    `json:"label"`
			CreatedAt time.Time `json:"createdAt"`
			From      string    `json:"from"`
		} `json:"fork,omitempty"`
	} `json:"collections"`
}

type UpdateCollectionRequest struct {
	Collection struct {
		Info struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Schema      string `json:"schema"`
		} `json:"info"`
		Item []struct {
			Name  string `json:"name"`
			Event []struct {
				Listen string `json:"listen"`
				Script struct {
					Id   string   `json:"id"`
					Exec []string `json:"exec"`
					Type string   `json:"type"`
				} `json:"script"`
			} `json:"event"`
			Request struct {
				Url    string `json:"url"`
				Method string `json:"method"`
				Header []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"header"`
			} `json:"request"`
		} `json:"item"`
	} `json:"collection"`
}

type UpdateCollectionResponse struct {
	Collection struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"collection"`
}

type DeleteCollectionResponse struct {
	Collection struct {
		Id  string `json:"id"`
		Uid string `json:"uid"`
	} `json:"collection"`
}

type CreateForkRequest struct {
	Label string `json:"label"`
}

type CreateForkResponse struct {
	Collection struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Fork struct {
			Label     string    `json:"label"`
			CreatedAt time.Time `json:"createdAt"`
			From      string    `json:"from"`
		} `json:"fork"`
		UID string `json:"uid"`
	} `json:"collection"`
}

type MergeForkRequest struct {
	Strategy    string `json:"strategy"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type MergeForkResponse struct {
	Collection struct {
		Id  string `json:"id"`
		Uid string `json:"uid"`
	} `json:"collection"`
}
