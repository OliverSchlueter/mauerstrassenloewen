package sloki_test

import (
	"context"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"testing"
)

func TestWrapContext(t *testing.T) {
	fn := func(ctx context.Context) string {
		return "test value"
	}

	sloki.RegisterContextFunc("testKey", fn)

	got := sloki.WrapContext(context.Background())
	if got.Key != "context" {
		t.Errorf("expected key 'context', got %s", got.Key)
	}
	if len(got.Value.Group()) != 1 {
		t.Errorf("expected 1 attribute, got %d", len(got.Value.Group()))
	}
}
