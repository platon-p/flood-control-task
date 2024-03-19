package main

import (
	"context"
	"testing"
	"time"
)

type AttemptsStorageMock struct {
	attempts map[int64][]time.Time
}

func NewAttemptsStorageMock() *AttemptsStorageMock {
	return &AttemptsStorageMock{
		attempts: make(map[int64][]time.Time),
	}
}

func (s *AttemptsStorageMock) GetAttemptsCount(ctx context.Context, userID int64, period time.Duration) (int, error) {
	attempts := s.attempts[userID]
	count := 0
	for _, t := range attempts {
		if time.Since(t) <= period {
			count++
		}
	}
	return count, nil
}

func (s *AttemptsStorageMock) RegisterAttempt(ctx context.Context, userID int64, timestamp time.Time) error {
	s.attempts[userID] = append(s.attempts[userID], timestamp)
	return nil
}

func TestFloodControl_Check(t *testing.T) {
	storage := NewAttemptsStorageMock()
	control := NewFloodControl(FloodControlOptions{
		AllowedPeriod:   time.Minute,
		AllowedAttempts: 2,
		Storage:         storage,
	})

	// New user should be allowed.
	userID := int64(1)
	allowed, err := control.Check(context.Background(), userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !allowed {
		t.Errorf("unexpected result: %v", allowed)
	}

	// User has 1 attempt, should be allowed.
	allowed, err = control.Check(context.Background(), userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !allowed {
		t.Errorf("unexpected result: %v", allowed)
	}

	// User already has 2 attempts, should not be allowed.
	allowed, err = control.Check(context.Background(), userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if allowed {
		t.Errorf("unexpected result: %v", allowed)
	}
}
