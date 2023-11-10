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

type CustomResponse struct {
	Context *gin.Context
}

func (cr *CustomResponse) Success(code ResCode, data interface{}) {
	cr.Context.JSON(http.StatusOK, Response{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	})
}

func (cr *CustomResponse) Fail(code ResCode, data interface{}) {
	cr.Context.JSON(http.StatusOK, Response{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	})
}
func (cr *CustomResponse) FailBasedCode(code ResCode, data interface{}) {
	cr.Context.JSON(http.StatusOK, Response{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	})
}
