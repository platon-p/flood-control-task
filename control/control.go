package control

import (
	"context"
	"log"
	"time"
)

type FloodControlImpl struct {
	allowedPeriod   time.Duration
	allowedAttempts int
	storage         AttemptsStorage
	logger          *log.Logger
}

type FloodControlOptions struct {
	AllowedPeriod   time.Duration
	AllowedAttempts int
	Storage         AttemptsStorage
	Logger          *log.Logger
}

func NewFloodControl(opts FloodControlOptions) *FloodControlImpl {
	return &FloodControlImpl{
		allowedPeriod:   opts.AllowedPeriod,
		allowedAttempts: opts.AllowedAttempts,
		storage:         opts.Storage,
		logger:          opts.Logger,
	}
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
		c.logger.Printf("user %d exceeded the allowed number of attempts", userID)
		return false, nil
	}
	// grant access.
	err = c.storage.RegisterAttempt(ctx, userID, time.Now())
	if err != nil {
		return false, err
	}
	c.logger.Printf("user %d granted access", userID)
	return true, nil
}
