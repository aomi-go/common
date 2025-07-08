package to

// Ptr returns a pointer to the provided value.
func Ptr[T any](v T) *T {
	return &v
}

// SliceOfPtrs returns a slice of *T from the specified values.
func SliceOfPtrs[T any](vv ...T) []*T {
	slc := make([]*T, len(vv))
	for i := range vv {
		slc[i] = Ptr(vv[i])
	}
	return slc
}

// Deref 指针变量转换为普通变量
func Deref[T any](v *T) T {
	if nil == v {
		var zero T
		return zero
	}
	return *v
}

func SliceOfDeref[T any](vv ...*T) []T {
	slc := make([]T, len(vv))
	for i := range vv {
		slc[i] = Deref(vv[i])
	}
	return slc
}
