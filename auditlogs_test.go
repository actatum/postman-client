package postman

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

var getAuditLogsSuccessResponse = GetAuditLogsResponse{
	Trails: []struct {
		Id        int       `json:"id"`
		Ip        string    `json:"ip"`
		UserAgent string    `json:"userAgent"`
		Action    string    `json:"action"`
		Timestamp time.Time `json:"timestamp"`
		Message   string    `json:"message"`
		Data      struct {
			Actor struct {
				Name     string `json:"name"`
				Username string `json:"username"`
				Email    string `json:"email"`
				Id       int    `json:"id"`
				Active   bool   `json:"active"`
			} `json:"actor"`
			User struct {
				Name     string `json:"name"`
				Username string `json:"username"`
				Email    string `json:"email"`
				Id       int    `json:"id"`
			} `json:"user"`
			Team struct {
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"team"`
		} `json:"data"`
	}{
		{
			Id:        12345678,
			Ip:        "192.0.2.0",
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
			Action:    "user.login_google_success",
			Timestamp: mustParseTime("2022-08-31T15:19:32.000Z"),
			Message:   "Taylor Lee successfully logged in using the Google OAuth.",
			Data: struct {
				Actor struct {
					Name     string `json:"name"`
					Username string `json:"username"`
					Email    string `json:"email"`
					Id       int    `json:"id"`
					Active   bool   `json:"active"`
				} `json:"actor"`
				User struct {
					Name     string `json:"name"`
					Username string `json:"username"`
					Email    string `json:"email"`
					Id       int    `json:"id"`
				} `json:"user"`
				Team struct {
					Name string `json:"name"`
					Id   int    `json:"id"`
				} `json:"team"`
			}{
				Actor: struct {
					Name     string `json:"name"`
					Username string `json:"username"`
					Email    string `json:"email"`
					Id       int    `json:"id"`
					Active   bool   `json:"active"`
				}{
					Name:     "Taylor Lee",
					Username: "taylor-lee",
					Email:    "taylor.lee@example.com",
					Id:       12345678,
					Active:   true,
				},
				User: struct {
					Name     string `json:"name"`
					Username string `json:"username"`
					Email    string `json:"email"`
					Id       int    `json:"id"`
				}{
					Name:     "Taylor Lee",
					Username: "taylor-lee",
					Email:    "taylor.lee@example.com",
					Id:       12345678,
				},
				Team: struct {
					Name string `json:"name"`
					Id   int    `json:"id"`
				}{
					Name: "Test-Team",
					Id:   1234,
				},
			},
		},
		{
			Id:        87654321,
			Ip:        "192.0.2.1",
			UserAgent: "PostmanRuntime/7.29.0 Postman/5.5.4 ChromeApp",
			Action:    "user.login_password_success",
			Timestamp: mustParseTime("2022-09-01T06:30:21.000Z"),
			Message:   "Alex Cruz successfully logged in using the Postman password.",
			Data: struct {
				Actor struct {
					Name     string `json:"name"`
					Username string `json:"username"`
					Email    string `json:"email"`
					Id       int    `json:"id"`
					Active   bool   `json:"active"`
				} `json:"actor"`
				User struct {
					Name     string `json:"name"`
					Username string `json:"username"`
					Email    string `json:"email"`
					Id       int    `json:"id"`
				} `json:"user"`
				Team struct {
					Name string `json:"name"`
					Id   int    `json:"id"`
				} `json:"team"`
			}{
				Actor: struct {
					Name     string `json:"name"`
					Username string `json:"username"`
					Email    string `json:"email"`
					Id       int    `json:"id"`
					Active   bool   `json:"active"`
				}{
					Name:     "Alex Cruz",
					Username: "alex-cruz",
					Email:    "alex.cruz@example.com",
					Id:       87654321,
					Active:   true,
				},
				User: struct {
					Name     string `json:"name"`
					Username string `json:"username"`
					Email    string `json:"email"`
					Id       int    `json:"id"`
				}{
					Name:     "Alex Cruz",
					Username: "alex-cruz",
					Email:    "alex.cruz@example.com",
					Id:       87654321,
				},
				Team: struct {
					Name string `json:"name"`
					Id   int    `json:"id"`
				}{
					Name: "Test-Team",
					Id:   1234,
				},
			},
		},
	},
}

func TestAuditLogsClient_Get(t *testing.T) {
	type fields struct {
		apiKey string
	}
	type args struct {
		ctx  context.Context
		req  GetAuditLogsRequest
		opts []RequestOption
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          GetAuditLogsResponse
		wantErr       bool
		postmanErr    *Error
		postmanStatus int
	}{
		{
			name: "success",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx: context.Background(),
				req: GetAuditLogsRequest{
					Limit: intPtr(2),
					Since: stringPtr("2022-08-30"),
				},
				opts: nil,
			},
			want:    getAuditLogsSuccessResponse,
			wantErr: false,
		},
		{
			name: "nil context",
			fields: fields{
				apiKey: "",
			},
			args: args{
				ctx:  nil,
				req:  GetAuditLogsRequest{},
				opts: nil,
			},
			want:    GetAuditLogsResponse{},
			wantErr: true,
		},
		{
			name: "unauthorized",
			fields: fields{
				apiKey: "bob",
			},
			args: args{
				ctx:  context.Background(),
				req:  GetAuditLogsRequest{},
				opts: nil,
			},
			want:    GetAuditLogsResponse{},
			wantErr: true,
			postmanErr: &Error{
				Name:    "AuthenticationError",
				Message: "Invalid API Key. Every request requires a valid API Key to be sent.",
			},
			postmanStatus: http.StatusUnauthorized,
		},
		{
			name: "forbidden",
			fields: fields{
				apiKey: "bob",
			},
			args: args{
				ctx:  context.Background(),
				req:  GetAuditLogsRequest{},
				opts: nil,
			},
			want:    GetAuditLogsResponse{},
			wantErr: true,
			postmanErr: &Error{
				Name:    "ForbiddenRequest",
				Message: "You do not have permissions to view team Audit logs",
			},
			postmanStatus: http.StatusForbidden,
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

				res := getAuditLogsSuccessResponse
				data, err := json.Marshal(res)
				if err != nil {
					t.Error(err)
				}

				w.WriteHeader(http.StatusOK)
				w.Write(data)
			}))
			t.Cleanup(srv.Close)

			c := &AuditLogsClient{
				httpClient: srv.Client(),
				apiKey:     tt.fields.apiKey,
				baseURL:    srv.URL,
			}
			got, err := c.Get(tt.args.ctx, tt.args.req, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func mustParseTime(timestamp string) time.Time {
	res, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		panic(err)
	}

	return res
}

func intPtr(n int) *int {
	return &n
}

func stringPtr(s string) *string {
	return &s
}
