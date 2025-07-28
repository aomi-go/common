package future

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestAllOf(t *testing.T) {

	values, err, errs := AllOf(context.Background(),
		func(ctx context.Context) (interface{}, error) {
			return "1", nil
		},
		func(ctx context.Context) (interface{}, error) {
			return 1, nil
		},
		func(ctx context.Context) (interface{}, error) {
			time.Sleep(time.Second)
			return nil, errors.New("error")
		},
	)

	fmt.Println(values)
	fmt.Println(errs)
	fmt.Println(err)
}
