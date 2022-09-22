package postman

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var getAuthenticatedUserSuccessResponse = GetAuthenticatedUserResponse{
	User: struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		FullName string `json:"fullName"`
		Avatar   string `json:"avatar"`
		IsPublic bool   `json:"isPublic"`
	}{
		Id:       12345678,
		Username: "taylor-lee",
		Email:    "taylor.lee@example.com",
		FullName: "Taylor Lee",
		Avatar:   "https://example.com/user/r5u9qpvmujfjf6lbqmga.jpg",
		IsPublic: true,
	},
	Operations: []struct {
		Name    string `json:"name"`
		Limit   int    `json:"limit"`
		Usage   int    `json:"usage"`
		Overage int    `json:"overage"`
	}{
		{
			Name:    "mock_usage",
			Limit:   1000000,
			Usage:   110276,
			Overage: 0,
		},
		{
			Name:    "monitor_request_runs",
			Limit:   10000000,
			Usage:   1141750,
			Overage: 0,
		},
		{
			Name:    "api_usage",
			Limit:   1000000,
			Usage:   16240,
			Overage: 0,
		},
		{
			Name:    "custom_domains",
			Limit:   25,
			Usage:   25,
			Overage: 0,
		},
		{
			Name:    "serverless_requests",
			Limit:   10000,
			Usage:   0,
			Overage: 0,
		},
		{
			Name:    "integrations",
			Limit:   5000,
			Usage:   1018,
			Overage: 0,
		},
		{
			Name:    "cloud_agent_requests",
			Limit:   1000000,
			Usage:   1615,
			Overage: 0,
		},
	},
}

func TestUsersClient_GetAuthenticatedUser(t *testing.T) {
	type fields struct {
		apiKey string
	}
	type args struct {
		ctx  context.Context
		opts []RequestOption
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          GetAuthenticatedUserResponse
		wantErr       bool
		postmanErr    *Error
		postmanStatus int
	}{
		{
			name: "success",
			fields: fields{
				apiKey: "abc",
			},
			args: args{
				ctx:  context.Background(),
				opts: nil,
			},
			want:    getAuthenticatedUserSuccessResponse,
			wantErr: false,
		},
		{
			name: "nil context",
			fields: fields{
				apiKey: "",
			},
			args: args{
				ctx:  nil,
				opts: nil,
			},
			want:    GetAuthenticatedUserResponse{},
			wantErr: true,
		},
		{
			name: "unauthorized",
			fields: fields{
				apiKey: "bob",
			},
			args: args{
				ctx:  context.Background(),
				opts: nil,
			},
			want:    GetAuthenticatedUserResponse{},
			wantErr: true,
			postmanErr: &Error{
				Name:    "AuthenticationError",
				Message: "Invalid API Key. Every request requires a valid API Key to be sent.",
			},
			postmanStatus: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("x-api-key") != tt.fields.apiKey {
					t.Errorf("expected X-API-Key: %s, got: %s", tt.fields.apiKey, r.Header.Get("x-api-key"))
				}
				if r.Method != http.MethodGet {
					t.Errorf("expected http method GET, got: %s", r.Method)
				}

				if tt.postmanErr != nil {
					data, err := json.Marshal(errorWrapper{Error: *tt.postmanErr})
					if err != nil {
						t.Error(err)
						return
					}

					w.WriteHeader(tt.postmanStatus)
					w.Write(data)
					return
				}

				res := getAuthenticatedUserSuccessResponse
				data, err := json.Marshal(res)
				if err != nil {
					t.Error(err)
				}

				w.WriteHeader(http.StatusOK)
				w.Write(data)
			}))
			t.Cleanup(srv.Close)

			c := &UsersClient{
				httpClient: &http.Client{},
				apiKey:     tt.fields.apiKey,
				baseURL:    srv.URL,
			}
			got, err := c.GetAuthenticatedUser(tt.args.ctx, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthenticatedUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthenticatedUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
