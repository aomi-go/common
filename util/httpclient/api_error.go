package httpclient

import "fmt"

type ApiError struct {
	// Code 错误代码
	Code string
	// Msg 错误描述
	Msg string
	// 堆栈
	Stack error

	HttpResponse any
}

// 实现 error 接口的 Error() 方法
func (e *ApiError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.Code, e.Msg)
}
