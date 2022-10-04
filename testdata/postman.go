// Package testdata provides constants for use in testing the postman client.
package testdata

import _ "embed"

const (
	// TestAPIKey is the API Key for the postman-client test account.
	TestAPIKey = "PMAK-63379469012d4655cd7a3531-97d68d5f3b3ff1ecd05010653bfc1f4d57"
	// TestWorkspaceID is the workspace id for the postman-client test workspace.
	TestWorkspaceID = "926daa5f-38db-405c-948f-3422b6300d4e"
	// TestCollectionInTestWorkspaceID is the collection id for the postman-client test collection in the postman-client test workspace.
	TestCollectionInTestWorkspaceID = "23646813-2e168694-3c47-4893-8b8c-3505db1b59a2"
	// TestCollectionID is the collection id for the postman-client test collection in the base workspace.
	TestCollectionID = "23646813-533b27c3-1838-458b-886f-f05bae867ac7"
	// TestEnvironmentID is the collection id for the postman-client test environment in the base workspace.
	TestEnvironmentID = "23646813-f1547ce0-2431-4ee0-8240-1a934fc74971"
)

const (
	AuthenticationErrorString = "AuthenticationError: Invalid API Key. Every request requires a valid API Key to be sent."
)

func IsUnauthorizedError(err error) bool {
	return err.Error() == AuthenticationErrorString
}
