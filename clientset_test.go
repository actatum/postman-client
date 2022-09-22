package postman

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		apiKey string
		opts   []Option
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "new",
			args: args{
				apiKey: "api-key",
				opts:   nil,
			},
		},
		{
			name: "new w/httpClient",
			args: args{
				apiKey: "api-key",
				opts:   []Option{WithHTTPClient(&http.Client{})},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.apiKey, tt.args.opts...)
			if got.collections == nil {
				t.Errorf("expected collections to not be nil")
			}
			if got.environments == nil {
				t.Errorf("expected environments to not be nil")
			}
			if got.monitors == nil {
				t.Errorf("expected monitors to not be nil")
			}
			if got.workspaces == nil {
				t.Errorf("expected workspaces to not be nil")
			}
			if got.users == nil {
				t.Errorf("expected users to not be nil")
			}
			if got.imports == nil {
				t.Errorf("expected imports to not be nil")
			}
			if got.webhooks == nil {
				t.Errorf("expected webhooks to not be nil")
			}
			if got.apiSecurity == nil {
				t.Errorf("expected apiSecurity to not be nil")
			}
			if got.auditLogs == nil {
				t.Errorf("expected auditLogs to not be nil")
			}

			if got.Collections() != got.collections {
				t.Errorf("expected Collections() to return collections")
			}
			if got.Environments() != got.environments {
				t.Errorf("expected Environments() to return environments")
			}
			if got.Monitors() != got.monitors {
				t.Errorf("expected Monitors() to return monitors")
			}
			if got.Collections() != got.collections {
				t.Errorf("expected Collections() to return collections")
			}
			if got.Workspaces() != got.workspaces {
				t.Errorf("expected Workspaces() to return workspaces")
			}
			if got.Users() != got.users {
				t.Errorf("expected Users() to return users")
			}
			if got.Import() != got.imports {
				t.Errorf("expected Import() to return imports")
			}
			if got.Webhooks() != got.webhooks {
				t.Errorf("expected Webhooks() to return webhooks")
			}
			if got.APISecurity() != got.apiSecurity {
				t.Errorf("expected APISecurity() to return apiSecurity")
			}
			if got.AuditLogs() != got.auditLogs {
				t.Errorf("expected AuditLogs() to return auditLogs")
			}
		})
	}
}
