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

func (u *UserApi) Test(c *gin.Context) {
	a := userService.Test(1)
	resp.Success(response.CodeSuccess, a)
}
