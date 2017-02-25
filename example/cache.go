package example

import "context"

type Cache interface {
	Set(ctx context.Context, key, value []byte) error
	Get(ctx context.Context, key []byte) (value []byte, err error)
}
