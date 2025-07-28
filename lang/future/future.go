package future

import (
	"context"
	"golang.org/x/sync/errgroup"
)

type TaskFunc[T any] func(ctx context.Context) (T, error)

/**
 * AllOf
 * @Description: 并发执行多个任务，等待所有任务完成后返回结果
 * @param tasks 任务列表
 * @return []any 所有任务的返回值
 * @return error 最后一个错误信息
 * @return []error 所有任务的错误
 */
func AllOf(ctx context.Context, tasks ...TaskFunc[any]) ([]any, error, []error) {
	eg, ctx := errgroup.WithContext(ctx)

	results := make([]any, len(tasks))
	errs := make([]error, len(tasks))

	for idx, item := range tasks {
		i, task := idx, item // 避免闭包变量问题
		eg.Go(func() error {
			val, err := task(ctx)
			if err != nil {
				errs[i] = err
				return err // 记录到 errgroup 中
			}
			results[i] = val
			return nil
		})
	}

	// 返回最后一个错误和所有错误
	if err := eg.Wait(); err != nil {
		return results, err, errs
	}
	return results, nil, nil
}
