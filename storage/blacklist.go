package storage

import (
	"context"
	"sync"
)

type BlackListMemoryStorage struct {
	blacklist map[int64]struct{}
	lock      sync.RWMutex
}

func NewBlackListMemoryStorage() *BlackListMemoryStorage {
	blacklist := make(map[int64]struct{})
	return &BlackListMemoryStorage{
		blacklist: blacklist,
		lock:      sync.RWMutex{},
	}
}

func (s *BlackListMemoryStorage) Contains(ctx context.Context, userID int64) (bool, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.blacklist[userID]
	return ok, nil
}
