// Package apisecurity provides types/client for making requests to /security/api-validation.
package apisecurity

import (
	"context"
	"testing"

	"github.com/actatum/postman-client/rest"
	"github.com/actatum/postman-client/testdata"
)

func TestClient_ValidateAPISchema(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	rc := rest.NewClient(testdata.TestAPIKey)
	c := NewClient(rc)

	t.Run("valid json schema", func(t *testing.T) {
		req := ValidateAPISchemaRequest{
			Schema: APISchema{
				Type:     OpenAPIV3,
				Language: LanguageJSON,
				Schema:   testdata.APISecurityValidationSchemaJSON,
			},
		}
		res, err := c.ValidateAPISchema(context.Background(), req)
		if err != nil {
			t.Error(err)
			return
		}

		if len(res) != 3 {
			t.Fatalf("len(res) got = %v, want %v", len(res), 3)
		}
	})

	t.Run("valid yaml schema", func(t *testing.T) {
		req := ValidateAPISchemaRequest{
			Schema: APISchema{
				Type:     OpenAPIV3,
				Language: LanguageYAML,
				Schema:   testdata.APISecurityValidationSchemaYAML,
			},
		}
		res, err := c.ValidateAPISchema(context.Background(), req)
		if err != nil {
			t.Error(err)
			return
		}

		if len(res) != 4 {
			t.Fatalf("len(res) got = %v, want %v", len(res), 4)
		}
	})

	t.Run("invalid schema type", func(t *testing.T) {
		req := ValidateAPISchemaRequest{
			Schema: APISchema{
				Type:     "openapi1",
				Language: LanguageJSON,
				Schema:   testdata.APISecurityValidationSchemaJSON,
			},
		}
		_, err := c.ValidateAPISchema(context.Background(), req)
		if err.Error() != "Invalid schema: Provided schema type is not supported." {
			t.Fatalf(
				"err.Error() got = %v, want %v",
				err.Error(),
				"Invalid schema: Provided schema type is not supported.",
			)
		}
	})

	t.Run("invalid schema type", func(t *testing.T) {
		req := ValidateAPISchemaRequest{
			Schema: APISchema{
				Type:     OpenAPIV3,
				Language: LanguageJSON,
				Schema:   "{}",
			},
		}
		_, err := c.ValidateAPISchema(context.Background(), req)
		if err.Error() != "Invalid Schema: Specification must contain a semantic version number of the OAS specification" {
			t.Fatalf(
				"err.Error() got = %v, want %v",
				err.Error(),
				"Invalid Schema: Specification must contain a semantic version number of the OAS specification",
			)
		}
	})
}
