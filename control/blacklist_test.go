package control

import (
	"context"
	"testing"
)

type BlackListStorageMock struct {
	blacklist map[int64]struct{}
}

func NewBlackListStorageMock() *BlackListStorageMock {
	return &BlackListStorageMock{
		blacklist: make(map[int64]struct{}),
	}
}

func (s *BlackListStorageMock) Contains(ctx context.Context, userID int64) (bool, error) {
	_, exists := s.blacklist[userID]
	return exists, nil
}

func (s *BlackListStorageMock) Add(userID int64) error {
	s.blacklist[userID] = struct{}{}
	return nil
}

func TestBlackListControl(t *testing.T) {
	storage := NewBlackListStorageMock()
	control := NewBlackListControl(storage)

	userID := int64(1)

	// User is not in the blacklist.
	allowed, err := control.Check(context.Background(), userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !allowed {
		t.Errorf("unexpected result: %v", allowed)
	}

	// Add user to the blacklist.
	err = storage.Add(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// User is in the blacklist.
	allowed, err = control.Check(context.Background(), userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if allowed {
		t.Errorf("unexpected result: %v", allowed)
	}
}
