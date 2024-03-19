package main

import (
	"context"
	"time"
)

type FloodControlImpl struct {
	allowedPeriod   time.Duration
	allowedAttempts int
	storage         AttemptsStorage
}

type AttemptsStorage interface {
	GetAttemptsCount(ctx context.Context, userID int64, period time.Duration) (int, error)
	RegisterAttempt(ctx context.Context, userID int64, timestamp time.Time) error
}

func (c *FloodControlImpl) Check(ctx context.Context, userID int64) (bool, error) {
	attempts, err := c.storage.GetAttemptsCount(ctx, userID, c.allowedPeriod)
	if err != nil {
		return false, err
	}
	if attempts >= c.allowedAttempts {
		return false, nil
	}
	// grant access.
	err = c.storage.RegisterAttempt(ctx, userID, time.Now())
	if err != nil {
		return false, err
	}
	return true, nil
}
