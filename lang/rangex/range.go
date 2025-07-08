package rangex

import "cmp"

func Closed[T cmp.Ordered](from, to T) Range[T] {
	return Range[T]{
		LowerBound: Inclusive(from),
		UpperBound: Inclusive(to),
	}
}

func Open[T cmp.Ordered](from, to T) Range[T] {
	return Range[T]{
		LowerBound: Exclusive(from),
		UpperBound: Exclusive(to),
	}
}

func LeftOpen[T cmp.Ordered](from, to T) Range[T] {
	return Range[T]{
		LowerBound: Exclusive(from),
		UpperBound: Inclusive(to),
	}
}

func RightOpen[T cmp.Ordered](from, to T) Range[T] {
	return Range[T]{
		LowerBound: Inclusive(from),
		UpperBound: Exclusive(to),
	}
}

func LeftUnbounded[T cmp.Ordered](to Bound[T]) Range[T] {
	return Range[T]{
		LowerBound: Unbounded[T](),
		UpperBound: to,
	}
}

func RightUnbounded[T cmp.Ordered](from Bound[T]) Range[T] {
	return Range[T]{
		LowerBound: from,
		UpperBound: Unbounded[T](),
	}
}

func Of[T cmp.Ordered](lowerBound Bound[T], upperBound Bound[T]) Range[T] {
	return Range[T]{
		LowerBound: lowerBound,
		UpperBound: upperBound,
	}
}

type Range[T cmp.Ordered] struct {
	LowerBound Bound[T]
	UpperBound Bound[T]
}

// Contains 判断值是否在范围内
func (r Range[T]) Contains(value T) bool {

	greaterThanLowerBound := true
	if nil != r.LowerBound.Value {
		it := *r.LowerBound.Value
		if r.LowerBound.Inclusive {
			greaterThanLowerBound = cmp.Compare(it, value) <= 0
		} else {
			greaterThanLowerBound = cmp.Compare(it, value) < 0
		}
	}
	lessThanUpperBound := true
	if nil != r.UpperBound.Value {
		it := *r.UpperBound.Value
		if r.UpperBound.Inclusive {
			lessThanUpperBound = cmp.Compare(it, value) >= 0
		} else {
			lessThanUpperBound = cmp.Compare(it, value) > 0
		}
	}
	return greaterThanLowerBound && lessThanUpperBound
}
