package errorx

import (
	"fmt"
)

type ServiceError struct {
	// Code 错误代码
	Code string
	// Msg 错误描述
	Msg string
	// 错误返回的数据
	Payload any
	// 堆栈
	Stack error
}

// 实现 error 接口的 Error() 方法
func (e *ServiceError) Error() string {
	return fmt.Sprintf("[%s][%s][%v]", e.Code, e.Msg, e.Payload)
}

func (e *ServiceError) WithCode(code string) *ServiceError {
	e.Code = code
	return e
}
func (e *ServiceError) WithMsg(msg string) *ServiceError {
	e.Msg = msg
	return e
}
func (e *ServiceError) WithPayload(payload any) *ServiceError {
	e.Payload = payload
	return e
}
func (e *ServiceError) WithStack(stack error) *ServiceError {
	e.Stack = stack
	return e
}
