package api

import (
	"github.com/gin-gonic/gin"
	"seat-service/response"
	service "seat-service/service/impl"
)

var userService service.UserService
var resp response.CustomResponse

type UserApi struct {
}

func (u *UserApi) Test(context *gin.Context) {
	a := userService.Test(1)
	resp.Success(context, response.CodeSuccess, a)
}
