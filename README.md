Tracegen
========

Generates interface decorators with opentracing support.

Installation
------------

```go get github.com/mkabischev/tracegen```

Example
-------

```go
type Cache interface {
	Set(ctx context.Context, key, value []byte) error
	Get(ctx context.Context, key []byte) (value []byte, err error)
}
```

```tracegen -i Cache -o example/cache_trace.go example```

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
```
