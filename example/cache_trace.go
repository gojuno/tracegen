/*
This code was automatically generated using github.com/gojuno/generator lib.
			Please DO NOT modify.
*/
package example

import (
	context "context"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
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

func (t *CacheTracer) Get(in context.Context, in1 []byte) (out []byte, out1 error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, t.prefix+".Cache.Get")
	defer func() {
		ext.Error.Set(span, out1 != nil)
		span.Finish()
	}()

	return t.next.Get(in, in1)
}

func (t *CacheTracer) Set(in context.Context, in1 []byte, in2 []byte) (out error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, t.prefix+".Cache.Set")
	defer func() {
		ext.Error.Set(span, out != nil)
		span.Finish()
	}()

	return t.next.Set(in, in1, in2)
}
