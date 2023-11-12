package middleware

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"seat-service/initialization"
	"seat-service/utils"
)

var Enforcer *casbin.Enforcer

func init() {
	//初始化 Casbin 和 Enforcer 适配器
	dsn := initialization.Config.Mysql.Dsn()
	a, err := gormadapter.NewAdapter("mysql", dsn, true)
	if err != nil {
		zap.L().Error("gormadapter.NewAdapter() is failed", zap.Error(err))
		return
	}
	Enforcer, err = casbin.NewEnforcer("../model.config", a)
	if err != nil {
		zap.L().Error("casbin.NewEnforcer() is failed", zap.Error(err))
		return
	}
	//将自定义权限匹配规则加入权限认证器
	Enforcer.AddFunction("my_func", utils.KeyMatchFunc)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取当前用户信息
		currentUser := getCurrentUser(context)

		//获取请求的路径,方法
		requestPath := context.Request.URL.Path
		requestMethod := context.Request.Method

		//检查权限
		ok, err := Enforcer.Enforce(currentUser, requestPath, requestMethod)
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
