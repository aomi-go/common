package assert

import "github.com/aomi-go/common/errorx"

func IsNil(v any, msg string) error {
	if v != nil {
		return errorx.NewParamsError().WithMsg(msg)
	}
	return nil
}

func NotNil(v any, msg string) error {
	if v == nil {
		return errorx.NewParamsError().WithMsg(msg)
	}
	return nil
}

// IsEmpty 判断是否为空
// 支持类型：数组、切片、字典、字符串
func IsEmpty(v any, msg string) error {
	if nil == v {
		return nil
	}
	switch v.(type) {
	case string:
		return StrIsEmpty(v.(string), msg)
	case []any:
		return SliceIsEmpty(v.([]any), msg)
	case map[any]any:
		return MapIsEmpty(v.(map[any]any), msg)
	}
	return nil
}

// IsNotEmpty 判断是否不为空
// 支持类型: 数组、切片、字典、字符串
func IsNotEmpty(v any, msg string) error {
	if nil == v {
		return errorx.NewParamsError().WithMsg(msg)
	}
	switch v.(type) {
	case string:
		return StrIsNotEmpty(v.(string), msg)
	case []any:
		return SliceIsNotEmpty(v.([]any), msg)
	case map[any]any:
		return MapIsNotEmpty(v.(map[any]any), msg)
	}
	return nil
}

func StrIsEmpty(v string, msg string) error {
	if v == "" {
		return nil
	}
	return errorx.NewParamsError().WithMsg(msg)
}

func StrIsNotEmpty(v string, msg string) error {
	if v == "" {
		return errorx.NewParamsError().WithMsg(msg)
	}
	return nil
}

func SliceIsEmpty[T any](v []T, msg string) error {
	if nil == v || len(v) == 0 {
		return nil
	}
	return errorx.NewParamsError().WithMsg(msg)
}
func SliceIsNotEmpty[T interface{}](v []T, msg string) error {
	if nil == v || len(v) == 0 {
		return errorx.NewParamsError().WithMsg(msg)
	}
	return nil
}

func MapIsEmpty[K comparable, V any](v map[K]V, msg string) error {
	if nil == v || len(v) == 0 {
		return nil
	}
	return errorx.NewParamsError().WithMsg(msg)
}

func MapIsNotEmpty[K comparable, V any](v map[K]V, msg string) error {
	if nil == v || len(v) == 0 {
		return errorx.NewParamsError().WithMsg(msg)
	}
	return nil
}
