// Package workspaces provides types/client for making requests to /workspaces.
package workspaces

import "time"

// Possible values for workspace types.
const (
	WorkspaceTypePersonal = "personal"
	WorkspaceTypeTeam     = "team"
)

// Possible values for workspace visibility.
const (
	WorkspaceVisibilityOnlyMe      = "only-me"
	WorkspaceVisibilityPersonal    = "personal"
	WorkspaceVisibilityTeam        = "team"
	WorkspaceVisibilityPrivateTeam = "private-team"
	WorkspaceVisibilityPublic      = "public"
)

// WorkspaceRequest is the request type for POST,PUT /workspaces.
type WorkspaceRequest struct {
	Workspace Workspace `json:"workspace"`
}

// GetAllWorkspacesRequest is the request type for GET /workspaces.
type GetAllWorkspacesRequest struct {
	// The type of workspace to filter the response by (team or personal).
	Type *string
}

// Workspace ...
type Workspace struct {
	ID           string        `json:"id,omitempty"` // Output only.
	Name         string        `json:"name,omitempty"`
	Type         string        `json:"type,omitempty"`
	Description  string        `json:"description,omitempty"`
	Visibility   string        `json:"visibility,omitempty"`   // Output only.
	CreatedBy    string        `json:"createdBy,omitempty"`    // Output only.
	UpdatedBy    string        `json:"updatedBy,omitempty"`    // Output only.
	CreatedAt    time.Time     `json:"createdAt,omitempty"`    // Output only.
	UpdatedAt    time.Time     `json:"updatedAt,omitempty"`    // Output only.
	Collections  []Collection  `json:"collections,omitempty"`  // Output only.
	Environments []Environment `json:"environments,omitempty"` // Output only.
	Mocks        []Mock        `json:"mocks,omitempty"`        // Output only.
	Monitors     []Monitor     `json:"monitors,omitempty"`     // Output only.
	APIs         []API         `json:"apis,omitempty"`         // Output only.
}

// Collection ...
type Collection struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	UID  string `json:"uid"`
}

// Environment ...
type Environment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	UID  string `json:"uid"`
}

// Mock ...
type Mock struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	UID  string `json:"uid"`
}

// Monitor ...
type Monitor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	UID  string `json:"uid"`
}

// API ...
type API struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	UID  string `json:"uid"`
}

type workspaceWrapper struct {
	Workspace  Workspace   `json:"workspace,omitempty"`
	Workspaces []Workspace `json:"workspaces,omitempty"`
}
