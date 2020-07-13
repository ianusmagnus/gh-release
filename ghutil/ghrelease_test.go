package ghutil

import (
	"context"
	"github.com/ianusmagnus/gh-release/ghutil/testclient"
	"testing"
)

func TestReleaseCreator(t *testing.T) {
	ctx := context.Background()
	client := testclient.New(ctx, "pat")
	creator := NewReleaseCreator(client, ctx, "username", "repo")

	err := creator.CreateNewRelease("test")
	if err != nil {
		t.Errorf("Failed: %o", err)
	}
}
