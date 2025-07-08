package numberx

import (
	"fmt"
	"strconv"
)

func ParseInt64Zero(value interface{}) int64 {
	v, _ := ParseInt64(value)
	return v
}

// ParseInt64 转换任意类型为int64
// v: 任意类型数字
func ParseInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case string:
		// 尝试解析字符串为浮点数
		num, err := strconv.ParseInt(v, 10, 64)
		if nil != err {
			return 0, err
		}
		return num, nil
	default:
		return 0, fmt.Errorf("unsupported type")
	}
}

// ParseFloat64Zero 转换任意类型为float64，失败返回0
func ParseFloat64Zero(value interface{}) float64 {
	v, _ := ParseFloat64(value)
	return v
}

// ParseFloat64 转换任意类型为float64
func ParseFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		// 尝试解析字符串为浮点数
		num, err := strconv.ParseFloat(v, 64)
		if nil != err {
			return 0, err
		}
		return num, nil
	default:
		return 0, fmt.Errorf("unsupported type")
	}
}
