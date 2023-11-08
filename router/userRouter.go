package router

import (
	"github.com/gin-gonic/gin"
	"seat-service/api"
)

type userRouter struct {
}

func (*userRouter) userRouterGroup(Router *gin.RouterGroup) *gin.RouterGroup {
	user := Router.Group("user")
	userApi := api.UserApi{}
	{
		user.POST("user", userApi.Test)
	}

	return user
}
