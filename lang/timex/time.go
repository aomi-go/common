package timex

import (
	"fmt"
	"strconv"
)

// ParseTimestamp 解析时间戳
// v: 时间戳 任意类型
func ParseTimestamp(value interface{}) (int64, error) {
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
