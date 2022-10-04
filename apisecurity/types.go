// Package apisecurity provides types/client for making requests to /security/api-validation.
package apisecurity

// Possible values for openapi versions.
const (
	OpenAPIV3 = "openapi3"
	OpenAPIV2 = "openapi2"
)

// Possible values for openapi spec languages.
const (
	LanguageJSON = "json"
	LanguageYAML = "yaml"
)

// ValidateAPISchemaRequest is the request type for POST /security/api-validation.
type ValidateAPISchemaRequest struct {
	Schema APISchema `json:"schema"`
}

// ValidateAPISchemaResponse is the response type for POST /security/api-validation.
type ValidateAPISchemaResponse struct {
	Warnings []Warning `json:"warnings"`
}

// APISchema ...
type APISchema struct {
	Type     string `json:"type"`
	Language string `json:"language"`
	Schema   string `json:"schema"`
}

// Warning ...
type Warning struct {
	Severity       string           `json:"severity"`
	Message        string           `json:"message"`
	Location       WarningLocations `json:"location"`
	DataPath       []string         `json:"dataPath"`
	PossibleFixURL string           `json:"possibleFixUrl"`
	Category       WarningCategory  `json:"category"`
}

// WarningCategory ...
type WarningCategory struct {
	Name string `json:"name"`
}

// WarningLocations ...
type WarningLocations struct {
	Start WarningLocation `json:"start"`
	End   WarningLocation `json:"end"`
}

// WarningLocation ...
type WarningLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
