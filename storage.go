package main

import (
	"sync"
	"time"
)

type AttemptsMemoryStorage struct {
	attempts map[int64][]time.Time
	lock     sync.RWMutex
}

func (s *AttemptsMemoryStorage) GetAttemptsCount(userID int64, period time.Duration) (int, error) {
	// RLock is used because we are reading the map.
	s.lock.RLock()
	defer s.lock.RUnlock()
	attempts, ok := s.attempts[userID]
	if !ok {
		return 0, nil
	}
	count := 0
	// select attempts that are not older than given period.
	for _, t := range attempts {
		if time.Since(t) < period {
			count++
		}
	}
	return count, nil
}

func (s *AttemptsMemoryStorage) RegisterAttempt(userID int64, timestamp time.Time) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	attempts, ok := s.attempts[userID]
	if !ok {
		attempts = make([]time.Time, 0)
	}
	s.attempts[userID] = append(attempts, timestamp)
	return nil
}
