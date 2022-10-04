// Package postman provides a client set with handles to all the different postman endpoints.
package postman

import (
	"net/http"

	"github.com/actatum/postman-client/apisecurity"
	"github.com/actatum/postman-client/auditlogs"
	"github.com/actatum/postman-client/collections"
	"github.com/actatum/postman-client/environments"
	"github.com/actatum/postman-client/monitors"
	"github.com/actatum/postman-client/rest"
	"github.com/actatum/postman-client/user"
	"github.com/actatum/postman-client/webhooks"
	"github.com/actatum/postman-client/workspaces"
)

// ClientSet holds handles to all the different postman endpoint clients.
type ClientSet struct {
	apisecurity  *apisecurity.Client
	auditlogs    *auditlogs.Client
	collections  *collections.Client
	environments *environments.Client
	monitors     *monitors.Client
	users        *user.Client
	webhooks     *webhooks.Client
	workspaces   *workspaces.Client
}

// NewClientSet returns a new instance of ClientSet.
func NewClientSet(apiKey string, opts ...Option) *ClientSet {
	options := options{
		httpClient: &http.Client{},
		debugLog:   nil,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	restClient := rest.NewClient(
		apiKey,
		rest.WithHTTPClient(options.httpClient),
		rest.WithDebugLog(options.debugLog),
	)
	return &ClientSet{
		apisecurity:  apisecurity.NewClient(restClient),
		auditlogs:    auditlogs.NewClient(restClient),
		collections:  collections.NewClient(restClient),
		environments: environments.NewClient(restClient),
		monitors:     monitors.NewClient(restClient),
		users:        user.NewClient(restClient),
		webhooks:     webhooks.NewClient(restClient),
		workspaces:   workspaces.NewClient(restClient),
	}
}

// APISecurity returns a handle to an apiesecurity.Client.
func (cs *ClientSet) APISecurity() *apisecurity.Client {
	return cs.apisecurity
}

// AuditLogs returns a handle to an auditlogs.Client.
func (cs *ClientSet) AuditLogs() *auditlogs.Client {
	return cs.auditlogs
}

// Collections returns a handle to a collections.Client.
func (cs *ClientSet) Collections() *collections.Client {
	return cs.collections
}

// Environments returns a handle to an environments.Client.
func (cs *ClientSet) Environments() *environments.Client {
	return cs.environments
}

// Monitors returns a handle to a monitors.Client.
func (cs *ClientSet) Monitors() *monitors.Client {
	return cs.monitors
}

// Users returns a handle to a user.Client.
func (cs *ClientSet) Users() *user.Client {
	return cs.users
}

// Webhooks returns a handle to a webhooks.Client.
func (cs *ClientSet) Webhooks() *webhooks.Client {
	return cs.webhooks
}

// Workspaces returns a handle to a workspaces.Client.
func (cs *ClientSet) Workspaces() *workspaces.Client {
	return cs.workspaces
}
