package rangex

import "cmp"

func Inclusive[T cmp.Ordered](value T) Bound[T] {
	return Bound[T]{
		Value:     &value,
		Inclusive: true,
	}
}
func Exclusive[T cmp.Ordered](value T) Bound[T] {
	return Bound[T]{
		Value:     &value,
		Inclusive: false,
	}
}

func Unbounded[T cmp.Ordered]() Bound[T] {
	return Bound[T]{
		Value:     nil,
		Inclusive: true,
	}
}

// Bound 表示区间端点类型
type Bound[T cmp.Ordered] struct {
	Value     *T
	Inclusive bool
}
