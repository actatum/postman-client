package testdata

import _ "embed"

// APISecurityValidationSchemaJSON is an example openapi schema in json format.
//
//go:embed apisecurity.schema.json
var APISecurityValidationSchemaJSON string

// APISecurityValidationSchemaYAML is an example openapi schema in yaml format.
//
//go:embed apisecurity.schema.yaml
var APISecurityValidationSchemaYAML string
