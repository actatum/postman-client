[![Go Report Card](https://goreportcard.com/badge/actatum/postman-client)](https://goreportcard.com/report/actatum/postman-client)
![Build Status](https://github.com/actatum/postman-client/actions/workflows/main.yaml/badge.svg)
[![codecov](https://codecov.io/gh/actatum/postman-client/branch/main/graph/badge.svg)](https://codecov.io/gh/actatum/postman-client)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/actatum/postman-client)
[![Release](https://img.shields.io/github/release/actatum/postman-client.svg)](https://github.com/actatum/postman-client/releases/latest)

# Go client for Postman REST API

## Coverage

### API

- [] POST /apis
- [] GET /apis
- [] GET /apis/:apiId
- [] PUT /apis/:apiId
- [] DELETE /apis/:apiId
- [] POST /apis/versions
- [] GET /apis/versions
- [] GET /apis/:apiId/versions/:versionId
- [] PUT /apis/:apiId/versions/:versionId
- [] DELETE /apis/:apiId/versions/:versionId
- [] POST /apis/versions/releases
- [] GET /apis/versions/releases
- [] GET /apis/:apiId/versions/:versionId/releases/:releaseId
- [] PUT /apis/:apiId/versions/:versionId/releases/:releaseId
- [] DELETE /apis/:apiId/versions/:versionId/releases/:releaseId
- [] POST /apis/:apiId/versions/:versionId/schemas
- [] GET /apis/:apiId/versions/:versionId/schemas/:schemaId
- [] PUT /apis/:apiId/versions/:versionId/schemas/:schemaId
- [] POST /apis/:apiId/versions/:versionId/schemas/:schemaId/collections
- [] POST /apis/:apiId/versions/:versionId/relations
- [] GET /apis/:apiId/versions/:versionId/relations
- [] GET /apis/:apiId/versions/:versionId/test
- [] GET /apis/:apiId/versions/:versionId/testsuite
- [] GET /apis/:apiId/versions/:versionId/contracttest
- [] GET /apis/:apiId/versions/:versionId/environment
- [] GET /apis/:apiId/versions/:versionId/integrationtest
- [] GET /apis/:apiId/versions/:versionId/documentation
- [] GET /apis/:apiId/versions/:versionId/monitor
- [] PUT /apis/:apiId/versions/:versionId/:relationType/:entityId/syncWithSchema

### API Security

- [x] POST /security/api-validation

### Audit Logs

- [x] GET /audit/logs

### Collections

- [x] POST /collections
- [x] GET /collections
- [x] GET /collections/:id
- [x] PUT /collections/:id
- [x] DELETE /collections/:id
- [x] POST /collections/fork/:id
- [x] POST /collections/merge

### Environments

- [x] POST /environments
- [x] GET /environments
- [x] GET /environments/:id
- [x] PUT /environments/:id
- [x] DELETE /environments/:id

### Mocks

- [] POST /mocks
- [] GET /mocks
- [] GET /mocks/:mockId
- [] PUT /mocks/:mockId
- [] DELETE /mocks/:mockId
- [] POST /mocks/:mockId/server-responses
- [] GET /mocks/:mockId/server-responses
- [] GET /mocks/:mockId/server-responses/:serverResponseId
- [] PUT /mocks/:mockId/server-respones/:serverResponseId
- [] DELETE /mocks/:mockId/server-responses/:serverResponseId
- [] POST /mocks/:mockId/publish
- [] DELETE /mocks/:mockId/unpublish 
- [] GET /mocks/:mockId/call-logs

### Monitors

- [x] POST /monitors
- [x] GET /monitors
- [x] GET /monitors/:id
- [x] PUT /monitors/:id
- [x] DELETE /monitors/:id
- [x] POST /monitors/:id/run

### Workspaces

- [x] POST /workspaces
- [x] GET /workspaces
- [x] GET /workspaces/:id
- [x] PUT /workspaces/:id
- [x] DELETE /workspaces/:id

### User

- [x] GET /me

### Import

- [] POST /import/openapi
- [] POST /import/exported

### Webhooks

- [x] POST /webhooks

### SCIM 2.0 - Identity

- [] GET /scim/v2/ResourceTypes
- [] GET /scim/v2/ServiceProviderConfig
- [] POST /scim/v2/Users
- [] GET /scim/v2/Users
- [] GET /scim/v2/Users/:id
- [] PUT /scim/v2/Users/:id
- [] PATCH /scim/v2/Users/:id
- [] POST /scim/v2/Groups
- [] GET /scim/v2/Groups
- [] GET /scim/v2/Groups/:id
- [] PATCH /scim/v2/Group/:id

## Usage

```go
package main

import (
	"context"
	"fmt"
	
	"github.com/actatum/postman-client"
	"github.com/actatum/postman-client/webhooks"
)

func main() {
	// Create a new client set.
	cs := postman.NewClientSet("api-key")
	// Get a handle to an individual client.
	webhookClient := cs.Webhooks()
	
	// Make a request with handle for individual client.
	webhook, err := webhookClient.Create(
		context.Background(),
		webhooks.Webhook{
			Name: "Test Webhook",
			Collection: "collection-id",
        },
	)
	if err != nil {
		panic(err)
    }
	
	fmt.Printf("%v\n", webhook)
}
```

### Missing endpoints

It's possible some endpoints may be missing from the client. You can use methods from the `rest.Client`
to call an endpoint. `rest.NewClient -> client.NewRequest -> client.DoRequest`.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	
	"github.com/actatum/postman-client/rest"
)

func main() {
	// Create a new rest client.
	rc := rest.NewClient("api-key")
	body := map[string]map[string]string{
		"workspace": {
			"name": "Test Workspace",
			"type": "personal",
			"description": "This is a test personal workspace.",
        },
    }
	
	// Create http request
	r, err := rc.NewRequest(
		context.Background(),
		http.MethodPost,
		"https://api.getpostman.com/workspaces",
		body,
	)
	if err != nil {
		panic(err)
    }
	
	// Do request and unmarshal into response object
	var response map[string]interface{}
	err = rc.DoRequest(r, &response)
	if err != nil {
		panic(err)
    }
	
	fmt.Printf("Response: %v\n", response)
}
```

## How to Contribute

* Fork this repository
* Add your change or fix
* Make sure tests pass
* Create a pull request

Current contributors:

- [Aaron Tatum](https://github.com/actatum)

## Tests

* Unit tests: `go test -v -race -cover -short ./...`
* Unit & Integration tests: `go test -v -race -cover ./...`