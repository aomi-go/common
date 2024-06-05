package controller

import (
	"errors"
	"fmt"
	"github.com/aomi-go/common/exception"
	"github.com/aomi-go/common/exception/errorcode"
	"github.com/aomi-go/common/web/dto"
	http2 "github.com/aomi-go/common/web/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func NewHandler(success SuccessHandler, err ErrorHandler) *Handler {
	return &Handler{
		success: success,
		err:     err,
	}
}

func DefaultSuccessHandler() SuccessHandler {
	return func(c *gin.Context, payload any) {
		reqId := c.GetHeader(http2.RequestId)

		c.JSON(http.StatusOK, dto.Result{
			Status:    errorcode.SUCCESS,
			Payload:   payload,
			RequestId: reqId,
		})
	}
}

func DefaultErrHandler() ErrorHandler {
	return func(c *gin.Context, err error) {

		reqId := c.GetHeader(http2.RequestId)

		var validErr validator.ValidationErrors
		if errors.As(err, &validErr) {
			errorMsgs := make(map[string]string)
			for _, e := range validErr {
				fieldName := e.Field()
				tagName := e.Tag()
				errorMessage := fmt.Sprintf("字段 %s 校验失败，条件：%s", fieldName, tagName)
				errorMsgs[fieldName] = errorMessage
			}

			c.JSON(http.StatusOK, &dto.Result{
				Status:    errorcode.ParamsError,
				Describe:  "params error",
				Payload:   errorMsgs,
				RequestId: reqId,
			})
			return
		}

		var e *exception.ServiceError
		if errors.As(err, &e) {
			c.JSON(http.StatusOK, dto.Result{
				Status:    e.Code,
				Describe:  e.Msg,
				Payload:   e.Payload,
				RequestId: reqId,
			})
			return
		}

		c.JSON(http.StatusOK, dto.Result{
			Status:    errorcode.EXCEPTION,
			Describe:  err.Error(),
			RequestId: reqId,
		})
	}
}

// SuccessHandler 响应处理器
type SuccessHandler func(c *gin.Context, payload any)

// ErrorHandler 错误处理器
// @return int 错误码 any 错误数据
type ErrorHandler func(c *gin.Context, err error)

type Handler struct {

	/*
		success 响应数据给客户端
	*/
	success SuccessHandler

	err ErrorHandler
}

/*
Wrapper 处理器包装
提供以下功能:
1. 支持返回任意类型的数据
2. 支持返回错误信息
*/
func (h *Handler) Wrapper(handler interface{}) func(c *gin.Context) {
	return func(c *gin.Context) {
		var res interface{}
		var err error

		switch handler := handler.(type) {
		case func(c *gin.Context):
			handler(c)
		case func(c *gin.Context) (interface{}, error):
			res, err = handler(c)
		}

		if nil == err {
			h.success(c, res)
		} else {
			h.err(c, err)
		}
	}
}
