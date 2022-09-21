package postman

import (
	"net/http"
)

const postmanBaseURL = "https://api.getpostman.com"

// ClientSet contains the clients for groups. Each group has a single client.
type ClientSet struct {
	collections  *CollectionsClient
	environments *EnvironmentsClient
	monitors     *MonitorsClient
	workspaces   *WorkspacesClient
	users        *UsersClient
	imports      *ImportClient
	webhooks     *WebhooksClient
	apiSecurity  *APISecurityClient
	auditLogs    *AuditLogsClient
}

// New returns a new instance of the ClientSet.
func New(apiKey string, opts ...Option) *ClientSet {
	options := options{
		httpClient: &http.Client{},
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &ClientSet{
		collections:  NewCollectionsClient(apiKey, options.httpClient),
		environments: NewEnvironmentsClient(apiKey, options.httpClient),
		monitors:     NewMonitorsClient(apiKey, options.httpClient),
		workspaces:   NewWorkspacesClient(apiKey, options.httpClient),
		users:        NewUsersClient(apiKey, options.httpClient),
		imports:      NewImportClient(apiKey, options.httpClient),
		webhooks:     NewWebhooksClient(apiKey, options.httpClient),
		apiSecurity:  NewAPISecurityClient(apiKey, options.httpClient),
		auditLogs:    NewAuditLogsClient(apiKey, options.httpClient),
	}
}

// Collections returns a handle to a CollectionsClient.
func (cs *ClientSet) Collections() *CollectionsClient {
	return cs.collections
}

// Environments returns a handle to an EnvironmentsClient.
func (cs *ClientSet) Environments() *EnvironmentsClient {
	return cs.environments
}

// Monitors returns a handle to a MonitorsClient.
func (cs *ClientSet) Monitors() *MonitorsClient {
	return cs.monitors
}

// Workspaces returns a handle to a WorkspacesClient.
func (cs *ClientSet) Workspaces() *WorkspacesClient {
	return cs.workspaces
}

// Users returns a handle to a UsersClient.
func (cs *ClientSet) Users() *UsersClient {
	return cs.users
}

// Import returns a handle to an ImportClient.
func (cs *ClientSet) Import() *ImportClient {
	return cs.imports
}

// Webhooks returns a handle to a WebhooksClient.
func (cs *ClientSet) Webhooks() *WebhooksClient {
	return cs.webhooks
}

// APISecurity returns a handle to an APISecurityClient.
func (cs *ClientSet) APISecurity() *APISecurityClient {
	return cs.apiSecurity
}

// AuditLogs returns a handle to an AuditLogsClient
func (cs *ClientSet) AuditLogs() *AuditLogsClient {
	return cs.auditLogs
}
