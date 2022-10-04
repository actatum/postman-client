package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/actatum/postman-client/rest"
	"github.com/actatum/postman-client/testdata"
)

func TestClient_GetAuthenticatedUser(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	rc := rest.NewClient(testdata.TestAPIKey)
	c := NewClient(rc)

	t.Run("me", func(t *testing.T) {
		got, _, err := c.GetAuthenticatedUser(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		want := User{
			ID:       23646813,
			Username: "avionics-pilot-21839645",
			Email:    "atatum0413@gmail.com",
			FullName: "Aaron Tatum",
			Avatar:   "",
			IsPublic: true,
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("GetAuthenticatedUser() got = %v, want %v", got, want)
		}
	})
}
