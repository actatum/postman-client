package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// MonitorsClient implements the functions that manage monitor resources.
type MonitorsClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewMonitorsClient returns a new instance of MonitorsClient.
func NewMonitorsClient(apiKey string, httpClient *http.Client) *MonitorsClient {
	return &MonitorsClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    postmanBaseURL + "/monitors",
	}
}

func (c *MonitorsClient) Create(ctx context.Context, req CreateMonitorRequest, opts ...RequestOption) (CreateMonitorResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return CreateMonitorResponse{}, err
	}

	return c.CreateWithJSON(ctx, data, opts...)
}

func (c *MonitorsClient) CreateWithJSON(ctx context.Context, jsonData []byte, opts ...RequestOption) (CreateMonitorResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return CreateMonitorResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return CreateMonitorResponse{}, err
	}

	var response CreateMonitorResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *MonitorsClient) Get(ctx context.Context, uid string, opts ...RequestOption) (GetMonitorResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+uid, nil)
	if err != nil {
		return GetMonitorResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetMonitorResponse{}, err
	}

	var response GetMonitorResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *MonitorsClient) GetAll(ctx context.Context, opts ...RequestOption) (GetAllMonitorsResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return GetAllMonitorsResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return GetAllMonitorsResponse{}, err
	}

	var response GetAllMonitorsResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *MonitorsClient) Update(ctx context.Context, uid string, req UpdateMonitorRequest, opts ...RequestOption) (UpdateMonitorResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return UpdateMonitorResponse{}, err
	}

	return c.UpdateWithJSON(ctx, uid, data, opts...)
}

func (c *MonitorsClient) UpdateWithJSON(ctx context.Context, uid string, jsonData []byte, opts ...RequestOption) (UpdateMonitorResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPut, c.baseURL+"/"+uid, bytes.NewBuffer(jsonData))
	if err != nil {
		return UpdateMonitorResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return UpdateMonitorResponse{}, err
	}

	var response UpdateMonitorResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *MonitorsClient) Delete(ctx context.Context, uid string, opts ...RequestOption) (DeleteMonitorResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+"/"+uid, nil)
	if err != nil {
		return DeleteMonitorResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return DeleteMonitorResponse{}, err
	}

	var response DeleteMonitorResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

func (c *MonitorsClient) Run(ctx context.Context, uid string, opts ...RequestOption) (RunMonitorResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/"+uid+"/run", nil)
	if err != nil {
		return RunMonitorResponse{}, err
	}

	respJSON, err := doRequest(r, c.apiKey, c.httpClient, opts...)
	if err != nil {
		return RunMonitorResponse{}, err
	}

	var response RunMonitorResponse
	err = json.Unmarshal(respJSON, &response)

	return response, err
}

type CreateMonitorRequest struct {
	Monitor struct {
		Name     string `json:"name"`
		Schedule struct {
			Cron     string `json:"cron"`
			Timezone string `json:"timezone"`
		} `json:"schedule"`
		Collection  string `json:"collection"`
		Environment string `json:"environment"`
	} `json:"monitor"`
}

