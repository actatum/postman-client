// Package environments provides types/client for making requests to /environments.
package environments

import "time"

// Possible values for environment value types.
const (
	EnvironmentValueTypeSecret  = "secret"
	EnvironmentValueTypeDefault = "default"
	EnvironmentValueTypeAny     = "any"
)

// Environment ...
type Environment struct {
	ID        string             `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Owner     string             `json:"owner,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty"`
	Values    []EnvironmentValue `json:"values,omitempty"`
	IsPublic  bool               `json:"isPublic,omitempty"`
	UID       string             `json:"uid,omitempty"`
}

// EnvironmentValue ...
type EnvironmentValue struct {
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Type    string `json:"type,omitempty"`
}

type environmentWrapper struct {
	Environment  Environment   `json:"environment,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
}
