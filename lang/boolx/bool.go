package boolx

import (
	"fmt"
	"strconv"
)

// ParseBoolDefault 解析bool值，任意失败都设置为false
func ParseBoolDefault(v interface{}) bool {
	if r, e := ParseBool(v); e == nil {
		return r
	}
	return false
}

// ParseBool 解析bool值
func ParseBool(v interface{}) (bool, error) {
	switch x := v.(type) {
	case bool:
		return x, nil
	case *bool:
		if x == nil {
			return false, nil
		}
		return *x, nil
	case string:
		if r, err := strconv.ParseBool(x); err == nil {
			return r, nil
		} else {
			return false, err
		}
	default:
		return false, fmt.Errorf("%T is not a bool", v)
	}
}
