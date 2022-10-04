// Package webhooks provides types/client for making requests to /webhooks.
package webhooks

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/actatum/postman-client/rest"
	"github.com/actatum/postman-client/testdata"
)

func TestClient_Create(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	rc := rest.NewClient(testdata.TestAPIKey)
	c := NewClient(rc)

	t.Run("create webhook invalid webhook", func(t *testing.T) {
		wh := Webhook{
			Name:       "",
			Collection: testdata.TestCollectionID,
		}
		_, err := c.Create(context.Background(), wh)
		if err == nil {
			t.Fatal("expected error got nil")
		}

		var e *rest.Error
		ok := errors.As(err, &e)
		if !ok {
			t.Fatalf("errors.As() got = %T, want %T", err, rest.Error{})
		}

		wantErr := &rest.Error{
			Name:    "validationError",
			Message: "name cannot be empty.",
		}

		if !reflect.DeepEqual(e, wantErr) {
			t.Fatalf("c.Create() error got = %v, want %v", e, wantErr)
		}
	})

	t.Run("create webhook success", func(t *testing.T) {
		wh := Webhook{
			Name:       "Test Webhook",
			Collection: testdata.TestCollectionID,
		}

		_, err := c.Create(context.Background(), wh)
		if err != nil {
			t.Fatal(err)
		}
	})
}
