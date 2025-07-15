package calculator

import (
	"github.com/shopspring/decimal"
)

func Add(v1, v2 interface{}) decimal.Decimal {
	if nil == v1 || nil == v2 {
		return decimal.Zero
	}
	d1 := ToDecimal(v1)
	d2 := ToDecimal(v2)
	return d1.Add(d2)
}

func Sub(v1, v2 interface{}) decimal.Decimal {
	if nil == v1 || nil == v2 {
		return decimal.Zero
	}
	d1 := ToDecimal(v1)
	d2 := ToDecimal(v2)
	return d1.Sub(d2)
}

func Mul(v1, v2 interface{}) decimal.Decimal {
	if nil == v1 || nil == v2 {
		return decimal.Zero
	}
	d1 := ToDecimal(v1)
	d2 := ToDecimal(v2)
	return d1.Mul(d2)
}

func Div(v1, v2 interface{}) decimal.Decimal {
	if nil == v1 || nil == v2 {
		return decimal.Zero
	}
	d1 := ToDecimal(v1)
	d2 := ToDecimal(v2)
	if d2.IsZero() {
		return decimal.Zero
	}
	return d1.Div(d2)
}

func ToDecimal(value interface{}) decimal.Decimal {
	if nil == value {
		return decimal.Zero
	}
	switch v := value.(type) {
	case int:
		return decimal.NewFromInt(int64(v))
	case int8:
		return decimal.NewFromInt(int64(v))
	case int16:
		return decimal.NewFromInt(int64(v))
	case int32:
		return decimal.NewFromInt32(v)
	case int64:
		return decimal.NewFromInt(v)
	case float32:
		return decimal.NewFromFloat32(v)
	case float64:
		return decimal.NewFromFloat(v)
	case string:
		d, err := decimal.NewFromString(v)
		if err != nil {
			return decimal.Zero
		}
		return d
	default:
		return decimal.Zero
	}
}
