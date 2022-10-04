// Package auditlogs provides types/client for making requests to /audit/logs.
package auditlogs

import (
	"context"
	"testing"

	"github.com/actatum/postman-client/rest"
	"github.com/actatum/postman-client/testdata"
)

func TestClient_GetAuditLogs(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	t.Run("forbidden", func(t *testing.T) {
		rc := rest.NewClient(testdata.TestAPIKey)
		c := NewClient(rc)

		req := GetAuditLogsRequest{}
		_, err := c.Get(context.Background(), req)
		if err.Error() != "ForbiddenRequest: You do not have permissions to view team Audit logs" {
			t.Errorf(
				"err.Error() got = %v, want %v",
				err.Error(),
				"ForbiddenRequest: You do not have permissions to view team Audit logs",
			)
		}
	})
}
