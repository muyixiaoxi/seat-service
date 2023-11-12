package router

import (
	"github.com/gin-gonic/gin"
	"seat-service/api"
)

type authRouter struct {
}

func (a *authRouter) authRouterGroup(Router *gin.RouterGroup) (auth *gin.RouterGroup) {
	auth = Router.Group("/auth")
	authApi := api.AuthApi{}
	{
		auth.GET("/casbin/users", authApi.GetUsers)                 //获取所有用户
		auth.GET("/casbin/roles", authApi.GetRoles)                 //获取所有角色组
		auth.GET("/casbin/getAllRolePolicy", authApi.GetRolePolicy) //获取所有角色组的策略
		auth.POST("/casbin/createRole", authApi.CreateRolePolicy)   //添加角色组
		auth.DELETE("/casbin/user-role", authApi.DeleteUserRole)    //从组中删除用户
		auth.POST("/casbin/updateRole", authApi.UpdateUserRole)     //添加用户组策略
	}
	return
}
