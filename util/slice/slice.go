package slice

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
	var result []T
	for _, item := range v {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T any, R any](s []T, f func(T) R) []R {
	var result []R
	for _, item := range s {
		result = append(result, f(item))
	}
	return result
}
