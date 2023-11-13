package router

import (
	"github.com/gin-gonic/gin"
	"seat-service/api"
)

type menuRouter struct {
}

func (*menuRouter) menuRouterGroup(Router *gin.RouterGroup) {
	menu := Router.Group("menu")
	menuApi := api.MenuApi{}
	{
		menu.GET("menu", menuApi.GetMenu)
	}
}
