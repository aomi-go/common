package ginx

import (
	"fmt"
	"github.com/aomi-go/common/exception"
	"github.com/aomi-go/common/exception/errorcode"
	lerrors "github.com/aomi-go/common/lang/errorx"
	"github.com/aomi-go/common/web/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// SuccessHandler 响应处理器
type SuccessHandler func(c *gin.Context, payload any)

// ErrorHandler 错误处理器
// @return int 错误码 any 错误数据
type ErrorHandler func(c *gin.Context, payload any, err error)

var (
	successHandler SuccessHandler = DefaultSuccessHandler()
	errorHandler   ErrorHandler   = CreateErrHandler(nil)
)

func SetSuccessHandler(h SuccessHandler) error {
	if nil == h {
		return fmt.Errorf("success handler cannot be nil")
	}
	successHandler = h
	return nil
}

func SetErrorHandler(h ErrorHandler) error {
	if nil == h {
		return fmt.Errorf("error handler cannot be nil")
	}
	errorHandler = h
	return nil
}

func HandlerResponse(c *gin.Context, payload any, err error) {
	if nil != err {
		errorHandler(c, payload, err)
	} else {
		successHandler(c, payload)
	}
}

func DefaultSuccessHandler() SuccessHandler {
	return func(c *gin.Context, payload any) {
		c.JSON(http.StatusOK, dto.Result{
			Status:  errorcode.SUCCESS,
			Payload: payload,
		})
	}
}

// CreateErrHandler 创建错误处理器
// @param err2msg 错误转换器, 返回值：http状态码、错误数据、是否成功处理
func CreateErrHandler(
	postHandler func(c *gin.Context, payload any, err error, httpStatus int, finalPayload any) (int, any),
) ErrorHandler {
	return func(c *gin.Context, payload any, err error) {

		var httpStatus int = http.StatusOK
		var p any

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
				Status:   errorcode.ParamsError,
				Describe: "params error",
				Payload:  errorMsgs,
			}
		case *lerrors.IllegalArgumentError:
			p = dto.Result{
				Status:   errorcode.ParamsError,
				Describe: e.Msg,
			}
		case *exception.ServiceError:
			p = dto.Result{
				Status:   e.Code,
				Describe: e.Msg,
				Payload:  p,
			}
		default:
			p = dto.Result{
				Status:   errorcode.EXCEPTION,
				Describe: "server internal error",
			}
		}

		if nil != postHandler {
			httpStatus, p = postHandler(c, payload, err, httpStatus, p)
		}

		c.JSON(httpStatus, p)
	}
}
