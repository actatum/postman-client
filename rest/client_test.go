// Package rest provides types/client for making requests to the postman REST api.
package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/actatum/postman-client/testdata"
)

func TestNewClient(t *testing.T) {
	type args struct {
		apiKey string
		opts   []Option
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "new client base",
			args: args{
				apiKey: "hi",
				opts:   nil,
			},
			want: &Client{
				httpClient: &http.Client{},
				apiKey:     "hi",
				baseURL:    baseURL,
			},
		},
		{
			name: "http client option",
			args: args{
				apiKey: "",
				opts: []Option{
					WithHTTPClient(&http.Client{
						Transport:     nil,
						CheckRedirect: nil,
						Jar:           nil,
						Timeout:       10 * time.Second,
					}),
				},
			},
			want: &Client{
				httpClient: &http.Client{
					Transport:     nil,
					CheckRedirect: nil,
					Jar:           nil,
					Timeout:       10 * time.Second,
				},
				apiKey:    "",
				baseURL:   baseURL,
				logOutput: nil,
			},
		},
		{
			name: "debug log option",
			args: args{
				apiKey: "",
				opts: []Option{
					WithDebugLog(os.Stdout),
				},
			},
			want: &Client{
				httpClient: &http.Client{},
				apiKey:     "",
				baseURL:    baseURL,
				logOutput:  os.Stdout,
			},
		},
		{
			name: "debug log option w/ nil writer",
			args: args{
				apiKey: "",
				opts: []Option{
					WithDebugLog(nil),
				},
			},
			want: &Client{
				httpClient: &http.Client{},
				apiKey:     "",
				baseURL:    baseURL,
				logOutput:  os.Stdout,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.apiKey, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_BaseURL(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		apiKey     string
		baseURL    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "baseURL",
			fields: fields{
				httpClient: &http.Client{},
				apiKey:     "",
				baseURL:    "howdy-doo-partner",
			},
			want: "howdy-doo-partner",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
				apiKey:     tt.fields.apiKey,
				baseURL:    tt.fields.baseURL,
			}
			if got := c.BaseURL(); got != tt.want {
				t.Errorf("BaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		apiKey     string
		baseURL    string
	}
	type args struct {
		ctx     context.Context
		method  string
		url     string
		payload interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "new request no body",
			fields: fields{
				httpClient: &http.Client{},
				apiKey:     "",
				baseURL:    "",
			},
			args: args{
				ctx:     context.Background(),
				method:  http.MethodGet,
				url:     "https://api.getpostman.com/collections",
				payload: nil,
			},
			want: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "https",
					Host:   "api.getpostman.com",
					Path:   "/collections",
				},
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     make(http.Header),
				Body:       nil,
				GetBody:    nil,
				Host:       "api.getpostman.com",
			},
			wantErr: false,
		},
		{
			name: "new request w/ unmarshallable body",
			fields: fields{
				httpClient: &http.Client{},
				apiKey:     "",
				baseURL:    "",
			},
			args: args{
				ctx:    context.Background(),
				method: http.MethodGet,
				url:    "https://api.getpostman.com/collections",
				payload: map[interface{}]interface{}{
					1: "10",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
				apiKey:     tt.fields.apiKey,
				baseURL:    tt.fields.baseURL,
			}
			got, err := c.NewRequest(tt.args.ctx, tt.args.method, tt.args.url, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			want := tt.want.WithContext(tt.args.ctx)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("NewRequest() got = %v\n, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DoRequest(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		apiKey     string
		logOutput  io.Writer
	}
	type args struct {
		method string
		body   io.Reader
		result interface{}
		opts   []RequestOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "get",
			fields: fields{
				httpClient: &http.Client{},
				apiKey:     "api-key",
				logOutput:  nil,
			},
			args: args{
				method: http.MethodGet,
				body:   nil,
				result: map[string]string{
					"response": "here",
				},
				opts: nil,
			},
			want: map[string]string{
				"response": "here",
			},
			wantErr: false,
		},
		{
			name: "get w/ workspace opt",
			fields: fields{
				httpClient: &http.Client{},
				apiKey:     "api-key",
				logOutput:  nil,
			},
			args: args{
				method: http.MethodGet,
				body:   nil,
				result: map[string]string{
					"response": "here",
				},
				opts: []RequestOption{
					WithWorkspace("123"),
				},
			},
			want: map[string]string{
				"response": "here",
			},
			wantErr: false,
		},
		{
			name: "post",
			fields: fields{
				httpClient: &http.Client{},
				apiKey:     "api-key",
				logOutput:  nil,
			},
			args: args{
				method: http.MethodPost,
				body:   bytes.NewBuffer([]byte(`{"message":"hello"}`)),
				result: map[string]string{
					"response": "hello",
				},
				opts: nil,
			},
			want: map[string]string{
				"response": "hello",
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				httpClient: &http.Client{},
				apiKey:     "api-key",
				logOutput:  nil,
			},
			args: args{
				method: http.MethodPost,
				body:   bytes.NewBuffer([]byte(`{"message":"hello"}`)),
				result: ErrorResponse{
					Error: &Error{
						Name:    "invalid request",
						Message: "missing field name",
					},
				},
				opts: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
				apiKey:     tt.fields.apiKey,
				logOutput:  tt.fields.logOutput,
			}
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("x-api-key") != tt.fields.apiKey {
					t.Errorf(
						"r.Header.Get(x-api-key) got = %v, want %v",
						r.Header.Get("X-API-Key"),
						tt.fields.apiKey,
					)
				}
				if r.Header.Get("Accept") != "application/json" {
					t.Errorf(
						"r.Header.Get(Accept) got = %v, want %v",
						r.Header.Get("Accept"),
						"application/json",
					)
				}
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf(
						"r.Header.Get(Content-Type) got = %v, want %v",
						r.Header.Get("Content-Type"),
						"application/json",
					)
				}

				for _, opt := range tt.args.opts {
					v, ok := opt.(workspaceOption)
					if ok {
						if r.URL.Query().Get("workspace") != string(v) {
							t.Errorf(
								"r.URL.Query.Get(workspace) got = %v, want %v",
								r.URL.Query().Get("workspace"),
								string(v),
							)
						}
					}
				}

				if tt.wantErr {
					responseJSON, err := json.Marshal(tt.args.result)
					if err != nil {
						t.Fatal(err)
					}
					w.WriteHeader(http.StatusBadRequest)
					w.Write(responseJSON)
				}

				responseJSON, err := json.Marshal(tt.args.result)
				if err != nil {
					t.Fatal(err)
				}

				w.Write(responseJSON)
			})
			srv := httptest.NewServer(h)
			t.Cleanup(srv.Close)

			c.httpClient = srv.Client()

			r, err := http.NewRequest(tt.args.method, srv.URL, tt.args.body)
			if err != nil {
				t.Fatal(err)
			}

			if err = c.DoRequest(r, &tt.args.result, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("DoRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			got, err := json.Marshal(tt.args.result)
			if err != nil {
				t.Fatal(err)
			}
			want, err := json.Marshal(tt.want)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(got, want) {
				t.Errorf("DoRequest result = %v, want %v", string(got), string(want))
			}
		})
	}
}

func TestClient_DoRequestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	rc := NewClient(testdata.TestAPIKey)

	t.Run("unauthorized", func(t *testing.T) {
		if testing.Short() {
			t.Skip()
		}

		rc = NewClient("")
		r, err := rc.NewRequest(
			context.Background(),
			http.MethodGet,
			fmt.Sprintf("%s%s", rc.BaseURL(), "/collections"),
			nil,
		)
		if err != nil {
			t.Fatal(err)
		}

		var result map[string]string
		err = rc.DoRequest(r, &result)
		if err == nil {
			t.Fatal("expected error got nil")
		}

		if !testdata.IsUnauthorizedError(err) {
			t.Fatalf("rc.DoRequest() error got = %v, want unauthorized error", err)
		}
	})
}
