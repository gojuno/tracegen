Tracegen
========

Generates interface decorators with [opentracing](http://opentracing.io) support.

Installation
------------

```
go get github.com/gojuno/tracegen
```

Example
-------

```go
type Cache interface {
	Set(ctx context.Context, key, value []byte) error
	Get(ctx context.Context, key []byte) (value []byte, err error)
}
```

```
tracegen -i Cache -o example/cache_trace.go example
```

Will generate:
```go
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
```
