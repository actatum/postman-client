// Package environments provides types/client for making requests to /environments.
package environments

import (
	"context"
	"testing"

	"github.com/actatum/postman-client/rest"
	"github.com/actatum/postman-client/testdata"
)

func TestClient_Environments(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	rc := rest.NewClient(testdata.TestAPIKey)
	c := NewClient(rc)

	t.Run("full lifecycle", func(t *testing.T) {
		ctx := context.Background()

		// Create failure (bad request)
		req := Environment{
			Values: []EnvironmentValue{
				{
					Key:     "apiKey",
					Value:   "PMAK-1234-5678-0987-6543",
					Enabled: true,
					Type:    EnvironmentValueTypeSecret,
				},
				{
					Key:     "variable_1",
					Value:   "The variable_1 value.",
					Enabled: false,
					Type:    EnvironmentValueTypeDefault,
				},
			},
		}
		_, err := c.Create(ctx, req)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		//Create success
		req = Environment{
			Name: "Test Environment",
			Values: []EnvironmentValue{
				{
					Key:     "apiKey",
					Value:   "PMAK-1234-5678-0987-6543",
					Enabled: true,
					Type:    EnvironmentValueTypeSecret,
				},
				{
					Key:     "variable_1",
					Value:   "The variable_1 value.",
					Enabled: false,
					Type:    EnvironmentValueTypeDefault,
				},
			},
		}
		environment, err := c.Create(ctx, req)
		if err != nil {
			t.Fatal(err)
		}

		// Get Not Found
		_, err = c.Get(ctx, "123")
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Get recently created environment
		environment, err = c.Get(ctx, environment.ID)
		if err != nil {
			t.Fatal(err)
		}

		// GetAll
		_, err = c.GetAll(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// Update Not Found
		req = Environment{
			Name: "Test Environment",
			Values: []EnvironmentValue{
				{
					Key:     "variable_a",
					Value:   "The variable_a value.",
					Enabled: true,
					Type:    EnvironmentValueTypeDefault,
				},
			},
		}
		_, err = c.Update(ctx, "123", req)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Update success
		req = Environment{
			Name: "Test Environment",
			Values: []EnvironmentValue{
				{
					Key:     "variable_a",
					Value:   "The variable_a value.",
					Enabled: true,
					Type:    EnvironmentValueTypeDefault,
				},
			},
		}
		_, err = c.Update(ctx, environment.ID, req)
		if err != nil {
			t.Fatal(err)
		}

		// Delete Not Found
		_, err = c.Delete(ctx, "123")
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Delete success
		_, err = c.Delete(ctx, environment.ID)
		if err != nil {
			t.Fatal(err)
		}

		// Verify Deletion
		_, err = c.Get(ctx, environment.ID)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}
	})

	t.Run("full lifecycle in workspace", func(t *testing.T) {
		ctx := context.Background()

		workspaceOpt := rest.WithWorkspace(testdata.TestWorkspaceID)

		// Create failure (bad request)
		req := Environment{
			Values: []EnvironmentValue{
				{
					Key:     "apiKey",
					Value:   "PMAK-1234-5678-0987-6543",
					Enabled: true,
					Type:    EnvironmentValueTypeSecret,
				},
				{
					Key:     "variable_1",
					Value:   "The variable_1 value.",
					Enabled: false,
					Type:    EnvironmentValueTypeDefault,
				},
			},
		}
		_, err := c.Create(ctx, req, workspaceOpt)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		//Create success
		req = Environment{
			Name: "Test Environment",
			Values: []EnvironmentValue{
				{
					Key:     "apiKey",
					Value:   "PMAK-1234-5678-0987-6543",
					Enabled: true,
					Type:    EnvironmentValueTypeSecret,
				},
				{
					Key:     "variable_1",
					Value:   "The variable_1 value.",
					Enabled: false,
					Type:    EnvironmentValueTypeDefault,
				},
			},
		}
		environment, err := c.Create(ctx, req, workspaceOpt)
		if err != nil {
			t.Fatal(err)
		}

		// Get Not Found
		_, err = c.Get(ctx, "123", workspaceOpt)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Get recently created environment
		environment, err = c.Get(ctx, environment.ID, workspaceOpt)
		if err != nil {
			t.Fatal(err)
		}

		// GetAll
		_, err = c.GetAll(ctx, workspaceOpt)
		if err != nil {
			t.Fatal(err)
		}

		// Update Not Found
		req = Environment{
			Name: "Test Environment",
			Values: []EnvironmentValue{
				{
					Key:     "variable_a",
					Value:   "The variable_a value.",
					Enabled: true,
					Type:    EnvironmentValueTypeDefault,
				},
			},
		}
		_, err = c.Update(ctx, "123", req, workspaceOpt)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Update success
		req = Environment{
			Name: "Test Environment",
			Values: []EnvironmentValue{
				{
					Key:     "variable_a",
					Value:   "The variable_a value.",
					Enabled: true,
					Type:    EnvironmentValueTypeAny,
				},
			},
		}
		_, err = c.Update(ctx, environment.ID, req, workspaceOpt)
		if err != nil {
			t.Fatal(err)
		}

		// Delete Not Found
		_, err = c.Delete(ctx, "123", workspaceOpt)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Delete success
		_, err = c.Delete(ctx, environment.ID, workspaceOpt)
		if err != nil {
			t.Fatal(err)
		}

		// Verify Deletion
		_, err = c.Get(ctx, environment.ID, workspaceOpt)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}
	})
}
