package future

import (
	"fmt"
	"sync"
)

type TaskFunc[T any] func(idx int, vc chan<- T, ec chan<- error, done chan<- struct{})

/**
 * AllOf
 * @Description: 并发执行多个任务，等待所有任务完成后返回结果
 * @param tasks 任务列表
 * @return []any 所有任务的返回值
 * @return error 最后一个错误信息
 * @return []error 所有任务的错误
 */
func AllOf(tasks ...TaskFunc[any]) ([]any, error, []error) {
	var wg sync.WaitGroup

	valueChans := make([]chan any, len(tasks))
	errChans := make([]chan error, len(tasks))
	done := make(chan struct{})

	for i, item := range tasks {
		valueChans[i] = make(chan any, 1)
		errChans[i] = make(chan error, 1)
		wg.Add(1)
		go func(idx int, task TaskFunc[any]) {
			defer wg.Done()
			select {
			case <-done:
				fmt.Println("done")
				return
			default:
				task(idx, valueChans[idx], errChans[idx], done)
			}
		}(i, item)
	}
	wg.Wait()

	// 读取 所有返回值和错误
	var results = make([]any, len(tasks))
	var errs = make([]error, len(tasks))
	var err error

	for i := range tasks {
		select {
		case e := <-errChans[i]:
			errs[i] = e
			err = e
		case val := <-valueChans[i]:
			results[i] = val
		default:
			// 如果某个通道没有数据，可以选择忽略或进行其他处理
		}
	}
	// 关闭所有通道
	for _, ch := range valueChans {
		close(ch)
	}
	for _, ch := range errChans {
		close(ch)
	}

	return results, err, errs
}
