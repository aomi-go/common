package assert

import "github.com/aomi-go/common/exception"

func IsNil(v any, msg string) error {
	if v != nil {
		return exception.NewParamsError(msg, nil)
	}
	return nil
}

func NotNil(v any, msg string) error {
	if v == nil {
		return exception.NewParamsError(msg, nil)
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
		return exception.NewParamsError(msg, nil)
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
	return exception.NewParamsError(msg, nil)
}

func StrIsNotEmpty(v string, msg string) error {
	if v == "" {
		return exception.NewParamsError(msg, nil)
	}
	return nil
}

func SliceIsEmpty(v []any, msg string) error {
	if nil == v || len(v) == 0 {
		return nil
	}
	return exception.NewParamsError(msg, nil)
}
func SliceIsNotEmpty(v []any, msg string) error {
	if nil == v || len(v) == 0 {
		return exception.NewParamsError(msg, nil)
	}
	return nil
}

func MapIsEmpty(v map[any]any, msg string) error {
	if nil == v || len(v) == 0 {
		return nil
	}
	return exception.NewParamsError(msg, nil)
}

func MapIsNotEmpty(v map[any]any, msg string) error {
	if nil == v || len(v) == 0 {
		return exception.NewParamsError(msg, nil)
	}
	return nil
}
