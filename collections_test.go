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

var createCollectionSuccessResponse = CreateCollectionResponse{
	Collection: struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		UID  string `json:"uid"`
	}{ID: "12ece9e1-2abf-4edc-8e34-de66e74114d2", Name: "Test Collection", UID: "12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2"},
}

var getCollectionSuccessResponse = GetCollectionResponse{
	Collection: struct {
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
	}{Info: struct {
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
	}(struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		PostmanId   string    `json:"_postman_id"`
		Schema      string    `json:"schema"`
		UpdatedAt   time.Time `json:"udpatedAt"`
		Fork        struct {
			Label     string    `json:"label"`
			CreatedAt time.Time `json:"createdAt"`
			From      string    `json:"from"`
		} `json:"fork"`
	}{
		Name:        "Test Collection",
		Description: "This is a test collection that makes a tiny request to Postman Echo service to get the list of request headers sent by a HTTP client.",
		PostmanId:   "12ece9e1-2abf-4edc-8e34-de66e74114d2",
		Schema:      "https://schema.getpostman.com/json/collection/v2.0.0/collection.json",
		UpdatedAt:   mustParseTime("2022-06-16T20:21:13.000Z"),
		Fork: struct {
			Label     string    `json:"label"`
			CreatedAt time.Time `json:"createdAt"`
			From      string    `json:"from"`
		}{
			Label:     "Test Fork",
			CreatedAt: mustParseTime("2022-06-16T19:51:44.069Z"),
			From:      "12345678-12ece9e1-2abf-4edc-8e34-de66e74114d2",
		},
	}),
		Item: []struct {
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
		}{
			{
				Name: "Test GET Response",
				Id:   "82ee981b-e19f-962a-401e-ea34ebfb4848",
				Event: []struct {
					Listen string `json:"listen"`
					Script struct {
						Id   string   `json:"id"`
						Exec []string `json:"exec"`
						Type string   `json:"type"`
					} `json:"script"`
				}{
					{
						Listen: "test",
						Script: struct {
							Id   string   `json:"id"`
							Exec []string `json:"exec"`
							Type string   `json:"type"`
						}{
							Id: "7d2334fc-a84a-4c3d-b26c-7529afa4c0a",
							Exec: []string{
								"pm.test(\"Status code is 200\", function () {",
								"    pm.response.to.have.status(200);",
								"});",
							},
							Type: "text/javascrip",
						},
					},
				},
				Request: struct {
					Url    string `json:"url"`
					Method string `json:"method"`
					Header []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"header"`
				}{
					Url:    "https://echo.getpostman.com/headers",
					Method: "GET",
					Header: []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					}{
						{
							Key:   "Content-Type",
							Value: "application/json",
						},
					},
				},
				Response: nil,
			},
		},
	},
}

func TestCollectionsClient_Create(t *testing.T) {
	type fields struct {
		apiKey string
	}
	type args struct {
		ctx  context.Context
		req  CreateCollectionRequest
		opts []RequestOption
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          CreateCollectionResponse
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
				req: CreateCollectionRequest{
					Collection: struct {
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
					}{Info: struct {
						Name        string `json:"name"`
						Description string `json:"description"`
						Schema      string `json:"schema"`
					}{
						Name:        "Test Collection",
						Description: "This collection makes a request to the Postman Echo service to get a list of request headers sent by an HTTP client.",
						Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
					}, Item: []struct {
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
					}{
						{
							Name: "Test GET Response",
							Event: []struct {
								Listen string `json:"listen"`
								Script struct {
									ID   string   `json:"id"`
									Exec []string `json:"exec"`
									Type string   `json:"type"`
								} `json:"script"`
							}{
								{
									Listen: "test",
									Script: struct {
										ID   string   `json:"id"`
										Exec []string `json:"exec"`
										Type string   `json:"type"`
									}{
										ID: "7d2334fc-a84a-4c3d-b26c-7529afa4c0ae",
										Exec: []string{
											"pm.test(\\\"Status code is 200\\\", function () {",
											"    pm.response.to.have.status(200);",
											"});",
										},
										Type: "text/javascript",
									},
								},
							},
							Request: struct {
								URL    string `json:"url"`
								Method string `json:"method"`
								Header []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								} `json:"header"`
							}{
								URL:    "https://echo.getpostman.com/headers",
								Method: "GET",
								Header: []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								}{
									{
										Key:   "Content-Type",
										Value: "application/json",
									},
								},
							},
						},
					}},
				},
			},
			want:    createCollectionSuccessResponse,
			wantErr: false,
		},
		{
			name: "success w/ workspace",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx: context.Background(),
				req: CreateCollectionRequest{
					Collection: struct {
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
					}{Info: struct {
						Name        string `json:"name"`
						Description string `json:"description"`
						Schema      string `json:"schema"`
					}{
						Name:        "Test Collection",
						Description: "This collection makes a request to the Postman Echo service to get a list of request headers sent by an HTTP client.",
						Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
					}, Item: []struct {
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
					}{
						{
							Name: "Test GET Response",
							Event: []struct {
								Listen string `json:"listen"`
								Script struct {
									ID   string   `json:"id"`
									Exec []string `json:"exec"`
									Type string   `json:"type"`
								} `json:"script"`
							}{
								{
									Listen: "test",
									Script: struct {
										ID   string   `json:"id"`
										Exec []string `json:"exec"`
										Type string   `json:"type"`
									}{
										ID: "7d2334fc-a84a-4c3d-b26c-7529afa4c0ae",
										Exec: []string{
											"pm.test(\\\"Status code is 200\\\", function () {",
											"    pm.response.to.have.status(200);",
											"});",
										},
										Type: "text/javascript",
									},
								},
							},
							Request: struct {
								URL    string `json:"url"`
								Method string `json:"method"`
								Header []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								} `json:"header"`
							}{
								URL:    "https://echo.getpostman.com/headers",
								Method: "GET",
								Header: []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								}{
									{
										Key:   "Content-Type",
										Value: "application/json",
									},
								},
							},
						},
					}},
				},
				opts: []RequestOption{WithWorkspace("abc")},
			},
			want:    createCollectionSuccessResponse,
			wantErr: false,
		},
		{
			name: "Bad Request",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx: context.Background(),
				req: CreateCollectionRequest{
					Collection: struct {
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
					}{Info: struct {
						Name        string `json:"name"`
						Description string `json:"description"`
						Schema      string `json:"schema"`
					}{
						Name:        "Test Collection",
						Description: "This collection makes a request to the Postman Echo service to get a list of request headers sent by an HTTP client.",
						Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
					}, Item: []struct {
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
					}{
						{
							Name: "Test GET Response",
							Event: []struct {
								Listen string `json:"listen"`
								Script struct {
									ID   string   `json:"id"`
									Exec []string `json:"exec"`
									Type string   `json:"type"`
								} `json:"script"`
							}{
								{
									Listen: "test",
									Script: struct {
										ID   string   `json:"id"`
										Exec []string `json:"exec"`
										Type string   `json:"type"`
									}{
										ID: "7d2334fc-a84a-4c3d-b26c-7529afa4c0ae",
										Exec: []string{
											"pm.test(\\\"Status code is 200\\\", function () {",
											"    pm.response.to.have.status(200);",
											"});",
										},
										Type: "text/javascript",
									},
								},
							},
							Request: struct {
								URL    string `json:"url"`
								Method string `json:"method"`
								Header []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								} `json:"header"`
							}{
								URL:    "https://echo.getpostman.com/headers",
								Method: "GET",
								Header: []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								}{
									{
										Key:   "Content-Type",
										Value: "application/json",
									},
								},
							},
						},
					}},
				},
				opts: []RequestOption{WithWorkspace("abc")},
			},
			want:    CreateCollectionResponse{},
			wantErr: true,
			postmanErr: &Error{
				Name:    "instanceFoundError",
				Message: "The specified item already exists.",
				Details: map[string]string{
					"item": "collection",
					"id":   "12ece9e1-2abf-4edc-8e34-de66e74114d2",
				},
			},
			postmanStatus: http.StatusBadRequest,
		},
		{
			name: "nil context",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx:  nil,
				req:  CreateCollectionRequest{},
				opts: nil,
			},
			want:    CreateCollectionResponse{},
			wantErr: true,
		},
		{
			name: "Bad Request",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx: context.Background(),
				req: CreateCollectionRequest{
					Collection: struct {
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
					}{Info: struct {
						Name        string `json:"name"`
						Description string `json:"description"`
						Schema      string `json:"schema"`
					}{}, Item: []struct {
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
					}{
						{
							Name: "Test GET Response",
							Event: []struct {
								Listen string `json:"listen"`
								Script struct {
									ID   string   `json:"id"`
									Exec []string `json:"exec"`
									Type string   `json:"type"`
								} `json:"script"`
							}{
								{
									Listen: "test",
									Script: struct {
										ID   string   `json:"id"`
										Exec []string `json:"exec"`
										Type string   `json:"type"`
									}{
										ID: "7d2334fc-a84a-4c3d-b26c-7529afa4c0ae",
										Exec: []string{
											"pm.test(\\\"Status code is 200\\\", function () {",
											"    pm.response.to.have.status(200);",
											"});",
										},
										Type: "text/javascript",
									},
								},
							},
							Request: struct {
								URL    string `json:"url"`
								Method string `json:"method"`
								Header []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								} `json:"header"`
							}{
								URL:    "https://echo.getpostman.com/headers",
								Method: "GET",
								Header: []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								}{
									{
										Key:   "Content-Type",
										Value: "application/json",
									},
								},
							},
						},
					}},
				},
				opts: []RequestOption{WithWorkspace("abc")},
			},
			want:    CreateCollectionResponse{},
			wantErr: true,
			postmanErr: &Error{
				Name:    "malformedRequestError",
				Message: "Found 1 errors with the supplied collection.",
			},
			postmanStatus: http.StatusBadRequest,
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

				var req CreateCollectionRequest
				if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
					t.Error(err)
				}

				if !reflect.DeepEqual(req, tt.args.req) {
					t.Errorf("expected: %v, got %v", tt.args.req, req)
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

				res := createCollectionSuccessResponse
				data, err := json.Marshal(res)
				if err != nil {
					t.Error(err)
				}

				w.WriteHeader(http.StatusOK)
				w.Write(data)
			}))
			t.Cleanup(srv.Close)

			c := &CollectionsClient{
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

func TestCollectionsClient_Get(t *testing.T) {
	type fields struct {
		apiKey string
	}
	type args struct {
		ctx  context.Context
		uid  string
		opts []RequestOption
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          GetCollectionResponse
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
				ctx:  context.Background(),
				uid:  "12ece9e1-2abf-4edc-8e34-de66e74114d2",
				opts: nil,
			},
			want:    getCollectionSuccessResponse,
			wantErr: false,
		},
		{
			name: "success w/workspace",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx:  context.Background(),
				uid:  "12ece9e1-2abf-4edc-8e34-de66e74114d2",
				opts: []RequestOption{WithWorkspace("abc")},
			},
			want:    getCollectionSuccessResponse,
			wantErr: false,
		},
		{
			name: "nil context",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx:  nil,
				uid:  "12ece9e1-2abf-4edc-8e34-de66e74114d2",
				opts: nil,
			},
			want:    GetCollectionResponse{},
			wantErr: true,
		},
		{
			name: "Not Found",
			fields: fields{
				apiKey: "123",
			},
			args: args{
				ctx:  context.Background(),
				uid:  "12ece9e1-2abf-4edc-8e34-de66e74114d2",
				opts: nil,
			},
			want:    GetCollectionResponse{},
			wantErr: true,
			postmanErr: &Error{
				Name:    "instanceNotFoundError",
				Message: "We could not find the collection you are looking for",
			},
			postmanStatus: http.StatusNotFound,
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
				if r.Method != http.MethodGet {
					t.Errorf("expected http method GET, got: %s", r.Method)
				}
				if r.URL.Path != "/"+tt.args.uid {
					t.Errorf("expected path to be: %s, got: %s", "/"+tt.args.uid, r.URL.Path)
				}

				for _, o := range tt.args.opts {
					switch v := o.(type) {
					case workspaceOption:
						if r.URL.Query().Get("workspace") != string(v) {
							t.Errorf("expected workspace query param = %s, got: %s", string(v), r.URL.Query().Get("workspace"))
						}
					}
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

				res := getCollectionSuccessResponse
				data, err := json.Marshal(res)
				if err != nil {
					t.Error(err)
				}

				w.WriteHeader(http.StatusOK)
				w.Write(data)
			}))
			t.Cleanup(srv.Close)

			c := &CollectionsClient{
				httpClient: srv.Client(),
				apiKey:     tt.fields.apiKey,
				baseURL:    srv.URL,
			}
			got, err := c.Get(tt.args.ctx, tt.args.uid, tt.args.opts...)
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
