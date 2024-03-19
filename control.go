package main

import "context"

type FloodControlImpl struct {
}

func (c *FloodControlImpl) Check(context context.Context, userID int64) (bool, error) {
    panic("not implemented")
}
