package controller

import (
	"fmt"
	"github.com/aomi-go/common/exception"
	"github.com/aomi-go/common/exception/errorcode"
	lerrors "github.com/aomi-go/common/lang/errorx"
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

// CreateErrHandler 创建错误处理器
// @param err2msg 错误转换器, 返回值：http状态码、错误数据、是否成功处理
func CreateErrHandler(
	err2msg func(c *gin.Context, payload any, err error) (int, any, bool),
	postHandler func(c *gin.Context, payload any, err error, httpStatus int, finalPayload any) (int, any),
) ErrorHandler {
	return func(c *gin.Context, payload any, err error) {

		reqId := c.GetHeader(http2.RequestId)

		var httpStatus int
		var p any
		var handled bool

		if nil != err2msg {
			httpStatus, p, handled = err2msg(c, payload, err)
		}

		if !handled {
			httpStatus = http.StatusOK

			switch e := err.(type) {
			case validator.ValidationErrors:
				errorMsgs := make(map[string]string)
				for _, item := range e {
					fieldName := item.Field()
					tagName := item.Tag()
					errorMessage := fmt.Sprintf("字段 %s 校验失败，条件：%s", fieldName, tagName)
					errorMsgs[fieldName] = errorMessage
				}

				p = dto.Result{
					Status:    errorcode.ParamsError,
					Describe:  "params error",
					Payload:   errorMsgs,
					RequestId: reqId,
				}
			case *lerrors.IllegalArgumentError:
				p = dto.Result{
					Status:    errorcode.ParamsError,
					Describe:  e.Msg,
					RequestId: reqId,
				}
			case *exception.ServiceError:
				p = dto.Result{
					Status:    e.Code,
					Describe:  e.Msg,
					Payload:   p,
					RequestId: reqId,
				}
			default:
				p = dto.Result{
					Status:    errorcode.EXCEPTION,
					Describe:  err.Error(),
					RequestId: reqId,
				}
			}
		}

		if nil != postHandler {
			httpStatus, p = postHandler(c, payload, err, httpStatus, p)
		}

		c.JSON(httpStatus, p)
	}
}

// SuccessHandler 响应处理器
type SuccessHandler func(c *gin.Context, payload any)

// ErrorHandler 错误处理器
// @return int 错误码 any 错误数据
type ErrorHandler func(c *gin.Context, payload any, err error)

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
		case func(c *gin.Context) error:
			err = handler(c)
		case func(c *gin.Context) (interface{}, error):
			res, err = handler(c)
		}

		if nil == err {
			h.success(c, res)
		} else {
			h.err(c, res, err)
		}
	}
}