type CreateMonitorResponse struct {
	Monitor struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"monitor"`
}

type GetMonitorResponse struct {
	Monitor struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		Uid            string `json:"uid"`
		Owner          int    `json:"owner"`
		CollectionUid  string `json:"collectionUid"`
		EnvironmentUid string `json:"environmentUid"`
		Options        struct {
			StrictSSL       bool `json:"strictSSL"`
			FollowRedirects bool `json:"followRedirects"`
			RequestTimeout  int  `json:"requestTimeout"`
			RequestDelay    int  `json:"requestDelay"`
		} `json:"options"`
		Notifications struct {
			OnError []struct {
				Email string `json:"email"`
			} `json:"onError"`
			OnFailure []struct {
				Email string `json:"email"`
			} `json:"onFailure"`
		} `json:"notifications"`
		Distribution []interface{} `json:"distribution"`
		Schedule     struct {
			Cron     string    `json:"cron"`
			Timezone string    `json:"timezone"`
			NextRun  time.Time `json:"nextRun"`
		} `json:"schedule"`
		LastRun struct {
			Status     string    `json:"status"`
			StartedAt  time.Time `json:"startedAt"`
			FinishedAt time.Time `json:"finishedAt"`
		} `json:"lastRun"`
		Stats struct {
			Assertions struct {
				Total  int `json:"total"`
				Failed int `json:"failed"`
			} `json:"assertions"`
			Requests struct {
				Total int `json:"total"`
			} `json:"requests"`
		} `json:"stats"`
	} `json:"monitor"`
}

type GetAllMonitorsResponse struct {
	Monitors []struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Uid   string `json:"uid"`
		Owner int    `json:"owner"`
	} `json:"monitors"`
}

type UpdateMonitorRequest struct {
	Monitor struct {
		Name     string `json:"name"`
		Schedule struct {
			Cron     string `json:"cron"`
			Timezone string `json:"timezone"`
		} `json:"schedule"`
	} `json:"monitor"`
}

type UpdateMonitorResponse struct {
	Monitor struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"monitor"`
}

type DeleteMonitorResponse struct {
	Monitor struct {
		Id  string `json:"id"`
		Uid string `json:"uid"`
	} `json:"monitor"`
}

type RunMonitorResponse struct {
	Run struct {
		Info struct {
			JobId          string    `json:"jobId"`
			MonitorId      string    `json:"monitorId"`
			Name           string    `json:"name"`
			CollectionUid  string    `json:"collectionUid"`
			EnvironmentUid string    `json:"environmentUid"`
			Status         string    `json:"status"`
			StartedAt      time.Time `json:"startedAt"`
			FinishedAt     time.Time `json:"finishedAt"`
		} `json:"info"`
		Stats struct {
			Assertions struct {
				Total  int `json:"total"`
				Failed int `json:"failed"`
			} `json:"assertions"`
			Requests struct {
				Total  int `json:"total"`
				Failed int `json:"failed"`
			} `json:"requests"`
		} `json:"stats"`
		Executions []struct {
			Id   int `json:"id"`
			Item struct {
				Name string `json:"name"`
			} `json:"item"`
			Request struct {
				Method  string `json:"method"`
				Url     string `json:"url"`
				Headers struct {
					UserAgent      string      `json:"User-Agent"`
					Accept         string      `json:"Accept"`
					CacheControl   string      `json:"Cache-Control"`
					PostmanToken   interface{} `json:"Postman-Token"`
					Host           string      `json:"Host"`
					AcceptEncoding string      `json:"Accept-Encoding"`
					Connection     string      `json:"Connection"`
				} `json:"headers"`
				Body struct {
					ContentLength int `json:"contentLength"`
				} `json:"body"`
				Timestamp time.Time `json:"timestamp"`
			} `json:"request"`
			Response struct {
				Code int `json:"code"`
				Body struct {
					ContentLength int `json:"contentLength"`
				} `json:"body"`
				ResponseTime int `json:"responseTime"`
				ResponseSize int `json:"responseSize"`
				Headers      struct {
					Server                    interface{} `json:"Server"`
					Date                      string      `json:"Date"`
					ContentType               string      `json:"Content-Type"`
					TransferEncoding          string      `json:"Transfer-Encoding"`
					Connection                string      `json:"Connection"`
					KeepAlive                 interface{} `json:"Keep-Alive"`
					AccessControlAllowOrigin  string      `json:"Access-Control-Allow-Origin"`
					AccessControlAllowMethods string      `json:"Access-Control-Allow-Methods"`
				} `json:"headers"`
			} `json:"response"`
			Errors []struct {
				Name    string `json:"name"`
				Message string `json:"message"`
			} `json:"errors"`
		} `json:"executions"`
		Failures []struct {
			ExecutionId int    `json:"executionId"`
			Name        string `json:"name"`
			Message     string `json:"message"`
		} `json:"failures"`
	} `json:"run"`
}
