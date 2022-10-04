// Package auditlogs provides types/client for making requests to /audit/logs.
package auditlogs

import "time"

// GetAuditLogsRequest is the request type for GET /audit/logs
type GetAuditLogsRequest struct {
	// Return logs created after the given time, in YYYY-MM-DD format.
	Since *string
	// Return logs created before the given time, in YYYY-MM-DD format.
	Until *string
	// The maximum number of audit events to return at once. The maximum value is 300.
	Limit *int
	// The cursor to get the next set of results.
	Cursor *int
	// Return the records in ascending ('ASC') or descending ('DESC') order. This value defaults to 'DESC' order.
	OrderBy *string
}

// AuditLogs is the response type for /audit/logs operations.
type AuditLogs struct {
	Trails []Trail `json:"trails"`
}

// Trail ...
type Trail struct {
	ID        int       `json:"id"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"userAgent"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      TrailData `json:"data"`
}

// TrailData ...
type TrailData struct {
	Actor Actor `json:"actor"`
	User  User  `json:"user"`
	Team  Team  `json:"team"`
}

// Actor ...
type Actor struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Active   bool   `json:"active"`
}

// Team ...
type Team struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// User ...
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
}
