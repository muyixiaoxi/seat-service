package api

import (
	"github.com/gin-gonic/gin"
	"seat-service/response"
	service "seat-service/service/impl"
)

var menuService service.MenuService

type MenuApi struct {
}

func (*MenuApi) GetMenu(c *gin.Context) {
	menus := menuService.GetMenu()
	resp.Success(c, response.CodeSuccess, menus)
}
