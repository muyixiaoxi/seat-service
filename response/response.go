package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    ResCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CustomResponse struct{}

func (cr *CustomResponse) Success(context *gin.Context, code ResCode, data interface{}) {
	context.JSON(http.StatusOK, Response{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	})
}

func (cr *CustomResponse) Fail(context *gin.Context, code ResCode, data interface{}) {
	context.JSON(http.StatusOK, Response{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	})
}
