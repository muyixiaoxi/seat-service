package router

import (
	"github.com/gin-gonic/gin"
	"seat-service/middleware"
)

func Router() {
	g := gin.Default()
	//跨域问题
	g.Use(middleware.Cors())
	//路由组声明
	userGroup := userRouter{}

	//路由组
	group := g.Group("")
	group.Use(middleware.Jwt())
	{
		//用户路由组
		userGroup.userRouterGroup(group)
	}

	g.Run(":8080")
}
