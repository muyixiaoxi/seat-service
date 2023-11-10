package middleware

import (
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var enforcer *casbin.Enforcer

func init() {
	//初始化 Casbin 和 Enforcer 适配器
	a, err := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/", true)
	if err != nil {
		zap.L().Error("初始化Gorm适配器失败")
		return
	}
	enforcer = casbin.NewEnforcer("../model.config", a)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取当前用户信息
		currentUser := getCurrentUser(context)

		//获取请求的路径,方法
		requestPath := context.Request.URL.Path
		requestMethod := context.Request.Method

		//检查权限
		if enforcer.Enforce(currentUser, requestPath, requestMethod) {
			//用户存在权限 , 继续请求处理
			context.Next()
		} else {
			//用户无权限 , 返回错误
			context.AbortWithStatus(403)
		}
	}
}

func getCurrentUser(context *gin.Context) string {
	//从会话或其它方式获取当前用户信息
	return "user"
}
