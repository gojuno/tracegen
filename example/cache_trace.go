/*
This code was automatically generated using github.com/gojuno/generator lib.
			Please DO NOT modify.
*/
package example

import (
	context "context"

	opentracing "github.com/opentracing/opentracing-go"
)

type CacheTracer struct {
	next   Cache
	prefix string
}

func NewCacheTracer(next Cache, prefix string) *CacheTracer {
	return &CacheTracer{
		next:   next,
		prefix: prefix,
	}
}

func (t *CacheTracer) Get(ctx context.Context, key []byte) (value []byte, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, t.prefix+".Cache.Get")
	defer span.Finish()

	return t.next.Get(ctx, key)
}

func (t *CacheTracer) Set(ctx context.Context, key []byte, value []byte) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, t.prefix+".Cache.Set")
	defer span.Finish()

	return t.next.Set(ctx, key, value)
}
