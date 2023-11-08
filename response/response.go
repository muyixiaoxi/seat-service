package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func Fail(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    500,
		Message: message,
		Data:    data,
	})
}
