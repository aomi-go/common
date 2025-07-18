package cache

import (
	"context"
)

func Wrapper[T any](ctx context.Context, c Cache, entry Entry, suffix string, f func() (*T, error)) (*T, error) {
	helper := NewHelper(c).
		WithContext(ctx).
		WithEntry(entry).
		WithSuffix(suffix)

	var r T
	err := helper.Get(&r)
	if nil == err {
		return &r, err
	}
	ur, err := f()
	if nil != err {
		return nil, err
	}
	go func() {
		_ = helper.WithContext(context.Background()).Value(ur).Set()
	}()
	return ur, nil
}
