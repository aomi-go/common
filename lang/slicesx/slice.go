package slicesx

func FindFirst[T any](v []T, predicate func(item T) bool) (T, bool) {
	for _, item := range v {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func Find[T any](v []T, predicate func(item T) bool) []T {
	var result = make([]T, 0)
	for _, item := range v {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T any, R any](s []T, f func(T) R) []R {
	var result = make([]R, len(s))
	for i, item := range s {
		result[i] = f(item)
	}
	return result
}
