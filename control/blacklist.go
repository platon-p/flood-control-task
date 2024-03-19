package control

import (
	"context"
	"log"
)

type BlackListControl struct {
	storage BlackListStorage
	logger  *log.Logger
}

func NewBlackListControl(storage BlackListStorage, logger *log.Logger) *BlackListControl {
	return &BlackListControl{
		storage: storage,
		logger:  logger,
	}
}

type BlackListStorage interface {
	Contains(ctx context.Context, userID int64) (bool, error)
}

func (c *BlackListControl) Check(ctx context.Context, userID int64) (bool, error) {
	contains, err := c.storage.Contains(ctx, userID)
	if err != nil {
		return false, err
	}
	if contains {
		c.logger.Printf("user %d is in the blacklist", userID)
		return false, nil
	}
	c.logger.Printf("user %d is not in the blacklist", userID)
	return true, nil
}
