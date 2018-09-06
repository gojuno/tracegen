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
	span, in := opentracing.StartSpanFromContext(in, t.prefix+".Cache.Get")
	defer func() {
		if out1 != nil {
			ext.Error.Set(span, true)
			span.LogFields(
				log.String("event", "error"),
				log.String("message", out1.Error()),
			)
		}
		span.Finish()
	}()

	return t.next.Get(in, in1)
}

func (t *CacheTracer) Set(in context.Context, in1 []byte, in2 []byte) (out error) {
	span, in := opentracing.StartSpanFromContext(in, t.prefix+".Cache.Set")
	defer func() {
		if out != nil {
			ext.Error.Set(span, true)
			span.LogFields(
				log.String("event", "error"),
				log.String("message", out.Error()),
			)
		}
		span.Finish()
	}()

	return t.next.Set(in, in1, in2)
}
```
