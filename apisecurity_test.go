package postman

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var validateSchemaSuccessResponse = APISecuritySchemaValidationResponse{
	Warnings: []struct {
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
	}{
		{
			Severity: "MEDIUM",
			Message:  "HTTP authentication scheme is using an unknown scheme.",
			Location: struct {
				Start struct {
					Line   int `json:"line"`
					Column int `json:"column"`
				} `json:"start"`
				End struct {
					Line   int `json:"line"`
					Column int `json:"column"`
				} `json:"end"`
			}{
				Start: struct {
					Line   int `json:"line"`
					Column int `json:"column"`
				}{
					Line:   1,
					Column: 1116,
				},
				End: struct {
					Line   int `json:"line"`
					Column int `json:"column"`
				}{
					Line:   1,
					Column: 1118,
				},
			},
			DataPath: []string{"components", "securitySchemes", "BasicAuth", "scheme"},
			PossibleFixUrl: `https://go.pstmn.io/
openapi3-security-warnings#http-authentication-scheme-is-using-an-unknown-scheme`,
			Category: struct {
				Name string `json:"name"`
			}{
				Name: "Broken User Authentication",
			},
		},
	},
}

func TestAPISecurityClient_ValidateSchema(t *testing.T) {
	type fields struct {
		apiKey string
	}
	type args struct {
		ctx  context.Context
		req  APISecuritySchemaValidationRequest
		opts []RequestOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    APISecuritySchemaValidationResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx: context.Background(),
				req: APISecuritySchemaValidationRequest{
					Schema: struct {
						Type     string `json:"type"`
						Language string `json:"language"`
						Schema   string `json:"schema"`
					}{
						Type:     "openapi3",
						Language: "json",
						Schema: `{\"openapi\":\"3.0.0\",
\"info\":{\"version\":\"1\",\"title\":\"temp\",\"license\":{\"name\":\"MIT\"}},
\"servers\":[{\"url\":\"https://petstore.swagger.io/v1\"}],\"paths\":{\"/user\"
:{\"get\":{\"summary\":\"Details about a user\",\"operationId\":\"listUser\",\"
tags\":[\"user\"],\"parameters\":[{\"name\":\"id\",\"in\":\"query\",\"descripti
on\":\"ID of the user\",\"required\":true,\"schema\":{\"type\":\"integer\",\"fo
rmat\":\"int32\"}}],\"responses\":{\"200\":{\"description\":\"Details about a u
ser\",\"headers\":{\"x-next\":{\"description\":\"A link to the next page of res
ponses\",\"schema\":{\"type\":\"string\"}}},\"content\":{\"application/json\":{
\"schema\":{\"$ref\":\"#/components/schemas/User\"}}}},\"default\":{\"descripti
on\":\"unexpected error\",\"content\":{\"application/json\":{\"schema\":{\"$ref
\":\"#/components/schemas/Error\"}}}}}}}},\"components\":{\"schemas\":{\"User\"
:{\"type\":\"object\",\"required\":[\"id\",\"name\"],\"properties\":{\"id\":{\"
type\":\"integer\",\"format\":\"int64\"},\"name\":{\"type\":\"string\"},\"tag\"
:{\"type\":\"string\"}}},\"Error\":{\"type\":\"object\",\"required\":[\"code\",
\"message\"],\"properties\":{\"code\":{\"type\":\"integer\",\"format\":\"int32\
"},\"message\":{\"type\":\"string\"}}}},\"securitySchemes\":{\"BasicAuth\":{\"t
ype\":\"http\",\"scheme\":\"\"}}},\"security\":[{\"BasicAuth\":[]}]}`,
					},
				},
				opts: nil,
			},
			want:    validateSchemaSuccessResponse,
			wantErr: false,
		},
		{
			name: "nil context",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx:  nil,
				req:  APISecuritySchemaValidationRequest{},
				opts: nil,
			},
			want:    APISecuritySchemaValidationResponse{},
			wantErr: true,
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

				var req APISecuritySchemaValidationRequest
				if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
					t.Error(err)
				}

				if req.Schema.Type != tt.args.req.Schema.Type {
					t.Errorf("expected Schema.Type = %v, got: %v", tt.args.req.Schema.Type, req.Schema.Type)
				}
				if req.Schema.Language != tt.args.req.Schema.Language {
					t.Errorf("expected Schema.Language = %v, got: %v", tt.args.req.Schema.Language, req.Schema.Language)
				}
				if req.Schema.Schema != tt.args.req.Schema.Schema {
					t.Errorf("expected Schema.Schema = %v, got: %v", tt.args.req.Schema.Schema, req.Schema.Schema)
				}

				res := validateSchemaSuccessResponse
				data, err := json.Marshal(res)
				if err != nil {
					t.Error(err)
				}

				w.WriteHeader(http.StatusOK)
				w.Write(data)
			}))
			t.Cleanup(srv.Close)

			c := &APISecurityClient{
				httpClient: srv.Client(),
				apiKey:     tt.fields.apiKey,
				baseURL:    srv.URL,
			}
			got, err := c.ValidateSchema(tt.args.ctx, tt.args.req, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateSchema() got = %v, want %v", got, tt.want)
			}
		})
	}
}
