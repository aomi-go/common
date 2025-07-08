package timex

import "testing"

func TestParseTimestamp(t *testing.T) {

	v, err := ParseTimestamp("1734080429")
	v, err = ParseTimestamp(1734080429)
	v, err = ParseTimestamp(1734080429.0)
	if err != nil {
		return
	}
	print(v)
}
