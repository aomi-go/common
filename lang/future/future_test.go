package future

import (
	"fmt"
	"testing"
)

func TestAllOf(t *testing.T) {

	values, errs, ok := AllOf(
		func(idx int, vc chan<- any, ec chan<- error, done chan<- struct{}) {
			vc <- 1
		},
	)

	fmt.Println(values)
	fmt.Println(errs)
	fmt.Println(ok)
}
