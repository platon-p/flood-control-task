package storage

import (
	"context"
	"testing"
)

func TestBlackListMemoryStorage(t *testing.T) {
	storage := NewBlackListMemoryStorage()
	storage.blacklist[1] = struct{}{}
	ctx := context.Background()
	if contains, err := storage.Contains(ctx, 1); err != nil {
		t.Error(err)
	} else if !contains {
		t.Error("user should be in the blacklist")
	}
	if contains, err := storage.Contains(ctx, 2); err != nil {
		t.Error(err)
	} else if contains {
		t.Error("user should not be in the blacklist")
	}
}
