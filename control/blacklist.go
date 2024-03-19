package control

import "context"

type BlackListControl struct {
	storage BlackListStorage
}

func NewBlackListControl(storage BlackListStorage) *BlackListControl {
	return &BlackListControl{
		storage: storage,
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
	return !contains, nil
}
