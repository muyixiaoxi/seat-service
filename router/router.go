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
	}

	g.Run(":8080")
}
