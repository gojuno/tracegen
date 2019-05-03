package tests

import "context"

type Example interface {
	Set(ctx context.Context, key, value []byte) error
	Get(ctx context.Context, key []byte) (value []byte, err error)
}
