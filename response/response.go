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
func FailBasedCode(code int, message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
