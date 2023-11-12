package middleware

import (
	"github.com/gin-gonic/gin"
	"seat-service/response"
	"seat-service/utils"
	"strconv"
	"strings"
	"time"
)

var resp response.CustomResponse

func Jwt() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		if token == "" {
			resp.Fail(context, response.CodeIllegalLogin, gin.H{"reload": true})
			context.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "seat-token") {
			resp.Fail(context, response.CodeLoginFailure, gin.H{"reload": true})
			context.Abort()
			return
		}
		jwt := utils.NewJWT()
		mc, err := jwt.ParseToken(parts[1])
		if err != nil && err.Error() != "Token is expired" {
			resp.Fail(context, response.CodeTokenInvalid, gin.H{"reload": true})
			context.Abort()
			return
		}
		//续签
		if mc.ExpiresAt.Unix()-time.Now().Unix() < time.Now().Add(time.Hour*time.Duration(mc.Buffer)).Unix() || err.Error() == "Token is expired" {
			newToken, _ := jwt.CreateTokenByOldToken(token, mc.UserClaims)
			newClaims, _ := jwt.ParseToken(newToken)
			context.Header("new-seat-token", newToken)
			context.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
		}

		// 将当前请求的username信息保存到请求的上下文c上
		context.Set("userClaims", mc.UserClaims)
		context.Set("id", mc.UserClaims.ID)
		context.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
