package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"seat-service/utils"
)

func AuthMiddlewareCasbin(srv *utils.CasbinService) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := srv.Enforcer.LoadPolicy()
		if err != nil {
			zap.L().Error("srv.Enforcer.LoadPolicy() is failed", zap.Error(err))
			return
		}
		//获取当前用户信息
		currentUser := getCurrentUser(context)

		//获取请求的路径,方法
		requestPath := context.Request.URL.Path
		requestMethod := context.Request.Method

		//检查权限
		ok, err := srv.Enforcer.Enforce(currentUser, requestPath, requestMethod)
		if err != nil {
			zap.L().Error("Enforcer.Enforce(sub , obj , act) if failed", zap.Error(err))
			return
		}
		if ok {
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
