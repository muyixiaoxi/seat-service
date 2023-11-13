package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"seat-service/initialization"
	"seat-service/model"
	"seat-service/response"
	service "seat-service/service/impl"
	"seat-service/utils"
)

var userService service.UserService
var resp response.CustomResponse
var jwt utils.JWT

type UserApi struct {
}

func (u *UserApi) Test(context *gin.Context) {
	a := userService.Test(1)
	resp.Success(context, response.CodeSuccess, a)
}

// Login 用户登录
func (u *UserApi) Login(c *gin.Context) {
	user := &model.User{}
	if err := c.ShouldBind(user); err != nil {
		initialization.SeatLogger.Error("c.ShouldBind(u) failed ", zap.Error(err))
		resp.Fail(c, response.CodeLoginFailure, nil)
		return
	}
	if err := userService.Login(user); err != nil {
		initialization.SeatLogger.Error("c.ShouldBind(u) failed ", zap.Error(err))
		resp.Fail(c, response.CodeServerBusy, nil)
		return
	}
	if user.ID == 0 {
		resp.Fail(c, response.CodeLoginFailure, nil)
		return
	}

	userClaims := utils.UserClaims{
		Username: user.Username,
		ID:       user.ID,
	}
	token, err := jwt.GenToken(userClaims)
	if err != nil {
		initialization.SeatLogger.Error("jwt.GenToken(userClaims) failed", zap.Error(err))
		resp.Fail(c, response.CodeServerBusy, nil)
		return
	}
	resp.Success(c, response.CodeSuccess, token)
}
