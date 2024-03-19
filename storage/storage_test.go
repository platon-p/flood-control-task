package storage

import (
	"context"
	"testing"
	"time"
)

func TestStorage_GetAttemptsCount(t *testing.T) {
	storage := NewAttemptsMemoryStorage()
	userID := int64(1)
	_, err := storage.GetAttemptsCount(context.Background(), userID, time.Minute)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestStorage_RegisterAttempt(t *testing.T) {
	storage := NewAttemptsMemoryStorage()
	userID := int64(1)
	err := storage.RegisterAttempt(context.Background(), userID, time.Now())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
