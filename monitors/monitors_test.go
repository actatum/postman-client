package monitors

import (
	"testing"
)

func TestClient_Monitors(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	//rc := rest.NewClient(testdata.TestAPIKey, rest.WithDebugLog(os.Stdout))
	//c := NewClient(rc)
	//
	//t.Run("full lifecycle", func(t *testing.T) {
	//	ctx := context.Background()
	//
	//	// Create failure (bad request)
	//	req := Monitor{
	//		Name: "Test Monitor",
	//		Schedule: Schedule{
	//			Cron:     "* * * * *",
	//			Timezone: "America/Chicago",
	//		},
	//		Collection:  testdata.TestCollectionID,
	//		Environment: testdata.TestEnvironmentID,
	//	}
	//	_, err := c.Create(ctx, req)
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//
	//	// Create success
	//	req = Monitor{
	//		Name: "Test Monitor",
	//		Schedule: Schedule{
	//			Cron:     "5 4 * * *",
	//			Timezone: "America/Chicago",
	//		},
	//		Collection:  testdata.TestCollectionID,
	//		Environment: testdata.TestEnvironmentID,
	//	}
	//	environment, err := c.Create(ctx, req)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// Get Not Found
	//	_, err = c.Get(ctx, "123")
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//
	//	// Get recently created environment
	//	environment, err = c.Get(ctx, environment.ID)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// GetAll
	//	_, err = c.GetAll(ctx)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// Update Not Found
	//	req = Monitor{
	//		Name: "Test Monitor",
	//		Schedule: Schedule{
	//			Cron:     "* * * * *",
	//			Timezone: "America/Chicago",
	//		},
	//		Collection:  testdata.TestCollectionID,
	//		Environment: testdata.TestEnvironmentID,
	//	}
	//	_, err = c.Update(ctx, "123", req)
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//
	//	// Update success
	//	req = Monitor{
	//		Name: "Test Monitor",
	//		Schedule: Schedule{
	//			Cron:     "5 4 * * *",
	//			Timezone: "America/Chicago",
	//		},
	//		Collection:  testdata.TestCollectionID,
	//		Environment: testdata.TestEnvironmentID,
	//	}
	//	_, err = c.Update(ctx, environment.ID, req)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// Delete Not Found
	//	_, err = c.Delete(ctx, "123")
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//
	//	// Delete success
	//	_, err = c.Delete(ctx, environment.ID)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// Verify Deletion
	//	_, err = c.Get(ctx, environment.ID)
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//})
	//
	//t.Run("full lifecycle in workspace", func(t *testing.T) {
	//	ctx := context.Background()
	//
	//	workspaceOpt := rest.WithWorkspace(testdata.TestWorkspaceID)
	//
	//	// Create failure (bad request)
	//	req := Monitor{
	//		Name: "Test Monitor",
	//		Schedule: Schedule{
	//			Cron:     "* * * * *",
	//			Timezone: "America/Chicago",
	//		},
	//		Collection:  testdata.TestCollectionID,
	//		Environment: testdata.TestEnvironmentID,
	//	}
	//	_, err := c.Create(ctx, req, workspaceOpt)
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//
	//	// Create success
	//	req = Monitor{
	//		Name: "Test Monitor",
	//		Schedule: Schedule{
	//			Cron:     "0 0 * * *",
	//			Timezone: "America/Chicago",
	//		},
	//		Collection:  testdata.TestCollectionID,
	//		Environment: testdata.TestEnvironmentID,
	//	}
	//	environment, err := c.Create(ctx, req, workspaceOpt)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// Get Not Found
	//	_, err = c.Get(ctx, "123", workspaceOpt)
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//
	//	// Get recently created environment
	//	environment, err = c.Get(ctx, environment.ID, workspaceOpt)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// GetAll
	//	_, err = c.GetAll(ctx, workspaceOpt)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// Update Not Found
	//	req = Monitor{
	//		Name: "Test Monitor",
	//		Schedule: Schedule{
	//			Cron:     "* * * * *",
	//			Timezone: "America/Chicago",
	//		},
	//		Collection:  testdata.TestCollectionID,
	//		Environment: testdata.TestEnvironmentID,
	//	}
	//	_, err = c.Update(ctx, "123", req, workspaceOpt)
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//
	//	// Update success
	//	req = Monitor{
	//		Name: "Test Monitor",
	//		Schedule: Schedule{
	//			Cron:     "0 0 * * *",
	//			Timezone: "America/Chicago",
	//		},
	//		Collection:  testdata.TestCollectionID,
	//		Environment: testdata.TestEnvironmentID,
	//	}
	//	_, err = c.Update(ctx, environment.ID, req, workspaceOpt)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// Delete Not Found
	//	_, err = c.Delete(ctx, "123", workspaceOpt)
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//
	//	// Delete success
	//	_, err = c.Delete(ctx, environment.ID, workspaceOpt)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	// Verify Deletion
	//	_, err = c.Get(ctx, environment.ID, workspaceOpt)
	//	if err == nil {
	//		t.Fatalf("expected error, got %v", err)
	//	}
	//})
}
