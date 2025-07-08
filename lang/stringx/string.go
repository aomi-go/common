package stringx

import "strings"

func EqualsIgnoreCase(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}
