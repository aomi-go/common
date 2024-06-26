package cache

import "context"

// NewHelper creates a new Helper.
func NewHelper(c Cache) Helper {
	return Helper{
		c: c,
	}
}

type Helper struct {
	c      Cache
	ctx    context.Context
	entry  Entry
	suffix string
	v      any
}

func (b Helper) WithContext(ctx context.Context) Helper {
	b.ctx = ctx
	return b
}

func (b Helper) WithEntry(entry Entry) Helper {
	b.entry = entry
	if "" != b.suffix {
		b.entry = b.entry.SetKeySuffix(b.suffix)
	}
	return b
}

func (b Helper) WithSuffix(suffix string) Helper {
	b.suffix = suffix
	b.entry = b.entry.SetKeySuffix(suffix)
	return b
}

func (b Helper) Value(v any) Helper {
	b.v = v
	return b
}

func (b Helper) Set() error {
	return b.c.Set(b.ctx, b.entry.Key, b.v, b.entry.Ttl)
}

func (b Helper) Get(value interface{}) error {
	return b.c.Get(b.ctx, b.entry.Key, value)
}
