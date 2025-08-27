package stringx

import "strings"

func EqualsIgnoreCase(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}

func IsEmpty(str string) bool {
	return len(str) == 0 || strings.Trim(str, " ") == ""
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}
