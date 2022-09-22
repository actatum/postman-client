package postman

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var createWebhookSuccessResponse = CreateWebhookResponse{
	Webhook: struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Collection string `json:"collection"`
		WebhookUrl string `json:"webhookUrl"`
		Uid        string `json:"uid"`
	}{
		Id:         "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
		Name:       "Test Webhook",
		Collection: "12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2",
		WebhookUrl: "https://newman-api.getpostman.com/run/12345678/267a6e99-b6da-407c-a96f-03be2d6282fb",
		Uid:        "12345678-1f0df51a-8658-4ee8-a2a1-d2567dfa09a9",
	},
}

func TestWebhooksClient_Create(t *testing.T) {
	type fields struct {
		apiKey string
	}
	type args struct {
		ctx  context.Context
		req  CreateWebhookRequest
		opts []RequestOption
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          CreateWebhookResponse
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
				ctx: context.Background(),
				req: CreateWebhookRequest{Webhook: struct {
					Name       string `json:"name"`
					Collection string `json:"collection"`
				}{
					Name:       createWebhookSuccessResponse.Webhook.Name,
					Collection: createWebhookSuccessResponse.Webhook.Collection,
				}},
				opts: nil,
			},
			want:    createWebhookSuccessResponse,
			wantErr: false,
		},
		{
			name: "success w/workspace",
			fields: fields{
				apiKey: "abc",
			},
			args: args{
				ctx: context.Background(),
				req: CreateWebhookRequest{Webhook: struct {
					Name       string `json:"name"`
					Collection string `json:"collection"`
				}{
					Name:       createWebhookSuccessResponse.Webhook.Name,
					Collection: createWebhookSuccessResponse.Webhook.Collection,
				}},
				opts: []RequestOption{WithWorkspace("123")},
			},
			want:    createWebhookSuccessResponse,
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
			want:    CreateWebhookResponse{},
			wantErr: true,
		},
		{
			name: "unauthorized",
			fields: fields{
				apiKey: "bob",
			},
			args: args{
				ctx:  context.Background(),
				req:  CreateWebhookRequest{},
				opts: nil,
			},
			want:    CreateWebhookResponse{},
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
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf("expected Content-Type: application/json, got: %s", r.Header.Get("Content-Type"))
				}
				if r.Method != http.MethodPost {
					t.Errorf("expected http method POST, got: %s", r.Method)
				}
				for _, o := range tt.args.opts {
					switch v := o.(type) {
					case workspaceOption:
						if r.URL.Query().Get("workspace") != string(v) {
							t.Errorf("expected workspace query param = %s, got: %s", string(v), r.URL.Query().Get("workspace"))
						}
					}
				}

				var req CreateWebhookRequest
				if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
					t.Error(err)
				}

				if req.Webhook.Name != tt.args.req.Webhook.Name {
					t.Errorf("expected Webhook.Name = %v, got: %v", tt.args.req.Webhook.Name, req.Webhook.Name)
				}
				if req.Webhook.Collection != tt.args.req.Webhook.Collection {
					t.Errorf("expected Webhook.Collection = %v, got: %v", tt.args.req.Webhook.Collection, req.Webhook.Collection)
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

				res := createWebhookSuccessResponse
				data, err := json.Marshal(res)
				if err != nil {
					t.Error(err)
				}

				w.WriteHeader(http.StatusOK)
				w.Write(data)
			}))
			t.Cleanup(srv.Close)

			c := &WebhooksClient{
				httpClient: srv.Client(),
				apiKey:     tt.fields.apiKey,
				baseURL:    srv.URL,
			}
			got, err := c.Create(tt.args.ctx, tt.args.req, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
