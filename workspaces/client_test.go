// Package workspaces provides types/client for making requests to /workspaces.
package workspaces

import (
	"context"
	"testing"

	"github.com/actatum/postman-client/rest"
	"github.com/actatum/postman-client/testdata"
)

func TestClient_Workspaces(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	rc := rest.NewClient(testdata.TestAPIKey)
	c := NewClient(rc)

	t.Run("full lifecycle", func(t *testing.T) {
		ctx := context.Background()

		// Create failure (bad request)
		req := Workspace{
			Name:        "",
			Description: "This is a test personal workspace",
			Type:        WorkspaceTypePersonal,
		}
		_, err := c.Create(ctx, req)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Create success
		req = Workspace{
			Name:        "Test Workspace",
			Description: "This is a test personal workspace",
			Type:        WorkspaceTypePersonal,
		}
		workspace, err := c.Create(ctx, req)
		if err != nil {
			t.Fatal(err)
		}

		// Get Not Found
		_, err = c.Get(ctx, "123")
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Get recently created workspace
		workspace, err = c.Get(ctx, workspace.ID)
		if err != nil {
			t.Fatal(err)
		}

		// GetAll
		_, err = c.GetAll(ctx, GetAllWorkspacesRequest{})
		if err != nil {
			t.Fatal(err)
		}

		// Update Not Found
		req = Workspace{
			Name:        "Test Workspace",
			Description: "This is a test personal workspace",
			Type:        WorkspaceTypePersonal,
		}
		_, err = c.Update(ctx, "123", req)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Update success
		req = Workspace{
			Name:        "Test Workspace",
			Description: "This is a test personal workspace :)",
			Type:        WorkspaceTypePersonal,
		}
		_, err = c.Update(ctx, workspace.ID, req)
		if err != nil {
			t.Fatal(err)
		}

		// Delete Not Found
		_, err = c.Delete(ctx, "123")
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		// Delete success
		_, err = c.Delete(ctx, workspace.ID)
		if err != nil {
			t.Fatal(err)
		}

		// Verify Deletion
		_, err = c.Get(ctx, workspace.ID)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}
	})
}
