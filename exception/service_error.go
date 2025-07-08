package exception

import (
	"fmt"
	"github.com/aomi-go/common/exception/errorcode"
)

var (
	ServerInternalError = NewError("server internal error", nil)
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

func NewFullError(code string, msg string, payload any, stack error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Msg:     msg,
		Payload: payload,
		Stack:   stack,
	}
}

func NewPartialSuccess(msg string, payload any, stack error) *ServiceError {
	return &ServiceError{
		Code:    errorcode.PartialSuccess,
		Msg:     msg,
		Payload: payload,
		Stack:   stack,
	}
}

func NewError(msg string, payload any) *ServiceError {
	return NewFullError(
		errorcode.EXCEPTION,
		msg,
		payload,
		nil,
	)
}

func NewErrorWithStack(msg string, payload any, stack error) *ServiceError {
	return NewFullError(
		errorcode.EXCEPTION,
		msg,
		payload,
		stack,
	)
}

func NewParamsError(msg string, payload any) *ServiceError {
	return NewParamsErrorWithStack(msg, payload, nil)
}

func NewParamsErrorWithStack(msg string, payload any, stack error) *ServiceError {
	return NewFullError(
		errorcode.ParamsError,
		msg,
		payload,
		stack,
	)
}

func NewUnauthorized(msg string, payload any, stack error) *ServiceError {
	return NewFullError(
		errorcode.Unauthorized,
		msg,
		payload,
		stack,
	)
}

func NewBadCredentials(msg string, payload any, stack error) *ServiceError {
	return NewFullError(
		errorcode.BadCredentials,
		msg,
		payload,
		stack,
	)
}

func NewAccessDenied(msg string, payload any, stack error) *ServiceError {
	return NewFullError(
		errorcode.AccessDenied,
		msg,
		payload,
		stack,
	)
}

func NewResourceExist(msg string, payload any, stack error) *ServiceError {
	return NewFullError(
		errorcode.ResourceExist,
		msg,
		payload,
		stack,
	)
}

func NewResourceNotExist(msg string, payload any, stack error) *ServiceError {
	return NewFullError(
		errorcode.ResourceNotExist,
		msg,
		payload,
		stack,
	)
}
