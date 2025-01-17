package router

import (
	"github.com/gin-gonic/gin"
	"seat-service/initialization"
	"seat-service/middleware"
	"seat-service/utils"
)

func Router() {
	g := gin.Default()
	//跨域问题
	g.Use(middleware.Cors())
	//路由组声明
	userGroup := userRouter{}
	authGroup := authRouter{}
	menuGroup := menuRouter{}

	//Casbin
	casbin, err := utils.InitCasbinGorm(initialization.DB)
	if err != nil {
		return
	}

	//路由组
	group := g.Group("")
	group.Use(middleware.Jwt(), middleware.AuthMiddlewareCasbin(casbin))
	{
		//用户路由组
		userGroup.userRouterGroup(group)
		//作者路由组
		authGroup.authRouterGroup(group)
		//菜单路由组
		menuGroup.menuRouterGroup(group)
	}

	g.Run(":8080")
}
